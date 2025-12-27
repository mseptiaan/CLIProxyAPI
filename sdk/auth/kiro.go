package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/router-for-me/CLIProxyAPI/v6/internal/config"
	"github.com/router-for-me/CLIProxyAPI/v6/internal/util"
	coreauth "github.com/router-for-me/CLIProxyAPI/v6/sdk/cliproxy/auth"
	log "github.com/sirupsen/logrus"
)

const (
	kiroRefreshTokenURL = "https://prod.us-east-1.auth.desktop.kiro.dev/refreshToken"
	kiroAPIEndpoint     = "https://cloudcode-pa.googleapis.com"
	kiroAPIVersion      = "v1internal"
	kiroAPIUserAgent    = "google-api-nodejs-client/9.15.1"
	kiroAPIClient       = "google-cloud-sdk vscode_cloudshelleditor/0.1"
	kiroClientMetadata  = `{"ideType":"IDE_UNSPECIFIED","platform":"PLATFORM_UNSPECIFIED","pluginType":"GEMINI"}`
)

// KiroAuthenticator implements OAuth login for the Kiro provider.
type KiroAuthenticator struct{}

// NewKiroAuthenticator constructs a new authenticator instance.
func NewKiroAuthenticator() Authenticator { return &KiroAuthenticator{} }

// Provider returns the provider key for kiro.
func (KiroAuthenticator) Provider() string { return "kiro" }

// RefreshLead instructs the manager to refresh five minutes before expiry.
func (KiroAuthenticator) RefreshLead() *time.Duration {
	lead := 5 * time.Minute
	return &lead
}

// Login performs authentication using a refresh token from Kiro IDE.
// The refresh token is typically obtained by intercepting Kiro IDE traffic.
func (KiroAuthenticator) Login(ctx context.Context, cfg *config.Config, opts *LoginOptions) (*coreauth.Auth, error) {
	if cfg == nil {
		return nil, fmt.Errorf("cliproxy auth: configuration is required")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if opts == nil {
		opts = &LoginOptions{}
	}

	// Get refresh token from user input
	if opts.Prompt == nil {
		return nil, fmt.Errorf("kiro: prompt function is required for refresh token input")
	}

	refreshToken, errPrompt := opts.Prompt("Enter your Kiro refresh token: ")
	if errPrompt != nil {
		return nil, fmt.Errorf("kiro: failed to get refresh token: %w", errPrompt)
	}

	refreshToken = strings.TrimSpace(refreshToken)
	if refreshToken == "" {
		return nil, fmt.Errorf("kiro: refresh token cannot be empty")
	}

	httpClient := util.SetProxy(&cfg.SDKConfig, &http.Client{})

	// Exchange refresh token for access token
	tokenResp, errToken := exchangeKiroRefreshToken(ctx, refreshToken, httpClient)
	if errToken != nil {
		return nil, fmt.Errorf("kiro: token exchange failed: %w", errToken)
	}

	// Fetch project ID via loadCodeAssist (same approach as Kiro gateway)
	projectID := ""
	if tokenResp.AccessToken != "" {
		fetchedProjectID, errProject := fetchKiroProjectID(ctx, tokenResp.AccessToken, httpClient)
		if errProject != nil {
			log.Warnf("kiro: failed to fetch project ID: %v", errProject)
		} else {
			projectID = fetchedProjectID
			log.Infof("kiro: obtained project ID %s", projectID)
		}
	}

	now := time.Now()
	expiresAt := now.Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
	
	metadata := map[string]any{
		"type":          "kiro",
		"access_token":  tokenResp.AccessToken,
		"refresh_token": refreshToken,
		"expires_in":    tokenResp.ExpiresIn,
		"timestamp":     now.UnixMilli(),
		"expired":       expiresAt.Format(time.RFC3339),
	}
	
	if tokenResp.ProfileArn != "" {
		metadata["profile_arn"] = tokenResp.ProfileArn
	}
	if tokenResp.Region != "" {
		metadata["region"] = tokenResp.Region
	}
	if projectID != "" {
		metadata["project_id"] = projectID
	}

	fileName := "kiro.json"
	label := "kiro"
	if projectID != "" {
		label = fmt.Sprintf("kiro (%s)", projectID)
	}

	fmt.Println("Kiro authentication successful")
	if projectID != "" {
		fmt.Printf("Using GCP project: %s\n", projectID)
	}
	
	return &coreauth.Auth{
		ID:       fileName,
		Provider: "kiro",
		FileName: fileName,
		Label:    label,
		Metadata: metadata,
	}, nil
}

type kiroTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken,omitempty"`
	ExpiresIn    int64  `json:"expiresIn"`
	ProfileArn   string `json:"profileArn,omitempty"`
	Region       string `json:"region,omitempty"`
}

// exchangeKiroRefreshToken exchanges a Kiro refresh token for an access token.
func exchangeKiroRefreshToken(ctx context.Context, refreshToken string, httpClient *http.Client) (*kiroTokenResponse, error) {
	reqBody := map[string]string{
		"refreshToken": refreshToken,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, kiroRefreshTokenURL, strings.NewReader(string(jsonBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, errDo := httpClient.Do(req)
	if errDo != nil {
		return nil, errDo
	}
	defer func() {
		if errClose := resp.Body.Close(); errClose != nil {
			log.Errorf("kiro token exchange: close body error: %v", errClose)
		}
	}()

	bodyBytes, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return nil, fmt.Errorf("read response: %w", errRead)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("token exchange failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var token kiroTokenResponse
	if errDecode := json.Unmarshal(bodyBytes, &token); errDecode != nil {
		return nil, fmt.Errorf("decode response: %w", errDecode)
	}

	return &token, nil
}

// fetchKiroProjectID retrieves the project ID for the authenticated user via loadCodeAssist.
// This uses the same approach as Kiro OpenAI Gateway to get the cloudaicompanionProject.
func fetchKiroProjectID(ctx context.Context, accessToken string, httpClient *http.Client) (string, error) {
	// Call loadCodeAssist to get the project
	loadReqBody := map[string]any{
		"metadata": map[string]string{
			"ideType":    "IDE_UNSPECIFIED",
			"platform":   "PLATFORM_UNSPECIFIED",
			"pluginType": "GEMINI",
		},
	}

	rawBody, errMarshal := json.Marshal(loadReqBody)
	if errMarshal != nil {
		return "", fmt.Errorf("marshal request body: %w", errMarshal)
	}

	endpointURL := fmt.Sprintf("%s/%s:loadCodeAssist", kiroAPIEndpoint, kiroAPIVersion)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, strings.NewReader(string(rawBody)))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", kiroAPIUserAgent)
	req.Header.Set("X-Goog-Api-Client", kiroAPIClient)
	req.Header.Set("Client-Metadata", kiroClientMetadata)

	resp, errDo := httpClient.Do(req)
	if errDo != nil {
		return "", fmt.Errorf("execute request: %w", errDo)
	}
	defer func() {
		if errClose := resp.Body.Close(); errClose != nil {
			log.Errorf("kiro loadCodeAssist: close body error: %v", errClose)
		}
	}()

	bodyBytes, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return "", fmt.Errorf("read response: %w", errRead)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return "", fmt.Errorf("request failed with status %d: %s", resp.StatusCode, strings.TrimSpace(string(bodyBytes)))
	}

	var loadResp map[string]any
	if errDecode := json.Unmarshal(bodyBytes, &loadResp); errDecode != nil {
		return "", fmt.Errorf("decode response: %w", errDecode)
	}

	// Extract projectID from response
	projectID := ""
	if id, ok := loadResp["cloudaicompanionProject"].(string); ok {
		projectID = strings.TrimSpace(id)
	}
	if projectID == "" {
		if projectMap, ok := loadResp["cloudaicompanionProject"].(map[string]any); ok {
			if id, okID := projectMap["id"].(string); okID {
				projectID = strings.TrimSpace(id)
			}
		}
	}

	if projectID == "" {
		return "", fmt.Errorf("no cloudaicompanionProject in response")
	}

	return projectID, nil
}

// FetchKiroProjectID exposes project discovery for external callers.
func FetchKiroProjectID(ctx context.Context, accessToken string, httpClient *http.Client) (string, error) {
	return fetchKiroProjectID(ctx, accessToken, httpClient)
}
