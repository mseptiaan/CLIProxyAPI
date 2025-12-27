// Package executor provides runtime execution capabilities for various AI service providers.
// This file implements the Kiro executor that proxies requests to the Kiro IDE
// upstream using refresh tokens obtained from the Kiro IDE.
package executor

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/router-for-me/CLIProxyAPI/v6/internal/config"
	"github.com/router-for-me/CLIProxyAPI/v6/internal/util"
	cliproxyauth "github.com/router-for-me/CLIProxyAPI/v6/sdk/cliproxy/auth"
	cliproxyexecutor "github.com/router-for-me/CLIProxyAPI/v6/sdk/cliproxy/executor"
	sdktranslator "github.com/router-for-me/CLIProxyAPI/v6/sdk/translator"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	kiroRefreshTokenURL    = "https://prod.us-east-1.auth.desktop.kiro.dev/refreshToken"
	kiroBaseURL            = "https://cloudcode-pa.googleapis.com"
	kiroAPIVersion         = "v1internal"
	kiroStreamPath         = "/v1internal:streamGenerateContent"
	kiroGeneratePath       = "/v1internal:generateContent"
	kiroCountTokensPath    = "/v1internal:countTokens"
	kiroAuthType           = "kiro"
	kiroRefreshSkew        = 300 * time.Second // 5 minutes before expiry
	kiroAPIUserAgent       = "google-api-nodejs-client/9.15.1"
	kiroAPIClient          = "google-cloud-sdk vscode_cloudshell_editor/0.1"
	kiroClientMetadata     = `{"ideType":"IDE_UNSPECIFIED","platform":"PLATFORM_UNSPECIFIED","pluginType":"GEMINI"}`
)

// KiroExecutor proxies requests to the Kiro upstream.
type KiroExecutor struct {
	cfg *config.Config
}

// NewKiroExecutor creates a new Kiro executor instance.
//
// Parameters:
//   - cfg: The application configuration
//
// Returns:
//   - *KiroExecutor: A new Kiro executor instance
func NewKiroExecutor(cfg *config.Config) *KiroExecutor {
	return &KiroExecutor{cfg: cfg}
}

// Identifier returns the executor identifier.
func (e *KiroExecutor) Identifier() string { return kiroAuthType }

// PrepareRequest prepares the HTTP request for execution (no-op for Kiro).
func (e *KiroExecutor) PrepareRequest(_ *http.Request, _ *cliproxyauth.Auth) error { return nil }

// Execute performs a non-streaming request to the Kiro API.
func (e *KiroExecutor) Execute(ctx context.Context, auth *cliproxyauth.Auth, req cliproxyexecutor.Request, opts cliproxyexecutor.Options) (resp cliproxyexecutor.Response, err error) {
	if strings.Contains(req.Model, "claude") {
		return e.executeClaudeNonStream(ctx, auth, req, opts)
	}

	token, updatedAuth, errToken := e.ensureAccessToken(ctx, auth)
	if errToken != nil {
		return resp, errToken
	}
	if updatedAuth != nil {
		auth = updatedAuth
	}

	reporter := newUsageReporter(ctx, e.Identifier(), req.Model, auth)
	defer reporter.trackFailure(ctx, &err)

	from := opts.SourceFormat
	to := sdktranslator.FromString("antigravity")
	translated := sdktranslator.TranslateRequest(from, to, req.Model, bytes.Clone(req.Payload), false)

	translated = applyThinkingMetadataCLI(translated, req.Metadata, req.Model)
	translated = util.ApplyGemini3ThinkingLevelFromMetadataCLI(req.Model, req.Metadata, translated)
	translated = util.ApplyDefaultThinkingIfNeededCLI(req.Model, translated)
	translated = normalizeAntigravityThinking(req.Model, translated)
	translated = applyPayloadConfigWithRoot(e.cfg, req.Model, "kiro", "request", translated)

	httpClient := newProxyAwareHTTPClient(ctx, e.cfg, auth, 0)
	httpReq, errReq := e.buildRequest(ctx, auth, token, req.Model, translated, false, opts.Alt)
	if errReq != nil {
		err = errReq
		return resp, err
	}

	httpResp, errDo := httpClient.Do(httpReq)
	if errDo != nil {
		recordAPIResponseError(ctx, e.cfg, errDo)
		err = errDo
		return resp, err
	}

	recordAPIResponseMetadata(ctx, e.cfg, httpResp.StatusCode, httpResp.Header.Clone())
	bodyBytes, errRead := io.ReadAll(httpResp.Body)
	if errClose := httpResp.Body.Close(); errClose != nil {
		log.Errorf("kiro executor: close response body error: %v", errClose)
	}
	if errRead != nil {
		recordAPIResponseError(ctx, e.cfg, errRead)
		err = errRead
		return resp, err
	}
	appendAPIResponseChunk(ctx, e.cfg, bodyBytes)

	if httpResp.StatusCode < http.StatusOK || httpResp.StatusCode >= http.StatusMultipleChoices {
		log.Debugf("kiro executor: upstream error status: %d, body: %s", httpResp.StatusCode, summarizeErrorBody(httpResp.Header.Get("Content-Type"), bodyBytes))
		err = statusErr{code: httpResp.StatusCode, msg: string(bodyBytes)}
		return resp, err
	}

	reporter.publish(ctx, parseAntigravityUsage(bodyBytes))
	var param any
	converted := sdktranslator.TranslateNonStream(ctx, to, from, req.Model, bytes.Clone(opts.OriginalRequest), translated, bodyBytes, &param)
	resp = cliproxyexecutor.Response{Payload: []byte(converted)}
	reporter.ensurePublished(ctx)
	return resp, nil
}

// executeClaudeNonStream performs a claude non-streaming request to the Kiro API.
func (e *KiroExecutor) executeClaudeNonStream(ctx context.Context, auth *cliproxyauth.Auth, req cliproxyexecutor.Request, opts cliproxyexecutor.Options) (resp cliproxyexecutor.Response, err error) {
	token, updatedAuth, errToken := e.ensureAccessToken(ctx, auth)
	if errToken != nil {
		return resp, errToken
	}
	if updatedAuth != nil {
		auth = updatedAuth
	}

	reporter := newUsageReporter(ctx, e.Identifier(), req.Model, auth)
	defer reporter.trackFailure(ctx, &err)

	from := opts.SourceFormat
	to := sdktranslator.FromString("antigravity")
	translated := sdktranslator.TranslateRequest(from, to, req.Model, bytes.Clone(req.Payload), true)

	translated = applyThinkingMetadataCLI(translated, req.Metadata, req.Model)
	translated = util.ApplyGemini3ThinkingLevelFromMetadataCLI(req.Model, req.Metadata, translated)
	translated = util.ApplyDefaultThinkingIfNeededCLI(req.Model, translated)
	translated = normalizeAntigravityThinking(req.Model, translated)
	translated = applyPayloadConfigWithRoot(e.cfg, req.Model, "kiro", "request", translated)

	httpClient := newProxyAwareHTTPClient(ctx, e.cfg, auth, 0)
	httpReq, errReq := e.buildRequest(ctx, auth, token, req.Model, translated, true, opts.Alt)
	if errReq != nil {
		err = errReq
		return resp, err
	}

	httpResp, errDo := httpClient.Do(httpReq)
	if errDo != nil {
		recordAPIResponseError(ctx, e.cfg, errDo)
		err = errDo
		return resp, err
	}

	recordAPIResponseMetadata(ctx, e.cfg, httpResp.StatusCode, httpResp.Header.Clone())
	if httpResp.StatusCode < http.StatusOK || httpResp.StatusCode >= http.StatusMultipleChoices {
		bodyBytes, errRead := io.ReadAll(httpResp.Body)
		if errClose := httpResp.Body.Close(); errClose != nil {
			log.Errorf("kiro executor: close response body error: %v", errClose)
		}
		if errRead != nil {
			recordAPIResponseError(ctx, e.cfg, errRead)
			err = errRead
			return resp, err
		}
		appendAPIResponseChunk(ctx, e.cfg, bodyBytes)
		err = statusErr{code: httpResp.StatusCode, msg: string(bodyBytes)}
		return resp, err
	}

	// Collect streaming data
	scanner := bufio.NewScanner(httpResp.Body)
	scanner.Buffer(nil, streamScannerBuffer)
	var collectedData bytes.Buffer
	for scanner.Scan() {
		line := scanner.Bytes()
		appendAPIResponseChunk(ctx, e.cfg, line)
		collectedData.Write(line)
		collectedData.WriteByte('\n')
	}
	if errClose := httpResp.Body.Close(); errClose != nil {
		log.Errorf("kiro executor: close response body error: %v", errClose)
	}
	if errScan := scanner.Err(); errScan != nil {
		recordAPIResponseError(ctx, e.cfg, errScan)
		err = errScan
		return resp, err
	}

	// Convert stream to non-stream response
	nonStreamBody := e.convertStreamToNonStream(collectedData.Bytes())
	reporter.publish(ctx, parseAntigravityUsage(nonStreamBody))

	var param any
	converted := sdktranslator.TranslateNonStream(ctx, to, from, req.Model, bytes.Clone(opts.OriginalRequest), translated, nonStreamBody, &param)
	resp = cliproxyexecutor.Response{Payload: []byte(converted)}
	reporter.ensurePublished(ctx)
	return resp, nil
}

// convertStreamToNonStream converts streaming responses to non-streaming format
func (e *KiroExecutor) convertStreamToNonStream(stream []byte) []byte {
	responseTemplate := ""
	var traceID string
	var finishReason string
	var modelVersion string
	var responseID string
	var role string
	var usageRaw string
	parts := make([]map[string]interface{}, 0)
	var pendingKind string
	var pendingText strings.Builder
	var pendingThoughtSig string

	flushPending := func() {
		if pendingKind == "" {
			return
		}
		text := pendingText.String()
		switch pendingKind {
		case "text":
			if strings.TrimSpace(text) == "" {
				pendingKind = ""
				pendingText.Reset()
				pendingThoughtSig = ""
				return
			}
			parts = append(parts, map[string]interface{}{"text": text})
		case "thought":
			if strings.TrimSpace(text) == "" && pendingThoughtSig == "" {
				pendingKind = ""
				pendingText.Reset()
				pendingThoughtSig = ""
				return
			}
			part := map[string]interface{}{"thought": true}
			part["text"] = text
			if pendingThoughtSig != "" {
				part["thoughtSignature"] = pendingThoughtSig
			}
			parts = append(parts, part)
		}
		pendingKind = ""
		pendingText.Reset()
		pendingThoughtSig = ""
	}

	normalizePart := func(partResult gjson.Result) map[string]interface{} {
		var m map[string]interface{}
		_ = json.Unmarshal([]byte(partResult.Raw), &m)
		if m == nil {
			m = map[string]interface{}{}
		}
		sig := partResult.Get("thoughtSignature").String()
		if sig == "" {
			sig = partResult.Get("thought_signature").String()
		}
		if sig != "" {
			m["thoughtSignature"] = sig
			delete(m, "thought_signature")
		}
		if inlineData, ok := m["inline_data"]; ok {
			m["inlineData"] = inlineData
			delete(m, "inline_data")
		}
		return m
	}

	for _, line := range bytes.Split(stream, []byte("\n")) {
		trimmed := bytes.TrimSpace(line)
		if len(trimmed) == 0 || !gjson.ValidBytes(trimmed) {
			continue
		}

		root := gjson.ParseBytes(trimmed)
		responseNode := root.Get("response")
		if !responseNode.Exists() {
			if root.Get("candidates").Exists() {
				responseNode = root
			} else {
				continue
			}
		}
		responseTemplate = responseNode.Raw

		if traceResult := root.Get("traceId"); traceResult.Exists() && traceResult.String() != "" {
			traceID = traceResult.String()
		}

		if roleResult := responseNode.Get("candidates.0.content.role"); roleResult.Exists() {
			role = roleResult.String()
		}

		if finishResult := responseNode.Get("candidates.0.finishReason"); finishResult.Exists() && finishResult.String() != "" {
			finishReason = finishResult.String()
		}

		if modelResult := responseNode.Get("modelVersion"); modelResult.Exists() && modelResult.String() != "" {
			modelVersion = modelResult.String()
		}
		if responseIDResult := responseNode.Get("responseId"); responseIDResult.Exists() && responseIDResult.String() != "" {
			responseID = responseIDResult.String()
		}
		if usageResult := responseNode.Get("usageMetadata"); usageResult.Exists() {
			usageRaw = usageResult.Raw
		} else if usageResult := root.Get("usageMetadata"); usageResult.Exists() {
			usageRaw = usageResult.Raw
		}

		if partsResult := responseNode.Get("candidates.0.content.parts"); partsResult.IsArray() {
			for _, part := range partsResult.Array() {
				hasFunctionCall := part.Get("functionCall").Exists()
				hasInlineData := part.Get("inlineData").Exists() || part.Get("inline_data").Exists()
				sig := part.Get("thoughtSignature").String()
				if sig == "" {
					sig = part.Get("thought_signature").String()
				}
				text := part.Get("text").String()
				thought := part.Get("thought").Bool()

				if hasFunctionCall || hasInlineData {
					flushPending()
					parts = append(parts, normalizePart(part))
					continue
				}

				if thought || part.Get("text").Exists() {
					kind := "text"
					if thought {
						kind = "thought"
					}
					if pendingKind != "" && pendingKind != kind {
						flushPending()
					}
					pendingKind = kind
					pendingText.WriteString(text)
					if kind == "thought" && sig != "" {
						pendingThoughtSig = sig
					}
					continue
				}

				flushPending()
				parts = append(parts, normalizePart(part))
			}
		}
	}
	flushPending()

	if responseTemplate == "" {
		responseTemplate = `{"candidates":[{"content":{"role":"model","parts":[]}}]}`
	}

	partsJSON, _ := json.Marshal(parts)
	responseTemplate, _ = sjson.SetRaw(responseTemplate, "candidates.0.content.parts", string(partsJSON))
	if role != "" {
		responseTemplate, _ = sjson.Set(responseTemplate, "candidates.0.content.role", role)
	}
	if finishReason != "" {
		responseTemplate, _ = sjson.Set(responseTemplate, "candidates.0.finishReason", finishReason)
	}
	if modelVersion != "" {
		responseTemplate, _ = sjson.Set(responseTemplate, "modelVersion", modelVersion)
	}
	if responseID != "" {
		responseTemplate, _ = sjson.Set(responseTemplate, "responseId", responseID)
	}
	if usageRaw != "" {
		responseTemplate, _ = sjson.SetRaw(responseTemplate, "usageMetadata", usageRaw)
	} else if !gjson.Get(responseTemplate, "usageMetadata").Exists() {
		responseTemplate, _ = sjson.Set(responseTemplate, "usageMetadata.promptTokenCount", 0)
		responseTemplate, _ = sjson.Set(responseTemplate, "usageMetadata.candidatesTokenCount", 0)
		responseTemplate, _ = sjson.Set(responseTemplate, "usageMetadata.totalTokenCount", 0)
	}

	output := `{"response":{},"traceId":""}`
	output, _ = sjson.SetRaw(output, "response", responseTemplate)
	if traceID != "" {
		output, _ = sjson.Set(output, "traceId", traceID)
	}
	return []byte(output)
}

// ExecuteStream performs a streaming request to the Kiro API.
func (e *KiroExecutor) ExecuteStream(ctx context.Context, auth *cliproxyauth.Auth, req cliproxyexecutor.Request, opts cliproxyexecutor.Options) (stream <-chan cliproxyexecutor.StreamChunk, err error) {
	ctx = context.WithValue(ctx, "alt", "")

	token, updatedAuth, errToken := e.ensureAccessToken(ctx, auth)
	if errToken != nil {
		return nil, errToken
	}
	if updatedAuth != nil {
		auth = updatedAuth
	}

	reporter := newUsageReporter(ctx, e.Identifier(), req.Model, auth)
	defer reporter.trackFailure(ctx, &err)

	from := opts.SourceFormat
	to := sdktranslator.FromString("antigravity")
	translated := sdktranslator.TranslateRequest(from, to, req.Model, bytes.Clone(req.Payload), true)

	translated = applyThinkingMetadataCLI(translated, req.Metadata, req.Model)
	translated = util.ApplyGemini3ThinkingLevelFromMetadataCLI(req.Model, req.Metadata, translated)
	translated = util.ApplyDefaultThinkingIfNeededCLI(req.Model, translated)
	translated = normalizeAntigravityThinking(req.Model, translated)
	translated = applyPayloadConfigWithRoot(e.cfg, req.Model, "kiro", "request", translated)

	httpClient := newProxyAwareHTTPClient(ctx, e.cfg, auth, 0)
	httpReq, errReq := e.buildRequest(ctx, auth, token, req.Model, translated, true, opts.Alt)
	if errReq != nil {
		err = errReq
		return nil, err
	}

	httpResp, errDo := httpClient.Do(httpReq)
	if errDo != nil {
		recordAPIResponseError(ctx, e.cfg, errDo)
		err = errDo
		return nil, err
	}

	recordAPIResponseMetadata(ctx, e.cfg, httpResp.StatusCode, httpResp.Header.Clone())
	if httpResp.StatusCode < http.StatusOK || httpResp.StatusCode >= http.StatusMultipleChoices {
		bodyBytes, errRead := io.ReadAll(httpResp.Body)
		if errClose := httpResp.Body.Close(); errClose != nil {
			log.Errorf("kiro executor: close response body error: %v", errClose)
		}
		if errRead != nil {
			recordAPIResponseError(ctx, e.cfg, errRead)
			err = errRead
			return nil, err
		}
		appendAPIResponseChunk(ctx, e.cfg, bodyBytes)
		err = statusErr{code: httpResp.StatusCode, msg: string(bodyBytes)}
		return nil, err
	}

	out := make(chan cliproxyexecutor.StreamChunk)
	stream = out
	go func(resp *http.Response) {
		defer close(out)
		defer func() {
			if errClose := resp.Body.Close(); errClose != nil {
				log.Errorf("kiro executor: close response body error: %v", errClose)
			}
		}()
		scanner := bufio.NewScanner(resp.Body)
		scanner.Buffer(nil, streamScannerBuffer)
		var param any
		for scanner.Scan() {
			line := scanner.Bytes()
			appendAPIResponseChunk(ctx, e.cfg, line)

			// Filter usage metadata for all models
			// Only retain usage statistics in the terminal chunk
			line = FilterSSEUsageMetadata(line)

			payload := jsonPayload(line)
			if payload == nil {
				continue
			}

			if detail, ok := parseAntigravityStreamUsage(payload); ok {
				reporter.publish(ctx, detail)
			}

			chunks := sdktranslator.TranslateStream(ctx, to, from, req.Model, bytes.Clone(opts.OriginalRequest), translated, bytes.Clone(payload), &param)
			for i := range chunks {
				out <- cliproxyexecutor.StreamChunk{Payload: []byte(chunks[i])}
			}
		}
		tail := sdktranslator.TranslateStream(ctx, to, from, req.Model, bytes.Clone(opts.OriginalRequest), translated, []byte("[DONE]"), &param)
		for i := range tail {
			out <- cliproxyexecutor.StreamChunk{Payload: []byte(tail[i])}
		}
		if errScan := scanner.Err(); errScan != nil {
			recordAPIResponseError(ctx, e.cfg, errScan)
			reporter.publishFailure(ctx)
			out <- cliproxyexecutor.StreamChunk{Err: errScan}
		} else {
			reporter.ensurePublished(ctx)
		}
	}(httpResp)
	return stream, nil
}

// Refresh refreshes the authentication credentials using the refresh token.
func (e *KiroExecutor) Refresh(ctx context.Context, auth *cliproxyauth.Auth) (*cliproxyauth.Auth, error) {
	if auth == nil {
		return auth, nil
	}
	updated, errRefresh := e.refreshToken(ctx, auth.Clone())
	if errRefresh != nil {
		return nil, errRefresh
	}
	return updated, nil
}

// CountTokens counts tokens for the given request using the Kiro API.
func (e *KiroExecutor) CountTokens(ctx context.Context, auth *cliproxyauth.Auth, req cliproxyexecutor.Request, opts cliproxyexecutor.Options) (cliproxyexecutor.Response, error) {
	token, updatedAuth, errToken := e.ensureAccessToken(ctx, auth)
	if errToken != nil {
		return cliproxyexecutor.Response{}, errToken
	}
	if updatedAuth != nil {
		auth = updatedAuth
	}
	if strings.TrimSpace(token) == "" {
		return cliproxyexecutor.Response{}, statusErr{code: http.StatusUnauthorized, msg: "missing access token"}
	}

	from := opts.SourceFormat
	to := sdktranslator.FromString("antigravity")
	respCtx := context.WithValue(ctx, "alt", opts.Alt)

	httpClient := newProxyAwareHTTPClient(ctx, e.cfg, auth, 0)

	payload := sdktranslator.TranslateRequest(from, to, req.Model, bytes.Clone(req.Payload), false)
	payload = applyThinkingMetadataCLI(payload, req.Metadata, req.Model)
	payload = util.ApplyDefaultThinkingIfNeededCLI(req.Model, payload)
	payload = normalizeAntigravityThinking(req.Model, payload)
	payload = deleteJSONField(payload, "project")
	payload = deleteJSONField(payload, "model")
	payload = deleteJSONField(payload, "request.safetySettings")

	httpReq, errReq := e.buildCountTokensRequest(ctx, auth, token, req.Model, payload)
	if errReq != nil {
		return cliproxyexecutor.Response{}, errReq
	}

	httpResp, errDo := httpClient.Do(httpReq)
	if errDo != nil {
		return cliproxyexecutor.Response{}, errDo
	}

	bodyBytes, errRead := io.ReadAll(httpResp.Body)
	if errClose := httpResp.Body.Close(); errClose != nil {
		log.Errorf("kiro executor: close response body error: %v", errClose)
	}
	if errRead != nil {
		return cliproxyexecutor.Response{}, errRead
	}

	if httpResp.StatusCode < http.StatusOK || httpResp.StatusCode >= http.StatusMultipleChoices {
		return cliproxyexecutor.Response{}, statusErr{code: httpResp.StatusCode, msg: string(bodyBytes)}
	}

	var param any
	converted := sdktranslator.TranslateNonStream(respCtx, to, from, req.Model, bytes.Clone(opts.OriginalRequest), payload, bodyBytes, &param)
	return cliproxyexecutor.Response{Payload: []byte(converted)}, nil
}

// ensureAccessToken ensures we have a valid access token, refreshing if necessary.
func (e *KiroExecutor) ensureAccessToken(ctx context.Context, auth *cliproxyauth.Auth) (string, *cliproxyauth.Auth, error) {
	if auth == nil || auth.Metadata == nil {
		return "", nil, fmt.Errorf("kiro: missing authentication metadata")
	}

	accessToken, _ := auth.Metadata["access_token"].(string)
	expiresAt, _ := auth.Metadata["expired"].(string)

	if accessToken != "" && expiresAt != "" {
		expTime, errParse := time.Parse(time.RFC3339, expiresAt)
		if errParse == nil && time.Now().Add(kiroRefreshSkew).Before(expTime) {
			return accessToken, nil, nil
		}
	}

	// Token expired or missing, refresh it
	updated, errRefresh := e.refreshToken(ctx, auth.Clone())
	if errRefresh != nil {
		return "", nil, errRefresh
	}

	newAccessToken, _ := updated.Metadata["access_token"].(string)
	return newAccessToken, updated, nil
}

// refreshToken refreshes the access token using the refresh token.
func (e *KiroExecutor) refreshToken(ctx context.Context, auth *cliproxyauth.Auth) (*cliproxyauth.Auth, error) {
	if auth == nil || auth.Metadata == nil {
		return nil, fmt.Errorf("kiro: missing authentication metadata")
	}

	refreshToken, ok := auth.Metadata["refresh_token"].(string)
	if !ok || strings.TrimSpace(refreshToken) == "" {
		return nil, fmt.Errorf("kiro: missing refresh token")
	}

	reqBody := map[string]string{
		"refreshToken": refreshToken,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	httpClient := util.SetProxy(&e.cfg.SDKConfig, &http.Client{})
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
			log.Errorf("kiro token refresh: close body error: %v", errClose)
		}
	}()

	bodyBytes, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return nil, fmt.Errorf("read response: %w", errRead)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("token refresh failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var tokenResp struct {
		AccessToken string `json:"accessToken"`
		ExpiresIn   int64  `json:"expiresIn"`
		ProfileArn  string `json:"profileArn,omitempty"`
		Region      string `json:"region,omitempty"`
	}
	if errDecode := json.Unmarshal(bodyBytes, &tokenResp); errDecode != nil {
		return nil, fmt.Errorf("decode response: %w", errDecode)
	}

	now := time.Now()
	expiresAt := now.Add(time.Duration(tokenResp.ExpiresIn) * time.Second)

	auth.Metadata["access_token"] = tokenResp.AccessToken
	auth.Metadata["expires_in"] = tokenResp.ExpiresIn
	auth.Metadata["timestamp"] = now.UnixMilli()
	auth.Metadata["expired"] = expiresAt.Format(time.RFC3339)

	if tokenResp.ProfileArn != "" {
		auth.Metadata["profile_arn"] = tokenResp.ProfileArn
	}
	if tokenResp.Region != "" {
		auth.Metadata["region"] = tokenResp.Region
	}

	return auth, nil
}

// buildRequest builds an HTTP request for Kiro API.
func (e *KiroExecutor) buildRequest(ctx context.Context, auth *cliproxyauth.Auth, token, modelName string, payload []byte, stream bool, alt string) (*http.Request, error) {
	projectID, _ := auth.Metadata["project_id"].(string)
	if projectID == "" {
		return nil, fmt.Errorf("kiro: missing project_id")
	}

	path := kiroGeneratePath
	if stream {
		path = kiroStreamPath
	}

	endpoint := fmt.Sprintf("%s%s", kiroBaseURL, path)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", kiroAPIUserAgent)
	req.Header.Set("X-Goog-Api-Client", kiroAPIClient)
	req.Header.Set("Client-Metadata", kiroClientMetadata)
	req.Header.Set("X-Gcp-Project", projectID)

	return req, nil
}

// buildCountTokensRequest builds a count tokens request.
func (e *KiroExecutor) buildCountTokensRequest(ctx context.Context, auth *cliproxyauth.Auth, token, modelName string, payload []byte) (*http.Request, error) {
	projectID, _ := auth.Metadata["project_id"].(string)
	if projectID == "" {
		return nil, fmt.Errorf("kiro: missing project_id")
	}

	endpoint := fmt.Sprintf("%s%s", kiroBaseURL, kiroCountTokensPath)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", kiroAPIUserAgent)
	req.Header.Set("X-Goog-Api-Client", kiroAPIClient)
	req.Header.Set("Client-Metadata", kiroClientMetadata)
	req.Header.Set("X-Gcp-Project", projectID)

	return req, nil
}
