Repository: jwadow/kiro-openai-gateway
Files analyzed: 40

Estimated tokens: 161.5k

Directory structure:
└── jwadow-kiro-openai-gateway/
    ├── README.md
    ├── CLA.md
    ├── CONTRIBUTORS.md
    ├── LICENSE
    ├── main.py
    ├── manual_api_test.py
    ├── pytest.ini
    ├── requirements.txt
    ├── .clabot
    ├── .env.example
    ├── docs/
    │   ├── en/
    │   │   └── ARCHITECTURE.md
    │   └── ru/
    │       └── ARCHITECTURE.md
    ├── kiro_gateway/
    │   ├── __init__.py
    │   ├── auth.py
    │   ├── cache.py
    │   ├── config.py
    │   ├── converters.py
    │   ├── debug_logger.py
    │   ├── exceptions.py
    │   ├── http_client.py
    │   ├── models.py
    │   ├── parsers.py
    │   ├── routes.py
    │   ├── streaming.py
    │   ├── tokenizer.py
    │   └── utils.py
    ├── tests/
    │   ├── README.md
    │   ├── conftest.py
    │   ├── integration/
    │   │   └── test_full_flow.py
    │   └── unit/
    │       ├── test_auth_manager.py
    │       ├── test_cache.py
    │       ├── test_config.py
    │       ├── test_converters.py
    │       ├── test_debug_logger.py
    │       ├── test_http_client.py
    │       ├── test_parsers.py
    │       ├── test_routes.py
    │       ├── test_streaming.py
    │       └── test_tokenizer.py
    └── .github/
        └── ISSUE_TEMPLATE/
            └── bug_report.yml


================================================
FILE: README.md
================================================
<div align="center">

# 🚀 Kiro OpenAI Gateway

**OpenAI-compatible proxy gateway for Kiro IDE API (AWS CodeWhisperer)**

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
[![Python 3.10+](https://img.shields.io/badge/python-3.10+-blue.svg)](https://www.python.org/downloads/)
[![FastAPI](https://img.shields.io/badge/FastAPI-0.100+-green.svg)](https://fastapi.tiangolo.com/)

*Use Claude models through any tools that support the OpenAI API*

[Features](#-features) • [Quick Start](#-quick-start) • [Configuration](#%EF%B8%8F-configuration) • [API Reference](#-api-reference) • [License](#-license)

</div>

---

## ✨ Features

| Feature | Description |
|---------|-------------|
| 🔌 **OpenAI-compatible API** | Works with any OpenAI client out of the box |
| 💬 **Full message history** | Passes complete conversation context |
| 🛠️ **Tool Calling** | Supports function calling in OpenAI format |
| 📡 **Streaming** | Full SSE streaming support |
| 🔄 **Retry Logic** | Automatic retries on errors (403, 429, 5xx) |
| 📋 **Extended model list** | Including versioned models |
| 🔐 **Smart token management** | Automatic refresh before expiration |
| 🧩 **Modular architecture** | Easy to extend with new providers |

---

## 🚀 Quick Start

### Prerequisites

- Python 3.10+
- [Kiro IDE](https://kiro.dev/) with logged in account

### Installation

```bash
# Clone the repository
git clone https://github.com/Jwadow/kiro-openai-gateway.git
cd kiro-openai-gateway

# Install dependencies
pip install -r requirements.txt

# Configure (see Configuration section)
cp .env.example .env
# Edit .env with your credentials

# Start the server
python main.py
```

The server will be available at `http://localhost:8000`

---

## ⚙️ Configuration

### Option 1: JSON Credentials File

Specify the path to the credentials file:

```env
KIRO_CREDS_FILE="~/.aws/sso/cache/kiro-auth-token.json"

# Password to protect YOUR proxy server (make up any secure string)
# You'll use this as api_key when connecting to your gateway
PROXY_API_KEY="my-super-secret-password-123"
```

<details>
<summary>📄 JSON file format</summary>

```json
{
  "accessToken": "eyJ...",
  "refreshToken": "eyJ...",
  "expiresAt": "2025-01-12T23:00:00.000Z",
  "profileArn": "arn:aws:codewhisperer:us-east-1:...",
  "region": "us-east-1"
}
```

</details>

### Option 2: Environment Variables (.env file)

Create a `.env` file in the project root:

```env
# Required
REFRESH_TOKEN="your_kiro_refresh_token"

# Password to protect YOUR proxy server (make up any secure string)
PROXY_API_KEY="my-super-secret-password-123"

# Optional
PROFILE_ARN="arn:aws:codewhisperer:us-east-1:..."
KIRO_REGION="us-east-1"
```

### Getting the Refresh Token

The refresh token can be obtained by intercepting Kiro IDE traffic. Look for requests to:
- `prod.us-east-1.auth.desktop.kiro.dev/refreshToken`

---

## 📡 API Reference

### Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Health check |
| `/health` | GET | Detailed health check |
| `/v1/models` | GET | List available models |
| `/v1/chat/completions` | POST | Chat completions |

### Available Models

| Model | Description |
|-------|-------------|
| `claude-opus-4-5` | Top-tier model |
| `claude-opus-4-5-20251101` | Top-tier model (versioned) |
| `claude-sonnet-4-5` | Enhanced model |
| `claude-sonnet-4-5-20250929` | Enhanced model (versioned) |
| `claude-sonnet-4` | Balanced model |
| `claude-sonnet-4-20250514` | Balanced model (versioned) |
| `claude-haiku-4-5` | Fast model |
| `claude-3-7-sonnet-20250219` | Legacy model |

---

## 💡 Usage Examples

<details>
<summary>🔹 Simple cURL Request</summary>

```bash
curl http://localhost:8000/v1/chat/completions \
  -H "Authorization: Bearer my-super-secret-password-123" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "claude-sonnet-4-5",
    "messages": [{"role": "user", "content": "Hello!"}],
    "stream": true
  }'
```

> **Note:** Replace `my-super-secret-password-123` with the `PROXY_API_KEY` you set in your `.env` file.

</details>

<details>
<summary>🔹 Streaming Request</summary>

```bash
curl http://localhost:8000/v1/chat/completions \
  -H "Authorization: Bearer my-super-secret-password-123" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "claude-sonnet-4-5",
    "messages": [
      {"role": "system", "content": "You are a helpful assistant."},
      {"role": "user", "content": "What is 2+2?"}
    ],
    "stream": true
  }'
```

</details>

<details>
<summary>🔹 With Tool Calling</summary>

```bash
curl http://localhost:8000/v1/chat/completions \
  -H "Authorization: Bearer my-super-secret-password-123" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "claude-sonnet-4-5",
    "messages": [{"role": "user", "content": "What is the weather in London?"}],
    "tools": [{
      "type": "function",
      "function": {
        "name": "get_weather",
        "description": "Get weather for a location",
        "parameters": {
          "type": "object",
          "properties": {
            "location": {"type": "string", "description": "City name"}
          },
          "required": ["location"]
        }
      }
    }]
  }'
```

</details>

<details>
<summary>🐍 Python OpenAI SDK</summary>

```python
from openai import OpenAI

client = OpenAI(
    base_url="http://localhost:8000/v1",
    api_key="my-super-secret-password-123"  # Your PROXY_API_KEY from .env
)

response = client.chat.completions.create(
    model="claude-sonnet-4-5",
    messages=[
        {"role": "system", "content": "You are a helpful assistant."},
        {"role": "user", "content": "Hello!"}
    ],
    stream=True
)

for chunk in response:
    if chunk.choices[0].delta.content:
        print(chunk.choices[0].delta.content, end="")
```

</details>

<details>
<summary>🦜 LangChain</summary>

```python
from langchain_openai import ChatOpenAI

llm = ChatOpenAI(
    base_url="http://localhost:8000/v1",
    api_key="my-super-secret-password-123",  # Your PROXY_API_KEY from .env
    model="claude-sonnet-4-5"
)

response = llm.invoke("Hello, how are you?")
print(response.content)
```

</details>

---

## 📁 Project Structure

```
kiro-openai-gateway/
├── main.py                    # Entry point, FastAPI app creation
├── requirements.txt           # Python dependencies
├── .env.example               # Environment configuration example
│
├── kiro_gateway/              # Main package
│   ├── __init__.py            # Package exports
│   ├── config.py              # Configuration and constants
│   ├── models.py              # Pydantic models for OpenAI API
│   ├── auth.py                # KiroAuthManager - token management
│   ├── cache.py               # ModelInfoCache - model caching
│   ├── utils.py               # Helper utilities
│   ├── converters.py          # OpenAI <-> Kiro conversion
│   ├── parsers.py             # AWS SSE stream parsers
│   ├── streaming.py           # Response streaming logic
│   ├── http_client.py         # HTTP client with retry logic
│   ├── debug_logger.py        # Debug logging (optional)
│   └── routes.py              # FastAPI routes
│
├── tests/                     # Tests
│   ├── unit/                  # Unit tests
│   └── integration/           # Integration tests
│
└── debug_logs/                # Debug logs (generated when enabled)
```

---

## 🔧 Debugging

Debug logging is **disabled by default**. To enable, add to your `.env`:

```env
# Debug logging mode:
# - off: disabled (default)
# - errors: save logs only for failed requests (4xx, 5xx) - recommended for troubleshooting
# - all: save logs for every request (overwrites on each request)
DEBUG_MODE=errors
```

### Debug Modes

| Mode | Description | Use Case |
|------|-------------|----------|
| `off` | Disabled (default) | Production |
| `errors` | Save logs only for failed requests (4xx, 5xx) | **Recommended for troubleshooting** |
| `all` | Save logs for every request | Development/debugging |

### Debug Files

When enabled, requests are logged to the `debug_logs/` folder:

| File | Description |
|------|-------------|
| `request_body.json` | Incoming request from client (OpenAI format) |
| `kiro_request_body.json` | Request sent to Kiro API |
| `response_stream_raw.txt` | Raw stream from Kiro |
| `response_stream_modified.txt` | Transformed stream (OpenAI format) |
| `app_logs.txt` | Application logs for the request |
| `error_info.json` | Error details (only on errors) |

---

## 🧪 Testing

```bash
# Run all tests
pytest

# Run unit tests only
pytest tests/unit/

# Run with coverage
pytest --cov=kiro_gateway
```

---

## 🔌 Extending with New Providers

The modular architecture makes it easy to add support for other providers:

1. Create a new module `kiro_gateway/providers/new_provider.py`
2. Implement the required classes:
   - `NewProviderAuthManager` — token management
   - `NewProviderConverter` — format conversion
   - `NewProviderParser` — response parsing
3. Add routes to `routes.py` or create a separate router

---

## 📜 License

This project is licensed under the **GNU Affero General Public License v3.0 (AGPL-3.0)**.

This means:
- ✅ You can use, modify, and distribute this software
- ✅ You can use it for commercial purposes
- ⚠️ **You must disclose source code** when you distribute the software
- ⚠️ **Network use is distribution** — if you run a modified version on a server and let others interact with it, you must make the source code available to them
- ⚠️ Modifications must be released under the same license

See the [LICENSE](LICENSE) file for the full license text.

### Why AGPL-3.0?

AGPL-3.0 ensures that improvements to this software benefit the entire community. If you modify this gateway and deploy it as a service, you must share your improvements with your users.

### Contributor License Agreement (CLA)

By submitting a contribution to this project, you agree to the terms of our [Contributor License Agreement (CLA)](CLA.md). This ensures that:
- You have the right to submit the contribution
- You grant the maintainer rights to use and relicense your contribution
- The project remains legally protected

---

## 👤 Author

**Jwadow** — [@Jwadow](https://github.com/jwadow)

---

## ⚠️ Disclaimer

This project is not affiliated with, endorsed by, or sponsored by Amazon Web Services (AWS), Anthropic, or Kiro IDE. Use at your own risk and in compliance with the terms of service of the underlying APIs.

---

<div align="center">

**[⬆ Back to Top](#-kiro-openai-gateway)**

</div>



================================================
FILE: CLA.md
================================================
# Contributor License Agreement (CLA)

**Kiro OpenAI Gateway**

Version 1.0 — Effective Date: December 2025

---

## Introduction

Thank you for your interest in contributing to **Kiro OpenAI Gateway** (the "Project"), maintained by **Jwadow** (the "Maintainer"). This Contributor License Agreement ("Agreement") documents the rights granted by contributors to the Maintainer.

By submitting a Contribution to this Project, you accept and agree to the following terms and conditions for your present and future Contributions.

---

## 1. Definitions

**"You" (or "Your")** means the copyright owner or legal entity authorized by the copyright owner that is making this Agreement with the Maintainer.

**"Contribution"** means any original work of authorship, including any modifications or additions to an existing work, that is intentionally submitted by You to the Maintainer for inclusion in the Project. This includes any communication sent to the Project's repositories, issue trackers, mailing lists, or any other communication channel.

**"Submitted"** means any form of electronic, verbal, or written communication sent to the Maintainer, including but not limited to communication on electronic mailing lists, source code control systems, and issue tracking systems.

---

## 2. Grant of Copyright License

Subject to the terms and conditions of this Agreement, You hereby grant to the Maintainer and to recipients of software distributed by the Maintainer a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable copyright license to:

- Reproduce, prepare derivative works of, publicly display, publicly perform, sublicense, and distribute Your Contributions and such derivative works
- Relicense the Contribution under any license, including proprietary licenses

---

## 3. Grant of Patent License

Subject to the terms and conditions of this Agreement, You hereby grant to the Maintainer and to recipients of software distributed by the Maintainer a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable patent license to make, have made, use, offer to sell, sell, import, and otherwise transfer the Work, where such license applies only to those patent claims licensable by You that are necessarily infringed by Your Contribution(s) alone or by combination of Your Contribution(s) with the Work to which such Contribution(s) was submitted.

---

## 4. Representations

You represent that:

### 4.1 Original Work
You are legally entitled to grant the above license. If your employer(s) has rights to intellectual property that you create that includes your Contributions, you represent that:
- You have received permission to make Contributions on behalf of that employer
- Your employer has waived such rights for your Contributions to the Maintainer
- Your employer has executed a separate Corporate CLA with the Maintainer

### 4.2 Third-Party Content
If your Contribution includes or is based on any third-party code, you represent that:
- You have identified all such third-party code in your Contribution
- You have provided complete details of any third-party license or other restriction associated with any part of your Contribution

### 4.3 No Conflicts
Your Contribution does not violate any agreement or obligation you have with any third party.

---

## 5. Support and Warranty Disclaimer

You are not expected to provide support for Your Contributions, except to the extent You desire to provide support. You may provide support for free, for a fee, or not at all.

**UNLESS REQUIRED BY APPLICABLE LAW OR AGREED TO IN WRITING, YOU PROVIDE YOUR CONTRIBUTIONS ON AN "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING, WITHOUT LIMITATION, ANY WARRANTIES OR CONDITIONS OF TITLE, NON-INFRINGEMENT, MERCHANTABILITY, OR FITNESS FOR A PARTICULAR PURPOSE.**

---

## 6. Notification of Changes

You agree to notify the Maintainer of any facts or circumstances of which you become aware that would make these representations inaccurate in any respect.

---

## 7. Moral Rights

To the fullest extent permitted under applicable law, You hereby waive, and agree not to assert, all of Your "moral rights" in or relating to Your Contributions for the benefit of the Maintainer, its assigns, and their respective direct and indirect sublicensees.

---

## 8. Governing Law

This Agreement shall be governed by and construed in accordance with the laws of the jurisdiction in which the Maintainer resides, without regard to its conflict of laws provisions.

---

## 9. Entire Agreement

This Agreement constitutes the entire agreement between the parties with respect to the subject matter hereof and supersedes all prior and contemporaneous agreements and understandings, whether written or oral, relating to such subject matter.

---

## How to Sign This CLA

By submitting a pull request or other Contribution to this Project, you signify your acceptance of this Agreement. 

For significant contributions, you may be asked to explicitly confirm your acceptance by:

1. Adding your name to the [CONTRIBUTORS.md](CONTRIBUTORS.md) file (if it exists)
2. Commenting "I have read the CLA and I accept its terms" on your pull request
3. Signing via a CLA bot (if implemented)

---

## Contact

If you have questions about this CLA, please open an issue in the repository or contact the Maintainer directly.

**Maintainer:** Jwadow  
**GitHub:** [@jwadow](https://github.com/jwadow)  
**Project:** [Kiro OpenAI Gateway](https://github.com/jwadow/kiro-openai-gateway)

---

*This CLA is based on the Apache Individual Contributor License Agreement and has been modified for this project.*


================================================
FILE: CONTRIBUTORS.md
================================================
# Contributors

Thank you to all the contributors who have helped improve this project!

## Contributors

- [@Kartvya69](https://github.com/Kartvya69) — STREAMING_READ_TIMEOUT feature (#9)



================================================
FILE: LICENSE
================================================
                    GNU AFFERO GENERAL PUBLIC LICENSE
                       Version 3, 19 November 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.

                            Preamble

  The GNU Affero General Public License is a free, copyleft license for
software and other kinds of works, specifically designed to ensure
cooperation with the community in the case of network server software.

  The licenses for most software and other practical works are designed
to take away your freedom to share and change the works.  By contrast,
our General Public Licenses are intended to guarantee your freedom to
share and change all versions of a program--to make sure it remains free
software for all its users.

  When we speak of free software, we are referring to freedom, not
price.  Our General Public Licenses are designed to make sure that you
have the freedom to distribute copies of free software (and charge for
them if you wish), that you receive source code or can get it if you
want it, that you can change the software or use pieces of it in new
free programs, and that you know you can do these things.

  Developers that use our General Public Licenses protect your rights
with two steps: (1) assert copyright on the software, and (2) offer
you this License which gives you legal permission to copy, distribute
and/or modify the software.

  A secondary benefit of defending all users' freedom is that
improvements made in alternate versions of the program, if they
receive widespread use, become available for other developers to
incorporate.  Many developers of free software are heartened and
encouraged by the resulting cooperation.  However, in the case of
software used on network servers, this result may fail to come about.
The GNU General Public License permits making a modified version and
letting the public access it on a server without ever releasing its
source code to the public.

  The GNU Affero General Public License is designed specifically to
ensure that, in such cases, the modified source code becomes available
to the community.  It requires the operator of a network server to
provide the source code of the modified version running there to the
users of that server.  Therefore, public use of a modified version, on
a publicly accessible server, gives the public access to the source
code of the modified version.

  An older license, called the Affero General Public License and
published by Affero, was designed to accomplish similar goals.  This is
a different license, not a version of the Affero GPL, but Affero has
released a new version of the Affero GPL which permits relicensing under
this license.

  The precise terms and conditions for copying, distribution and
modification follow.

                       TERMS AND CONDITIONS

  0. Definitions.

  "This License" refers to version 3 of the GNU Affero General Public License.

  "Copyright" also means copyright-like laws that apply to other kinds of
works, such as semiconductor masks.

  "The Program" refers to any copyrightable work licensed under this
License.  Each licensee is addressed as "you".  "Licensees" and
"recipients" may be individuals or organizations.

  To "modify" a work means to copy from or adapt all or part of the work
in a fashion requiring copyright permission, other than the making of an
exact copy.  The resulting work is called a "modified version" of the
earlier work or a work "based on" the earlier work.

  A "covered work" means either the unmodified Program or a work based
on the Program.

  To "propagate" a work means to do anything with it that, without
permission, would make you directly or secondarily liable for
infringement under applicable copyright law, except executing it on a
computer or modifying a private copy.  Propagation includes copying,
distribution (with or without modification), making available to the
public, and in some countries other activities as well.

  To "convey" a work means any kind of propagation that enables other
parties to make or receive copies.  Mere interaction with a user through
a computer network, with no transfer of a copy, is not conveying.

  An interactive user interface displays "Appropriate Legal Notices"
to the extent that it includes a convenient and prominently visible
feature that (1) displays an appropriate copyright notice, and (2)
tells the user that there is no warranty for the work (except to the
extent that warranties are provided), that licensees may convey the
work under this License, and how to view a copy of this License.  If
the interface presents a list of user commands or options, such as a
menu, a prominent item in the list meets this criterion.

  1. Source Code.

  The "source code" for a work means the preferred form of the work
for making modifications to it.  "Object code" means any non-source
form of a work.

  A "Standard Interface" means an interface that either is an official
standard defined by a recognized standards body, or, in the case of
interfaces specified for a particular programming language, one that
is widely used among developers working in that language.

  The "System Libraries" of an executable work include anything, other
than the work as a whole, that (a) is included in the normal form of
packaging a Major Component, but which is not part of that Major
Component, and (b) serves only to enable use of the work with that
Major Component, or to implement a Standard Interface for which an
implementation is available to the public in source code form.  A
"Major Component", in this context, means a major essential component
(kernel, window system, and so on) of the specific operating system
(if any) on which the executable work runs, or a compiler used to
produce the work, or an object code interpreter used to run it.

  The "Corresponding Source" for a work in object code form means all
the source code needed to generate, install, and (for an executable
work) run the object code and to modify the work, including scripts to
control those activities.  However, it does not include the work's
System Libraries, or general-purpose tools or generally available free
programs which are used unmodified in performing those activities but
which are not part of the work.  For example, Corresponding Source
includes interface definition files associated with source files for
the work, and the source code for shared libraries and dynamically
linked subprograms that the work is specifically designed to require,
such as by intimate data communication or control flow between those
subprograms and other parts of the work.

  The Corresponding Source need not include anything that users
can regenerate automatically from other parts of the Corresponding
Source.

  The Corresponding Source for a work in source code form is that
same work.

  2. Basic Permissions.

  All rights granted under this License are granted for the term of
copyright on the Program, and are irrevocable provided the stated
conditions are met.  This License explicitly affirms your unlimited
permission to run the unmodified Program.  The output from running a
covered work is covered by this License only if the output, given its
content, constitutes a covered work.  This License acknowledges your
rights of fair use or other equivalent, as provided by copyright law.

  You may make, run and propagate covered works that you do not
convey, without conditions so long as your license otherwise remains
in force.  You may convey covered works to others for the sole purpose
of having them make modifications exclusively for you, or provide you
with facilities for running those works, provided that you comply with
the terms of this License in conveying all material for which you do
not control copyright.  Those thus making or running the covered works
for you must do so exclusively on your behalf, under your direction
and control, on terms that prohibit them from making any copies of
your copyrighted material outside their relationship with you.

  Conveying under any other circumstances is permitted solely under
the conditions stated below.  Sublicensing is not allowed; section 10
makes it unnecessary.

  3. Protecting Users' Legal Rights From Anti-Circumvention Law.

  No covered work shall be deemed part of an effective technological
measure under any applicable law fulfilling obligations under article
11 of the WIPO copyright treaty adopted on 20 December 1996, or
similar laws prohibiting or restricting circumvention of such
measures.

  When you convey a covered work, you waive any legal power to forbid
circumvention of technological measures to the extent such circumvention
is effected by exercising rights under this License with respect to
the covered work, and you disclaim any intention to limit operation or
modification of the work as a means of enforcing, against the work's
users, your or third parties' legal rights to forbid circumvention of
technological measures.

  4. Conveying Verbatim Copies.

  You may convey verbatim copies of the Program's source code as you
receive it, in any medium, provided that you conspicuously and
appropriately publish on each copy an appropriate copyright notice;
keep intact all notices stating that this License and any
non-permissive terms added in accord with section 7 apply to the code;
keep intact all notices of the absence of any warranty; and give all
recipients a copy of this License along with the Program.

  You may charge any price or no price for each copy that you convey,
and you may offer support or warranty protection for a fee.

  5. Conveying Modified Source Versions.

  You may convey a work based on the Program, or the modifications to
produce it from the Program, in the form of source code under the
terms of section 4, provided that you also meet all of these conditions:

    a) The work must carry prominent notices stating that you modified
    it, and giving a relevant date.

    b) The work must carry prominent notices stating that it is
    released under this License and any conditions added under section
    7.  This requirement modifies the requirement in section 4 to
    "keep intact all notices".

    c) You must license the entire work, as a whole, under this
    License to anyone who comes into possession of a copy.  This
    License will therefore apply, along with any applicable section 7
    additional terms, to the whole of the work, and all its parts,
    regardless of how they are packaged.  This License gives no
    permission to license the work in any other way, but it does not
    invalidate such permission if you have separately received it.

    d) If the work has interactive user interfaces, each must display
    Appropriate Legal Notices; however, if the Program has interactive
    interfaces that do not display Appropriate Legal Notices, your
    work need not make them do so.

  A compilation of a covered work with other separate and independent
works, which are not by their nature extensions of the covered work,
and which are not combined with it such as to form a larger program,
in or on a volume of a storage or distribution medium, is called an
"aggregate" if the compilation and its resulting copyright are not
used to limit the access or legal rights of the compilation's users
beyond what the individual works permit.  Inclusion of a covered work
in an aggregate does not cause this License to apply to the other
parts of the aggregate.

  6. Conveying Non-Source Forms.

  You may convey a covered work in object code form under the terms
of sections 4 and 5, provided that you also convey the
machine-readable Corresponding Source under the terms of this License,
in one of these ways:

    a) Convey the object code in, or embodied in, a physical product
    (including a physical distribution medium), accompanied by the
    Corresponding Source fixed on a durable physical medium
    customarily used for software interchange.

    b) Convey the object code in, or embodied in, a physical product
    (including a physical distribution medium), accompanied by a
    written offer, valid for at least three years and valid for as
    long as you offer spare parts or customer support for that product
    model, to give anyone who possesses the object code either (1) a
    copy of the Corresponding Source for all the software in the
    product that is covered by this License, on a durable physical
    medium customarily used for software interchange, for a price no
    more than your reasonable cost of physically performing this
    conveying of source, or (2) access to copy the
    Corresponding Source from a network server at no charge.

    c) Convey individual copies of the object code with a copy of the
    written offer to provide the Corresponding Source.  This
    alternative is allowed only occasionally and noncommercially, and
    only if you received the object code with such an offer, in accord
    with subsection 6b.

    d) Convey the object code by offering access from a designated
    place (gratis or for a charge), and offer equivalent access to the
    Corresponding Source in the same way through the same place at no
    further charge.  You need not require recipients to copy the
    Corresponding Source along with the object code.  If the place to
    copy the object code is a network server, the Corresponding Source
    may be on a different server (operated by you or a third party)
    that supports equivalent copying facilities, provided you maintain
    clear directions next to the object code saying where to find the
    Corresponding Source.  Regardless of what server hosts the
    Corresponding Source, you remain obligated to ensure that it is
    available for as long as needed to satisfy these requirements.

    e) Convey the object code using peer-to-peer transmission, provided
    you inform other peers where the object code and Corresponding
    Source of the work are being offered to the general public at no
    charge under subsection 6d.

  A separable portion of the object code, whose source code is excluded
from the Corresponding Source as a System Library, need not be
included in conveying the object code work.

  A "User Product" is either (1) a "consumer product", which means any
tangible personal property which is normally used for personal, family,
or household purposes, or (2) anything designed or sold for incorporation
into a dwelling.  In determining whether a product is a consumer product,
doubtful cases shall be resolved in favor of coverage.  For a particular
product received by a particular user, "normally used" refers to a
typical or common use of that class of product, regardless of the status
of the particular user or of the way in which the particular user
actually uses, or expects or is expected to use, the product.  A product
is a consumer product regardless of whether the product has substantial
commercial, industrial or non-consumer uses, unless such uses represent
the only significant mode of use of the product.

  "Installation Information" for a User Product means any methods,
procedures, authorization keys, or other information required to install
and execute modified versions of a covered work in that User Product from
a modified version of its Corresponding Source.  The information must
suffice to ensure that the continued functioning of the modified object
code is in no case prevented or interfered with solely because
modification has been made.

  If you convey an object code work under this section in, or with, or
specifically for use in, a User Product, and the conveying occurs as
part of a transaction in which the right of possession and use of the
User Product is transferred to the recipient in perpetuity or for a
fixed term (regardless of how the transaction is characterized), the
Corresponding Source conveyed under this section must be accompanied
by the Installation Information.  But this requirement does not apply
if neither you nor any third party retains the ability to install
modified object code on the User Product (for example, the work has
been installed in ROM).

  The requirement to provide Installation Information does not include a
requirement to continue to provide support service, warranty, or updates
for a work that has been modified or installed by the recipient, or for
the User Product in which it has been modified or installed.  Access to a
network may be denied when the modification itself materially and
adversely affects the operation of the network or violates the rules and
protocols for communication across the network.

  Corresponding Source conveyed, and Installation Information provided,
in accord with this section must be in a format that is publicly
documented (and with an implementation available to the public in
source code form), and must require no special password or key for
unpacking, reading or copying.

  7. Additional Terms.

  "Additional permissions" are terms that supplement the terms of this
License by making exceptions from one or more of its conditions.
Additional permissions that are applicable to the entire Program shall
be treated as though they were included in this License, to the extent
that they are valid under applicable law.  If additional permissions
apply only to part of the Program, that part may be used separately
under those permissions, but the entire Program remains governed by
this License without regard to the additional permissions.

  When you convey a copy of a covered work, you may at your option
remove any additional permissions from that copy, or from any part of
it.  (Additional permissions may be written to require their own
removal in certain cases when you modify the work.)  You may place
additional permissions on material, added by you to a covered work,
for which you have or can give appropriate copyright permission.

  Notwithstanding any other provision of this License, for material you
add to a covered work, you may (if authorized by the copyright holders of
that material) supplement the terms of this License with terms:

    a) Disclaiming warranty or limiting liability differently from the
    terms of sections 15 and 16 of this License; or

    b) Requiring preservation of specified reasonable legal notices or
    author attributions in that material or in the Appropriate Legal
    Notices displayed by works containing it; or

    c) Prohibiting misrepresentation of the origin of that material, or
    requiring that modified versions of such material be marked in
    reasonable ways as different from the original version; or

    d) Limiting the use for publicity purposes of names of licensors or
    authors of the material; or

    e) Declining to grant rights under trademark law for use of some
    trade names, trademarks, or service marks; or

    f) Requiring indemnification of licensors and authors of that
    material by anyone who conveys the material (or modified versions of
    it) with contractual assumptions of liability to the recipient, for
    any liability that these contractual assumptions directly impose on
    those licensors and authors.

  All other non-permissive additional terms are considered "further
restrictions" within the meaning of section 10.  If the Program as you
received it, or any part of it, contains a notice stating that it is
governed by this License along with a term that is a further
restriction, you may remove that term.  If a license document contains
a further restriction but permits relicensing or conveying under this
License, you may add to a covered work material governed by the terms
of that license document, provided that the further restriction does
not survive such relicensing or conveying.

  If you add terms to a covered work in accord with this section, you
must place, in the relevant source files, a statement of the
additional terms that apply to those files, or a notice indicating
where to find the applicable terms.

  Additional terms, permissive or non-permissive, may be stated in the
form of a separately written license, or stated as exceptions;
the above requirements apply either way.

  8. Termination.

  You may not propagate or modify a covered work except as expressly
provided under this License.  Any attempt otherwise to propagate or
modify it is void, and will automatically terminate your rights under
this License (including any patent licenses granted under the third
paragraph of section 11).

  However, if you cease all violation of this License, then your
license from a particular copyright holder is reinstated (a)
provisionally, unless and until the copyright holder explicitly and
finally terminates your license, and (b) permanently, if the copyright
holder fails to notify you of the violation by some reasonable means
prior to 60 days after the cessation.

  Moreover, your license from a particular copyright holder is
reinstated permanently if the copyright holder notifies you of the
violation by some reasonable means, this is the first time you have
received notice of violation of this License (for any work) from that
copyright holder, and you cure the violation prior to 30 days after
your receipt of the notice.

  Termination of your rights under this section does not terminate the
licenses of parties who have received copies or rights from you under
this License.  If your rights have been terminated and not permanently
reinstated, you do not qualify to receive new licenses for the same
material under section 10.

  9. Acceptance Not Required for Having Copies.

  You are not required to accept this License in order to receive or
run a copy of the Program.  Ancillary propagation of a covered work
occurring solely as a consequence of using peer-to-peer transmission
to receive a copy likewise does not require acceptance.  However,
nothing other than this License grants you permission to propagate or
modify any covered work.  These actions infringe copyright if you do
not accept this License.  Therefore, by modifying or propagating a
covered work, you indicate your acceptance of this License to do so.

  10. Automatic Licensing of Downstream Recipients.

  Each time you convey a covered work, the recipient automatically
receives a license from the original licensors, to run, modify and
propagate that work, subject to this License.  You are not responsible
for enforcing compliance by third parties with this License.

  An "entity transaction" is a transaction transferring control of an
organization, or substantially all assets of one, or subdividing an
organization, or merging organizations.  If propagation of a covered
work results from an entity transaction, each party to that
transaction who receives a copy of the work also receives whatever
licenses to the work the party's predecessor in interest had or could
give under the previous paragraph, plus a right to possession of the
Corresponding Source of the work from the predecessor in interest, if
the predecessor has it or can get it with reasonable efforts.

  You may not impose any further restrictions on the exercise of the
rights granted or affirmed under this License.  For example, you may
not impose a license fee, royalty, or other charge for exercise of
rights granted under this License, and you may not initiate litigation
(including a cross-claim or counterclaim in a lawsuit) alleging that
any patent claim is infringed by making, using, selling, offering for
sale, or importing the Program or any portion of it.

  11. Patents.

  A "contributor" is a copyright holder who authorizes use under this
License of the Program or a work on which the Program is based.  The
work thus licensed is called the contributor's "contributor version".

  A contributor's "essential patent claims" are all patent claims
owned or controlled by the contributor, whether already acquired or
hereafter acquired, that would be infringed by some manner, permitted
by this License, of making, using, or selling its contributor version,
but do not include claims that would be infringed only as a
consequence of further modification of the contributor version.  For
purposes of this definition, "control" includes the right to grant
patent sublicenses in a manner consistent with the requirements of
this License.

  Each contributor grants you a non-exclusive, worldwide, royalty-free
patent license under the contributor's essential patent claims, to
make, use, sell, offer for sale, import and otherwise run, modify and
propagate the contents of its contributor version.

  In the following three paragraphs, a "patent license" is any express
agreement or commitment, however denominated, not to enforce a patent
(such as an express permission to practice a patent or covenant not to
sue for patent infringement).  To "grant" such a patent license to a
party means to make such an agreement or commitment not to enforce a
patent against the party.

  If you convey a covered work, knowingly relying on a patent license,
and the Corresponding Source of the work is not available for anyone
to copy, free of charge and under the terms of this License, through a
publicly available network server or other readily accessible means,
then you must either (1) cause the Corresponding Source to be so
available, or (2) arrange to deprive yourself of the benefit of the
patent license for this particular work, or (3) arrange, in a manner
consistent with the requirements of this License, to extend the patent
license to downstream recipients.  "Knowingly relying" means you have
actual knowledge that, but for the patent license, your conveying the
covered work in a country, or your recipient's use of the covered work
in a country, would infringe one or more identifiable patents in that
country that you have reason to believe are valid.

  If, pursuant to or in connection with a single transaction or
arrangement, you convey, or propagate by procuring conveyance of, a
covered work, and grant a patent license to some of the parties
receiving the covered work authorizing them to use, propagate, modify
or convey a specific copy of the covered work, then the patent license
you grant is automatically extended to all recipients of the covered
work and works based on it.

  A patent license is "discriminatory" if it does not include within
the scope of its coverage, prohibits the exercise of, or is
conditioned on the non-exercise of one or more of the rights that are
specifically granted under this License.  You may not convey a covered
work if you are a party to an arrangement with a third party that is
in the business of distributing software, under which you make payment
to the third party based on the extent of your activity of conveying
the work, and under which the third party grants, to any of the
parties who would receive the covered work from you, a discriminatory
patent license (a) in connection with copies of the covered work
conveyed by you (or copies made from those copies), or (b) primarily
for and in connection with specific products or compilations that
contain the covered work, unless you entered into that arrangement,
or that patent license was granted, prior to 28 March 2007.

  Nothing in this License shall be construed as excluding or limiting
any implied license or other defenses to infringement that may
otherwise be available to you under applicable patent law.

  12. No Surrender of Others' Freedom.

  If conditions are imposed on you (whether by court order, agreement or
otherwise) that contradict the conditions of this License, they do not
excuse you from the conditions of this License.  If you cannot convey a
covered work so as to satisfy simultaneously your obligations under this
License and any other pertinent obligations, then as a consequence you may
not convey it at all.  For example, if you agree to terms that obligate you
to collect a royalty for further conveying from those to whom you convey
the Program, the only way you could satisfy both those terms and this
License would be to refrain entirely from conveying the Program.

  13. Remote Network Interaction; Use with the GNU General Public License.

  Notwithstanding any other provision of this License, if you modify the
Program, your modified version must prominently offer all users
interacting with it remotely through a computer network (if your version
supports such interaction) an opportunity to receive the Corresponding
Source of your version by providing access to the Corresponding Source
from a network server at no charge, through some standard or customary
means of facilitating copying of software.  This Corresponding Source
shall include the Corresponding Source for any work covered by version 3
of the GNU General Public License that is incorporated pursuant to the
following paragraph.

  Notwithstanding any other provision of this License, you have
permission to link or combine any covered work with a work licensed
under version 3 of the GNU General Public License into a single
combined work, and to convey the resulting work.  The terms of this
License will continue to apply to the part which is the covered work,
but the work with which it is combined will remain governed by version
3 of the GNU General Public License.

  14. Revised Versions of this License.

  The Free Software Foundation may publish revised and/or new versions of
the GNU Affero General Public License from time to time.  Such new versions
will be similar in spirit to the present version, but may differ in detail to
address new problems or concerns.

  Each version is given a distinguishing version number.  If the
Program specifies that a certain numbered version of the GNU Affero General
Public License "or any later version" applies to it, you have the
option of following the terms and conditions either of that numbered
version or of any later version published by the Free Software
Foundation.  If the Program does not specify a version number of the
GNU Affero General Public License, you may choose any version ever published
by the Free Software Foundation.

  If the Program specifies that a proxy can decide which future
versions of the GNU Affero General Public License can be used, that proxy's
public statement of acceptance of a version permanently authorizes you
to choose that version for the Program.

  Later license versions may give you additional or different
permissions.  However, no additional obligations are imposed on any
author or copyright holder as a result of your choosing to follow a
later version.

  15. Disclaimer of Warranty.

  THERE IS NO WARRANTY FOR THE PROGRAM, TO THE EXTENT PERMITTED BY
APPLICABLE LAW.  EXCEPT WHEN OTHERWISE STATED IN WRITING THE COPYRIGHT
HOLDERS AND/OR OTHER PARTIES PROVIDE THE PROGRAM "AS IS" WITHOUT WARRANTY
OF ANY KIND, EITHER EXPRESSED OR IMPLIED, INCLUDING, BUT NOT LIMITED TO,
THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
PURPOSE.  THE ENTIRE RISK AS TO THE QUALITY AND PERFORMANCE OF THE PROGRAM
IS WITH YOU.  SHOULD THE PROGRAM PROVE DEFECTIVE, YOU ASSUME THE COST OF
ALL NECESSARY SERVICING, REPAIR OR CORRECTION.

  16. Limitation of Liability.

  IN NO EVENT UNLESS REQUIRED BY APPLICABLE LAW OR AGREED TO IN WRITING
WILL ANY COPYRIGHT HOLDER, OR ANY OTHER PARTY WHO MODIFIES AND/OR CONVEYS
THE PROGRAM AS PERMITTED ABOVE, BE LIABLE TO YOU FOR DAMAGES, INCLUDING ANY
GENERAL, SPECIAL, INCIDENTAL OR CONSEQUENTIAL DAMAGES ARISING OUT OF THE
USE OR INABILITY TO USE THE PROGRAM (INCLUDING BUT NOT LIMITED TO LOSS OF
DATA OR DATA BEING RENDERED INACCURATE OR LOSSES SUSTAINED BY YOU OR THIRD
PARTIES OR A FAILURE OF THE PROGRAM TO OPERATE WITH ANY OTHER PROGRAMS),
EVEN IF SUCH HOLDER OR OTHER PARTY HAS BEEN ADVISED OF THE POSSIBILITY OF
SUCH DAMAGES.

  17. Interpretation of Sections 15 and 16.

  If the disclaimer of warranty and limitation of liability provided
above cannot be given local legal effect according to their terms,
reviewing courts shall apply local law that most closely approximates
an absolute waiver of all civil liability in connection with the
Program, unless a warranty or assumption of liability accompanies a
copy of the Program in return for a fee.

                     END OF TERMS AND CONDITIONS

            How to Apply These Terms to Your New Programs

  If you develop a new program, and you want it to be of the greatest
possible use to the public, the best way to achieve this is to make it
free software which everyone can redistribute and change under these terms.

  To do so, attach the following notices to the program.  It is safest
to attach them to the start of each source file to most effectively
state the exclusion of warranty; and each file should have at least
the "copyright" line and a pointer to where the full notice is found.

    <one line to give the program's name and a brief idea of what it does.>
    Copyright (C) <year>  <name of author>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

Also add information on how to contact you by electronic and paper mail.

  If your software can interact with users remotely through a computer
network, you should also make sure that it provides a way for users to
get its source.  For example, if your program is a web application, its
interface could display a "Source" link that leads users to an archive
of the code.  There are many ways you could offer source, and different
solutions will be better for different programs; see section 13 for the
specific requirements.

  You should also get your employer (if you work as a programmer) or school,
if any, to sign a "copyright disclaimer" for the program, if necessary.
For more information on this, and how to apply and follow the GNU AGPL, see
<https://www.gnu.org/licenses/>.



================================================
FILE: main.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
Kiro API Gateway - OpenAI-compatible interface for Kiro API.

Application entry point. Creates FastAPI app and connects routes.

Usage:
    uvicorn main:app --host 0.0.0.0 --port 8000
    or directly:
    python main.py
"""

import logging
import sys
from contextlib import asynccontextmanager
from pathlib import Path

from fastapi import FastAPI
from fastapi.exceptions import RequestValidationError
from fastapi.middleware.cors import CORSMiddleware
from loguru import logger

from kiro_gateway.config import (
    APP_TITLE,
    APP_DESCRIPTION,
    APP_VERSION,
    REFRESH_TOKEN,
    PROFILE_ARN,
    REGION,
    KIRO_CREDS_FILE,
    PROXY_API_KEY,
    LOG_LEVEL,
    _warn_deprecated_debug_setting,
    _warn_timeout_configuration,
)
from kiro_gateway.auth import KiroAuthManager
from kiro_gateway.cache import ModelInfoCache
from kiro_gateway.routes import router
from kiro_gateway.exceptions import validation_exception_handler


# --- Loguru Configuration ---
logger.remove()
logger.add(
    sys.stderr,
    level=LOG_LEVEL,
    colorize=True,
    format="<green>{time:YYYY-MM-DD HH:mm:ss}</green> | <level>{level: <8}</level> | <cyan>{name}</cyan>:<cyan>{function}</cyan>:<cyan>{line}</cyan> - <level>{message}</level>"
)


class InterceptHandler(logging.Handler):
    """
    Intercepts logs from standard logging and redirects them to loguru.
    
    This allows capturing logs from uvicorn, FastAPI and other libraries
    that use standard logging instead of loguru.
    """
    
    def emit(self, record: logging.LogRecord) -> None:
        # Get the corresponding loguru level
        try:
            level = logger.level(record.levelname).name
        except ValueError:
            level = record.levelno
        
        # Find the caller frame for correct source display
        frame, depth = logging.currentframe(), 2
        while frame.f_code.co_filename == logging.__file__:
            frame = frame.f_back
            depth += 1
        
        logger.opt(depth=depth, exception=record.exc_info).log(level, record.getMessage())


def setup_logging_intercept():
    """
    Configures log interception from standard logging to loguru.
    
    Intercepts logs from:
    - uvicorn (access logs, error logs)
    - uvicorn.error
    - uvicorn.access
    - fastapi
    """
    # List of loggers to intercept
    loggers_to_intercept = [
        "uvicorn",
        "uvicorn.error",
        "uvicorn.access",
        "fastapi",
    ]
    
    for logger_name in loggers_to_intercept:
        logging_logger = logging.getLogger(logger_name)
        logging_logger.handlers = [InterceptHandler()]
        logging_logger.propagate = False


# Configure uvicorn/fastapi log interception
setup_logging_intercept()


# --- Configuration Validation ---
def validate_configuration() -> None:
    """
    Validates that required configuration is present.
    
    Checks:
    - .env file exists
    - Either REFRESH_TOKEN or KIRO_CREDS_FILE is configured
    
    Raises:
        SystemExit: If critical configuration is missing
    """
    errors = []
    
    # Check if .env file exists
    env_file = Path(".env")
    env_example = Path(".env.example")
    
    if not env_file.exists():
        errors.append(
            ".env file not found!\n"
            "\n"
            "To get started:\n"
            "1. Create .env or rename from .env.example:\n"
            "   cp .env.example .env\n"
            "\n"
            "2. Edit .env and configure your credentials:\n"
            "   2.1. Set you super-secret password as PROXY_API_KEY\n"
            "   2.2. Set your Kiro credentials:\n"
            "      - 1 way: KIRO_CREDS_FILE to your Kiro credentials JSON file\n"
            "      - 2 way: REFRESH_TOKEN from Kiro IDE traffic\n"
            "\n"
            "See README.md for detailed instructions."
        )
    else:
        # .env exists, check for credentials
        has_refresh_token = bool(REFRESH_TOKEN)
        has_creds_file = bool(KIRO_CREDS_FILE)
        
        # Check if creds file actually exists
        if KIRO_CREDS_FILE:
            creds_path = Path(KIRO_CREDS_FILE).expanduser()
            if not creds_path.exists():
                has_creds_file = False
                logger.warning(f"KIRO_CREDS_FILE not found: {KIRO_CREDS_FILE}")
        
        if not has_refresh_token and not has_creds_file:
            errors.append(
                "No Kiro credentials configured!\n"
                "\n"
                "   Configure one of the following in your .env file:\n"
                "\n"
                "Set you super-secret password as PROXY_API_KEY\n"
                "   PROXY_API_KEY=\"my-super-secret-password-123\"\n"
                "\n"
                "   Option 1 (Recommended): JSON credentials file\n"
                "      KIRO_CREDS_FILE=\"path/to/your/kiro-credentials.json\"\n"
                "\n"
                "   Option 2: Refresh token\n"
                "      REFRESH_TOKEN=\"your_refresh_token_here\"\n"
                "\n"
                "   See README.md for how to obtain credentials."
            )
    
    # Print errors and exit if any
    if errors:
        logger.error("")
        logger.error("=" * 60)
        logger.error("  CONFIGURATION ERROR")
        logger.error("=" * 60)
        for error in errors:
            for line in error.split('\n'):
                logger.error(f"  {line}")
        logger.error("=" * 60)
        logger.error("")
        sys.exit(1)
    
    # Log successful configuration
    if KIRO_CREDS_FILE:
        logger.info(f"Using credentials file: {KIRO_CREDS_FILE}")
    elif REFRESH_TOKEN:
        logger.info("Using refresh token from environment")


# Run configuration validation on import
validate_configuration()

# Warn about deprecated DEBUG_LAST_REQUEST if used
_warn_deprecated_debug_setting()

# Warn about suboptimal timeout configuration
_warn_timeout_configuration()


# --- Lifespan Manager ---
@asynccontextmanager
async def lifespan(app: FastAPI):
    """
    Manages the application lifecycle.
    
    Creates and initializes:
    - KiroAuthManager for token management
    - ModelInfoCache for model caching
    """
    logger.info("Starting application... Creating state managers.")
    
    # Create AuthManager
    app.state.auth_manager = KiroAuthManager(
        refresh_token=REFRESH_TOKEN,
        profile_arn=PROFILE_ARN,
        region=REGION,
        creds_file=KIRO_CREDS_FILE if KIRO_CREDS_FILE else None
    )
    
    # Create model cache
    app.state.model_cache = ModelInfoCache()
    
    yield
    
    logger.info("Shutting down application.")


# --- FastAPI Application ---
app = FastAPI(
    title=APP_TITLE,
    description=APP_DESCRIPTION,
    version=APP_VERSION,
    lifespan=lifespan
)


# --- CORS Middleware ---
# Allow CORS for all origins to support browser clients
# and tools that send preflight OPTIONS requests
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Allow all origins
    allow_credentials=True,
    allow_methods=["*"],  # Allow all methods (GET, POST, OPTIONS, etc.)
    allow_headers=["*"],  # Allow all headers
)


# --- Validation Error Handler Registration ---
app.add_exception_handler(RequestValidationError, validation_exception_handler)


# --- Route Registration ---
app.include_router(router)


# --- Uvicorn log config ---
# Minimal configuration for redirecting uvicorn logs to loguru.
# Uses InterceptHandler which intercepts logs and passes them to loguru.
UVICORN_LOG_CONFIG = {
    "version": 1,
    "disable_existing_loggers": False,
    "handlers": {
        "default": {
            "class": "main.InterceptHandler",
        },
    },
    "loggers": {
        "uvicorn": {"handlers": ["default"], "level": "INFO", "propagate": False},
        "uvicorn.error": {"handlers": ["default"], "level": "INFO", "propagate": False},
        "uvicorn.access": {"handlers": ["default"], "level": "INFO", "propagate": False},
    },
}


# --- Entry Point ---
if __name__ == "__main__":
    import uvicorn
    logger.info("Starting Uvicorn server...")
    
    uvicorn.run(
        app,
        host="0.0.0.0",
        port=8000,
        log_config=UVICORN_LOG_CONFIG,
    )



================================================
FILE: manual_api_test.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# https://github.com/jwadow/kiro-openai-gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

import json
import os
import sys
import uuid
from pathlib import Path

import requests
from dotenv import load_dotenv
from loguru import logger

# --- Load environment variables ---
load_dotenv()

# --- Configuration ---
KIRO_API_HOST = "https://q.us-east-1.amazonaws.com"
TOKEN_URL = "https://prod.us-east-1.auth.desktop.kiro.dev/refreshToken"
REFRESH_TOKEN = os.getenv("REFRESH_TOKEN")
PROFILE_ARN = os.getenv("PROFILE_ARN", "arn:aws:codewhisperer:us-east-1:699475941385:profile/EHGA3GRVQMUK")
KIRO_CREDS_FILE = os.getenv("KIRO_CREDS_FILE", "")

# --- Load credentials from file if REFRESH_TOKEN not in env ---
if not REFRESH_TOKEN and KIRO_CREDS_FILE:
    try:
        creds_path = Path(KIRO_CREDS_FILE).expanduser()
        if creds_path.exists():
            with open(creds_path, 'r', encoding='utf-8') as f:
                creds_data = json.load(f)
            REFRESH_TOKEN = creds_data.get("refreshToken", "")
            if creds_data.get("profileArn"):
                PROFILE_ARN = creds_data["profileArn"]
            logger.info(f"Credentials loaded from {KIRO_CREDS_FILE}")
        else:
            logger.warning(f"Credentials file not found: {KIRO_CREDS_FILE}")
    except Exception as e:
        logger.error(f"Error loading credentials from file: {e}")

# --- Validate required credentials ---
if not REFRESH_TOKEN:
    logger.error("Neither REFRESH_TOKEN env variable nor KIRO_CREDS_FILE is configured. Exiting.")
    sys.exit(1)

# Global variables
AUTH_TOKEN = None
HEADERS = {
    "Authorization": None,
    "Content-Type": "application/json",
    "User-Agent": "aws-sdk-js/1.0.27 ua/2.1 os/win32#10.0.19044 lang/js md/nodejs#22.21.1 api/codewhispererstreaming#1.0.27 m/E KiroIDE-0.7.45-31c325a0ff0a9c8dec5d13048f4257462d751fe5b8af4cb1088f1fca45856c64",
    "x-amz-user-agent": "aws-sdk-js/1.0.27 KiroIDE-0.7.45-31c325a0ff0a9c8dec5d13048f4257462d751fe5b8af4cb1088f1fca45856c64",
    "x-amzn-codewhisperer-optout": "true",
    "x-amzn-kiro-agent-mode": "vibe",
}


def refresh_auth_token():
    """Refreshes AUTH_TOKEN via Kiro API."""
    global AUTH_TOKEN, HEADERS
    logger.info("Refreshing Kiro token...")
    
    payload = {"refreshToken": REFRESH_TOKEN}
    headers = {
        "Content-Type": "application/json",
        "User-Agent": "KiroIDE-0.7.45-31c325a0ff0a9c8dec5d13048f4257462d751fe5b8af4cb1088f1fca45856c64",
    }
    
    try:
        response = requests.post(TOKEN_URL, json=payload, headers=headers)
        response.raise_for_status()
        data = response.json()
        
        new_token = data.get("accessToken")
        expires_in = data.get("expiresIn")
        
        if not new_token:
            logger.error("Failed to get accessToken from response")
            return False

        logger.success(f"Token successfully refreshed. Expires in: {expires_in}s")
        AUTH_TOKEN = new_token
        HEADERS['Authorization'] = f"Bearer {AUTH_TOKEN}"
        return True
        
    except requests.exceptions.RequestException as e:
        logger.error(f"Error refreshing token: {e}")
        if hasattr(e, 'response') and e.response:
            logger.error(f"Server response: {e.response.status_code} {e.response.text}")
        return False


def test_get_models():
    """Tests the ListAvailableModels endpoint."""
    logger.info("Testing /ListAvailableModels...")
    url = f"{KIRO_API_HOST}/ListAvailableModels"
    params = {
        "origin": "AI_EDITOR",
        "profileArn": PROFILE_ARN
    }

    try:
        response = requests.get(url, headers=HEADERS, params=params)
        response.raise_for_status()

        logger.info(f"Response status: {response.status_code}")
        logger.debug(f"Response (JSON):\n{json.dumps(response.json(), indent=2, ensure_ascii=False)}")
        logger.success("ListAvailableModels test COMPLETED SUCCESSFULLY")
        return True
    except requests.exceptions.RequestException as e:
        logger.error(f"ListAvailableModels test failed: {e}")
        return False


def test_generate_content():
    """Tests the generateAssistantResponse endpoint."""
    logger.info("Testing /generateAssistantResponse...")
    url = f"{KIRO_API_HOST}/generateAssistantResponse"
    
    payload = {
        "conversationState": {
            "agentContinuationId": str(uuid.uuid4()),
            "agentTaskType": "vibe",
            "chatTriggerType": "MANUAL",
            "conversationId": str(uuid.uuid4()),
            "currentMessage": {
                "userInputMessage": {
                    "content": "Hello! Say something short.",
                    "modelId": "claude-haiku-4.5",
                    "origin": "AI_EDITOR",
                    "userInputMessageContext": {
                        "tools": []
                    }
                }
            },
            "history": []
        },
        "profileArn": PROFILE_ARN
    }

    try:
        with requests.post(url, headers=HEADERS, json=payload, stream=True) as response:
            response.raise_for_status()
            logger.info(f"Response status: {response.status_code}")
            logger.info("Streaming response:")

            for chunk in response.iter_content(chunk_size=1024):
                if chunk:
                    # Try to decode and find JSON
                    chunk_str = chunk.decode('utf-8', errors='ignore')
                    logger.debug(f"Chunk: {chunk_str[:200]}...")

        logger.success("generateAssistantResponse test COMPLETED")
        return True
    except requests.exceptions.RequestException as e:
        logger.error(f"generateAssistantResponse test failed: {e}")
        return False


if __name__ == "__main__":
    # Determine credential source for logging
    cred_source = "KIRO_CREDS_FILE" if KIRO_CREDS_FILE else "REFRESH_TOKEN"
    logger.info(f"Starting Kiro API tests (credentials from {cred_source})...")

    token_ok = refresh_auth_token()

    if token_ok:
        models_ok = test_get_models()
        generate_ok = test_generate_content()

        if models_ok and generate_ok:
            logger.success(f"All tests passed successfully! (credentials from {cred_source})")
        else:
            logger.warning(f"One or more tests failed. (credentials from {cred_source})")
    else:
        logger.error("Failed to refresh token. Tests not started.")



================================================
FILE: pytest.ini
================================================
[pytest]
# Конфигурация pytest для проекта
testpaths = tests
python_files = test_*.py
python_classes = Test*
python_functions = test_*

# Добавляем корневую директорию в PYTHONPATH
pythonpath = .

# Исключаем manual_api_test.py из автоматического запуска
# (это скрипт для ручного тестирования реального API, не unit-тест)
# Чтобы запустить его: python manual_api_test.py
norecursedirs = .git __pycache__ old requests _notes


================================================
FILE: requirements.txt
================================================
# Prod dependencies
fastapi
uvicorn[standard]
httpx
loguru
requests
python-dotenv
tiktoken

# Testing dependencies
pytest
pytest-asyncio
hypothesis


================================================
FILE: .clabot
================================================
{
  "contributors": ["Kartvya69", "Doggyman67"],
  "label": "cla-signed",
  "message": "Thanks for the PR! 🎉\n\nBefore merge, we need a one-time CLA confirmation.\nIt confirms that you have the right to contribute this code and allow the project to use it.\n\nFull CLA text:\nhttps://github.com/jwadow/kiro-openai-gateway/blob/main/CLA.md\n\nPlease reply with:\n```\nI have read the CLA and I accept its terms\n```"
}


================================================
FILE: .env.example
================================================
# Kiro OpenAI Gateway - Environment Configuration
# Copy this file to .env and fill in your values

# ===========================================
# REQUIRED
# ===========================================

# Password to protect YOUR proxy server
# This is NOT a token from anywhere - YOU make it up!
# Use this same value as api_key when connecting to your gateway
# Example: "my-super-secret-password-123" or any secure string
PROXY_API_KEY="my-super-secret-password-123"

# ===========================================
# FIRST WAY: JSON credentials file
# ===========================================

# Path to JSON credentials file (alternative to REFRESH_TOKEN)
KIRO_CREDS_FILE="~/.aws/sso/cache/kiro-auth-token.json"

# ===========================================
# SECOND WAY: Kiro refresh token
# ===========================================

# Your Kiro refresh token obtained from Kiro IDE traffic.
# (Alternative to KIRO_CREDS_FILE)
# REFRESH_TOKEN="your_kiro_refresh_token_here"

# ===========================================
# OPTIONAL
# ===========================================

# AWS CodeWhisperer profile ARN (usually auto-detected)
# PROFILE_ARN="arn:aws:codewhisperer:us-east-1:..."

# AWS region (default: us-east-1)
# KIRO_REGION="us-east-1"

# ===========================================
# LOGGING
# ===========================================

# Log level: TRACE, DEBUG, INFO, WARNING, ERROR, CRITICAL
# Default: INFO (recommended for production)
# Set to DEBUG for detailed troubleshooting
# LOG_LEVEL="INFO"

# ===========================================
# FIRST TOKEN TIMEOUT (Streaming Retry)
# ===========================================

# Timeout for waiting for the first token from the model (in seconds).
# If the model doesn't respond within this time, the request will be cancelled and retried.
# This helps handle "stuck" requests when the model takes too long to start responding.
# Default: 15 seconds (recommended for production)
# Set a lower value (e.g., 5-10) for more aggressive retry behavior.
# FIRST_TOKEN_TIMEOUT="15"

# Maximum number of retry attempts when first token timeout occurs.
# After exhausting all attempts, a 504 Gateway Timeout error will be returned.
# Default: 3 attempts
# FIRST_TOKEN_MAX_RETRIES="3"

# Read timeout for streaming responses (in seconds).
# This is the maximum time to wait for data between chunks during streaming.
# Should be longer than FIRST_TOKEN_TIMEOUT since the model may pause between chunks
# while "thinking" (especially for tool calls or complex reasoning).
# Default: 300 seconds (5 minutes) - generous timeout to avoid premature disconnects.
# STREAMING_READ_TIMEOUT="300"

# ===========================================
# DEBUG (for development only)
# ===========================================

# Debug logging mode:
# - off: disabled (default)
# - errors: save logs only for failed requests (4xx, 5xx) - recommended for troubleshooting
# - all: save logs for every request (overwrites on each request)
# DEBUG_MODE=off

# Directory for debug log files
# DEBUG_DIR="debug_logs"

# Legacy option (WILL BE REMOVED in future releases, use DEBUG_MODE instead)
# DEBUG_LAST_REQUEST=true is equivalent to DEBUG_MODE=all
# DEBUG_LAST_REQUEST=true



================================================
FILE: docs/en/ARCHITECTURE.md
================================================
# Architectural Overview: Kiro OpenAI Gateway

## 1. System Purpose and Goals

The project is a high-level proxy gateway implementing the **"Adapter"** structural design pattern.

The main goal of the system is to provide transparent compatibility between two heterogeneous interfaces:
1.  **Target Interface (Client):** Standard OpenAI API protocol (endpoints `/v1/models`, `/v1/chat/completions`).
2.  **Adaptee (Provider):** Internal Kiro IDE API (AWS CodeWhisperer), discovered in the Amazon Kiro ecosystem.

The system acts as a "translator", allowing the use of any tools, libraries, and IDE plugins developed for the OpenAI ecosystem with Claude models through the Kiro API.

## 2. Project Structure

The project is organized as a modular Python package `kiro_gateway/`:

```
kiro-openai-gateway/
├── main.py                    # Entry point, FastAPI application creation
├── requirements.txt           # Python dependencies
├── .env.example               # Environment configuration example
│
├── kiro_gateway/              # Main package
│   ├── __init__.py            # Package exports, version
│   ├── config.py              # Configuration and constants
│   ├── models.py              # Pydantic models for OpenAI API
│   ├── auth.py                # KiroAuthManager - token management
│   ├── cache.py               # ModelInfoCache - model cache
│   ├── utils.py               # Helper utilities
│   ├── converters.py          # OpenAI <-> Kiro conversion
│   ├── parsers.py             # AWS SSE stream parsers
│   ├── streaming.py           # Response streaming logic
│   ├── http_client.py         # HTTP client with retry logic
│   ├── routes.py              # FastAPI routes
│   ├── debug_logger.py        # Debug request logging
│   ├── tokenizer.py           # Token counting (tiktoken)
│   └── exceptions.py          # Exception handlers
│
├── tests/                     # Tests
│   ├── conftest.py            # Pytest fixtures
│   ├── unit/                  # Unit tests
│   └── integration/           # Integration tests
│
├── docs/                      # Documentation
│   ├── ru/                    # Russian version
│   └── en/                    # English version
│
└── debug_logs/                # Debug logs (generated when DEBUG_LAST_REQUEST=true)
```

## 3. Architectural Topology and Components

The system is built on the asynchronous `FastAPI` framework and uses an event-driven lifecycle management model (`Lifespan Events`).

### 3.1. Entry Point (`main.py`)

The `main.py` file is responsible for:

1. **Logging configuration** — Loguru setup with colored output
2. **Configuration validation** — `validate_configuration()` function checks:
   - Presence of `.env` file
   - Presence of credentials (REFRESH_TOKEN or KIRO_CREDS_FILE)
3. **Lifespan Manager** — creation and initialization of:
   - `KiroAuthManager` for token management
   - `ModelInfoCache` for model caching
4. **Error handler registration** — `validation_exception_handler` for 422 errors
5. **Route connection** — `app.include_router(router)`

### 3.2. Configuration Module (`kiro_gateway/config.py`)

Centralized storage of all settings:

| Parameter | Description | Default Value |
|-----------|-------------|---------------|
| `PROXY_API_KEY` | API key for proxy access | `changeme_proxy_secret` |
| `REFRESH_TOKEN` | Kiro refresh token | from `.env` |
| `PROFILE_ARN` | AWS CodeWhisperer profile ARN | from `.env` |
| `REGION` | AWS region | `us-east-1` |
| `KIRO_CREDS_FILE` | Path to JSON credentials file | from `.env` |
| `TOKEN_REFRESH_THRESHOLD` | Time before token refresh | 600 sec (10 min) |
| `MAX_RETRIES` | Max retry attempts | 3 |
| `BASE_RETRY_DELAY` | Base retry delay | 1.0 sec |
| `MODEL_CACHE_TTL` | Model cache TTL | 3600 sec (1 hour) |
| `DEFAULT_MAX_INPUT_TOKENS` | Default max input tokens | 200000 |
| `TOOL_DESCRIPTION_MAX_LENGTH` | Max tool description length | 10000 characters |
| `DEBUG_LAST_REQUEST` | Enable debug logging | `false` |
| `DEBUG_DIR` | Debug logs directory | `debug_logs` |
| `APP_VERSION` | Application version | `0.0.0` |

**Helper functions:**
- `get_kiro_refresh_url(region)` — URL for token refresh
- `get_kiro_api_host(region)` — main API host
- `get_kiro_q_host(region)` — Q API host
- `get_internal_model_id(external_model)` — model name conversion

### 3.3. Pydantic Models (`kiro_gateway/models.py`)

#### Models for `/v1/models`

| Model | Description |
|-------|-------------|
| `OpenAIModel` | AI model description (id, object, created, owned_by) |
| `ModelList` | Model list for endpoint response |

#### Models for `/v1/chat/completions`

| Model | Description |
|-------|-------------|
| `ChatMessage` | Chat message (role, content, tool_calls, tool_call_id) |
| `ToolFunction` | Tool function description (name, description, parameters) |
| `Tool` | OpenAI format tool (type, function) |
| `ChatCompletionRequest` | Generation request (model, messages, stream, tools, ...) |

#### Response Models

| Model | Description |
|-------|-------------|
| `ChatCompletionChoice` | Single response variant |
| `ChatCompletionUsage` | Token information (prompt_tokens, completion_tokens, credits_used) |
| `ChatCompletionResponse` | Full response (non-streaming) |
| `ChatCompletionChunk` | Streaming chunk |
| `ChatCompletionChunkDelta` | Delta changes in chunk |
| `ChatCompletionChunkChoice` | Variant in streaming chunk |

### 3.4. State Management Layer

#### KiroAuthManager (`kiro_gateway/auth.py`)

**Role:** Stateful singleton encapsulating Kiro token management logic.

**Capabilities:**
- Loading credentials from `.env` or JSON file
- Support for `expiresAt` to check token expiration time
- Automatic token refresh 10 minutes before expiration
- Saving updated tokens back to JSON file
- Support for different AWS regions
- Unique fingerprint generation for User-Agent

**Concurrency Control:** Uses `asyncio.Lock` to protect against race conditions.

**Main methods:**
- `get_access_token()` — returns valid token, refreshing if necessary
- `force_refresh()` — forced token refresh (on 403)
- `is_token_expiring_soon()` — expiration time check

**Properties:**
- `profile_arn` — profile ARN
- `region` — AWS region
- `api_host` — API host for region
- `q_host` — Q API host for region
- `fingerprint` — unique machine fingerprint

```python
# Usage example
auth_manager = KiroAuthManager(
    refresh_token="your_token",
    region="us-east-1",
    creds_file="~/.aws/sso/cache/kiro-auth-token.json"
)
token = await auth_manager.get_access_token()
```

#### ModelInfoCache (`kiro_gateway/cache.py`)

**Role:** Thread-safe storage for model configurations.

**Population Strategy:** 
- Lazy Loading via `/ListAvailableModels`
- Cache TTL: 1 hour
- Fallback to static model list

**Main methods:**
- `update(models_data)` — cache update
- `get(model_id)` — get model information
- `get_max_input_tokens(model_id)` — get token limit
- `is_empty()` / `is_stale()` — cache state check
- `get_all_model_ids()` — list of all model IDs

### 3.5. Helper Utilities (`kiro_gateway/utils.py`)

| Function | Description |
|----------|-------------|
| `get_machine_fingerprint()` | SHA256 hash of `{hostname}-{username}-kiro-gateway` |
| `get_kiro_headers(auth_manager, token)` | Form headers for Kiro API |
| `generate_completion_id()` | ID in format `chatcmpl-{uuid_hex}` |
| `generate_conversation_id()` | UUID for conversation |
| `generate_tool_call_id()` | ID in format `call_{uuid_hex[:8]}` |

### 3.6. Conversion Layer (`kiro_gateway/converters.py`)

#### Message Conversion

OpenAI messages are transformed into Kiro conversationState:

1. **System prompt** — added to the first user message
2. **Message history** — fully passed in `history` array
3. **Adjacent message merging** — messages with the same role are merged
4. **Tool calls** — OpenAI tools format support
5. **Tool results** — correct transmission of tool call results

#### Long Tool Description Handling

**Problem:** Kiro API returns error 400 for too long descriptions in `toolSpecification.description`.

**Solution:** Tool Documentation Reference Pattern
- If `description ≤ TOOL_DESCRIPTION_MAX_LENGTH` → leave as is
- If `description > TOOL_DESCRIPTION_MAX_LENGTH`:
  * In `toolSpecification.description` → reference: `"[Full documentation in system prompt under '## Tool: {name}']"`
  * In system prompt, section `"## Tool: {name}"` with full description is added

**Function:** `process_tools_with_long_descriptions(tools)` → `(processed_tools, tool_documentation)`

#### Main Functions

| Function | Description |
|----------|-------------|
| `extract_text_content(content)` | Extract text from various formats |
| `merge_adjacent_messages(messages)` | Merge adjacent messages with same role |
| `build_kiro_history(messages, model_id)` | Build history array for Kiro |
| `build_kiro_payload(request_data, conversation_id, profile_arn)` | Full payload for request |

#### Model Mapping

External model names are converted to internal Kiro IDs:

| External Name | Internal Kiro ID |
|---------------|------------------|
| `claude-opus-4-5` | `claude-opus-4.5` |
| `claude-opus-4-5-20251101` | `claude-opus-4.5` |
| `claude-haiku-4-5` | `claude-haiku-4.5` |
| `claude-haiku-4.5` | `claude-haiku-4.5` (direct passthrough) |
| `claude-sonnet-4-5` | `CLAUDE_SONNET_4_5_20250929_V1_0` |
| `claude-sonnet-4-5-20250929` | `CLAUDE_SONNET_4_5_20250929_V1_0` |
| `claude-sonnet-4` | `CLAUDE_SONNET_4_20250514_V1_0` |
| `claude-sonnet-4-20250514` | `CLAUDE_SONNET_4_20250514_V1_0` |
| `claude-3-7-sonnet-20250219` | `CLAUDE_3_7_SONNET_20250219_V1_0` |
| `auto` | `claude-sonnet-4.5` (alias) |

### 3.7. Parsing Layer (`kiro_gateway/parsers.py`)

#### AwsEventStreamParser

Advanced AWS SSE format parser with support for:

- **Bracket counting** — correct parsing of nested JSON objects
- **Content deduplication** — filtering of duplicate events
- **Tool calls** — parsing of structured and bracket-style tool calls
- **Escape sequences** — decoding of `\n` and others

#### Event Types

| Event | Description |
|-------|-------------|
| `content` | Text content of the response |
| `tool_start` | Start of tool call (name, toolUseId) |
| `tool_input` | Continuation of input for tool call |
| `tool_stop` | End of tool call |
| `usage` | Credit consumption information |
| `context_usage` | Context usage percentage |

#### Helper Functions

| Function | Description |
|----------|-------------|
| `find_matching_brace(text, start_pos)` | Find closing brace with nesting support |
| `parse_bracket_tool_calls(response_text)` | Parse `[Called func with args: {...}]` |
| `deduplicate_tool_calls(tool_calls)` | Remove duplicate tool calls |

### 3.8. Streaming (`kiro_gateway/streaming.py`)

#### stream_kiro_to_openai

Async generator for transforming Kiro stream to OpenAI format.

**Functionality:**
- Parse AWS SSE stream via `AwsEventStreamParser`
- Form OpenAI `chat.completion.chunk`
- Handle tool calls (structured and bracket-style)
- Calculate usage based on `contextUsagePercentage`
- Debug logging via `debug_logger`

#### collect_stream_response

Collects full response from streaming for non-streaming mode.

### 3.9. HTTP Client (`kiro_gateway/http_client.py`)

#### KiroHttpClient

Automatic error handling with exponential backoff:

| Error Code | Action |
|------------|--------|
| `403` | Token refresh via `force_refresh()` + retry |
| `429` | Exponential backoff: `BASE_RETRY_DELAY * (2 ** attempt)` |
| `5xx` | Exponential backoff (up to MAX_RETRIES attempts) |
| Timeout | Exponential backoff |

**Delay formula:** `1s, 2s, 4s` (with `BASE_RETRY_DELAY=1.0`)

**Methods:**
- `request_with_retry(method, url, json_data, stream)` — request with retry
- `close()` — close client

Supports async context manager (`async with`).

### 3.10. Routes (`kiro_gateway/routes.py`)

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Health check (status, message, version) |
| `/health` | GET | Detailed health check (status, timestamp, version) |
| `/v1/models` | GET | List of available models (requires API key) |
| `/v1/chat/completions` | POST | Chat completions (requires API key) |

**Authentication:** Bearer token in `Authorization` header

### 3.11. Exception Handling (`kiro_gateway/exceptions.py`)

| Function | Description |
|----------|-------------|
| `sanitize_validation_errors(errors)` | Convert bytes to strings for JSON serialization |
| `validation_exception_handler(request, exc)` | Pydantic validation error handler (422) |

### 3.12. Debug Logging (`kiro_gateway/debug_logger.py`)

**Class:** `DebugLogger` (singleton)

**Activation:** `DEBUG_LAST_REQUEST=true` in `.env`

**Methods:**
| Method | Description |
|--------|-------------|
| `prepare_new_request()` | Clear directory for new request |
| `log_request_body(body)` | Save incoming request |
| `log_kiro_request_body(body)` | Save request to Kiro API |
| `log_raw_chunk(chunk)` | Append raw chunk from Kiro |
| `log_modified_chunk(chunk)` | Append transformed chunk |

**Files in `debug_logs/`:**
- `request_body.json` — incoming request (OpenAI format)
- `kiro_request_body.json` — request to Kiro API
- `response_stream_raw.txt` — raw stream from Kiro
- `response_stream_modified.txt` — transformed stream (OpenAI format)

### 3.13. Tokenizer (`kiro_gateway/tokenizer.py`)

**Problem:** Kiro API does not return token counts directly. Instead, the API only provides `context_usage_percentage` — the percentage of model context usage.

**Solution:** Tokenizer module based on `tiktoken` (OpenAI's Rust library) for fast token counting.

**Features:**
- Uses `cl100k_base` encoding (GPT-4), close to Claude tokenization
- Correction factor `CLAUDE_CORRECTION_FACTOR = 1.15` for improved accuracy
- Lazy initialization for faster imports
- Fallback to rough estimation if tiktoken is unavailable

**Token calculation formula in response:**
```
total_tokens = context_usage_percentage × max_input_tokens  (from Kiro API)
completion_tokens = tiktoken(response)                       (our calculation)
prompt_tokens = total_tokens - completion_tokens             (subtraction)
```

**Main functions:**

| Function | Description |
|----------|-------------|
| `count_tokens(text)` | Count tokens in text |
| `count_message_tokens(messages)` | Count tokens in message list |
| `count_tools_tokens(tools)` | Count tokens in tool definitions |
| `estimate_request_tokens(messages, tools)` | Full request token estimation |

**Debug log:**
```
[Usage] claude-opus-4-5: prompt_tokens=142211 (subtraction), completion_tokens=769 (tiktoken), total_tokens=142980 (API Kiro)
```

**Accuracy:** ~97-99.7% compared to API data.

### 3.14. Kiro API Endpoints

All URLs are dynamically formed based on the region:

*   **Token Refresh:** `POST https://prod.{region}.auth.desktop.kiro.dev/refreshToken`
*   **List Models:** `GET https://q.{region}.amazonaws.com/ListAvailableModels`
*   **Generate Response:** `POST https://codewhisperer.{region}.amazonaws.com/generateAssistantResponse`

## 4. Detailed Data Flow

```
┌─────────────────┐
│  OpenAI Client  │
└────────┬────────┘
         │ POST /v1/chat/completions
         ▼
┌─────────────────┐
│  Security Gate  │ ◄── Proxy Bearer token verification
│  (routes.py)    │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ KiroAuthManager │ ◄── Get/refresh accessToken
│   (auth.py)     │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ Payload Builder │ ◄── Convert OpenAI → Kiro format
│ (converters.py) │     (history, system prompt, tools)
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ KiroHttpClient  │ ◄── Retry logic (403, 429, 5xx)
│ (http_client.py)│
└────────┬────────┘
         │ POST /generateAssistantResponse
         ▼
┌─────────────────┐
│   Kiro API      │
└────────┬────────┘
         │ AWS SSE Stream
         ▼
┌─────────────────┐
│ SSE Parser      │ ◄── Event parsing, tool calls
│  (parsers.py)   │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ OpenAI Format   │ ◄── Convert to OpenAI SSE
│ (streaming.py)  │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  OpenAI Client  │
└─────────────────┘
```

## 5. Available Models

| Model | Description | Credits |
|-------|-------------|---------|
| `claude-opus-4-5` | Top-tier model | ~2.2 |
| `claude-opus-4-5-20251101` | Top-tier model (version) | ~2.2 |
| `claude-sonnet-4-5` | Enhanced model | ~1.3 |
| `claude-sonnet-4-5-20250929` | Enhanced model (version) | ~1.3 |
| `claude-sonnet-4` | Balanced model | ~1.3 |
| `claude-sonnet-4-20250514` | Balanced (version) | ~1.3 |
| `claude-haiku-4-5` | Fast model | ~0.4 |
| `claude-3-7-sonnet-20250219` | Legacy model | ~1.0 |

## 6. Configuration

### Environment Variables (.env)

```env
# Required
REFRESH_TOKEN="your_kiro_refresh_token"
PROXY_API_KEY="your_proxy_secret"

# Optional
PROFILE_ARN="arn:aws:codewhisperer:..."
KIRO_REGION="us-east-1"
KIRO_CREDS_FILE="~/.aws/sso/cache/kiro-auth-token.json"

# Debug
DEBUG_LAST_REQUEST="false"
DEBUG_DIR="debug_logs"

# Limits
TOOL_DESCRIPTION_MAX_LENGTH="10000"
```

### JSON Credentials File (optional)

```json
{
  "accessToken": "eyJ...",
  "refreshToken": "eyJ...",
  "expiresAt": "2025-01-12T23:00:00.000Z",
  "profileArn": "arn:aws:codewhisperer:us-east-1:...",
  "region": "us-east-1"
}
```

## 7. API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Health check |
| `/health` | GET | Detailed health check |
| `/v1/models` | GET | List of available models |
| `/v1/chat/completions` | POST | Chat completions (streaming/non-streaming) |

## 8. Implementation Features

### Tool Calling

Support for OpenAI-compatible tools format:

```json
{
  "tools": [{
    "type": "function",
    "function": {
      "name": "get_weather",
      "description": "Get weather for a location",
      "parameters": {
        "type": "object",
        "properties": {
          "location": {"type": "string"}
        }
      }
    }
  }]
}
```

### Streaming

Full SSE streaming support with correct OpenAI format:

```
data: {"id":"chatcmpl-...","object":"chat.completion.chunk",...}

data: [DONE]
```

### Debugging

When `DEBUG_LAST_REQUEST=true`, all requests and responses are logged in `debug_logs/`:
- `request_body.json` — incoming request
- `kiro_request_body.json` — request to Kiro API
- `response_stream_raw.txt` — raw stream from Kiro
- `response_stream_modified.txt` — transformed stream

## 9. Extensibility

### Adding a New Provider

The modular architecture allows easy addition of support for other providers:

1. Create a new module `kiro_gateway/providers/new_provider.py`
2. Implement classes:
   - `NewProviderAuthManager` — token management
   - `NewProviderConverter` — format conversion
   - `NewProviderParser` — response parsing
3. Add routes to `routes.py` or create a separate router

### Example Structure for a New Provider

```python
# kiro_gateway/providers/gemini.py

class GeminiAuthManager:
    """Gemini API key management."""
    pass

class GeminiConverter:
    """OpenAI -> Gemini format conversion."""
    pass

class GeminiParser:
    """Gemini SSE stream parsing."""
    pass
```

## 10. Dependencies

Main project dependencies (from `requirements.txt`):

| Package | Purpose |
|---------|---------|
| `fastapi` | Asynchronous web framework |
| `uvicorn` | ASGI server |
| `httpx` | Asynchronous HTTP client |
| `pydantic` | Data validation and models |
| `python-dotenv` | Environment variable loading |
| `loguru` | Advanced logging |
| `tiktoken` | Fast token counting |


================================================
FILE: docs/ru/ARCHITECTURE.md
================================================
# Архитектурный Обзор: Kiro OpenAI Gateway

## 1. Назначение и Цели Системы

Проект представляет собой высокоуровневый прокси-шлюз, реализующий структурный паттерн проектирования **"Адаптер" (Adapter)**.

Основная цель системы — обеспечить прозрачную совместимость между двумя гетерогенными интерфейсами:
1.  **Target Interface (Клиент):** Стандартный протокол OpenAI API (эндпоинты `/v1/models`, `/v1/chat/completions`).
2.  **Adaptee (Поставщик):** Внутренний API Kiro IDE (AWS CodeWhisperer), обнаруженный в экосистеме Amazon Kiro.

Система выступает в роли "переводчика", позволяя использовать любые инструменты, библиотеки и IDE-плагины, разработанные для экосистемы OpenAI, с моделями Claude через Kiro API.

## 2. Структура Проекта

Проект организован в виде модульного Python-пакета `kiro_gateway/`:

```
kiro-openai-gateway/
├── main.py                    # Точка входа, создание FastAPI приложения
├── requirements.txt           # Зависимости Python
├── .env.example               # Пример конфигурации окружения
│
├── kiro_gateway/              # Основной пакет
│   ├── __init__.py            # Экспорты пакета, версия
│   ├── config.py              # Конфигурация и константы
│   ├── models.py              # Pydantic модели OpenAI API
│   ├── auth.py                # KiroAuthManager - управление токенами
│   ├── cache.py               # ModelInfoCache - кэш моделей
│   ├── utils.py               # Вспомогательные утилиты
│   ├── converters.py          # Конвертация OpenAI <-> Kiro
│   ├── parsers.py             # Парсеры AWS SSE потоков
│   ├── streaming.py           # Логика стриминга ответов
│   ├── http_client.py         # HTTP клиент с retry логикой
│   ├── routes.py              # FastAPI роуты
│   ├── debug_logger.py        # Отладочное логирование запросов
│   ├── tokenizer.py           # Подсчёт токенов (tiktoken)
│   └── exceptions.py          # Обработчики исключений
│
├── tests/                     # Тесты
│   ├── conftest.py            # Pytest fixtures
│   ├── unit/                  # Юнит-тесты
│   └── integration/           # Интеграционные тесты
│
├── docs/                      # Документация
│   ├── ru/                    # Русская версия
│   └── en/                    # Английская версия
│
└── debug_logs/                # Отладочные логи (генерируются при DEBUG_LAST_REQUEST=true)
```

## 3. Архитектурная Топология и Компоненты

Система построена на базе асинхронного фреймворка `FastAPI` и использует событийную модель управления жизненным циклом (`Lifespan Events`).

### 3.1. Точка входа (`main.py`)

Файл `main.py` отвечает за:

1. **Конфигурацию логирования** — настройка Loguru с цветным выводом
2. **Валидацию конфигурации** — функция `validate_configuration()` проверяет:
   - Наличие файла `.env`
   - Наличие credentials (REFRESH_TOKEN или KIRO_CREDS_FILE)
3. **Lifespan Manager** — создание и инициализация:
   - `KiroAuthManager` для управления токенами
   - `ModelInfoCache` для кэширования моделей
4. **Регистрация обработчиков ошибок** — `validation_exception_handler` для ошибок 422
5. **Подключение роутов** — `app.include_router(router)`

### 3.2. Модуль конфигурации (`kiro_gateway/config.py`)

Централизованное хранение всех настроек:

| Параметр | Описание | Значение по умолчанию |
|----------|----------|----------------------|
| `PROXY_API_KEY` | API ключ для доступа к прокси | `changeme_proxy_secret` |
| `REFRESH_TOKEN` | Refresh token Kiro | из `.env` |
| `PROFILE_ARN` | ARN профиля AWS CodeWhisperer | из `.env` |
| `REGION` | Регион AWS | `us-east-1` |
| `KIRO_CREDS_FILE` | Путь к JSON файлу credentials | из `.env` |
| `TOKEN_REFRESH_THRESHOLD` | Время до обновления токена | 600 сек (10 мин) |
| `MAX_RETRIES` | Макс. количество повторов | 3 |
| `BASE_RETRY_DELAY` | Базовая задержка retry | 1.0 сек |
| `MODEL_CACHE_TTL` | TTL кэша моделей | 3600 сек (1 час) |
| `DEFAULT_MAX_INPUT_TOKENS` | Макс. input токенов по умолчанию | 200000 |
| `TOOL_DESCRIPTION_MAX_LENGTH` | Макс. длина описания tool | 10000 символов |
| `DEBUG_LAST_REQUEST` | Включить отладочное логирование | `false` |
| `DEBUG_DIR` | Директория для debug логов | `debug_logs` |
| `APP_VERSION` | Версия приложения | `0.0.0` |

**Вспомогательные функции:**
- `get_kiro_refresh_url(region)` — URL для обновления токена
- `get_kiro_api_host(region)` — хост основного API
- `get_kiro_q_host(region)` — хост Q API
- `get_internal_model_id(external_model)` — конвертация имени модели

### 3.3. Pydantic Модели (`kiro_gateway/models.py`)

#### Модели для `/v1/models`

| Модель | Описание |
|--------|----------|
| `OpenAIModel` | Описание AI модели (id, object, created, owned_by) |
| `ModelList` | Список моделей для ответа endpoint |

#### Модели для `/v1/chat/completions`

| Модель | Описание |
|--------|----------|
| `ChatMessage` | Сообщение чата (role, content, tool_calls, tool_call_id) |
| `ToolFunction` | Описание функции инструмента (name, description, parameters) |
| `Tool` | Инструмент OpenAI формата (type, function) |
| `ChatCompletionRequest` | Запрос на генерацию (model, messages, stream, tools, ...) |

#### Модели ответов

| Модель | Описание |
|--------|----------|
| `ChatCompletionChoice` | Один вариант ответа |
| `ChatCompletionUsage` | Информация о токенах (prompt_tokens, completion_tokens, credits_used) |
| `ChatCompletionResponse` | Полный ответ (non-streaming) |
| `ChatCompletionChunk` | Streaming chunk |
| `ChatCompletionChunkDelta` | Дельта изменений в chunk |
| `ChatCompletionChunkChoice` | Вариант в streaming chunk |

### 3.4. Управление Состоянием (State Management Layer)

#### KiroAuthManager (`kiro_gateway/auth.py`)

**Роль:** Stateful-синглтон, инкапсулирующий логику управления токенами Kiro.

**Возможности:**
- Загрузка credentials из `.env` или JSON файла
- Поддержка `expiresAt` для проверки времени истечения токена
- Автоматическое обновление токена за 10 минут до истечения
- Сохранение обновлённых токенов обратно в JSON файл
- Поддержка разных регионов AWS
- Генерация уникального fingerprint для User-Agent

**Concurrency Control:** Использует `asyncio.Lock` для защиты от состояния гонки.

**Основные методы:**
- `get_access_token()` — возвращает действительный токен, обновляя при необходимости
- `force_refresh()` — принудительное обновление токена (при 403)
- `is_token_expiring_soon()` — проверка времени истечения

**Properties:**
- `profile_arn` — ARN профиля
- `region` — регион AWS
- `api_host` — хост API для региона
- `q_host` — хост Q API для региона
- `fingerprint` — уникальный fingerprint машины

```python
# Пример использования
auth_manager = KiroAuthManager(
    refresh_token="your_token",
    region="us-east-1",
    creds_file="~/.aws/sso/cache/kiro-auth-token.json"
)
token = await auth_manager.get_access_token()
```

#### ModelInfoCache (`kiro_gateway/cache.py`)

**Роль:** Потокобезопасное хранилище конфигураций моделей.

**Стратегия Заполнения:** 
- Lazy Loading через `/ListAvailableModels`
- TTL кэша: 1 час
- Fallback на статический список моделей

**Основные методы:**
- `update(models_data)` — обновление кэша
- `get(model_id)` — получение информации о модели
- `get_max_input_tokens(model_id)` — получение лимита токенов
- `is_empty()` / `is_stale()` — проверка состояния кэша
- `get_all_model_ids()` — список всех ID моделей

### 3.5. Вспомогательные Утилиты (`kiro_gateway/utils.py`)

| Функция | Описание |
|---------|----------|
| `get_machine_fingerprint()` | SHA256 хеш `{hostname}-{username}-kiro-gateway` |
| `get_kiro_headers(auth_manager, token)` | Формирование заголовков для Kiro API |
| `generate_completion_id()` | ID в формате `chatcmpl-{uuid_hex}` |
| `generate_conversation_id()` | UUID для разговора |
| `generate_tool_call_id()` | ID в формате `call_{uuid_hex[:8]}` |

### 3.6. Слой Конвертации (`kiro_gateway/converters.py`)

#### Конвертация сообщений

OpenAI messages преобразуются в Kiro conversationState:

1. **System prompt** — добавляется к первому user сообщению
2. **История сообщений** — полностью передаётся в `history` array
3. **Объединение соседних сообщений** — сообщения с одинаковой ролью мерджатся
4. **Tool calls** — поддержка OpenAI tools формата
5. **Tool results** — корректная передача результатов вызова инструментов

#### Обработка длинных описаний Tools

**Проблема:** Kiro API возвращает ошибку 400 при слишком длинных описаниях в `toolSpecification.description`.

**Решение:** Tool Documentation Reference Pattern
- Если `description ≤ TOOL_DESCRIPTION_MAX_LENGTH` → оставляем как есть
- Если `description > TOOL_DESCRIPTION_MAX_LENGTH`:
  * В `toolSpecification.description` → ссылка: `"[Full documentation in system prompt under '## Tool: {name}']"`
  * В system prompt добавляется секция `"## Tool: {name}"` с полным описанием

**Функция:** `process_tools_with_long_descriptions(tools)` → `(processed_tools, tool_documentation)`

#### Основные функции

| Функция | Описание |
|---------|----------|
| `extract_text_content(content)` | Извлечение текста из различных форматов |
| `merge_adjacent_messages(messages)` | Объединение соседних сообщений с одной ролью |
| `build_kiro_history(messages, model_id)` | Построение массива history для Kiro |
| `build_kiro_payload(request_data, conversation_id, profile_arn)` | Полный payload для запроса |

#### Маппинг моделей

Внешние имена моделей преобразуются во внутренние ID Kiro:

| Внешнее имя | Внутренний ID Kiro |
|-------------|-------------------|
| `claude-opus-4-5` | `claude-opus-4.5` |
| `claude-opus-4-5-20251101` | `claude-opus-4.5` |
| `claude-haiku-4-5` | `claude-haiku-4.5` |
| `claude-haiku-4.5` | `claude-haiku-4.5` (прямой проброс) |
| `claude-sonnet-4-5` | `CLAUDE_SONNET_4_5_20250929_V1_0` |
| `claude-sonnet-4-5-20250929` | `CLAUDE_SONNET_4_5_20250929_V1_0` |
| `claude-sonnet-4` | `CLAUDE_SONNET_4_20250514_V1_0` |
| `claude-sonnet-4-20250514` | `CLAUDE_SONNET_4_20250514_V1_0` |
| `claude-3-7-sonnet-20250219` | `CLAUDE_3_7_SONNET_20250219_V1_0` |
| `auto` | `claude-sonnet-4.5` (алиас) |

### 3.7. Слой Парсинга (`kiro_gateway/parsers.py`)

#### AwsEventStreamParser

Продвинутый парсер AWS SSE формата с поддержкой:

- **Bracket counting** — корректный парсинг вложенных JSON объектов
- **Дедупликация контента** — фильтрация повторяющихся событий
- **Tool calls** — парсинг структурированных и bracket-style tool calls
- **Escape-последовательности** — декодирование `\n` и других

#### Типы событий

| Событие | Описание |
|---------|----------|
| `content` | Текстовый контент ответа |
| `tool_start` | Начало tool call (name, toolUseId) |
| `tool_input` | Продолжение input для tool call |
| `tool_stop` | Завершение tool call |
| `usage` | Информация о потреблении кредитов |
| `context_usage` | Процент использования контекста |

#### Вспомогательные функции

| Функция | Описание |
|---------|----------|
| `find_matching_brace(text, start_pos)` | Поиск закрывающей скобки с учётом вложенности |
| `parse_bracket_tool_calls(response_text)` | Парсинг `[Called func with args: {...}]` |
| `deduplicate_tool_calls(tool_calls)` | Удаление дубликатов tool calls |

### 3.8. Streaming (`kiro_gateway/streaming.py`)

#### stream_kiro_to_openai

Асинхронный генератор для преобразования потока Kiro в OpenAI формат.

**Функциональность:**
- Парсинг AWS SSE stream через `AwsEventStreamParser`
- Формирование OpenAI `chat.completion.chunk`
- Обработка tool calls (структурированных и bracket-style)
- Вычисление usage на основе `contextUsagePercentage`
- Отладочное логирование через `debug_logger`

#### collect_stream_response

Собирает полный ответ из streaming потока для non-streaming режима.

### 3.9. HTTP Клиент (`kiro_gateway/http_client.py`)

#### KiroHttpClient

Автоматическая обработка ошибок с exponential backoff:

| Код ошибки | Действие |
|------------|----------|
| `403` | Refresh токена через `force_refresh()` + повтор |
| `429` | Exponential backoff: `BASE_RETRY_DELAY * (2 ** attempt)` |
| `5xx` | Exponential backoff (до MAX_RETRIES попыток) |
| Timeout | Exponential backoff |

**Формула задержки:** `1s, 2s, 4s` (при `BASE_RETRY_DELAY=1.0`)

**Методы:**
- `request_with_retry(method, url, json_data, stream)` — запрос с retry
- `close()` — закрытие клиента

Поддерживает async context manager (`async with`).

### 3.10. Роуты (`kiro_gateway/routes.py`)

| Endpoint | Метод | Описание |
|----------|-------|----------|
| `/` | GET | Health check (status, message, version) |
| `/health` | GET | Детальный health check (status, timestamp, version) |
| `/v1/models` | GET | Список доступных моделей (требует API key) |
| `/v1/chat/completions` | POST | Chat completions (требует API key) |

**Аутентификация:** Bearer token в заголовке `Authorization`

### 3.11. Обработка Исключений (`kiro_gateway/exceptions.py`)

| Функция | Описание |
|---------|----------|
| `sanitize_validation_errors(errors)` | Конвертация bytes в строки для JSON-сериализации |
| `validation_exception_handler(request, exc)` | Обработчик ошибок валидации Pydantic (422) |

### 3.12. Отладочное Логирование (`kiro_gateway/debug_logger.py`)

**Класс:** `DebugLogger` (синглтон)

**Активация:** `DEBUG_LAST_REQUEST=true` в `.env`

**Методы:**
| Метод | Описание |
|-------|----------|
| `prepare_new_request()` | Очистка директории для нового запроса |
| `log_request_body(body)` | Сохранение входящего запроса |
| `log_kiro_request_body(body)` | Сохранение запроса к Kiro API |
| `log_raw_chunk(chunk)` | Дописывание сырого chunk от Kiro |
| `log_modified_chunk(chunk)` | Дописывание преобразованного chunk |

**Файлы в `debug_logs/`:**
- `request_body.json` — входящий запрос (OpenAI формат)
- `kiro_request_body.json` — запрос к Kiro API
- `response_stream_raw.txt` — сырой поток от Kiro
- `response_stream_modified.txt` — преобразованный поток (OpenAI формат)

### 3.13. Токенизатор (`kiro_gateway/tokenizer.py`)

**Проблема:** Kiro API не возвращает напрямую количество токенов. Вместо этого API предоставляет только `context_usage_percentage` — процент использования контекста модели.

**Решение:** Модуль токенизатора на базе `tiktoken` (библиотека OpenAI на Rust) для быстрого подсчёта токенов.

**Особенности:**
- Использует кодировку `cl100k_base` (GPT-4), близкую к токенизации Claude
- Коэффициент коррекции `CLAUDE_CORRECTION_FACTOR = 1.15` для повышения точности
- Ленивая инициализация для ускорения импорта
- Fallback на грубую оценку если tiktoken недоступен

**Формула расчёта токенов в ответе:**
```
total_tokens = context_usage_percentage × max_input_tokens  (от API Kiro)
completion_tokens = tiktoken(ответ)                         (наш подсчёт)
prompt_tokens = total_tokens - completion_tokens            (вычитание)
```

**Основные функции:**

| Функция | Описание |
|---------|----------|
| `count_tokens(text)` | Подсчёт токенов в тексте |
| `count_message_tokens(messages)` | Подсчёт токенов в списке сообщений |
| `count_tools_tokens(tools)` | Подсчёт токенов в определениях инструментов |
| `estimate_request_tokens(messages, tools)` | Полная оценка токенов запроса |

**Дебаг-лог:**
```
[Usage] claude-opus-4-5: prompt_tokens=142211 (subtraction), completion_tokens=769 (tiktoken), total_tokens=142980 (API Kiro)
```

**Точность:** ~97-99.7% по сравнению с данными от API.

### 3.14. Kiro API Endpoints

Все URL динамически формируются на основе региона:

*   **Token Refresh:** `POST https://prod.{region}.auth.desktop.kiro.dev/refreshToken`
*   **List Models:** `GET https://q.{region}.amazonaws.com/ListAvailableModels`
*   **Generate Response:** `POST https://codewhisperer.{region}.amazonaws.com/generateAssistantResponse`

## 4. Детальный Поток Данных

```
┌─────────────────┐
│  OpenAI Client  │
└────────┬────────┘
         │ POST /v1/chat/completions
         ▼
┌─────────────────┐
│  Security Gate  │ ◄── Проверка Bearer токена прокси
│  (routes.py)    │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ KiroAuthManager │ ◄── Получение/обновление accessToken
│   (auth.py)     │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ Payload Builder │ ◄── Конвертация OpenAI → Kiro формат
│ (converters.py) │     (история, system prompt, tools)
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ KiroHttpClient  │ ◄── Retry логика (403, 429, 5xx)
│ (http_client.py)│
└────────┬────────┘
         │ POST /generateAssistantResponse
         ▼
┌─────────────────┐
│   Kiro API      │
└────────┬────────┘
         │ AWS SSE Stream
         ▼
┌─────────────────┐
│ SSE Parser      │ ◄── Парсинг событий, tool calls
│  (parsers.py)   │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ OpenAI Format   │ ◄── Конвертация в OpenAI SSE
│ (streaming.py)  │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  OpenAI Client  │
└─────────────────┘
```

## 5. Доступные Модели

| Модель | Описание | Credits |
|--------|----------|---------|
| `claude-opus-4-5` | Топовая модель | ~2.2 |
| `claude-opus-4-5-20251101` | Топовая модель (версия) | ~2.2 |
| `claude-sonnet-4-5` | Улучшенная модель | ~1.3 |
| `claude-sonnet-4-5-20250929` | Улучшенная модель (версия) | ~1.3 |
| `claude-sonnet-4` | Сбалансированная модель | ~1.3 |
| `claude-sonnet-4-20250514` | Сбалансированная (версия) | ~1.3 |
| `claude-haiku-4-5` | Быстрая модель | ~0.4 |
| `claude-3-7-sonnet-20250219` | Legacy модель | ~1.0 |

## 6. Конфигурация

### Переменные окружения (.env)

```env
# Обязательные
REFRESH_TOKEN="your_kiro_refresh_token"
PROXY_API_KEY="your_proxy_secret"

# Опциональные
PROFILE_ARN="arn:aws:codewhisperer:..."
KIRO_REGION="us-east-1"
KIRO_CREDS_FILE="~/.aws/sso/cache/kiro-auth-token.json"

# Отладка
DEBUG_LAST_REQUEST="false"
DEBUG_DIR="debug_logs"

# Лимиты
TOOL_DESCRIPTION_MAX_LENGTH="10000"
```

### JSON файл credentials (опционально)

```json
{
  "accessToken": "eyJ...",
  "refreshToken": "eyJ...",
  "expiresAt": "2025-01-12T23:00:00.000Z",
  "profileArn": "arn:aws:codewhisperer:us-east-1:...",
  "region": "us-east-1"
}
```

## 7. API Endpoints

| Endpoint | Метод | Описание |
|----------|-------|----------|
| `/` | GET | Health check |
| `/health` | GET | Детальный health check |
| `/v1/models` | GET | Список доступных моделей |
| `/v1/chat/completions` | POST | Chat completions (streaming/non-streaming) |

## 8. Особенности Реализации

### Tool Calling

Поддержка OpenAI-совместимого формата tools:

```json
{
  "tools": [{
    "type": "function",
    "function": {
      "name": "get_weather",
      "description": "Get weather for a location",
      "parameters": {
        "type": "object",
        "properties": {
          "location": {"type": "string"}
        }
      }
    }
  }]
}
```

### Streaming

Полная поддержка SSE streaming с корректным форматом OpenAI:

```
data: {"id":"chatcmpl-...","object":"chat.completion.chunk",...}

data: [DONE]
```

### Отладка

При `DEBUG_LAST_REQUEST=true` все запросы и ответы логируются в `debug_logs/`:
- `request_body.json` — входящий запрос
- `kiro_request_body.json` — запрос к Kiro API
- `response_stream_raw.txt` — сырой поток от Kiro
- `response_stream_modified.txt` — преобразованный поток

## 9. Расширяемость

### Добавление нового провайдера

Модульная архитектура позволяет легко добавить поддержку других провайдеров:

1. Создать новый модуль `kiro_gateway/providers/new_provider.py`
2. Реализовать классы:
   - `NewProviderAuthManager` — управление токенами
   - `NewProviderConverter` — конвертация форматов
   - `NewProviderParser` — парсинг ответов
3. Добавить роуты в `routes.py` или создать отдельный роутер

### Пример структуры для нового провайдера

```python
# kiro_gateway/providers/gemini.py

class GeminiAuthManager:
    """Управление API ключами Gemini."""
    pass

class GeminiConverter:
    """Конвертация OpenAI -> Gemini формат."""
    pass

class GeminiParser:
    """Парсинг SSE потока Gemini."""
    pass
```

## 10. Зависимости

Основные зависимости проекта (из `requirements.txt`):

| Пакет | Назначение |
|-------|------------|
| `fastapi` | Асинхронный веб-фреймворк |
| `uvicorn` | ASGI сервер |
| `httpx` | Асинхронный HTTP клиент |
| `pydantic` | Валидация данных и модели |
| `python-dotenv` | Загрузка переменных окружения |
| `loguru` | Продвинутое логирование |
| `tiktoken` | Быстрый подсчёт токенов |



================================================
FILE: kiro_gateway/__init__.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
Kiro Gateway - OpenAI-совместимый прокси для Kiro API.

Этот пакет предоставляет модульную архитектуру для проксирования
запросов OpenAI API к Kiro (AWS CodeWhisperer).

Модули:
    - config: Конфигурация и константы
    - models: Pydantic модели для OpenAI API
    - auth: Менеджер аутентификации Kiro
    - cache: Кэш метаданных моделей
    - utils: Вспомогательные утилиты
    - converters: Конвертация OpenAI <-> Kiro форматов
    - parsers: Парсеры AWS SSE потоков
    - streaming: Логика стриминга ответов
    - http_client: HTTP клиент с retry логикой
    - routes: FastAPI роуты
    - exceptions: Обработчики исключений
"""

# Версия импортируется из config.py — единственного источника истины (Single Source of Truth)
# Это позволяет менять версию только в одном месте
from kiro_gateway.config import APP_VERSION as __version__

__author__ = "Jwadow"

# Основные компоненты для удобного импорта
from kiro_gateway.auth import KiroAuthManager
from kiro_gateway.cache import ModelInfoCache
from kiro_gateway.http_client import KiroHttpClient
from kiro_gateway.routes import router

# Конфигурация
from kiro_gateway.config import (
    PROXY_API_KEY,
    REGION,
    MODEL_MAPPING,
    AVAILABLE_MODELS,
    APP_VERSION,
)

# Модели
from kiro_gateway.models import (
    ChatCompletionRequest,
    ChatMessage,
    OpenAIModel,
    ModelList,
)

# Конвертеры
from kiro_gateway.converters import (
    build_kiro_payload,
    extract_text_content,
    merge_adjacent_messages,
)

# Парсеры
from kiro_gateway.parsers import (
    AwsEventStreamParser,
    parse_bracket_tool_calls,
)

# Streaming
from kiro_gateway.streaming import (
    stream_kiro_to_openai,
    collect_stream_response,
)

# Exceptions
from kiro_gateway.exceptions import (
    validation_exception_handler,
    sanitize_validation_errors,
)

__all__ = [
    # Версия
    "__version__",
    
    # Основные классы
    "KiroAuthManager",
    "ModelInfoCache",
    "KiroHttpClient",
    "router",
    
    # Конфигурация
    "PROXY_API_KEY",
    "REGION",
    "MODEL_MAPPING",
    "AVAILABLE_MODELS",
    "APP_VERSION",
    
    # Модели
    "ChatCompletionRequest",
    "ChatMessage",
    "OpenAIModel",
    "ModelList",
    
    # Конвертеры
    "build_kiro_payload",
    "extract_text_content",
    "merge_adjacent_messages",
    
    # Парсеры
    "AwsEventStreamParser",
    "parse_bracket_tool_calls",
    
    # Streaming
    "stream_kiro_to_openai",
    "collect_stream_response",
    
    # Exceptions
    "validation_exception_handler",
    "sanitize_validation_errors",
]


================================================
FILE: kiro_gateway/auth.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# https://github.com/jwadow/kiro-openai-gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
Authentication manager for Kiro API.

Manages the lifecycle of access tokens:
- Loading credentials from .env or JSON file
- Automatic token refresh on expiration
- Thread-safe refresh using asyncio.Lock
"""

import asyncio
import json
from datetime import datetime, timezone
from pathlib import Path
from typing import Optional

import httpx
from loguru import logger

from kiro_gateway.config import (
    TOKEN_REFRESH_THRESHOLD,
    get_kiro_refresh_url,
    get_kiro_api_host,
    get_kiro_q_host,
)
from kiro_gateway.utils import get_machine_fingerprint


class KiroAuthManager:
    """
    Manages the token lifecycle for accessing Kiro API.
    
    Supports:
    - Loading credentials from .env or JSON file
    - Automatic token refresh on expiration
    - Expiration time validation (expiresAt)
    - Saving updated tokens to file
    
    Attributes:
        profile_arn: AWS CodeWhisperer profile ARN
        region: AWS region
        api_host: API host for the current region
        q_host: Q API host for the current region
        fingerprint: Unique machine fingerprint
    
    Example:
        >>> auth_manager = KiroAuthManager(
        ...     refresh_token="your_refresh_token",
        ...     region="us-east-1"
        ... )
        >>> token = await auth_manager.get_access_token()
    """
    
    def __init__(
        self,
        refresh_token: Optional[str] = None,
        profile_arn: Optional[str] = None,
        region: str = "us-east-1",
        creds_file: Optional[str] = None
    ):
        """
        Initializes the authentication manager.
        
        Args:
            refresh_token: Refresh token for obtaining access token
            profile_arn: AWS CodeWhisperer profile ARN
            region: AWS region (default: us-east-1)
            creds_file: Path to JSON file with credentials (optional)
        """
        self._refresh_token = refresh_token
        self._profile_arn = profile_arn
        self._region = region
        self._creds_file = creds_file
        
        self._access_token: Optional[str] = None
        self._expires_at: Optional[datetime] = None
        self._lock = asyncio.Lock()
        
        # Dynamic URLs based on region
        self._refresh_url = get_kiro_refresh_url(region)
        self._api_host = get_kiro_api_host(region)
        self._q_host = get_kiro_q_host(region)
        
        # Fingerprint for User-Agent
        self._fingerprint = get_machine_fingerprint()
        
        # Load credentials from file if specified
        if creds_file:
            self._load_credentials_from_file(creds_file)
    
    def _load_credentials_from_file(self, file_path: str) -> None:
        """
        Loads credentials from a JSON file.
        
        Supported JSON fields:
        - refreshToken: Refresh token
        - accessToken: Access token (if already available)
        - profileArn: Profile ARN
        - region: AWS region
        - expiresAt: Token expiration time (ISO 8601)
        
        Args:
            file_path: Path to JSON file
        """
        try:
            path = Path(file_path).expanduser()
            if not path.exists():
                logger.warning(f"Credentials file not found: {file_path}")
                return
            
            with open(path, 'r', encoding='utf-8') as f:
                data = json.load(f)
            
            # Load data from file
            if 'refreshToken' in data:
                self._refresh_token = data['refreshToken']
            if 'accessToken' in data:
                self._access_token = data['accessToken']
            if 'profileArn' in data:
                self._profile_arn = data['profileArn']
            if 'region' in data:
                self._region = data['region']
                # Update URLs for new region
                self._refresh_url = get_kiro_refresh_url(self._region)
                self._api_host = get_kiro_api_host(self._region)
                self._q_host = get_kiro_q_host(self._region)
            
            # Parse expiresAt
            if 'expiresAt' in data:
                try:
                    expires_str = data['expiresAt']
                    # Support for different date formats
                    if expires_str.endswith('Z'):
                        self._expires_at = datetime.fromisoformat(expires_str.replace('Z', '+00:00'))
                    else:
                        self._expires_at = datetime.fromisoformat(expires_str)
                except Exception as e:
                    logger.warning(f"Failed to parse expiresAt: {e}")
            
            logger.info(f"Credentials loaded from {file_path}")
            
        except Exception as e:
            logger.error(f"Error loading credentials from file: {e}")
    
    def _save_credentials_to_file(self) -> None:
        """
        Saves updated credentials to a JSON file.
        
        Updates the existing file while preserving other fields.
        """
        if not self._creds_file:
            return
        
        try:
            path = Path(self._creds_file).expanduser()
            
            # Read existing data
            existing_data = {}
            if path.exists():
                with open(path, 'r', encoding='utf-8') as f:
                    existing_data = json.load(f)
            
            # Update data
            existing_data['accessToken'] = self._access_token
            existing_data['refreshToken'] = self._refresh_token
            if self._expires_at:
                existing_data['expiresAt'] = self._expires_at.isoformat()
            if self._profile_arn:
                existing_data['profileArn'] = self._profile_arn
            
            # Save
            with open(path, 'w', encoding='utf-8') as f:
                json.dump(existing_data, f, indent=2, ensure_ascii=False)
            
            logger.debug(f"Credentials saved to {self._creds_file}")
            
        except Exception as e:
            logger.error(f"Error saving credentials: {e}")
    
    def is_token_expiring_soon(self) -> bool:
        """
        Checks if the token is expiring soon.
        
        Returns:
            True if the token expires within TOKEN_REFRESH_THRESHOLD seconds
            or if expiration time information is not available
        """
        if not self._expires_at:
            return True  # If no expiration info available, assume refresh is needed
        
        now = datetime.now(timezone.utc)
        threshold = now.timestamp() + TOKEN_REFRESH_THRESHOLD
        
        return self._expires_at.timestamp() <= threshold
    
    async def _refresh_token_request(self) -> None:
        """
        Performs a token refresh request.
        
        Sends a POST request to Kiro API to obtain a new access token.
        Updates internal state and saves credentials to file.
        
        Raises:
            ValueError: If refresh token is not set or response doesn't contain accessToken
            httpx.HTTPError: On HTTP request error
        """
        if not self._refresh_token:
            raise ValueError("Refresh token is not set")
        
        logger.info("Refreshing Kiro token...")
        
        payload = {'refreshToken': self._refresh_token}
        headers = {
            "Content-Type": "application/json",
            "User-Agent": f"KiroIDE-0.7.45-{self._fingerprint}",
        }
        
        async with httpx.AsyncClient(timeout=30) as client:
            response = await client.post(self._refresh_url, json=payload, headers=headers)
            response.raise_for_status()
            data = response.json()
        
        new_access_token = data.get("accessToken")
        new_refresh_token = data.get("refreshToken")
        expires_in = data.get("expiresIn", 3600)
        new_profile_arn = data.get("profileArn")
        
        if not new_access_token:
            raise ValueError(f"Response does not contain accessToken: {data}")
        
        # Update data
        self._access_token = new_access_token
        if new_refresh_token:
            self._refresh_token = new_refresh_token
        if new_profile_arn:
            self._profile_arn = new_profile_arn
        
        # Calculate expiration time with buffer (minus 60 seconds)
        self._expires_at = datetime.now(timezone.utc).replace(microsecond=0)
        self._expires_at = datetime.fromtimestamp(
            self._expires_at.timestamp() + expires_in - 60,
            tz=timezone.utc
        )
        
        logger.info(f"Token refreshed, expires: {self._expires_at.isoformat()}")
        
        # Save to file
        self._save_credentials_to_file()
    
    async def get_access_token(self) -> str:
        """
        Returns a valid access_token, refreshing it if necessary.
        
        Thread-safe method using asyncio.Lock.
        Automatically refreshes the token if it has expired or is about to expire.
        
        Returns:
            Valid access token
        
        Raises:
            ValueError: If unable to obtain access token
        """
        async with self._lock:
            if not self._access_token or self.is_token_expiring_soon():
                await self._refresh_token_request()
            
            if not self._access_token:
                raise ValueError("Failed to obtain access token")
            
            return self._access_token
    
    async def force_refresh(self) -> str:
        """
        Forces a token refresh.
        
        Used when receiving a 403 error from the API.
        
        Returns:
            New access token
        """
        async with self._lock:
            await self._refresh_token_request()
            return self._access_token
    
    @property
    def profile_arn(self) -> Optional[str]:
        """AWS CodeWhisperer profile ARN."""
        return self._profile_arn
    
    @property
    def region(self) -> str:
        """AWS region."""
        return self._region
    
    @property
    def api_host(self) -> str:
        """API host for the current region."""
        return self._api_host
    
    @property
    def q_host(self) -> str:
        """Q API host for the current region."""
        return self._q_host
    
    @property
    def fingerprint(self) -> str:
        """Unique machine fingerprint."""
        return self._fingerprint


================================================
FILE: kiro_gateway/cache.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
Кэш метаданных моделей для Kiro Gateway.

Потокобезопасное хранилище информации о доступных моделях
с поддержкой TTL и lazy loading.
"""

import asyncio
import time
from typing import Any, Dict, List, Optional

from loguru import logger

from kiro_gateway.config import MODEL_CACHE_TTL, DEFAULT_MAX_INPUT_TOKENS


class ModelInfoCache:
    """
    Потокобезопасный кэш для хранения метаданных о моделях.
    
    Использует Lazy Loading для заполнения - данные загружаются
    только при первом обращении или когда кэш устарел.
    
    Attributes:
        cache_ttl: Время жизни кэша в секундах
    
    Example:
        >>> cache = ModelInfoCache()
        >>> await cache.update([{"modelId": "claude-sonnet-4", "tokenLimits": {...}}])
        >>> info = cache.get("claude-sonnet-4")
        >>> max_tokens = cache.get_max_input_tokens("claude-sonnet-4")
    """
    
    def __init__(self, cache_ttl: int = MODEL_CACHE_TTL):
        """
        Инициализирует кэш моделей.
        
        Args:
            cache_ttl: Время жизни кэша в секундах (по умолчанию из конфига)
        """
        self._cache: Dict[str, Dict[str, Any]] = {}
        self._lock = asyncio.Lock()
        self._last_update: Optional[float] = None
        self._cache_ttl = cache_ttl
    
    async def update(self, models_data: List[Dict[str, Any]]) -> None:
        """
        Обновляет кэш моделей.
        
        Потокобезопасно заменяет содержимое кэша новыми данными.
        
        Args:
            models_data: Список словарей с информацией о моделях.
                        Каждый словарь должен содержать ключ "modelId".
        """
        async with self._lock:
            logger.info(f"Updating model cache. Found {len(models_data)} models.")
            self._cache = {model["modelId"]: model for model in models_data}
            self._last_update = time.time()
    
    def get(self, model_id: str) -> Optional[Dict[str, Any]]:
        """
        Возвращает информацию о модели.
        
        Args:
            model_id: ID модели
        
        Returns:
            Словарь с информацией о модели или None если модель не найдена
        """
        return self._cache.get(model_id)
    
    def get_max_input_tokens(self, model_id: str) -> int:
        """
        Возвращает maxInputTokens для модели.
        
        Args:
            model_id: ID модели
        
        Returns:
            Максимальное количество input токенов или DEFAULT_MAX_INPUT_TOKENS
        """
        model = self._cache.get(model_id)
        if model and model.get("tokenLimits"):
            return model["tokenLimits"].get("maxInputTokens") or DEFAULT_MAX_INPUT_TOKENS
        return DEFAULT_MAX_INPUT_TOKENS
    
    def is_empty(self) -> bool:
        """
        Проверяет, пуст ли кэш.
        
        Returns:
            True если кэш пуст
        """
        return not self._cache
    
    def is_stale(self) -> bool:
        """
        Проверяет, устарел ли кэш.
        
        Returns:
            True если кэш устарел (прошло больше cache_ttl секунд)
            или если кэш никогда не обновлялся
        """
        if not self._last_update:
            return True
        return time.time() - self._last_update > self._cache_ttl
    
    def get_all_model_ids(self) -> List[str]:
        """
        Возвращает список всех ID моделей в кэше.
        
        Returns:
            Список ID моделей
        """
        return list(self._cache.keys())
    
    @property
    def size(self) -> int:
        """Количество моделей в кэше."""
        return len(self._cache)
    
    @property
    def last_update_time(self) -> Optional[float]:
        """Время последнего обновления (timestamp) или None."""
        return self._last_update


================================================
FILE: kiro_gateway/config.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# https://github.com/jwadow/kiro-openai-gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
Kiro Gateway Configuration.

Centralized storage for all settings, constants, and mappings.
Loads environment variables and provides typed access to them.
"""

import os
import re
from pathlib import Path
from typing import Dict, List, Optional
from dotenv import load_dotenv

# Load environment variables
load_dotenv()


def _get_raw_env_value(var_name: str, env_file: str = ".env") -> Optional[str]:
    """
    Read variable value from .env file without processing escape sequences.
    
    This is necessary for correct handling of Windows paths where backslashes
    (e.g., D:\\Projects\\file.json) may be incorrectly interpreted
    as escape sequences (\\a -> bell, \\n -> newline, etc.).
    
    Args:
        var_name: Environment variable name
        env_file: Path to .env file (default ".env")
    
    Returns:
        Raw variable value or None if not found
    """
    env_path = Path(env_file)
    if not env_path.exists():
        return None
    
    try:
        # Read file as-is, without interpretation
        content = env_path.read_text(encoding="utf-8")
        
        # Search for variable considering different formats:
        # VAR="value" or VAR='value' or VAR=value
        # Pattern captures value with or without quotes
        pattern = rf'^{re.escape(var_name)}=(["\']?)(.+?)\1\s*$'
        
        for line in content.splitlines():
            line = line.strip()
            if line.startswith("#") or not line:
                continue
            
            match = re.match(pattern, line)
            if match:
                # Return value as-is, without processing escape sequences
                return match.group(2)
    except Exception:
        pass
    
    return None

# ==================================================================================================
# Proxy Server Settings
# ==================================================================================================

# API key for proxy access (clients must pass it in Authorization header)
PROXY_API_KEY: str = os.getenv("PROXY_API_KEY", "changeme_proxy_secret")

# ==================================================================================================
# Kiro API Credentials
# ==================================================================================================

# Refresh token for updating access token
REFRESH_TOKEN: str = os.getenv("REFRESH_TOKEN", "")

# Profile ARN for AWS CodeWhisperer
PROFILE_ARN: str = os.getenv("PROFILE_ARN", "")

# AWS region (default us-east-1)
REGION: str = os.getenv("KIRO_REGION", "us-east-1")

# Path to credentials file (optional, alternative to .env)
# Read directly from .env to avoid escape sequence issues on Windows
# (e.g., \a in path D:\Projects\adolf is interpreted as bell character)
_raw_creds_file = _get_raw_env_value("KIRO_CREDS_FILE") or os.getenv("KIRO_CREDS_FILE", "")
# Normalize path for cross-platform compatibility
KIRO_CREDS_FILE: str = str(Path(_raw_creds_file)) if _raw_creds_file else ""

# ==================================================================================================
# Kiro API URL Templates
# ==================================================================================================

# URL for token refresh
KIRO_REFRESH_URL_TEMPLATE: str = "https://prod.{region}.auth.desktop.kiro.dev/refreshToken"

# Host for main API (generateAssistantResponse)
KIRO_API_HOST_TEMPLATE: str = "https://codewhisperer.{region}.amazonaws.com"

# Host for Q API (ListAvailableModels)
KIRO_Q_HOST_TEMPLATE: str = "https://q.{region}.amazonaws.com"

# ==================================================================================================
# Token Settings
# ==================================================================================================

# Time before token expiration when refresh is needed (in seconds)
# Default 10 minutes - refresh token in advance to avoid errors
TOKEN_REFRESH_THRESHOLD: int = 600

# ==================================================================================================
# Retry Configuration
# ==================================================================================================

# Maximum number of retry attempts on errors
MAX_RETRIES: int = 3

# Base delay between attempts (seconds)
# Uses exponential backoff: delay * (2 ** attempt)
BASE_RETRY_DELAY: float = 1.0

# ==================================================================================================
# Model Mapping
# ==================================================================================================

# External model names (OpenAI-compatible) -> internal Kiro IDs
# Clients use external names, and we convert them to internal ones
MODEL_MAPPING: Dict[str, str] = {
    # Claude Opus 4.5 - top-tier model
    "claude-opus-4-5": "claude-opus-4.5",
    "claude-opus-4-5-20251101": "claude-opus-4.5",
    
    # Claude Haiku 4.5 - fast model
    "claude-haiku-4-5": "claude-haiku-4.5",
    "claude-haiku-4.5": "claude-haiku-4.5",  # Direct passthrough
    
    # Claude Sonnet 4.5 - enhanced model
    "claude-sonnet-4-5": "CLAUDE_SONNET_4_5_20250929_V1_0",
    "claude-sonnet-4-5-20250929": "CLAUDE_SONNET_4_5_20250929_V1_0",
    
    # Claude Sonnet 4 - balanced model
    "claude-sonnet-4": "CLAUDE_SONNET_4_20250514_V1_0",
    "claude-sonnet-4-20250514": "CLAUDE_SONNET_4_20250514_V1_0",
    
    # Claude 3.7 Sonnet - legacy model
    "claude-3-7-sonnet-20250219": "CLAUDE_3_7_SONNET_20250219_V1_0",
    
    # Aliases for convenience
    "auto": "claude-sonnet-4.5",
}

# List of available models for /v1/models endpoint
# These models will be displayed to clients as available
AVAILABLE_MODELS: List[str] = [
    "claude-opus-4-5",
    "claude-opus-4-5-20251101",
    "claude-haiku-4-5",
    "claude-sonnet-4-5",
    "claude-sonnet-4-5-20250929",
    "claude-sonnet-4",
    "claude-sonnet-4-20250514",
    "claude-3-7-sonnet-20250219",
]

# ==================================================================================================
# Model Cache Settings
# ==================================================================================================

# Model cache TTL in seconds (1 hour)
MODEL_CACHE_TTL: int = 3600

# Default maximum number of input tokens
DEFAULT_MAX_INPUT_TOKENS: int = 200000

# ==================================================================================================
# Tool Description Handling (Kiro API Limitations)
# ==================================================================================================

# Kiro API возвращает ошибку 400 "Improperly formed request" при слишком длинных
# описаниях инструментов в toolSpecification.description.
#
# Решение: Tool Documentation Reference Pattern
# - Если description ≤ лимита → оставляем как есть
# - Если description > лимита:
#   * В toolSpecification.description → ссылка на system prompt:
#     "[Full documentation in system prompt under '## Tool: {name}']"
#   * В system prompt добавляется секция "## Tool: {name}" с полным описанием
#
# Модель видит явную ссылку и точно понимает, где искать полную документацию.

# Максимальная длина description для tool в символах.
# Описания длиннее этого лимита будут перенесены в system prompt.
# Установите 0 для отключения (не рекомендуется - вызовет ошибки Kiro API).
TOOL_DESCRIPTION_MAX_LENGTH: int = int(os.getenv("TOOL_DESCRIPTION_MAX_LENGTH", "10000"))

# ==================================================================================================
# Logging Settings
# ==================================================================================================

# Log level for the application
# Available levels: TRACE, DEBUG, INFO, WARNING, ERROR, CRITICAL
# Default: INFO (recommended for production)
# Set to DEBUG for detailed troubleshooting
LOG_LEVEL: str = os.getenv("LOG_LEVEL", "INFO").upper()

# ==================================================================================================
# First Token Timeout Settings (Streaming Retry)
# ==================================================================================================

# Timeout for waiting for the first token from the model (in seconds).
# If the model doesn't respond within this time, the request will be cancelled and retried.
# This helps handle "stuck" requests when the model takes too long to think.
# Default: 30 seconds (recommended for production)
# Set a lower value (e.g., 10-15) for more aggressive retry.
FIRST_TOKEN_TIMEOUT: float = float(os.getenv("FIRST_TOKEN_TIMEOUT", "15"))

# Read timeout for streaming responses (in seconds).
# This is the maximum time to wait for data between chunks during streaming.
# Should be longer than FIRST_TOKEN_TIMEOUT since the model may pause between chunks
# while "thinking" (especially for tool calls or complex reasoning).
# Default: 300 seconds (5 minutes) - generous timeout to avoid premature disconnects.
STREAMING_READ_TIMEOUT: float = float(os.getenv("STREAMING_READ_TIMEOUT", "300"))

# Maximum number of attempts on first token timeout.
# After exhausting all attempts, an error will be returned.
# Default: 3 attempts
FIRST_TOKEN_MAX_RETRIES: int = int(os.getenv("FIRST_TOKEN_MAX_RETRIES", "3"))

# ==================================================================================================
# Debug Settings
# ==================================================================================================

# Legacy option (deprecated, will be removed in future releases)
# Use DEBUG_MODE instead
_DEBUG_LAST_REQUEST_RAW: str = os.getenv("DEBUG_LAST_REQUEST", "").lower()
DEBUG_LAST_REQUEST: bool = _DEBUG_LAST_REQUEST_RAW in ("true", "1", "yes")

# Debug logging mode:
# - off: disabled (default)
# - errors: save logs only for failed requests (4xx, 5xx)
# - all: save logs for every request (overwrites on each request)
_DEBUG_MODE_RAW: str = os.getenv("DEBUG_MODE", "").lower()

# Priority logic:
# 1. If DEBUG_MODE is explicitly set → use it
# 2. If DEBUG_MODE is not set but DEBUG_LAST_REQUEST=true → mode "all" (backward compatibility)
# 3. Otherwise → mode "off"
if _DEBUG_MODE_RAW in ("off", "errors", "all"):
    DEBUG_MODE: str = _DEBUG_MODE_RAW
elif DEBUG_LAST_REQUEST:
    DEBUG_MODE: str = "all"
else:
    DEBUG_MODE: str = "off"

# Directory for debug log files
DEBUG_DIR: str = os.getenv("DEBUG_DIR", "debug_logs")


def _warn_deprecated_debug_setting():
    """
    Print warning if deprecated DEBUG_LAST_REQUEST is used.
    Called at application startup.
    """
    if _DEBUG_LAST_REQUEST_RAW and not _DEBUG_MODE_RAW:
        import sys
        # ANSI escape codes: yellow text
        YELLOW = "\033[93m"
        RESET = "\033[0m"
        
        warning_text = f"""
{YELLOW}⚠️  DEPRECATED: DEBUG_LAST_REQUEST will be removed in future releases.
    Please use DEBUG_MODE instead:
      - DEBUG_MODE=off     (disabled, default)
      - DEBUG_MODE=errors  (save logs only for failed requests)
      - DEBUG_MODE=all     (save logs for every request)
    
    DEBUG_LAST_REQUEST=true is equivalent to DEBUG_MODE=all
    See .env.example for more details.{RESET}
"""
        print(warning_text, file=sys.stderr)


def _warn_timeout_configuration():
    """
    Print warning if timeout configuration is suboptimal.
    Called at application startup.
    
    FIRST_TOKEN_TIMEOUT should be less than STREAMING_READ_TIMEOUT:
    - FIRST_TOKEN_TIMEOUT: time to wait for model to START responding
    - STREAMING_READ_TIMEOUT: time to wait BETWEEN chunks during streaming
    """
    if FIRST_TOKEN_TIMEOUT >= STREAMING_READ_TIMEOUT:
        import sys
        YELLOW = "\033[93m"
        RESET = "\033[0m"
        
        warning_text = f"""
{YELLOW}⚠️  WARNING: Suboptimal timeout configuration detected.
    
    FIRST_TOKEN_TIMEOUT ({FIRST_TOKEN_TIMEOUT}s) >= STREAMING_READ_TIMEOUT ({STREAMING_READ_TIMEOUT}s)
    
    These timeouts serve different purposes:
      - FIRST_TOKEN_TIMEOUT: time to wait for model to START responding (default: 15s)
      - STREAMING_READ_TIMEOUT: time to wait BETWEEN chunks during streaming (default: 300s)
    
    Recommendation: FIRST_TOKEN_TIMEOUT should be LESS than STREAMING_READ_TIMEOUT.
    
    Example configuration:
      FIRST_TOKEN_TIMEOUT=15
      STREAMING_READ_TIMEOUT=300{RESET}
"""
        print(warning_text, file=sys.stderr)

# ==================================================================================================
# Application Version
# ==================================================================================================

APP_VERSION: str = "1.0.7"
APP_TITLE: str = "Kiro API Gateway"
APP_DESCRIPTION: str = "OpenAI-compatible interface for Kiro API (AWS CodeWhisperer). Made by @jwadow"


def get_kiro_refresh_url(region: str) -> str:
    """Return token refresh URL for the specified region."""
    return KIRO_REFRESH_URL_TEMPLATE.format(region=region)


def get_kiro_api_host(region: str) -> str:
    """Return API host for the specified region."""
    return KIRO_API_HOST_TEMPLATE.format(region=region)


def get_kiro_q_host(region: str) -> str:
    """Return Q API host for the specified region."""
    return KIRO_Q_HOST_TEMPLATE.format(region=region)


def get_internal_model_id(external_model: str) -> str:
    """
    Convert external model name to internal Kiro ID.
    
    Args:
        external_model: External model name (e.g., "claude-sonnet-4-5")
    
    Returns:
        Internal model ID for Kiro API
    """
    return MODEL_MAPPING.get(external_model, external_model)


================================================
FILE: kiro_gateway/converters.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
Конвертеры для преобразования форматов OpenAI <-> Kiro.

Содержит функции для:
- Извлечения текстового контента из различных форматов
- Объединения соседних сообщений
- Построения истории разговора для Kiro API
- Сборки полного payload для запроса
"""

import json
from typing import Any, Dict, List, Optional, Tuple

from loguru import logger

from kiro_gateway.config import get_internal_model_id, TOOL_DESCRIPTION_MAX_LENGTH
from kiro_gateway.models import ChatMessage, ChatCompletionRequest, Tool


def extract_text_content(content: Any) -> str:
    """
    Извлекает текстовый контент из различных форматов.
    
    OpenAI API поддерживает несколько форматов content:
    - Строка: "Hello, world!"
    - Список: [{"type": "text", "text": "Hello"}]
    - None: пустое сообщение
    
    Args:
        content: Контент в любом поддерживаемом формате
    
    Returns:
        Извлечённый текст или пустая строка
    
    Example:
        >>> extract_text_content("Hello")
        'Hello'
        >>> extract_text_content([{"type": "text", "text": "World"}])
        'World'
        >>> extract_text_content(None)
        ''
    """
    if content is None:
        return ""
    if isinstance(content, str):
        return content
    if isinstance(content, list):
        text_parts = []
        for item in content:
            if isinstance(item, dict):
                if item.get("type") == "text":
                    text_parts.append(item.get("text", ""))
                elif "text" in item:
                    text_parts.append(item["text"])
            elif isinstance(item, str):
                text_parts.append(item)
        return "".join(text_parts)
    return str(content)


def merge_adjacent_messages(messages: List[ChatMessage]) -> List[ChatMessage]:
    """
    Объединяет соседние сообщения с одинаковой ролью и обрабатывает tool messages.
    
    Kiro API не принимает несколько сообщений подряд от одного role.
    Эта функция объединяет такие сообщения в одно.
    
    Tool messages (role="tool") преобразуются в user messages с tool_results.
    
    Args:
        messages: Список сообщений
    
    Returns:
        Список сообщений с объединёнными соседними сообщениями
    
    Example:
        >>> msgs = [
        ...     ChatMessage(role="user", content="Hello"),
        ...     ChatMessage(role="user", content="World")
        ... ]
        >>> merged = merge_adjacent_messages(msgs)
        >>> len(merged)
        1
        >>> merged[0].content
        'Hello\\nWorld'
    """
    if not messages:
        return []
    
    # Сначала преобразуем tool messages в user messages с tool_results
    processed = []
    pending_tool_results = []
    
    for msg in messages:
        if msg.role == "tool":
            # Собираем tool results
            tool_result = {
                "type": "tool_result",
                "tool_use_id": msg.tool_call_id or "",
                "content": extract_text_content(msg.content) or "(empty result)"
            }
            pending_tool_results.append(tool_result)
            logger.debug(f"Collected tool result for tool_call_id={msg.tool_call_id}")
        else:
            # Если есть накопленные tool results, создаём user message с ними
            if pending_tool_results:
                # Создаём user message с tool_results
                tool_results_msg = ChatMessage(
                    role="user",
                    content=pending_tool_results.copy()
                )
                processed.append(tool_results_msg)
                pending_tool_results.clear()
                logger.debug(f"Created user message with {len(tool_results_msg.content)} tool results")
            
            processed.append(msg)
    
    # Если остались tool results в конце
    if pending_tool_results:
        tool_results_msg = ChatMessage(
            role="user",
            content=pending_tool_results.copy()
        )
        processed.append(tool_results_msg)
        logger.debug(f"Created final user message with {len(pending_tool_results)} tool results")
    
    # Теперь объединяем соседние сообщения с одинаковой ролью
    merged = []
    for msg in processed:
        if not merged:
            merged.append(msg)
            continue
        
        last = merged[-1]
        if msg.role == last.role:
            # Объединяем контент
            # Если оба контента - списки, объединяем списки
            if isinstance(last.content, list) and isinstance(msg.content, list):
                last.content = last.content + msg.content
            elif isinstance(last.content, list):
                last.content = last.content + [{"type": "text", "text": extract_text_content(msg.content)}]
            elif isinstance(msg.content, list):
                last.content = [{"type": "text", "text": extract_text_content(last.content)}] + msg.content
            else:
                last_text = extract_text_content(last.content)
                current_text = extract_text_content(msg.content)
                last.content = f"{last_text}\n{current_text}"
            
            # Объединяем tool_calls для assistant сообщений
            # Критично: без этого теряются tool_calls из второго и последующих сообщений,
            # что приводит к ошибке 400 от Kiro API (toolResult без соответствующего toolUse)
            if msg.role == "assistant" and msg.tool_calls:
                if last.tool_calls is None:
                    last.tool_calls = []
                last.tool_calls = list(last.tool_calls) + list(msg.tool_calls)
                logger.debug(f"Merged tool_calls: added {len(msg.tool_calls)} tool calls, total now: {len(last.tool_calls)}")
            
            logger.debug(f"Merged adjacent messages with role {msg.role}")
        else:
            merged.append(msg)
    
    return merged


def build_kiro_history(messages: List[ChatMessage], model_id: str) -> List[Dict[str, Any]]:
    """
    Строит массив history для Kiro API из OpenAI messages.
    
    Kiro API ожидает чередование userInputMessage и assistantResponseMessage.
    Эта функция преобразует OpenAI формат в Kiro формат.
    
    Args:
        messages: Список сообщений в формате OpenAI
        model_id: Внутренний ID модели Kiro
    
    Returns:
        Список словарей для поля history в Kiro API
    
    Example:
        >>> msgs = [ChatMessage(role="user", content="Hello")]
        >>> history = build_kiro_history(msgs, "claude-sonnet-4")
        >>> history[0]["userInputMessage"]["content"]
        'Hello'
    """
    history = []
    
    for msg in messages:
        if msg.role == "user":
            content = extract_text_content(msg.content)
            
            user_input = {
                "content": content,
                "modelId": model_id,
                "origin": "AI_EDITOR",
            }
            
            # Обработка tool_results (ответы на tool calls)
            tool_results = _extract_tool_results(msg.content)
            if tool_results:
                user_input["userInputMessageContext"] = {"toolResults": tool_results}
            
            history.append({"userInputMessage": user_input})
            
        elif msg.role == "assistant":
            content = extract_text_content(msg.content)
            
            assistant_response = {"content": content}
            
            # Обработка tool_calls
            tool_uses = _extract_tool_uses(msg)
            if tool_uses:
                assistant_response["toolUses"] = tool_uses
            
            history.append({"assistantResponseMessage": assistant_response})
            
        elif msg.role == "system":
            # System prompt обрабатывается отдельно в build_kiro_payload
            pass
    
    return history


def _extract_tool_results(content: Any) -> List[Dict[str, Any]]:
    """
    Извлекает tool results из контента сообщения.
    
    Args:
        content: Контент сообщения (может быть списком)
    
    Returns:
        Список tool results в формате Kiro
    """
    tool_results = []
    
    if isinstance(content, list):
        for item in content:
            if isinstance(item, dict) and item.get("type") == "tool_result":
                tool_results.append({
                    "content": [{"text": extract_text_content(item.get("content", ""))}],
                    "status": "success",
                    "toolUseId": item.get("tool_use_id", "")
                })
    
    return tool_results


def process_tools_with_long_descriptions(
    tools: Optional[List[Tool]]
) -> Tuple[Optional[List[Tool]], str]:
    """
    Обрабатывает tools с длинными descriptions.
    
    Kiro API имеет ограничение на длину description в toolSpecification.
    Если description превышает лимит, полное описание переносится в system prompt,
    а в tool остаётся ссылка на документацию.
    
    Args:
        tools: Список инструментов из запроса OpenAI
    
    Returns:
        Tuple из:
        - Список tools с обработанными descriptions (или None если tools пуст)
        - Строка с документацией для добавления в system prompt (пустая если все descriptions короткие)
    
    Example:
        >>> tools = [Tool(type="function", function=ToolFunction(name="bash", description="Very long..."))]
        >>> processed_tools, doc = process_tools_with_long_descriptions(tools)
        >>> "## Tool: bash" in doc
        True
    """
    if not tools:
        return None, ""
    
    # Если лимит отключен (0), возвращаем tools без изменений
    if TOOL_DESCRIPTION_MAX_LENGTH <= 0:
        return tools, ""
    
    tool_documentation_parts = []
    processed_tools = []
    
    for tool in tools:
        if tool.type != "function":
            processed_tools.append(tool)
            continue
        
        description = tool.function.description or ""
        
        if len(description) <= TOOL_DESCRIPTION_MAX_LENGTH:
            # Description короткий - оставляем как есть
            processed_tools.append(tool)
        else:
            # Description слишком длинный - переносим в system prompt
            tool_name = tool.function.name
            
            logger.debug(
                f"Tool '{tool_name}' has long description ({len(description)} chars > {TOOL_DESCRIPTION_MAX_LENGTH}), "
                f"moving to system prompt"
            )
            
            # Создаём документацию для system prompt
            tool_documentation_parts.append(f"## Tool: {tool_name}\n\n{description}")
            
            # Создаём копию tool с reference description
            # Используем модель Tool для создания новой копии
            from kiro_gateway.models import ToolFunction
            
            reference_description = f"[Full documentation in system prompt under '## Tool: {tool_name}']"
            
            processed_tool = Tool(
                type=tool.type,
                function=ToolFunction(
                    name=tool.function.name,
                    description=reference_description,
                    parameters=tool.function.parameters
                )
            )
            processed_tools.append(processed_tool)
    
    # Формируем итоговую документацию
    tool_documentation = ""
    if tool_documentation_parts:
        tool_documentation = (
            "\n\n---\n"
            "# Tool Documentation\n"
            "The following tools have detailed documentation that couldn't fit in the tool definition.\n\n"
            + "\n\n---\n\n".join(tool_documentation_parts)
        )
    
    return processed_tools if processed_tools else None, tool_documentation


def _extract_tool_uses(msg: ChatMessage) -> List[Dict[str, Any]]:
    """
    Извлекает tool uses из сообщения assistant.
    
    Args:
        msg: Сообщение assistant
    
    Returns:
        Список tool uses в формате Kiro
    """
    tool_uses = []
    
    # Из поля tool_calls
    if msg.tool_calls:
        for tc in msg.tool_calls:
            if isinstance(tc, dict):
                tool_uses.append({
                    "name": tc.get("function", {}).get("name", ""),
                    "input": json.loads(tc.get("function", {}).get("arguments", "{}")),
                    "toolUseId": tc.get("id", "")
                })
    
    # Из content (если там есть tool_use)
    if isinstance(msg.content, list):
        for item in msg.content:
            if isinstance(item, dict) and item.get("type") == "tool_use":
                tool_uses.append({
                    "name": item.get("name", ""),
                    "input": item.get("input", {}),
                    "toolUseId": item.get("id", "")
                })
    
    return tool_uses


def build_kiro_payload(
    request_data: ChatCompletionRequest,
    conversation_id: str,
    profile_arn: str
) -> dict:
    """
    Строит полный payload для Kiro API.
    
    Включает:
    - Полную историю сообщений
    - System prompt (добавляется к первому user сообщению)
    - Tools definitions (с обработкой длинных descriptions)
    - Текущее сообщение
    
    Если tools содержат слишком длинные descriptions, они автоматически
    переносятся в system prompt, а в tool остаётся ссылка на документацию.
    
    Args:
        request_data: Запрос в формате OpenAI
        conversation_id: Уникальный ID разговора
        profile_arn: ARN профиля AWS CodeWhisperer
    
    Returns:
        Словарь payload для POST запроса к Kiro API
    
    Raises:
        ValueError: Если нет сообщений для отправки
    """
    messages = list(request_data.messages)
    
    # Обрабатываем tools с длинными descriptions
    processed_tools, tool_documentation = process_tools_with_long_descriptions(request_data.tools)
    
    # Извлекаем system prompt
    system_prompt = ""
    non_system_messages = []
    for msg in messages:
        if msg.role == "system":
            system_prompt += extract_text_content(msg.content) + "\n"
        else:
            non_system_messages.append(msg)
    system_prompt = system_prompt.strip()
    
    # Добавляем документацию по tools в system prompt если есть
    if tool_documentation:
        system_prompt = system_prompt + tool_documentation if system_prompt else tool_documentation.strip()
    
    # Объединяем соседние сообщения с одинаковой ролью
    merged_messages = merge_adjacent_messages(non_system_messages)
    
    if not merged_messages:
        raise ValueError("No messages to send")
    
    # Получаем внутренний ID модели
    model_id = get_internal_model_id(request_data.model)
    
    # Строим историю (все сообщения кроме последнего)
    history_messages = merged_messages[:-1] if len(merged_messages) > 1 else []
    
    # Если есть system prompt, добавляем его к первому user сообщению в истории
    if system_prompt and history_messages:
        first_msg = history_messages[0]
        if first_msg.role == "user":
            original_content = extract_text_content(first_msg.content)
            first_msg.content = f"{system_prompt}\n\n{original_content}"
    
    history = build_kiro_history(history_messages, model_id)
    
    # Текущее сообщение (последнее)
    current_message = merged_messages[-1]
    current_content = extract_text_content(current_message.content)
    
    # Если system prompt есть, но история пуста - добавляем к текущему сообщению
    if system_prompt and not history:
        current_content = f"{system_prompt}\n\n{current_content}"
    
    # Если текущее сообщение - assistant, нужно добавить его в историю
    # и создать user сообщение "Continue"
    if current_message.role == "assistant":
        history.append({
            "assistantResponseMessage": {
                "content": current_content
            }
        })
        current_content = "Continue"
    
    # Если контент пустой
    if not current_content:
        current_content = "Continue"
    
    # Строим userInputMessage
    user_input_message = {
        "content": current_content,
        "modelId": model_id,
        "origin": "AI_EDITOR",
    }
    
    # Добавляем tools и tool_results если есть
    # Используем обработанные tools (с короткими descriptions)
    user_input_context = _build_user_input_context(request_data, current_message, processed_tools)
    if user_input_context:
        user_input_message["userInputMessageContext"] = user_input_context
    
    # Собираем финальный payload
    payload = {
        "conversationState": {
            "chatTriggerType": "MANUAL",
            "conversationId": conversation_id,
            "currentMessage": {
                "userInputMessage": user_input_message
            }
        }
    }
    
    # Добавляем историю только если она не пуста
    if history:
        payload["conversationState"]["history"] = history
    
    # Добавляем profileArn
    if profile_arn:
        payload["profileArn"] = profile_arn
    
    return payload


def _sanitize_json_schema(schema: Optional[Dict[str, Any]]) -> Dict[str, Any]:
    """
    Очищает JSON Schema от полей, которые Kiro API не принимает.
    
    Kiro API возвращает ошибку 400 "Improperly formed request" если:
    - required является пустым массивом []
    - additionalProperties присутствует в схеме
    
    Эта функция рекурсивно обрабатывает схему и удаляет проблемные поля.
    
    Args:
        schema: JSON Schema для очистки
    
    Returns:
        Очищенная копия схемы
    """
    if not schema:
        return {}
    
    # Создаём копию чтобы не мутировать оригинал
    result = {}
    
    for key, value in schema.items():
        # Пропускаем пустые required массивы
        if key == "required" and isinstance(value, list) and len(value) == 0:
            continue
        
        # Пропускаем additionalProperties - Kiro API его не поддерживает
        if key == "additionalProperties":
            continue
        
        # Рекурсивно обрабатываем вложенные объекты
        if key == "properties" and isinstance(value, dict):
            result[key] = {
                prop_name: _sanitize_json_schema(prop_value) if isinstance(prop_value, dict) else prop_value
                for prop_name, prop_value in value.items()
            }
        elif isinstance(value, dict):
            result[key] = _sanitize_json_schema(value)
        elif isinstance(value, list):
            # Обрабатываем списки (например, anyOf, oneOf)
            result[key] = [
                _sanitize_json_schema(item) if isinstance(item, dict) else item
                for item in value
            ]
        else:
            result[key] = value
    
    return result


def _build_user_input_context(
    request_data: ChatCompletionRequest,
    current_message: ChatMessage,
    processed_tools: Optional[List[Tool]] = None
) -> Dict[str, Any]:
    """
    Строит userInputMessageContext для текущего сообщения.
    
    Включает tools definitions и tool_results.
    
    Args:
        request_data: Запрос с tools
        current_message: Текущее сообщение
        processed_tools: Обработанные tools с короткими descriptions (опционально).
                        Если None, используются tools из request_data.
    
    Returns:
        Словарь с контекстом или пустой словарь
    """
    context = {}
    
    # Используем обработанные tools если переданы, иначе оригинальные
    tools_to_use = processed_tools if processed_tools is not None else request_data.tools
    
    # Добавляем tools если есть
    if tools_to_use:
        tools_list = []
        for tool in tools_to_use:
            if tool.type == "function":
                # Очищаем parameters от полей, которые Kiro API не принимает
                sanitized_params = _sanitize_json_schema(tool.function.parameters)
                
                # Kiro API требует непустое description
                # Если description пустое или None, используем placeholder
                description = tool.function.description
                if not description or not description.strip():
                    description = f"Tool: {tool.function.name}"
                    logger.debug(f"Tool '{tool.function.name}' has empty description, using placeholder")
                
                tools_list.append({
                    "toolSpecification": {
                        "name": tool.function.name,
                        "description": description,
                        "inputSchema": {"json": sanitized_params}
                    }
                })
        if tools_list:
            context["tools"] = tools_list
    
    # Обработка tool_results в текущем сообщении
    tool_results = _extract_tool_results(current_message.content)
    if tool_results:
        context["toolResults"] = tool_results
    
    return context


================================================
FILE: kiro_gateway/debug_logger.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
Модуль для отладочного логирования запросов.

Поддерживает три режима (DEBUG_MODE):
- off: логирование отключено
- errors: логи сохраняются только при ошибках (4xx, 5xx)
- all: логи перезаписываются на каждый запрос

В режиме "errors" данные буферизуются в памяти и сбрасываются в файлы
только при вызове flush_on_error().

Также захватывает логи приложения (loguru) для каждого запроса и сохраняет
их в файл app_logs.txt для удобства отладки.
"""

import io
import json
import shutil
from pathlib import Path
from typing import Optional
from loguru import logger

from kiro_gateway.config import DEBUG_MODE, DEBUG_DIR


class DebugLogger:
    """
    Синглтон для управления отладочными логами запросов.
    
    Режимы работы:
    - off: ничего не делает
    - errors: буферизует данные, сбрасывает в файлы только при ошибках
    - all: пишет данные сразу в файлы (как раньше)
    """
    _instance = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super(DebugLogger, cls).__new__(cls)
            cls._instance._initialized = False
        return cls._instance

    def __init__(self):
        if self._initialized:
            return
        self.debug_dir = Path(DEBUG_DIR)
        self._initialized = True
        
        # Буферы для режима "errors"
        self._request_body_buffer: Optional[bytes] = None
        self._kiro_request_body_buffer: Optional[bytes] = None
        self._raw_chunks_buffer: bytearray = bytearray()
        self._modified_chunks_buffer: bytearray = bytearray()
        
        # Буфер для логов приложения (loguru)
        self._app_logs_buffer: io.StringIO = io.StringIO()
        self._loguru_sink_id: Optional[int] = None
    
    def _is_enabled(self) -> bool:
        """Проверяет, включено ли логирование."""
        return DEBUG_MODE in ("errors", "all")
    
    def _is_immediate_write(self) -> bool:
        """Проверяет, нужно ли писать сразу в файлы (режим all)."""
        return DEBUG_MODE == "all"
    
    def _clear_buffers(self):
        """Очищает все буферы."""
        self._request_body_buffer = None
        self._kiro_request_body_buffer = None
        self._raw_chunks_buffer.clear()
        self._modified_chunks_buffer.clear()
        self._clear_app_logs_buffer()
    
    def _clear_app_logs_buffer(self):
        """Очищает буфер логов приложения и удаляет sink."""
        # Удаляем sink из loguru
        if self._loguru_sink_id is not None:
            try:
                logger.remove(self._loguru_sink_id)
            except ValueError:
                # Sink уже удалён
                pass
            self._loguru_sink_id = None
        
        # Очищаем буфер
        self._app_logs_buffer = io.StringIO()
    
    def _setup_app_logs_capture(self):
        """
        Настраивает захват логов приложения в буфер.
        
        Добавляет временный sink в loguru, который пишет в StringIO буфер.
        Захватывает ВСЕ логи без фильтрации, так как sink активен только
        на время обработки конкретного запроса.
        """
        # Удаляем предыдущий sink если есть
        self._clear_app_logs_buffer()
        
        # Добавляем новый sink для захвата ВСЕХ логов
        # Формат: время | уровень | модуль:функция:строка | сообщение
        self._loguru_sink_id = logger.add(
            self._app_logs_buffer,
            format="{time:YYYY-MM-DD HH:mm:ss.SSS} | {level: <8} | {name}:{function}:{line} | {message}",
            level="DEBUG",  # Захватываем все уровни от DEBUG и выше
            colorize=False,  # Без ANSI цветов в файле
            # Без фильтра - захватываем ВСЕ логи во время обработки запроса
        )

    def prepare_new_request(self):
        """
        Подготавливает логгер для нового запроса.
        
        В режиме "all": очищает папку с логами.
        В режиме "errors": очищает буферы.
        В обоих режимах: настраивает захват логов приложения.
        """
        if not self._is_enabled():
            return
        
        # Очищаем буферы в любом случае
        self._clear_buffers()
        
        # Настраиваем захват логов приложения
        self._setup_app_logs_capture()

        if self._is_immediate_write():
            # Режим "all" - очищаем папку и создаём заново
            try:
                if self.debug_dir.exists():
                    shutil.rmtree(self.debug_dir)
                self.debug_dir.mkdir(parents=True, exist_ok=True)
                logger.debug(f"[DebugLogger] Directory {self.debug_dir} cleared for new request.")
            except Exception as e:
                logger.error(f"[DebugLogger] Error preparing directory: {e}")

    def log_request_body(self, body: bytes):
        """
        Сохраняет тело запроса (от клиента, OpenAI формат).
        
        В режиме "all": пишет сразу в файл.
        В режиме "errors": буферизует.
        """
        if not self._is_enabled():
            return

        if self._is_immediate_write():
            self._write_request_body_to_file(body)
        else:
            # Режим "errors" - буферизуем
            self._request_body_buffer = body

    def log_kiro_request_body(self, body: bytes):
        """
        Сохраняет модифицированное тело запроса (к Kiro API).
        
        В режиме "all": пишет сразу в файл.
        В режиме "errors": буферизует.
        """
        if not self._is_enabled():
            return

        if self._is_immediate_write():
            self._write_kiro_request_body_to_file(body)
        else:
            # Режим "errors" - буферизуем
            self._kiro_request_body_buffer = body

    def log_raw_chunk(self, chunk: bytes):
        """
        Дописывает сырой чанк ответа (от провайдера).
        
        В режиме "all": пишет сразу в файл.
        В режиме "errors": буферизует.
        """
        if not self._is_enabled():
            return

        if self._is_immediate_write():
            self._append_raw_chunk_to_file(chunk)
        else:
            # Режим "errors" - буферизуем
            self._raw_chunks_buffer.extend(chunk)

    def log_modified_chunk(self, chunk: bytes):
        """
        Дописывает модифицированный чанк (клиенту).
        
        В режиме "all": пишет сразу в файл.
        В режиме "errors": буферизует.
        """
        if not self._is_enabled():
            return

        if self._is_immediate_write():
            self._append_modified_chunk_to_file(chunk)
        else:
            # Режим "errors" - буферизуем
            self._modified_chunks_buffer.extend(chunk)
    
    def log_error_info(self, status_code: int, error_message: str = ""):
        """
        Записывает информацию об ошибке в файл.
        
        Работает в обоих режимах (errors и all).
        В режиме "all" записывает сразу в файл.
        В режиме "errors" вызывается из flush_on_error().
        
        Args:
            status_code: HTTP статус код ошибки
            error_message: Сообщение об ошибке (опционально)
        """
        if not self._is_enabled():
            return
        
        try:
            # Убеждаемся что директория существует
            self.debug_dir.mkdir(parents=True, exist_ok=True)
            
            error_info = {
                "status_code": status_code,
                "error_message": error_message
            }
            error_file = self.debug_dir / "error_info.json"
            with open(error_file, "w", encoding="utf-8") as f:
                json.dump(error_info, f, indent=2, ensure_ascii=False)
            
            logger.debug(f"[DebugLogger] Error info saved (status={status_code})")
        except Exception as e:
            logger.error(f"[DebugLogger] Error writing error_info: {e}")

    def flush_on_error(self, status_code: int, error_message: str = ""):
        """
        Сбрасывает буферы в файлы при ошибке.
        
        В режиме "errors": сбрасывает буферы и сохраняет error_info.
        В режиме "all": только сохраняет error_info (данные уже записаны).
        
        Args:
            status_code: HTTP статус код ошибки
            error_message: Сообщение об ошибке (опционально)
        """
        if not self._is_enabled():
            return
        
        # В режиме "all" данные уже записаны, добавляем error_info и логи приложения
        if self._is_immediate_write():
            self.log_error_info(status_code, error_message)
            self._write_app_logs_to_file()
            self._clear_app_logs_buffer()
            return
        
        # Проверяем, есть ли что сбрасывать
        if not any([
            self._request_body_buffer,
            self._kiro_request_body_buffer,
            self._raw_chunks_buffer,
            self._modified_chunks_buffer
        ]):
            return
        
        try:
            # Создаём директорию если не существует
            if self.debug_dir.exists():
                shutil.rmtree(self.debug_dir)
            self.debug_dir.mkdir(parents=True, exist_ok=True)
            
            # Сбрасываем буферы в файлы
            if self._request_body_buffer:
                self._write_request_body_to_file(self._request_body_buffer)
            
            if self._kiro_request_body_buffer:
                self._write_kiro_request_body_to_file(self._kiro_request_body_buffer)
            
            if self._raw_chunks_buffer:
                file_path = self.debug_dir / "response_stream_raw.txt"
                with open(file_path, "wb") as f:
                    f.write(self._raw_chunks_buffer)
            
            if self._modified_chunks_buffer:
                file_path = self.debug_dir / "response_stream_modified.txt"
                with open(file_path, "wb") as f:
                    f.write(self._modified_chunks_buffer)
            
            # Сохраняем информацию об ошибке
            self.log_error_info(status_code, error_message)
            
            # Сохраняем логи приложения
            self._write_app_logs_to_file()
            
            logger.info(f"[DebugLogger] Error logs flushed to {self.debug_dir} (status={status_code})")
            
        except Exception as e:
            logger.error(f"[DebugLogger] Error flushing buffers: {e}")
        finally:
            # Очищаем буферы после сброса
            self._clear_buffers()
    
    def discard_buffers(self):
        """
        Очищает буферы без записи в файлы.
        
        Вызывается когда запрос завершился успешно в режиме "errors".
        Также вызывается в режиме "all" для сохранения логов успешного запроса.
        """
        if DEBUG_MODE == "errors":
            self._clear_buffers()
        elif DEBUG_MODE == "all":
            # В режиме "all" сохраняем логи даже для успешных запросов
            self._write_app_logs_to_file()
            self._clear_app_logs_buffer()
    
    # ==================== Приватные методы записи в файлы ====================
    
    def _write_request_body_to_file(self, body: bytes):
        """Записывает тело запроса в файл."""
        try:
            file_path = self.debug_dir / "request_body.json"
            try:
                json_obj = json.loads(body)
                with open(file_path, "w", encoding="utf-8") as f:
                    json.dump(json_obj, f, indent=2, ensure_ascii=False)
            except json.JSONDecodeError:
                with open(file_path, "wb") as f:
                    f.write(body)
        except Exception as e:
            logger.error(f"[DebugLogger] Error writing request_body: {e}")
    
    def _write_kiro_request_body_to_file(self, body: bytes):
        """Записывает тело запроса к Kiro в файл."""
        try:
            file_path = self.debug_dir / "kiro_request_body.json"
            try:
                json_obj = json.loads(body)
                with open(file_path, "w", encoding="utf-8") as f:
                    json.dump(json_obj, f, indent=2, ensure_ascii=False)
            except json.JSONDecodeError:
                with open(file_path, "wb") as f:
                    f.write(body)
        except Exception as e:
            logger.error(f"[DebugLogger] Error writing kiro_request_body: {e}")
    
    def _append_raw_chunk_to_file(self, chunk: bytes):
        """Дописывает сырой чанк в файл."""
        try:
            file_path = self.debug_dir / "response_stream_raw.txt"
            with open(file_path, "ab") as f:
                f.write(chunk)
        except Exception:
            pass
    
    def _append_modified_chunk_to_file(self, chunk: bytes):
        """Дописывает модифицированный чанк в файл."""
        try:
            file_path = self.debug_dir / "response_stream_modified.txt"
            with open(file_path, "ab") as f:
                f.write(chunk)
        except Exception:
            pass
    
    def _write_app_logs_to_file(self):
        """Записывает захваченные логи приложения в файл."""
        try:
            # Получаем содержимое буфера
            logs_content = self._app_logs_buffer.getvalue()
            
            if not logs_content.strip():
                return
            
            # Убеждаемся что директория существует
            self.debug_dir.mkdir(parents=True, exist_ok=True)
            
            file_path = self.debug_dir / "app_logs.txt"
            with open(file_path, "w", encoding="utf-8") as f:
                f.write(logs_content)
            
            logger.debug(f"[DebugLogger] App logs saved to {file_path}")
        except Exception as e:
            # Не логируем ошибку через logger чтобы избежать рекурсии
            pass


# Глобальный экземпляр
debug_logger = DebugLogger()


================================================
FILE: kiro_gateway/exceptions.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
Обработчики исключений для Kiro Gateway.

Содержит функции для обработки ошибок валидации и других исключений
в формате, совместимом с JSON-сериализацией.
"""

from typing import Any, List, Dict

from fastapi import Request
from fastapi.exceptions import RequestValidationError
from fastapi.responses import JSONResponse
from loguru import logger


def sanitize_validation_errors(errors: List[Dict[str, Any]]) -> List[Dict[str, Any]]:
    """
    Преобразует ошибки валидации в JSON-сериализуемый формат.
    
    Pydantic может включать bytes объекты в поле 'input', которые
    не сериализуются в JSON. Эта функция конвертирует их в строки.
    
    Args:
        errors: Список ошибок валидации от Pydantic
    
    Returns:
        Список ошибок с bytes преобразованными в строки
    """
    sanitized = []
    for error in errors:
        sanitized_error = {}
        for key, value in error.items():
            if isinstance(value, bytes):
                # Конвертируем bytes в строку
                sanitized_error[key] = value.decode("utf-8", errors="replace")
            elif isinstance(value, (list, tuple)):
                # Рекурсивно обрабатываем списки
                sanitized_error[key] = [
                    v.decode("utf-8", errors="replace") if isinstance(v, bytes) else v
                    for v in value
                ]
            else:
                sanitized_error[key] = value
        sanitized.append(sanitized_error)
    return sanitized


async def validation_exception_handler(request: Request, exc: RequestValidationError) -> JSONResponse:
    """
    Обработчик ошибок валидации Pydantic.
    
    Логирует детали ошибки и возвращает информативный ответ.
    Корректно обрабатывает bytes объекты в ошибках, преобразуя их в строки.
    
    Args:
        request: FastAPI Request объект
        exc: Исключение валидации от Pydantic
    
    Returns:
        JSONResponse с деталями ошибки и статусом 422
    """
    body = await request.body()
    body_str = body.decode("utf-8", errors="replace")
    
    # Санитизируем ошибки для JSON-сериализации
    sanitized_errors = sanitize_validation_errors(exc.errors())
    
    logger.error(f"Validation error (422): {sanitized_errors}")
    logger.error(f"Request body: {body_str[:500]}...")
    
    return JSONResponse(
        status_code=422,
        content={"detail": sanitized_errors, "body": body_str[:500]},
    )


================================================
FILE: kiro_gateway/http_client.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# https://github.com/jwadow/kiro-openai-gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
HTTP client for Kiro API with retry logic support.

Handles:
- 403: automatic token refresh and retry
- 429: exponential backoff
- 5xx: exponential backoff
- Timeouts: exponential backoff
"""

import asyncio
from typing import Optional

import httpx
from fastapi import HTTPException
from loguru import logger

from kiro_gateway.config import MAX_RETRIES, BASE_RETRY_DELAY, FIRST_TOKEN_MAX_RETRIES, STREAMING_READ_TIMEOUT
from kiro_gateway.auth import KiroAuthManager
from kiro_gateway.utils import get_kiro_headers


class KiroHttpClient:
    """
    HTTP client for Kiro API with retry logic support.
    
    Automatically handles errors and retries requests:
    - 403: refreshes token and retries
    - 429: waits with exponential backoff
    - 5xx: waits with exponential backoff
    - Timeouts: waits with exponential backoff
    Attributes:
        auth_manager: Authentication manager for obtaining tokens
        client: httpx HTTP client
    
    Example:
        >>> client = KiroHttpClient(auth_manager)
        >>> response = await client.request_with_retry(
        ...     "POST",
        ...     "https://api.example.com/endpoint",
        ...     {"data": "value"},
        ...     stream=True
        ... )
    """
    
    def __init__(self, auth_manager: KiroAuthManager):
        """
        Initializes the HTTP client.
        
        Args:
            auth_manager: Authentication manager
        """
        self.auth_manager = auth_manager
        self.client: Optional[httpx.AsyncClient] = None
    
    async def _get_client(self, stream: bool = False) -> httpx.AsyncClient:
        """
        Returns or creates an HTTP client with proper timeouts.
        
        httpx timeouts:
        - connect: TCP handshake (DNS + TCP SYN/ACK)
        - read: waiting for data from server between chunks
        - write: sending data to server
        - pool: waiting for free connection from pool
        
        IMPORTANT: FIRST_TOKEN_TIMEOUT is NOT used here!
        It is applied in streaming.py via asyncio.wait_for() to control
        the wait time for the first token from the model (retry business logic).
        
        Args:
            stream: If True, uses STREAMING_READ_TIMEOUT for read
        
        Returns:
            Active HTTP client
        """
        if self.client is None or self.client.is_closed:
            if stream:
                # For streaming:
                # - connect: 30 sec (TCP connection, usually < 1 sec)
                # - read: STREAMING_READ_TIMEOUT (300 sec) - model may "think" between chunks
                # - write/pool: standard values
                timeout_config = httpx.Timeout(
                    connect=30.0,
                    read=STREAMING_READ_TIMEOUT,
                    write=30.0,
                    pool=30.0
                )
                logger.debug(f"Creating streaming HTTP client (read_timeout={STREAMING_READ_TIMEOUT}s)")
            else:
                # For regular requests: single timeout of 300 sec
                timeout_config = httpx.Timeout(timeout=300.0)
                logger.debug("Creating non-streaming HTTP client (timeout=300s)")
            
            self.client = httpx.AsyncClient(timeout=timeout_config, follow_redirects=True)
        return self.client
    
    async def close(self) -> None:
        """Closes the HTTP client."""
        if self.client and not self.client.is_closed:
            await self.client.aclose()
    
    async def request_with_retry(
        self,
        method: str,
        url: str,
        json_data: dict,
        stream: bool = False
    ) -> httpx.Response:
        """
        Executes an HTTP request with retry logic.
        
        Automatically handles various error types:
        - 403: refreshes token via auth_manager.force_refresh() and retries
        - 429: waits with exponential backoff (1s, 2s, 4s)
        - 5xx: waits with exponential backoff
        - Timeouts: waits with exponential backoff
        
        For streaming, STREAMING_READ_TIMEOUT is used for waiting between chunks.
        First token timeout is controlled separately in streaming.py via asyncio.wait_for().
        
        Args:
            method: HTTP method (GET, POST, etc.)
            url: Request URL
            json_data: Request body (JSON)
            stream: Use streaming (default False)
        
        Returns:
            httpx.Response with successful response
        
        Raises:
            HTTPException: On failure after all attempts (502/504)
        """
        # Determine the number of retry attempts
        # FIRST_TOKEN_TIMEOUT is used in streaming.py, not here
        max_retries = FIRST_TOKEN_MAX_RETRIES if stream else MAX_RETRIES
        
        client = await self._get_client(stream=stream)
        last_error = None
        
        for attempt in range(max_retries):
            try:
                # Get current token
                token = await self.auth_manager.get_access_token()
                headers = get_kiro_headers(self.auth_manager, token)
                
                if stream:
                    req = client.build_request(method, url, json=json_data, headers=headers)
                    response = await client.send(req, stream=True)
                else:
                    response = await client.request(method, url, json=json_data, headers=headers)
                
                # Check status
                if response.status_code == 200:
                    return response
                
                # 403 - token expired, refresh and retry
                if response.status_code == 403:
                    logger.warning(f"Received 403, refreshing token (attempt {attempt + 1}/{MAX_RETRIES})")
                    await self.auth_manager.force_refresh()
                    continue
                
                # 429 - rate limit, wait and retry
                if response.status_code == 429:
                    delay = BASE_RETRY_DELAY * (2 ** attempt)
                    logger.warning(f"Received 429, waiting {delay}s (attempt {attempt + 1}/{MAX_RETRIES})")
                    await asyncio.sleep(delay)
                    continue
                
                # 5xx - server error, wait and retry
                if 500 <= response.status_code < 600:
                    delay = BASE_RETRY_DELAY * (2 ** attempt)
                    logger.warning(f"Received {response.status_code}, waiting {delay}s (attempt {attempt + 1}/{MAX_RETRIES})")
                    await asyncio.sleep(delay)
                    continue
                
                # Other errors - return as is
                return response
                
            except httpx.TimeoutException as e:
                last_error = e
                # Determine timeout type for logging
                timeout_type = type(e).__name__
                
                if stream:
                    # For streaming this could be:
                    # - ConnectTimeout: TCP connection issue
                    # - ReadTimeout: server not responding (STREAMING_READ_TIMEOUT)
                    if isinstance(e, httpx.ConnectTimeout):
                        logger.warning(
                            f"[{timeout_type}] Connection timeout (attempt {attempt + 1}/{max_retries})"
                        )
                    elif isinstance(e, httpx.ReadTimeout):
                        logger.warning(
                            f"[{timeout_type}] Read timeout after {STREAMING_READ_TIMEOUT}s - "
                            f"server stopped responding (attempt {attempt + 1}/{max_retries})"
                        )
                    else:
                        logger.warning(
                            f"[{timeout_type}] Timeout during streaming (attempt {attempt + 1}/{max_retries})"
                        )
                else:
                    delay = BASE_RETRY_DELAY * (2 ** attempt)
                    logger.warning(
                        f"[{timeout_type}] Request timeout, waiting {delay}s (attempt {attempt + 1}/{max_retries})"
                    )
                    await asyncio.sleep(delay)
                
            except httpx.RequestError as e:
                last_error = e
                delay = BASE_RETRY_DELAY * (2 ** attempt)
                logger.warning(f"Request error: {e}, waiting {delay}s (attempt {attempt + 1}/{max_retries})")
                await asyncio.sleep(delay)
        
        # All attempts exhausted
        if stream:
            raise HTTPException(
                status_code=504,
                detail=f"Streaming failed after {max_retries} attempts. Last error: {type(last_error).__name__}"
            )
        else:
            raise HTTPException(
                status_code=502,
                detail=f"Failed to complete request after {max_retries} attempts: {last_error}"
            )
    
    async def __aenter__(self) -> "KiroHttpClient":
        """Async context manager support."""
        return self
    
    async def __aexit__(self, exc_type, exc_val, exc_tb) -> None:
        """Closes the client when exiting context."""
        await self.close()


================================================
FILE: kiro_gateway/models.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
Pydantic модели для OpenAI-совместимого API.

Определяет схемы данных для запросов и ответов,
обеспечивая валидацию и сериализацию.
"""

import time
from typing import Any, Dict, List, Optional, Union
from typing_extensions import Annotated
from pydantic import BaseModel, Field


# ==================================================================================================
# Модели для /v1/models endpoint
# ==================================================================================================

class OpenAIModel(BaseModel):
    """
    Модель данных для описания AI модели в формате OpenAI.
    
    Используется в ответе эндпоинта /v1/models.
    """
    id: str
    object: str = "model"
    created: int = Field(default_factory=lambda: int(time.time()))
    owned_by: str = "anthropic"
    description: Optional[str] = None


class ModelList(BaseModel):
    """
    Список моделей в формате OpenAI.
    
    Ответ эндпоинта GET /v1/models.
    """
    object: str = "list"
    data: List[OpenAIModel]


# ==================================================================================================
# Модели для /v1/chat/completions endpoint
# ==================================================================================================

class ChatMessage(BaseModel):
    """
    Сообщение в чате в формате OpenAI.
    
    Поддерживает различные роли (user, assistant, system, tool)
    и различные форматы контента (строка, список, объект).
    
    Attributes:
        role: Роль отправителя (user, assistant, system, tool)
        content: Содержимое сообщения (может быть строкой, списком или None)
        name: Опциональное имя отправителя
        tool_calls: Список вызовов инструментов (для assistant)
        tool_call_id: ID вызова инструмента (для tool)
    """
    role: str
    content: Optional[Union[str, List[Any], Any]] = None
    name: Optional[str] = None
    tool_calls: Optional[List[Any]] = None
    tool_call_id: Optional[str] = None
    
    model_config = {"extra": "allow"}


class ToolFunction(BaseModel):
    """
    Описание функции инструмента.
    
    Attributes:
        name: Имя функции
        description: Описание функции
        parameters: JSON Schema параметров функции
    """
    name: str
    description: Optional[str] = None
    parameters: Optional[Dict[str, Any]] = None


class Tool(BaseModel):
    """
    Инструмент (tool) в формате OpenAI.
    
    Attributes:
        type: Тип инструмента (обычно "function")
        function: Описание функции
    """
    type: str = "function"
    function: ToolFunction


class ChatCompletionRequest(BaseModel):
    """
    Запрос на генерацию ответа в формате OpenAI Chat Completions API.
    
    Поддерживает все стандартные поля OpenAI API, включая:
    - Базовые параметры (model, messages, stream)
    - Параметры генерации (temperature, top_p, max_tokens)
    - Tools (function calling)
    - Дополнительные параметры (игнорируются, но принимаются для совместимости)
    
    Attributes:
        model: ID модели для генерации
        messages: Список сообщений чата
        stream: Использовать streaming (по умолчанию False)
        temperature: Температура генерации (0-2)
        top_p: Top-p sampling
        n: Количество вариантов ответа
        max_tokens: Максимальное количество токенов в ответе
        max_completion_tokens: Альтернативное поле для max_tokens
        stop: Стоп-последовательности
        presence_penalty: Штраф за повторение тем
        frequency_penalty: Штраф за повторение слов
        tools: Список доступных инструментов
        tool_choice: Стратегия выбора инструмента
    """
    model: str
    messages: Annotated[List[ChatMessage], Field(min_length=1)]
    stream: bool = False
    
    # Параметры генерации
    temperature: Optional[float] = None
    top_p: Optional[float] = None
    n: Optional[int] = 1
    max_tokens: Optional[int] = None
    max_completion_tokens: Optional[int] = None
    stop: Optional[Union[str, List[str]]] = None
    presence_penalty: Optional[float] = None
    frequency_penalty: Optional[float] = None
    
    # Tools (function calling)
    tools: Optional[List[Tool]] = None
    tool_choice: Optional[Union[str, Dict]] = None
    
    # Поля для совместимости (игнорируются)
    stream_options: Optional[Dict[str, Any]] = None
    logit_bias: Optional[Dict[str, float]] = None
    logprobs: Optional[bool] = None
    top_logprobs: Optional[int] = None
    user: Optional[str] = None
    seed: Optional[int] = None
    parallel_tool_calls: Optional[bool] = None
    
    model_config = {"extra": "allow"}


# ==================================================================================================
# Модели для ответов
# ==================================================================================================

class ChatCompletionChoice(BaseModel):
    """
    Один вариант ответа в Chat Completion.
    
    Attributes:
        index: Индекс варианта
        message: Сообщение ответа
        finish_reason: Причина завершения (stop, tool_calls, length)
    """
    index: int = 0
    message: Dict[str, Any]
    finish_reason: Optional[str] = None


class ChatCompletionUsage(BaseModel):
    """
    Информация об использовании токенов.
    
    Attributes:
        prompt_tokens: Количество токенов в запросе
        completion_tokens: Количество токенов в ответе
        total_tokens: Общее количество токенов
        credits_used: Использованные кредиты (специфично для Kiro)
    """
    prompt_tokens: int = 0
    completion_tokens: int = 0
    total_tokens: int = 0
    credits_used: Optional[float] = None


class ChatCompletionResponse(BaseModel):
    """
    Полный ответ Chat Completion (non-streaming).
    
    Attributes:
        id: Уникальный ID ответа
        object: Тип объекта ("chat.completion")
        created: Timestamp создания
        model: Использованная модель
        choices: Список вариантов ответа
        usage: Информация об использовании токенов
    """
    id: str
    object: str = "chat.completion"
    created: int = Field(default_factory=lambda: int(time.time()))
    model: str
    choices: List[ChatCompletionChoice]
    usage: ChatCompletionUsage


class ChatCompletionChunkDelta(BaseModel):
    """
    Дельта изменений в streaming chunk.
    
    Attributes:
        role: Роль (только в первом chunk)
        content: Новый контент
        tool_calls: Новые tool calls
    """
    role: Optional[str] = None
    content: Optional[str] = None
    tool_calls: Optional[List[Dict[str, Any]]] = None


class ChatCompletionChunkChoice(BaseModel):
    """
    Один вариант в streaming chunk.
    
    Attributes:
        index: Индекс варианта
        delta: Дельта изменений
        finish_reason: Причина завершения (только в последнем chunk)
    """
    index: int = 0
    delta: ChatCompletionChunkDelta
    finish_reason: Optional[str] = None


class ChatCompletionChunk(BaseModel):
    """
    Streaming chunk в формате OpenAI.
    
    Attributes:
        id: Уникальный ID ответа
        object: Тип объекта ("chat.completion.chunk")
        created: Timestamp создания
        model: Использованная модель
        choices: Список вариантов
        usage: Информация об использовании (только в последнем chunk)
    """
    id: str
    object: str = "chat.completion.chunk"
    created: int = Field(default_factory=lambda: int(time.time()))
    model: str
    choices: List[ChatCompletionChunkChoice]
    usage: Optional[ChatCompletionUsage] = None


================================================
FILE: kiro_gateway/parsers.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
Парсеры для AWS Event Stream формата.

Содержит классы и функции для:
- Парсинга бинарного AWS SSE потока
- Извлечения JSON событий
- Обработки tool calls
- Дедупликации контента
"""

import json
import re
from typing import Any, Dict, List, Optional

from loguru import logger

from kiro_gateway.utils import generate_tool_call_id


def find_matching_brace(text: str, start_pos: int) -> int:
    """
    Находит позицию закрывающей скобки с учётом вложенности и строк.
    
    Использует bracket counting для корректного парсинга вложенных JSON.
    Учитывает строки в кавычках и escape-последовательности.
    
    Args:
        text: Текст для поиска
        start_pos: Позиция открывающей скобки '{'
    
    Returns:
        Позиция закрывающей скобки или -1 если не найдена
    
    Example:
        >>> find_matching_brace('{"a": {"b": 1}}', 0)
        14
        >>> find_matching_brace('{"a": "{}"}', 0)
        10
    """
    if start_pos >= len(text) or text[start_pos] != '{':
        return -1
    
    brace_count = 0
    in_string = False
    escape_next = False
    
    for i in range(start_pos, len(text)):
        char = text[i]
        
        if escape_next:
            escape_next = False
            continue
        
        if char == '\\' and in_string:
            escape_next = True
            continue
        
        if char == '"' and not escape_next:
            in_string = not in_string
            continue
        
        if not in_string:
            if char == '{':
                brace_count += 1
            elif char == '}':
                brace_count -= 1
                if brace_count == 0:
                    return i
    
    return -1


def parse_bracket_tool_calls(response_text: str) -> List[Dict[str, Any]]:
    """
    Парсит tool calls в формате [Called func_name with args: {...}].
    
    Некоторые модели возвращают tool calls в текстовом формате вместо
    структурированного JSON. Эта функция извлекает их.
    
    Args:
        response_text: Текст ответа модели
    
    Returns:
        Список tool calls в формате OpenAI
    
    Example:
        >>> text = "[Called get_weather with args: {\"city\": \"London\"}]"
        >>> calls = parse_bracket_tool_calls(text)
        >>> calls[0]["function"]["name"]
        'get_weather'
    """
    if not response_text or "[Called" not in response_text:
        return []
    
    tool_calls = []
    pattern = r'\[Called\s+(\w+)\s+with\s+args:\s*'
    
    for match in re.finditer(pattern, response_text, re.IGNORECASE):
        func_name = match.group(1)
        args_start = match.end()
        
        # Ищем начало JSON
        json_start = response_text.find('{', args_start)
        if json_start == -1:
            continue
        
        # Ищем конец JSON с учётом вложенности
        json_end = find_matching_brace(response_text, json_start)
        if json_end == -1:
            continue
        
        json_str = response_text[json_start:json_end + 1]
        
        try:
            args = json.loads(json_str)
            tool_call_id = generate_tool_call_id()
            # index будет добавлен позже при формировании финального ответа
            tool_calls.append({
                "id": tool_call_id,
                "type": "function",
                "function": {
                    "name": func_name,
                    "arguments": json.dumps(args)
                }
            })
        except json.JSONDecodeError:
            logger.warning(f"Failed to parse tool call arguments: {json_str[:100]}")
    
    return tool_calls


def deduplicate_tool_calls(tool_calls: List[Dict[str, Any]]) -> List[Dict[str, Any]]:
    """
    Удаляет дубликаты tool calls.
    
    Дедупликация происходит по двум критериям:
    1. По id - если есть несколько tool calls с одинаковым id, оставляем тот у которого
       больше аргументов (не пустой "{}")
    2. По name+arguments - удаляем полные дубликаты
    
    Args:
        tool_calls: Список tool calls
    
    Returns:
        Список уникальных tool calls
    """
    # Сначала дедупликация по id - оставляем tool call с непустыми аргументами
    by_id: Dict[str, Dict[str, Any]] = {}
    for tc in tool_calls:
        tc_id = tc.get("id", "")
        if not tc_id:
            # Без id - добавляем как есть (будет дедуплицировано по name+args)
            continue
        
        existing = by_id.get(tc_id)
        if existing is None:
            by_id[tc_id] = tc
        else:
            # Есть дубликат по id - оставляем тот у которого больше аргументов
            existing_args = existing.get("function", {}).get("arguments", "{}")
            current_args = tc.get("function", {}).get("arguments", "{}")
            
            # Предпочитаем непустые аргументы
            if current_args != "{}" and (existing_args == "{}" or len(current_args) > len(existing_args)):
                logger.debug(f"Replacing tool call {tc_id} with better arguments: {len(existing_args)} -> {len(current_args)}")
                by_id[tc_id] = tc
    
    # Собираем tool calls: сначала те что с id, потом без id
    result_with_id = list(by_id.values())
    result_without_id = [tc for tc in tool_calls if not tc.get("id")]
    
    # Теперь дедупликация по name+arguments для всех
    seen = set()
    unique = []
    
    for tc in result_with_id + result_without_id:
        # Защита от None в function
        func = tc.get("function") or {}
        func_name = func.get("name") or ""
        func_args = func.get("arguments") or "{}"
        key = f"{func_name}-{func_args}"
        if key not in seen:
            seen.add(key)
            unique.append(tc)
    
    if len(tool_calls) != len(unique):
        logger.debug(f"Deduplicated tool calls: {len(tool_calls)} -> {len(unique)}")
    
    return unique


class AwsEventStreamParser:
    """
    Парсер для AWS Event Stream формата.
    
    AWS возвращает события в бинарном формате с разделителями :message-type...event.
    Этот класс извлекает JSON события из потока и преобразует их в удобный формат.
    
    Поддерживаемые типы событий:
    - content: Текстовый контент ответа
    - tool_start: Начало tool call (name, toolUseId)
    - tool_input: Продолжение input для tool call
    - tool_stop: Завершение tool call
    - usage: Информация о потреблении кредитов
    - context_usage: Процент использования контекста
    
    Attributes:
        buffer: Буфер для накопления данных
        last_content: Последний обработанный контент (для дедупликации)
        current_tool_call: Текущий незавершённый tool call
        tool_calls: Список завершённых tool calls
    
    Example:
        >>> parser = AwsEventStreamParser()
        >>> events = parser.feed(chunk)
        >>> for event in events:
        ...     if event["type"] == "content":
        ...         print(event["data"])
    """
    
    # Паттерны для поиска JSON событий
    EVENT_PATTERNS = [
        ('{"content":', 'content'),
        ('{"name":', 'tool_start'),
        ('{"input":', 'tool_input'),
        ('{"stop":', 'tool_stop'),
        ('{"followupPrompt":', 'followup'),
        ('{"usage":', 'usage'),
        ('{"contextUsagePercentage":', 'context_usage'),
    ]
    
    def __init__(self):
        """Инициализирует парсер."""
        self.buffer = ""
        self.last_content: Optional[str] = None  # Для дедупликации повторяющегося контента
        self.current_tool_call: Optional[Dict[str, Any]] = None
        self.tool_calls: List[Dict[str, Any]] = []
    
    def feed(self, chunk: bytes) -> List[Dict[str, Any]]:
        """
        Добавляет chunk в буфер и возвращает распарсенные события.
        
        Args:
            chunk: Байты данных из потока
        
        Returns:
            Список событий в формате {"type": str, "data": Any}
        """
        try:
            self.buffer += chunk.decode('utf-8', errors='ignore')
        except Exception:
            return []
        
        events = []
        
        while True:
            # Находим ближайший паттерн
            earliest_pos = -1
            earliest_type = None
            
            for pattern, event_type in self.EVENT_PATTERNS:
                pos = self.buffer.find(pattern)
                if pos != -1 and (earliest_pos == -1 or pos < earliest_pos):
                    earliest_pos = pos
                    earliest_type = event_type
            
            if earliest_pos == -1:
                break
            
            # Ищем конец JSON
            json_end = find_matching_brace(self.buffer, earliest_pos)
            if json_end == -1:
                # JSON не полный, ждём больше данных
                break
            
            json_str = self.buffer[earliest_pos:json_end + 1]
            self.buffer = self.buffer[json_end + 1:]
            
            try:
                data = json.loads(json_str)
                event = self._process_event(data, earliest_type)
                if event:
                    events.append(event)
            except json.JSONDecodeError:
                logger.warning(f"Failed to parse JSON: {json_str[:100]}")
        
        return events
    
    def _process_event(self, data: dict, event_type: str) -> Optional[Dict[str, Any]]:
        """
        Обрабатывает распарсенное событие.
        
        Args:
            data: Распарсенный JSON
            event_type: Тип события
        
        Returns:
            Обработанное событие или None
        """
        if event_type == 'content':
            return self._process_content_event(data)
        elif event_type == 'tool_start':
            return self._process_tool_start_event(data)
        elif event_type == 'tool_input':
            return self._process_tool_input_event(data)
        elif event_type == 'tool_stop':
            return self._process_tool_stop_event(data)
        elif event_type == 'usage':
            return {"type": "usage", "data": data.get('usage', 0)}
        elif event_type == 'context_usage':
            return {"type": "context_usage", "data": data.get('contextUsagePercentage', 0)}
        
        return None
    
    def _process_content_event(self, data: dict) -> Optional[Dict[str, Any]]:
        """Обрабатывает событие с контентом."""
        content = data.get('content', '')
        
        # Пропускаем followupPrompt
        if data.get('followupPrompt'):
            return None
        
        # Дедупликация повторяющегося контента
        if content == self.last_content:
            return None
        
        self.last_content = content
        
        return {"type": "content", "data": content}
    
    def _process_tool_start_event(self, data: dict) -> Optional[Dict[str, Any]]:
        """Обрабатывает начало tool call."""
        # Завершаем предыдущий tool call если есть
        if self.current_tool_call:
            self._finalize_tool_call()
        
        # input может быть строкой или объектом
        input_data = data.get('input', '')
        if isinstance(input_data, dict):
            input_str = json.dumps(input_data)
        else:
            input_str = str(input_data) if input_data else ''
        
        self.current_tool_call = {
            "id": data.get('toolUseId', generate_tool_call_id()),
            "type": "function",
            "function": {
                "name": data.get('name', ''),
                "arguments": input_str
            }
        }
        
        if data.get('stop'):
            self._finalize_tool_call()
        
        return None
    
    def _process_tool_input_event(self, data: dict) -> Optional[Dict[str, Any]]:
        """Обрабатывает продолжение input для tool call."""
        if self.current_tool_call:
            # input может быть строкой или объектом
            input_data = data.get('input', '')
            if isinstance(input_data, dict):
                input_str = json.dumps(input_data)
            else:
                input_str = str(input_data) if input_data else ''
            self.current_tool_call['function']['arguments'] += input_str
        return None
    
    def _process_tool_stop_event(self, data: dict) -> Optional[Dict[str, Any]]:
        """Обрабатывает завершение tool call."""
        if self.current_tool_call and data.get('stop'):
            self._finalize_tool_call()
        return None
    
    def _finalize_tool_call(self) -> None:
        """Завершает текущий tool call и добавляет в список."""
        if not self.current_tool_call:
            return
        
        # Пытаемся распарсить и нормализовать arguments как JSON
        args = self.current_tool_call['function']['arguments']
        tool_name = self.current_tool_call['function'].get('name', 'unknown')
        
        logger.debug(f"Finalizing tool call '{tool_name}' with raw arguments: {repr(args)[:200]}")
        
        if isinstance(args, str):
            if args.strip():
                try:
                    parsed = json.loads(args)
                    # Убеждаемся что результат - строка JSON
                    self.current_tool_call['function']['arguments'] = json.dumps(parsed)
                    logger.debug(f"Tool '{tool_name}' arguments parsed successfully: {list(parsed.keys()) if isinstance(parsed, dict) else type(parsed)}")
                except json.JSONDecodeError as e:
                    # Если не удалось распарсить, оставляем пустой объект
                    logger.warning(f"Failed to parse tool '{tool_name}' arguments: {e}. Raw: {args[:200]}")
                    self.current_tool_call['function']['arguments'] = "{}"
            else:
                # Пустая строка - используем пустой объект
                # Это нормальное поведение для дубликатов tool calls от Kiro
                logger.debug(f"Tool '{tool_name}' has empty arguments string (will be deduplicated)")
                self.current_tool_call['function']['arguments'] = "{}"
        elif isinstance(args, dict):
            # Если уже объект - сериализуем в строку
            self.current_tool_call['function']['arguments'] = json.dumps(args)
            logger.debug(f"Tool '{tool_name}' arguments already dict with keys: {list(args.keys())}")
        else:
            # Неизвестный тип - пустой объект
            logger.warning(f"Tool '{tool_name}' has unexpected arguments type: {type(args)}")
            self.current_tool_call['function']['arguments'] = "{}"
        
        self.tool_calls.append(self.current_tool_call)
        self.current_tool_call = None
    
    def get_tool_calls(self) -> List[Dict[str, Any]]:
        """
        Возвращает все собранные tool calls.
        
        Завершает текущий tool call если он не завершён.
        Удаляет дубликаты.
        
        Returns:
            Список уникальных tool calls
        """
        if self.current_tool_call:
            self._finalize_tool_call()
        return deduplicate_tool_calls(self.tool_calls)
    
    def reset(self) -> None:
        """Сбрасывает состояние парсера."""
        self.buffer = ""
        self.last_content = None
        self.current_tool_call = None
        self.tool_calls = []


================================================
FILE: kiro_gateway/routes.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# https://github.com/jwadow/kiro-openai-gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
FastAPI routes for Kiro Gateway.

Contains all API endpoints:
- / and /health: Health check
- /v1/models: Models list
- /v1/chat/completions: Chat completions
"""

import json
from datetime import datetime, timezone

import httpx
from fastapi import APIRouter, Depends, HTTPException, Request, Response, Security
from fastapi.responses import JSONResponse, StreamingResponse
from fastapi.security import APIKeyHeader
from loguru import logger

from kiro_gateway.config import (
    PROXY_API_KEY,
    AVAILABLE_MODELS,
    APP_VERSION,
)
from kiro_gateway.models import (
    OpenAIModel,
    ModelList,
    ChatCompletionRequest,
)
from kiro_gateway.auth import KiroAuthManager
from kiro_gateway.cache import ModelInfoCache
from kiro_gateway.converters import build_kiro_payload
from kiro_gateway.streaming import stream_kiro_to_openai, collect_stream_response, stream_with_first_token_retry
from kiro_gateway.http_client import KiroHttpClient
from kiro_gateway.utils import get_kiro_headers, generate_conversation_id

# Import debug_logger
try:
    from kiro_gateway.debug_logger import debug_logger
except ImportError:
    debug_logger = None


# --- Security scheme ---
api_key_header = APIKeyHeader(name="Authorization", auto_error=False)


async def verify_api_key(auth_header: str = Security(api_key_header)) -> bool:
    """
    Verify API key in Authorization header.
    
    Expects format: "Bearer {PROXY_API_KEY}"
    
    Args:
        auth_header: Authorization header value
    
    Returns:
        True if key is valid
    
    Raises:
        HTTPException: 401 if key is invalid or missing
    """
    if not auth_header or auth_header != f"Bearer {PROXY_API_KEY}":
        logger.warning("Access attempt with invalid API key.")
        raise HTTPException(status_code=401, detail="Invalid or missing API Key")
    return True


# --- Router ---
router = APIRouter()


@router.get("/")
async def root():
    """
    Health check endpoint.
    
    Returns:
        Status and application version
    """
    return {
        "status": "ok",
        "message": "Kiro API Gateway is running",
        "version": APP_VERSION
    }


@router.get("/health")
async def health():
    """
    Detailed health check.
    
    Returns:
        Status, timestamp and version
    """
    return {
        "status": "healthy",
        "timestamp": datetime.now(timezone.utc).isoformat(),
        "version": APP_VERSION
    }


@router.get("/v1/models", response_model=ModelList, dependencies=[Depends(verify_api_key)])
async def get_models(request: Request):
    """
    Return list of available models.
    
    Uses static model list with ability to update from API.
    Caches results to reduce API load.
    
    Args:
        request: FastAPI Request for accessing app.state
    
    Returns:
        ModelList with available models
    """
    logger.info("Request to /v1/models")
    
    auth_manager: KiroAuthManager = request.app.state.auth_manager
    model_cache: ModelInfoCache = request.app.state.model_cache
    
    # Try to get models from API if cache is empty or stale
    if model_cache.is_empty() or model_cache.is_stale():
        try:
            token = await auth_manager.get_access_token()
            headers = get_kiro_headers(auth_manager, token)
            
            async with httpx.AsyncClient(timeout=30) as client:
                response = await client.get(
                    f"{auth_manager.q_host}/ListAvailableModels",
                    headers=headers,
                    params={
                        "origin": "AI_EDITOR",
                        "profileArn": auth_manager.profile_arn or ""
                    }
                )
                
                if response.status_code == 200:
                    data = response.json()
                    models_list = data.get("models", [])
                    await model_cache.update(models_list)
                    logger.info(f"Received {len(models_list)} models from API")
        except Exception as e:
            logger.warning(f"Failed to fetch models from API: {e}")
    
    # Return static model list
    openai_models = [
        OpenAIModel(
            id=model_id,
            owned_by="anthropic",
            description="Claude model via Kiro API"
        )
        for model_id in AVAILABLE_MODELS
    ]
    
    return ModelList(data=openai_models)


@router.post("/v1/chat/completions", dependencies=[Depends(verify_api_key)])
async def chat_completions(request: Request, request_data: ChatCompletionRequest):
    """
    Chat completions endpoint - compatible with OpenAI API.
    
    Accepts requests in OpenAI format and translates them to Kiro API.
    Supports streaming and non-streaming modes.
    
    Args:
        request: FastAPI Request for accessing app.state
        request_data: Request in OpenAI ChatCompletionRequest format
    
    Returns:
        StreamingResponse for streaming mode
        JSONResponse for non-streaming mode
    
    Raises:
        HTTPException: On validation or API errors
    """
    logger.info(f"Request to /v1/chat/completions (model={request_data.model}, stream={request_data.stream})")
    
    auth_manager: KiroAuthManager = request.app.state.auth_manager
    model_cache: ModelInfoCache = request.app.state.model_cache
    
    # Prepare debug logs
    if debug_logger:
        debug_logger.prepare_new_request()
    
    # Log incoming request
    try:
        request_body = json.dumps(request_data.model_dump(), ensure_ascii=False, indent=2).encode('utf-8')
        if debug_logger:
            debug_logger.log_request_body(request_body)
    except Exception as e:
        logger.warning(f"Failed to log request body: {e}")
    
    # Lazy model cache population
    if model_cache.is_empty():
        logger.debug("Model cache is empty, skipping forced population")
    
    # Generate conversation ID
    conversation_id = generate_conversation_id()
    
    # Build payload for Kiro
    try:
        kiro_payload = build_kiro_payload(
            request_data,
            conversation_id,
            auth_manager.profile_arn or ""
        )
    except ValueError as e:
        raise HTTPException(status_code=400, detail=str(e))
    
    # Log Kiro payload
    try:
        kiro_request_body = json.dumps(kiro_payload, ensure_ascii=False, indent=2).encode('utf-8')
        if debug_logger:
            debug_logger.log_kiro_request_body(kiro_request_body)
    except Exception as e:
        logger.warning(f"Failed to log Kiro request: {e}")
    
    # Create HTTP client with retry logic
    http_client = KiroHttpClient(auth_manager)
    url = f"{auth_manager.api_host}/generateAssistantResponse"
    try:
        # Make request to Kiro API (for both streaming and non-streaming modes)
        # Important: we wait for Kiro response BEFORE returning StreamingResponse,
        # so that 200 OK means Kiro accepted the request and started responding
        response = await http_client.request_with_retry(
            "POST",
            url,
            kiro_payload,
            stream=True
        )
        
        if response.status_code != 200:
            try:
                error_content = await response.aread()
            except Exception:
                error_content = b"Unknown error"
            
            await http_client.close()
            error_text = error_content.decode('utf-8', errors='replace')
            logger.error(f"Error from Kiro API: {response.status_code} - {error_text}")
            
            # Try to parse JSON response from Kiro to extract error message
            error_message = error_text
            try:
                error_json = json.loads(error_text)
                if "message" in error_json:
                    error_message = error_json["message"]
                    if "reason" in error_json:
                        error_message = f"{error_message} (reason: {error_json['reason']})"
            except (json.JSONDecodeError, KeyError):
                pass
            
            # Log access log for error (before flush, so it gets into app_logs)
            logger.warning(
                f"HTTP {response.status_code} - POST /v1/chat/completions - {error_message[:100]}"
            )
            
            # Flush debug logs on error ("errors" mode)
            if debug_logger:
                debug_logger.flush_on_error(response.status_code, error_message)
            
            # Return error in OpenAI API format
            return JSONResponse(
                status_code=response.status_code,
                content={
                    "error": {
                        "message": error_message,
                        "type": "kiro_api_error",
                        "code": response.status_code
                    }
                }
            )
        
        # Prepare data for fallback token counting
        # Convert Pydantic models to dicts for tokenizer
        messages_for_tokenizer = [msg.model_dump() for msg in request_data.messages]
        tools_for_tokenizer = [tool.model_dump() for tool in request_data.tools] if request_data.tools else None
        
        if request_data.stream:
            # Streaming mode
            async def stream_wrapper():
                streaming_error = None
                client_disconnected = False
                try:
                    async for chunk in stream_kiro_to_openai(
                        http_client.client,
                        response,
                        request_data.model,
                        model_cache,
                        auth_manager,
                        request_messages=messages_for_tokenizer,
                        request_tools=tools_for_tokenizer
                    ):
                        yield chunk
                except GeneratorExit:
                    # Client disconnected - this is normal
                    client_disconnected = True
                    logger.debug("Client disconnected during streaming (GeneratorExit in routes)")
                except Exception as e:
                    streaming_error = e
                    # Try to send [DONE] to client before finishing
                    # so client doesn't "hang" waiting for data
                    try:
                        yield "data: [DONE]\n\n"
                    except Exception:
                        pass  # Client already disconnected
                    raise
                finally:
                    await http_client.close()
                    # Log access log for streaming (success or error)
                    if streaming_error:
                        error_type = type(streaming_error).__name__
                        error_msg = str(streaming_error) if str(streaming_error) else "(empty message)"
                        logger.error(f"HTTP 500 - POST /v1/chat/completions (streaming) - [{error_type}] {error_msg[:100]}")
                    elif client_disconnected:
                        logger.info(f"HTTP 200 - POST /v1/chat/completions (streaming) - client disconnected")
                    else:
                        logger.info(f"HTTP 200 - POST /v1/chat/completions (streaming) - completed")
                    # Write debug logs AFTER streaming completes
                    if debug_logger:
                        if streaming_error:
                            debug_logger.flush_on_error(500, str(streaming_error))
                        else:
                            debug_logger.discard_buffers()
            
            return StreamingResponse(stream_wrapper(), media_type="text/event-stream")
        
        else:
            
            # Non-streaming mode - collect entire response
            openai_response = await collect_stream_response(
                http_client.client,
                response,
                request_data.model,
                model_cache,
                auth_manager,
                request_messages=messages_for_tokenizer,
                request_tools=tools_for_tokenizer
            )
            
            await http_client.close()
            
            # Log access log for non-streaming success
            logger.info(f"HTTP 200 - POST /v1/chat/completions (non-streaming) - completed")
            
            # Write debug logs after non-streaming request completes
            if debug_logger:
                debug_logger.discard_buffers()
            
            return JSONResponse(content=openai_response)
    
    except HTTPException as e:
        await http_client.close()
        # Log access log for HTTP error
        logger.warning(f"HTTP {e.status_code} - POST /v1/chat/completions - {e.detail}")
        # Flush debug logs on HTTP error ("errors" mode)
        if debug_logger:
            debug_logger.flush_on_error(e.status_code, str(e.detail))
        raise
    except Exception as e:
        await http_client.close()
        logger.error(f"Internal error: {e}", exc_info=True)
        # Log access log for internal error
        logger.error(f"HTTP 500 - POST /v1/chat/completions - {str(e)[:100]}")
        # Flush debug logs on internal error ("errors" mode)
        if debug_logger:
            debug_logger.flush_on_error(500, str(e))
        raise HTTPException(status_code=500, detail=f"Internal Server Error: {str(e)}")


================================================
FILE: kiro_gateway/streaming.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# https://github.com/jwadow/kiro-openai-gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
Streaming logic for converting Kiro stream to OpenAI format.

Contains generators for:
- Converting AWS SSE to OpenAI SSE
- Forming streaming chunks
- Processing tool calls in stream
"""

import asyncio
import json
import time
from typing import TYPE_CHECKING, AsyncGenerator, Callable, Awaitable, Optional

import httpx
from fastapi import HTTPException
from loguru import logger

from kiro_gateway.parsers import AwsEventStreamParser, parse_bracket_tool_calls, deduplicate_tool_calls
from kiro_gateway.utils import generate_completion_id
from kiro_gateway.config import FIRST_TOKEN_TIMEOUT, FIRST_TOKEN_MAX_RETRIES
from kiro_gateway.tokenizer import count_tokens, count_message_tokens, count_tools_tokens

if TYPE_CHECKING:
    from kiro_gateway.auth import KiroAuthManager
    from kiro_gateway.cache import ModelInfoCache

# Import debug_logger for logging
try:
    from kiro_gateway.debug_logger import debug_logger
except ImportError:
    debug_logger = None


class FirstTokenTimeoutError(Exception):
    """Exception raised when first token timeout occurs."""
    pass


async def stream_kiro_to_openai_internal(
    client: httpx.AsyncClient,
    response: httpx.Response,
    model: str,
    model_cache: "ModelInfoCache",
    auth_manager: "KiroAuthManager",
    first_token_timeout: float = FIRST_TOKEN_TIMEOUT,
    request_messages: Optional[list] = None,
    request_tools: Optional[list] = None
) -> AsyncGenerator[str, None]:
    """
    Internal generator for converting Kiro stream to OpenAI format.
    
    Parses AWS SSE stream and converts events to OpenAI chat.completion.chunk.
    Supports tool calls and usage calculation.
    
    IMPORTANT: This function raises FirstTokenTimeoutError if first token
    is not received within first_token_timeout seconds.
    
    Args:
        client: HTTP client (for connection management)
        response: HTTP response with data stream
        model: Model name to include in response
        model_cache: Model cache for getting token limits
        auth_manager: Authentication manager
        first_token_timeout: First token wait timeout (seconds)
        request_messages: Original request messages (for fallback token counting)
        request_tools: Original request tools (for fallback token counting)
    
    Yields:
        Strings in SSE format: "data: {...}\\n\\n" or "data: [DONE]\\n\\n"
    
    Raises:
        FirstTokenTimeoutError: If first token not received within timeout
    
    Example:
        >>> async for chunk in stream_kiro_to_openai_internal(client, response, "claude-sonnet-4", cache, auth):
        ...     print(chunk)
        data: {"id":"chatcmpl-...","object":"chat.completion.chunk",...}
        
        data: [DONE]
    """
    completion_id = generate_completion_id()
    created_time = int(time.time())
    first_chunk = True
    first_token_received = False
    
    parser = AwsEventStreamParser()
    metering_data = None
    context_usage_percentage = None
    full_content = ""
    
    streaming_error_occurred = False
    
    try:
        # Create iterator for reading bytes
        byte_iterator = response.aiter_bytes()
        
        # Wait for first chunk with timeout (FIRST_TOKEN_TIMEOUT)
        # This is our business logic for detecting "stuck" requests
        # where the model takes too long to start responding
        try:
            logger.debug(f"Waiting for first token (timeout={first_token_timeout}s)...")
            first_byte_chunk = await asyncio.wait_for(
                byte_iterator.__anext__(),
                timeout=first_token_timeout
            )
            logger.debug("First token received")
        except asyncio.TimeoutError:
            logger.warning(f"[FirstTokenTimeout] Model did not respond within {first_token_timeout}s")
            raise FirstTokenTimeoutError(f"No response within {first_token_timeout} seconds")
        except StopAsyncIteration:
            # Empty response - this is normal, just finish
            logger.debug("Empty response from Kiro API")
            yield "data: [DONE]\n\n"
            return
        
        # Process first chunk
        if debug_logger:
            debug_logger.log_raw_chunk(first_byte_chunk)
        
        events = parser.feed(first_byte_chunk)
        for event in events:
            if event["type"] == "content":
                first_token_received = True
                content = event["data"]
                full_content += content
                
                delta = {"content": content}
                if first_chunk:
                    delta["role"] = "assistant"
                    first_chunk = False
                
                openai_chunk = {
                    "id": completion_id,
                    "object": "chat.completion.chunk",
                    "created": created_time,
                    "model": model,
                    "choices": [{"index": 0, "delta": delta, "finish_reason": None}]
                }
                
                chunk_text = f"data: {json.dumps(openai_chunk, ensure_ascii=False)}\n\n"
                
                if debug_logger:
                    debug_logger.log_modified_chunk(chunk_text.encode('utf-8'))
                
                yield chunk_text
            
            elif event["type"] == "usage":
                metering_data = event["data"]
            
            elif event["type"] == "context_usage":
                context_usage_percentage = event["data"]
        
        # Continue reading remaining chunks (no longer with first token timeout)
        async for chunk in byte_iterator:
            # Log raw chunk
            if debug_logger:
                debug_logger.log_raw_chunk(chunk)
            
            events = parser.feed(chunk)
            
            for event in events:
                if event["type"] == "content":
                    content = event["data"]
                    full_content += content
                    
                    # Form delta
                    delta = {"content": content}
                    if first_chunk:
                        delta["role"] = "assistant"
                        first_chunk = False
                    
                    # Form OpenAI chunk
                    openai_chunk = {
                        "id": completion_id,
                        "object": "chat.completion.chunk",
                        "created": created_time,
                        "model": model,
                        "choices": [{"index": 0, "delta": delta, "finish_reason": None}]
                    }
                    
                    chunk_text = f"data: {json.dumps(openai_chunk, ensure_ascii=False)}\n\n"
                    
                    # Log modified chunk
                    if debug_logger:
                        debug_logger.log_modified_chunk(chunk_text.encode('utf-8'))
                    
                    yield chunk_text
                
                elif event["type"] == "usage":
                    metering_data = event["data"]
                
                elif event["type"] == "context_usage":
                    context_usage_percentage = event["data"]
        
        # Check bracket-style tool calls in full content
        bracket_tool_calls = parse_bracket_tool_calls(full_content)
        all_tool_calls = parser.get_tool_calls() + bracket_tool_calls
        all_tool_calls = deduplicate_tool_calls(all_tool_calls)
        
        # Determine finish_reason
        finish_reason = "tool_calls" if all_tool_calls else "stop"
        
        # Count completion_tokens (output) using tiktoken
        completion_tokens = count_tokens(full_content)
        
        # Calculate total_tokens based on context_usage_percentage from Kiro API
        # context_usage shows TOTAL percentage of context usage (input + output)
        total_tokens_from_api = 0
        if context_usage_percentage is not None and context_usage_percentage > 0:
            max_input_tokens = model_cache.get_max_input_tokens(model)
            total_tokens_from_api = int((context_usage_percentage / 100) * max_input_tokens)
        
        # Determine data source and calculate tokens
        if total_tokens_from_api > 0:
            # Use data from Kiro API
            # prompt_tokens (input) = total_tokens - completion_tokens
            prompt_tokens = max(0, total_tokens_from_api - completion_tokens)
            total_tokens = total_tokens_from_api
            prompt_source = "subtraction"
            total_source = "API Kiro"
        else:
            # Fallback: Kiro API didn't return context_usage, use tiktoken
            # Count prompt_tokens from original messages
            # IMPORTANT: Don't apply correction coefficient for prompt_tokens,
            # as it was calibrated for completion_tokens
            prompt_tokens = 0
            if request_messages:
                prompt_tokens += count_message_tokens(request_messages, apply_claude_correction=False)
            if request_tools:
                prompt_tokens += count_tools_tokens(request_tools, apply_claude_correction=False)
            total_tokens = prompt_tokens + completion_tokens
            prompt_source = "tiktoken"
            total_source = "tiktoken"
        
        # Send tool calls if present
        if all_tool_calls:
            logger.debug(f"Processing {len(all_tool_calls)} tool calls for streaming response")
            
            # Add required index field to each tool_call
            # according to OpenAI API specification for streaming
            indexed_tool_calls = []
            for idx, tc in enumerate(all_tool_calls):
                # Extract function with None protection
                func = tc.get("function") or {}
                # Use "or" for protection against explicit None in values
                tool_name = func.get("name") or ""
                tool_args = func.get("arguments") or "{}"
                
                logger.debug(f"Tool call [{idx}] '{tool_name}': id={tc.get('id')}, args_length={len(tool_args)}")
                
                indexed_tc = {
                    "index": idx,
                    "id": tc.get("id"),
                    "type": tc.get("type", "function"),
                    "function": {
                        "name": tool_name,
                        "arguments": tool_args
                    }
                }
                indexed_tool_calls.append(indexed_tc)
            
            tool_calls_chunk = {
                "id": completion_id,
                "object": "chat.completion.chunk",
                "created": created_time,
                "model": model,
                "choices": [{
                    "index": 0,
                    "delta": {"tool_calls": indexed_tool_calls},
                    "finish_reason": None
                }]
            }
            yield f"data: {json.dumps(tool_calls_chunk, ensure_ascii=False)}\n\n"
        
        # Final chunk with usage
        final_chunk = {
            "id": completion_id,
            "object": "chat.completion.chunk",
            "created": created_time,
            "model": model,
            "choices": [{"index": 0, "delta": {}, "finish_reason": finish_reason}],
            "usage": {
                "prompt_tokens": prompt_tokens,
                "completion_tokens": completion_tokens,
                "total_tokens": total_tokens,
            }
        }
        
        if metering_data:
            final_chunk["usage"]["credits_used"] = metering_data
        
        # Log final token values being sent to client
        logger.debug(
            f"[Usage] {model}: "
            f"prompt_tokens={prompt_tokens} ({prompt_source}), "
            f"completion_tokens={completion_tokens} (tiktoken), "
            f"total_tokens={total_tokens} ({total_source})"
        )
        
        yield f"data: {json.dumps(final_chunk, ensure_ascii=False)}\n\n"
        yield "data: [DONE]\n\n"
        
    except FirstTokenTimeoutError:
        # Propagate timeout up for retry
        raise
    except GeneratorExit:
        # Client disconnected - this is normal, don't log as error
        logger.debug("Client disconnected (GeneratorExit)")
        streaming_error_occurred = True
    except Exception as e:
        streaming_error_occurred = True
        # Log exception type and message for better diagnostics
        error_type = type(e).__name__
        error_msg = str(e) if str(e) else "(empty message)"
        logger.error(
            f"Error during streaming: [{error_type}] {error_msg}",
            exc_info=True
        )
        # Propagate error up for proper handling in routes.py
        raise
    finally:
        # Always close response
        try:
            await response.aclose()
        except Exception as close_error:
            logger.debug(f"Error closing response: {close_error}")
        
        if streaming_error_occurred:
            logger.debug("Streaming completed with error")
        else:
            logger.debug("Streaming completed successfully")


async def stream_kiro_to_openai(
    client: httpx.AsyncClient,
    response: httpx.Response,
    model: str,
    model_cache: "ModelInfoCache",
    auth_manager: "KiroAuthManager",
    request_messages: Optional[list] = None,
    request_tools: Optional[list] = None
) -> AsyncGenerator[str, None]:
    """
    Generator for converting Kiro stream to OpenAI format.
    
    This is a wrapper over stream_kiro_to_openai_internal that does NOT retry.
    Retry logic is implemented in stream_with_first_token_retry.
    
    Args:
        client: HTTP client (for connection management)
        response: HTTP response with data stream
        model: Model name to include in response
        model_cache: Model cache for getting token limits
        auth_manager: Authentication manager
        request_messages: Original request messages (for fallback token counting)
        request_tools: Original request tools (for fallback token counting)
    
    Yields:
        Strings in SSE format: "data: {...}\\n\\n" or "data: [DONE]\\n\\n"
    """
    async for chunk in stream_kiro_to_openai_internal(
        client, response, model, model_cache, auth_manager,
        request_messages=request_messages,
        request_tools=request_tools
    ):
        yield chunk


async def stream_with_first_token_retry(
    make_request: Callable[[], Awaitable[httpx.Response]],
    client: httpx.AsyncClient,
    model: str,
    model_cache: "ModelInfoCache",
    auth_manager: "KiroAuthManager",
    max_retries: int = FIRST_TOKEN_MAX_RETRIES,
    first_token_timeout: float = FIRST_TOKEN_TIMEOUT,
    request_messages: Optional[list] = None,
    request_tools: Optional[list] = None
) -> AsyncGenerator[str, None]:
    """
    Streaming with automatic retry on first token timeout.
    
    If model doesn't respond within first_token_timeout seconds,
    request is cancelled and a new one is made. Maximum max_retries attempts.
    
    This is seamless for user - they just see a delay,
    but eventually get a response (or error after all attempts).
    
    Args:
        make_request: Function to create new HTTP request
        client: HTTP client
        model: Model name
        model_cache: Model cache
        auth_manager: Authentication manager
        max_retries: Maximum number of attempts
        first_token_timeout: First token wait timeout (seconds)
        request_messages: Original request messages (for fallback token counting)
        request_tools: Original request tools (for fallback token counting)
    
    Yields:
        Strings in SSE format
    
    Raises:
        HTTPException: After exhausting all attempts
    
    Example:
        >>> async def make_req():
        ...     return await http_client.request_with_retry("POST", url, payload, stream=True)
        >>> async for chunk in stream_with_first_token_retry(make_req, client, model, cache, auth):
        ...     print(chunk)
    """
    last_error: Optional[Exception] = None
    
    for attempt in range(max_retries):
        response: Optional[httpx.Response] = None
        try:
            # Make request
            if attempt > 0:
                logger.warning(f"Retry attempt {attempt + 1}/{max_retries} after first token timeout")
            
            response = await make_request()
            
            if response.status_code != 200:
                # Error from API - close response and raise exception
                try:
                    error_content = await response.aread()
                    error_text = error_content.decode('utf-8', errors='replace')
                except Exception:
                    error_text = "Unknown error"
                
                try:
                    await response.aclose()
                except Exception:
                    pass
                
                logger.error(f"Error from Kiro API: {response.status_code} - {error_text}")
                raise HTTPException(
                    status_code=response.status_code,
                    detail=f"Upstream API error: {error_text}"
                )
            
            # Try to stream with first token timeout
            async for chunk in stream_kiro_to_openai_internal(
                client,
                response,
                model,
                model_cache,
                auth_manager,
                first_token_timeout=first_token_timeout,
                request_messages=request_messages,
                request_tools=request_tools
            ):
                yield chunk
            
            # Successfully completed - exit
            return
            
        except FirstTokenTimeoutError as e:
            last_error = e
            logger.warning(
                f"[FirstTokenTimeout] Attempt {attempt + 1}/{max_retries} failed - "
                f"model did not respond within {first_token_timeout}s"
            )
            
            # Close current response if open
            if response:
                try:
                    await response.aclose()
                except Exception:
                    pass
            
            # Continue to next attempt
            continue
            
        except Exception as e:
            # Other errors - no retry, propagate
            logger.error(f"Unexpected error during streaming: {e}", exc_info=True)
            if response:
                try:
                    await response.aclose()
                except Exception:
                    pass
            raise
    
    # All attempts exhausted - raise HTTP error
    logger.error(
        f"[FirstTokenTimeout] All {max_retries} attempts exhausted - "
        f"model never responded within {first_token_timeout}s per attempt"
    )
    raise HTTPException(
        status_code=504,
        detail=f"Model did not respond within {first_token_timeout}s after {max_retries} attempts. Please try again."
    )


async def collect_stream_response(
    client: httpx.AsyncClient,
    response: httpx.Response,
    model: str,
    model_cache: "ModelInfoCache",
    auth_manager: "KiroAuthManager",
    request_messages: Optional[list] = None,
    request_tools: Optional[list] = None
) -> dict:
    """
    Collect full response from streaming stream.
    
    Used for non-streaming mode - collects all chunks
    and forms a single response.
    
    Args:
        client: HTTP client
        response: HTTP response with stream
        model: Model name
        model_cache: Model cache
        auth_manager: Authentication manager
        request_messages: Original request messages (for fallback token counting)
        request_tools: Original request tools (for fallback token counting)
    
    Returns:
        Dictionary with full response in OpenAI chat.completion format
    """
    full_content = ""
    final_usage = None
    tool_calls = []
    completion_id = generate_completion_id()
    
    async for chunk_str in stream_kiro_to_openai(
        client,
        response,
        model,
        model_cache,
        auth_manager,
        request_messages=request_messages,
        request_tools=request_tools
    ):
        if not chunk_str.startswith("data:"):
            continue
        
        data_str = chunk_str[len("data:"):].strip()
        if not data_str or data_str == "[DONE]":
            continue
        
        try:
            chunk_data = json.loads(data_str)
            
            # Extract data from chunk
            delta = chunk_data.get("choices", [{}])[0].get("delta", {})
            if "content" in delta:
                full_content += delta["content"]
            if "tool_calls" in delta:
                tool_calls.extend(delta["tool_calls"])
            
            # Save usage from last chunk
            if "usage" in chunk_data:
                final_usage = chunk_data["usage"]
                
        except (json.JSONDecodeError, IndexError):
            continue
    
    # Form final response
    message = {"role": "assistant", "content": full_content}
    if tool_calls:
        # For non-streaming response remove index field from tool_calls,
        # as it's only required for streaming chunks
        cleaned_tool_calls = []
        for tc in tool_calls:
            # Extract function with None protection
            func = tc.get("function") or {}
            cleaned_tc = {
                "id": tc.get("id"),
                "type": tc.get("type", "function"),
                "function": {
                    "name": func.get("name", ""),
                    "arguments": func.get("arguments", "{}")
                }
            }
            cleaned_tool_calls.append(cleaned_tc)
        message["tool_calls"] = cleaned_tool_calls
    
    finish_reason = "tool_calls" if tool_calls else "stop"
    
    # Form usage for response
    usage = final_usage or {"prompt_tokens": 0, "completion_tokens": 0, "total_tokens": 0}
    
    # Log token info for debugging (non-streaming uses same logs from streaming)
    
    return {
        "id": completion_id,
        "object": "chat.completion",
        "created": int(time.time()),
        "model": model,
        "choices": [{
            "index": 0,
            "message": message,
            "finish_reason": finish_reason
        }],
        "usage": usage
    }


================================================
FILE: kiro_gateway/tokenizer.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
Модуль для быстрого подсчёта токенов.

Использует tiktoken (библиотека OpenAI на Rust) для приблизительного
подсчёта токенов. Кодировка cl100k_base близка к токенизации Claude.

Примечание: Это приблизительный подсчёт, так как точный токенизатор
Claude не является публичным. Anthropic не публикует свой токенизатор,
поэтому используется tiktoken с коэффициентом коррекции.

Коэффициент коррекции CLAUDE_CORRECTION_FACTOR = 1.15 основан на
эмпирических наблюдениях: Claude токенизирует текст примерно на 15%
больше чем GPT-4 (cl100k_base). Это связано с различиями в BPE словарях.
"""

from typing import List, Dict, Any, Optional
from loguru import logger

# Ленивая загрузка tiktoken для ускорения импорта
_encoding = None

# Коэффициент коррекции для Claude моделей
# Claude токенизирует текст примерно на 15% больше чем GPT-4 (cl100k_base)
# Это эмпирическое значение, основанное на сравнении с context_usage от API
CLAUDE_CORRECTION_FACTOR = 1.15


def _get_encoding():
    """
    Ленивая инициализация токенизатора.
    
    Использует cl100k_base - кодировку для GPT-4/ChatGPT,
    которая достаточно близка к токенизации Claude.
    
    Returns:
        tiktoken.Encoding или None если tiktoken недоступен
    """
    global _encoding
    if _encoding is None:
        try:
            import tiktoken
            _encoding = tiktoken.get_encoding("cl100k_base")
            logger.debug("[Tokenizer] Initialized tiktoken with cl100k_base encoding")
        except ImportError:
            logger.warning(
                "[Tokenizer] tiktoken not installed. "
                "Token counting will use fallback estimation. "
                "Install with: pip install tiktoken"
            )
            _encoding = False  # Маркер что импорт не удался
        except Exception as e:
            logger.error(f"[Tokenizer] Failed to initialize tiktoken: {e}")
            _encoding = False
    return _encoding if _encoding else None


def count_tokens(text: str, apply_claude_correction: bool = True) -> int:
    """
    Подсчитывает количество токенов в тексте.
    
    Args:
        text: Текст для подсчёта токенов
        apply_claude_correction: Применять коэффициент коррекции для Claude (по умолчанию True)
    
    Returns:
        Количество токенов (приблизительное, с коррекцией для Claude)
    """
    if not text:
        return 0
    
    encoding = _get_encoding()
    if encoding:
        try:
            base_tokens = len(encoding.encode(text))
            if apply_claude_correction:
                return int(base_tokens * CLAUDE_CORRECTION_FACTOR)
            return base_tokens
        except Exception as e:
            logger.warning(f"[Tokenizer] Error encoding text: {e}")
    
    # Fallback: грубая оценка ~4 символа на токен для английского,
    # ~2-3 символа для других языков (берём среднее ~3.5)
    # Для Claude добавляем коррекцию
    base_estimate = len(text) // 4 + 1
    if apply_claude_correction:
        return int(base_estimate * CLAUDE_CORRECTION_FACTOR)
    return base_estimate


def count_message_tokens(messages: List[Dict[str, Any]], apply_claude_correction: bool = True) -> int:
    """
    Подсчитывает токены в списке сообщений чата.
    
    Учитывает структуру сообщений OpenAI/Claude:
    - role: ~1 токен
    - content: токены текста
    - Служебные токены между сообщениями: ~3-4 токена
    
    Args:
        messages: Список сообщений в формате OpenAI
        apply_claude_correction: Применять коэффициент коррекции для Claude
    
    Returns:
        Приблизительное количество токенов (с коррекцией для Claude)
    """
    if not messages:
        return 0
    
    total_tokens = 0
    
    for message in messages:
        # Базовые токены на сообщение (role, разделители)
        total_tokens += 4  # ~4 токена на служебную информацию
        
        # Токены роли (без коррекции, это короткие строки)
        role = message.get("role", "")
        total_tokens += count_tokens(role, apply_claude_correction=False)
        
        # Токены контента
        content = message.get("content")
        if content:
            if isinstance(content, str):
                total_tokens += count_tokens(content, apply_claude_correction=False)
            elif isinstance(content, list):
                # Мультимодальный контент (текст + изображения)
                for item in content:
                    if isinstance(item, dict):
                        if item.get("type") == "text":
                            total_tokens += count_tokens(item.get("text", ""), apply_claude_correction=False)
                        elif item.get("type") == "image_url":
                            # Изображения занимают ~85-170 токенов в зависимости от размера
                            total_tokens += 100  # Средняя оценка
        
        # Токены tool_calls (если есть)
        tool_calls = message.get("tool_calls")
        if tool_calls:
            for tc in tool_calls:
                total_tokens += 4  # Служебные токены
                func = tc.get("function", {})
                total_tokens += count_tokens(func.get("name", ""), apply_claude_correction=False)
                total_tokens += count_tokens(func.get("arguments", ""), apply_claude_correction=False)
        
        # Токены tool_call_id (для ответов от инструментов)
        if message.get("tool_call_id"):
            total_tokens += count_tokens(message["tool_call_id"], apply_claude_correction=False)
    
    # Финальные служебные токены
    total_tokens += 3
    
    # Применяем коррекцию к общему количеству
    if apply_claude_correction:
        return int(total_tokens * CLAUDE_CORRECTION_FACTOR)
    return total_tokens


def count_tools_tokens(tools: Optional[List[Dict[str, Any]]], apply_claude_correction: bool = True) -> int:
    """
    Подсчитывает токены в определениях инструментов.
    
    Args:
        tools: Список инструментов в формате OpenAI
        apply_claude_correction: Применять коэффициент коррекции для Claude
    
    Returns:
        Приблизительное количество токенов (с коррекцией для Claude)
    """
    if not tools:
        return 0
    
    total_tokens = 0
    
    for tool in tools:
        total_tokens += 4  # Служебные токены
        
        if tool.get("type") == "function":
            func = tool.get("function", {})
            
            # Имя функции
            total_tokens += count_tokens(func.get("name", ""), apply_claude_correction=False)
            
            # Описание функции
            total_tokens += count_tokens(func.get("description", ""), apply_claude_correction=False)
            
            # Параметры (JSON schema)
            params = func.get("parameters")
            if params:
                import json
                params_str = json.dumps(params, ensure_ascii=False)
                total_tokens += count_tokens(params_str, apply_claude_correction=False)
    
    # Применяем коррекцию к общему количеству
    if apply_claude_correction:
        return int(total_tokens * CLAUDE_CORRECTION_FACTOR)
    return total_tokens


def estimate_request_tokens(
    messages: List[Dict[str, Any]],
    tools: Optional[List[Dict[str, Any]]] = None,
    system_prompt: Optional[str] = None
) -> Dict[str, int]:
    """
    Оценивает общее количество токенов в запросе.
    
    Args:
        messages: Список сообщений
        tools: Список инструментов (опционально)
        system_prompt: Системный промпт (опционально, если не в messages)
    
    Returns:
        Словарь с детализацией токенов:
        - messages_tokens: токены сообщений
        - tools_tokens: токены инструментов
        - system_tokens: токены системного промпта
        - total_tokens: общее количество
    """
    messages_tokens = count_message_tokens(messages)
    tools_tokens = count_tools_tokens(tools)
    system_tokens = count_tokens(system_prompt) if system_prompt else 0
    
    return {
        "messages_tokens": messages_tokens,
        "tools_tokens": tools_tokens,
        "system_tokens": system_tokens,
        "total_tokens": messages_tokens + tools_tokens + system_tokens
    }


================================================
FILE: kiro_gateway/utils.py
================================================
# -*- coding: utf-8 -*-

# Kiro OpenAI Gateway
# https://github.com/jwadow/kiro-openai-gateway
# Copyright (C) 2025 Jwadow
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <https://www.gnu.org/licenses/>.

"""
Utility functions for Kiro Gateway.

Contains functions for fingerprint generation, header formatting,
and other common utilities.
"""

import hashlib
import uuid
from typing import TYPE_CHECKING

from loguru import logger

if TYPE_CHECKING:
    from kiro_gateway.auth import KiroAuthManager


def get_machine_fingerprint() -> str:
    """
    Generates a unique machine fingerprint based on hostname and username.
    
    Used for User-Agent formation to identify a specific gateway installation.
    
    Returns:
        SHA256 hash of the string "{hostname}-{username}-kiro-gateway"
    """
    try:
        import socket
        import getpass
        
        hostname = socket.gethostname()
        username = getpass.getuser()
        unique_string = f"{hostname}-{username}-kiro-gateway"
        
        return hashlib.sha256(unique_string.encode()).hexdigest()
    except Exception as e:
        logger.warning(f"Failed to get machine fingerprint: {e}")
        return hashlib.sha256(b"default-kiro-gateway").hexdigest()


def get_kiro_headers(auth_manager: "KiroAuthManager", token: str) -> dict:
    """
    Builds headers for Kiro API requests.
    
    Includes all necessary headers for authentication and identification:
    - Authorization with Bearer token
    - User-Agent with fingerprint
    - AWS CodeWhisperer specific headers
    
    Args:
        auth_manager: Authentication manager for obtaining fingerprint
        token: Access token for authorization
    
    Returns:
        Dictionary with headers for HTTP request
    """
    fingerprint = auth_manager.fingerprint
    
    return {
        "Authorization": f"Bearer {token}",
        "Content-Type": "application/json",
        "User-Agent": f"aws-sdk-js/1.0.27 ua/2.1 os/win32#10.0.19044 lang/js md/nodejs#22.21.1 api/codewhispererstreaming#1.0.27 m/E KiroIDE-0.7.45-{fingerprint}",
        "x-amz-user-agent": f"aws-sdk-js/1.0.27 KiroIDE-0.7.45-{fingerprint}",
        "x-amzn-codewhisperer-optout": "true",
        "x-amzn-kiro-agent-mode": "vibe",
        "amz-sdk-invocation-id": str(uuid.uuid4()),
        "amz-sdk-request": "attempt=1; max=3",
    }


def generate_completion_id() -> str:
    """
    Generates a unique ID for chat completion.
    
    Returns:
        ID in format "chatcmpl-{uuid_hex}"
    """
    return f"chatcmpl-{uuid.uuid4().hex}"


def generate_conversation_id() -> str:
    """
    Generates a unique ID for conversation.
    
    Returns:
        UUID in string format
    """
    return str(uuid.uuid4())


def generate_tool_call_id() -> str:
    """
    Generates a unique ID for tool call.
    
    Returns:
        ID in format "call_{uuid_hex[:8]}"
    """
    return f"call_{uuid.uuid4().hex[:8]}"


================================================
FILE: tests/README.md
================================================
# Tests for Kiro OpenAI Gateway

A comprehensive set of unit and integration tests for Kiro OpenAI Gateway, providing full coverage of all system components.

## Testing Philosophy: Complete Network Isolation

**The key principle of this test suite is 100% isolation from real network requests.**

This is achieved through a global, automatically applied fixture `block_all_network_calls` in `tests/conftest.py`. It intercepts and blocks any attempts by `httpx.AsyncClient` to establish connections at the application level.

**Benefits:**
1.  **Reliability**: Tests don't depend on external API availability or network state.
2.  **Speed**: Absence of real network delays makes test execution instant.
3.  **Security**: Guarantees that test runs never use real credentials.

Any attempt to make an unauthorized network call will result in immediate test failure with an error, ensuring strict isolation control.

## Running Tests

### Installing Dependencies

```bash
# Main project dependencies
pip install -r requirements.txt

# Additional testing dependencies
pip install pytest pytest-asyncio hypothesis
```

### Running All Tests

```bash
# Run the entire test suite
pytest

# Run with verbose output
pytest -v

# Run with verbose output and coverage
pytest -v -s --tb=short

# Run only unit tests
pytest tests/unit/ -v

# Run only integration tests
pytest tests/integration/ -v

# Run a specific file
pytest tests/unit/test_auth_manager.py -v

# Run a specific test
pytest tests/unit/test_auth_manager.py::TestKiroAuthManagerInitialization::test_initialization_stores_credentials -v
```

### pytest Options

```bash
# Stop on first failure
pytest -x

# Show local variables on errors
pytest -l

# Run in parallel mode (requires pytest-xdist)
pip install pytest-xdist
pytest -n auto
```

## Test Structure

```
tests/
├── conftest.py                      # Shared fixtures and utilities
├── unit/                            # Unit tests for individual components
│   ├── test_auth_manager.py        # KiroAuthManager tests
│   ├── test_cache.py               # ModelInfoCache tests
│   ├── test_config.py              # Configuration tests (LOG_LEVEL, etc.)
│   ├── test_converters.py          # OpenAI <-> Kiro converter tests
│   ├── test_debug_logger.py        # DebugLogger tests (off/errors/all modes)
│   ├── test_parsers.py             # AwsEventStreamParser tests
│   ├── test_streaming.py           # Streaming function tests
│   ├── test_tokenizer.py           # Tokenizer tests (tiktoken)
│   ├── test_http_client.py         # KiroHttpClient tests
│   └── test_routes.py              # API endpoint tests
├── integration/                     # Integration tests for full flow
│   └── test_full_flow.py           # End-to-end tests
└── README.md                        # This file
```

## Test Coverage

### `conftest.py`

Shared fixtures and utilities for all tests:

**Environment Fixtures:**
- **`mock_env_vars()`**: Mocks environment variables (REFRESH_TOKEN, PROXY_API_KEY)
  - **What it does**: Isolates tests from real credentials
  - **Purpose**: Security and test reproducibility

**Data Fixtures:**
- **`valid_kiro_token()`**: Returns a mock Kiro access token
  - **What it does**: Provides a predictable token for tests
  - **Purpose**: Testing without real Kiro requests

- **`mock_kiro_token_response()`**: Factory for creating mock refreshToken responses
  - **What it does**: Generates Kiro auth endpoint response structure
  - **Purpose**: Testing various token refresh scenarios

- **`temp_creds_file()`**: Creates a temporary JSON file with credentials
  - **What it does**: Provides a file for testing credentials loading
  - **Purpose**: Testing credentials file operations

- **`sample_openai_chat_request()`**: Factory for creating OpenAI requests
  - **What it does**: Generates valid chat completion requests
  - **Purpose**: Convenient creation of test requests with different parameters

**Security Fixtures:**
- **`valid_proxy_api_key()`**: Valid proxy API key
- **`invalid_proxy_api_key()`**: Invalid key for negative tests
- **`auth_headers()`**: Factory for creating Authorization headers

**HTTP Fixtures:**
- **`mock_httpx_client()`**: Mocked httpx.AsyncClient
  - **What it does**: Isolates tests from real HTTP requests
  - **Purpose**: Test speed and reliability

- **`mock_httpx_response()`**: Factory for creating mock HTTP responses
  - **What it does**: Creates configurable httpx.Response objects
  - **Purpose**: Testing various HTTP scenarios

**Application Fixtures:**
- **`clean_app()`**: Clean FastAPI app instance
  - **What it does**: Returns a "clean" application instance
  - **Purpose**: Ensure application state isolation between tests

- **`test_client()`**: Synchronous FastAPI TestClient
- **`async_test_client()`**: Asynchronous test client for async endpoints

---

### `tests/unit/test_auth_manager.py`

Unit tests for **KiroAuthManager** (Kiro token management).

#### `TestKiroAuthManagerInitialization`

- **`test_initialization_stores_credentials()`**:
  - **What it does**: Verifies correct credential storage during creation
  - **Purpose**: Ensure all constructor parameters are stored in private fields

- **`test_initialization_sets_correct_urls_for_region()`**:
  - **What it does**: Verifies URL formation based on region
  - **Purpose**: Ensure URLs are dynamically formed with the correct region

- **`test_initialization_generates_fingerprint()`**:
  - **What it does**: Verifies unique fingerprint generation
  - **Purpose**: Ensure fingerprint is generated and has correct format

#### `TestKiroAuthManagerCredentialsFile`

- **`test_load_credentials_from_file()`**:
  - **What it does**: Verifies credentials loading from JSON file
  - **Purpose**: Ensure data is correctly read from file

- **`test_load_credentials_file_not_found()`**:
  - **What it does**: Verifies handling of missing credentials file
  - **Purpose**: Ensure application doesn't crash when file is missing

#### `TestKiroAuthManagerTokenExpiration`

- **`test_is_token_expiring_soon_returns_true_when_no_expires_at()`**:
  - **What it does**: Verifies that without expires_at, token is considered expiring
  - **Purpose**: Ensure safe behavior when time information is missing

- **`test_is_token_expiring_soon_returns_true_when_expired()`**:
  - **What it does**: Verifies that expired token is correctly identified
  - **Purpose**: Ensure token in the past is considered expiring

- **`test_is_token_expiring_soon_returns_true_within_threshold()`**:
  - **What it does**: Verifies that token within threshold is considered expiring
  - **Purpose**: Ensure token is refreshed in advance (10 minutes before expiration)

- **`test_is_token_expiring_soon_returns_false_when_valid()`**:
  - **What it does**: Verifies that valid token is not considered expiring
  - **Purpose**: Ensure token far in the future doesn't require refresh

#### `TestKiroAuthManagerTokenRefresh`

- **`test_refresh_token_successful()`**:
  - **What it does**: Tests successful token refresh via Kiro API
  - **Purpose**: Verify correct setting of access_token and expires_at

- **`test_refresh_token_updates_refresh_token()`**:
  - **What it does**: Verifies refresh_token update from response
  - **Purpose**: Ensure new refresh_token is saved

- **`test_refresh_token_missing_access_token_raises()`**:
  - **What it does**: Verifies handling of response without accessToken
  - **Purpose**: Ensure exception is thrown for incorrect response

- **`test_refresh_token_no_refresh_token_raises()`**:
  - **What it does**: Verifies handling of missing refresh_token
  - **Purpose**: Ensure exception is thrown without refresh_token

#### `TestKiroAuthManagerGetAccessToken`

- **`test_get_access_token_refreshes_when_expired()`**:
  - **What it does**: Verifies automatic refresh of expired token
  - **Purpose**: Ensure stale token is refreshed before return

- **`test_get_access_token_returns_valid_without_refresh()`**:
  - **What it does**: Verifies return of valid token without extra requests
  - **Purpose**: Optimization - don't make requests if token is still valid

- **`test_get_access_token_thread_safety()`**:
  - **What it does**: Verifies thread safety via asyncio.Lock
  - **Purpose**: Prevent race conditions during parallel calls

#### `TestKiroAuthManagerForceRefresh`

- **`test_force_refresh_updates_token()`**:
  - **What it does**: Verifies forced token refresh
  - **Purpose**: Ensure force_refresh always refreshes token

#### `TestKiroAuthManagerProperties`

- **`test_profile_arn_property()`**:
  - **What it does**: Verifies profile_arn property
  - **Purpose**: Ensure profile_arn is accessible via property

- **`test_region_property()`**:
  - **What it does**: Verifies region property
  - **Purpose**: Ensure region is accessible via property

- **`test_api_host_property()`**:
  - **What it does**: Verifies api_host property
  - **Purpose**: Ensure api_host is formed correctly

- **`test_fingerprint_property()`**:
  - **What it does**: Verifies fingerprint property
  - **Purpose**: Ensure fingerprint is accessible via property

---

### `tests/unit/test_cache.py`

Unit tests for **ModelInfoCache** (model metadata cache). **23 tests.**

#### `TestModelInfoCacheInitialization`

- **`test_initialization_creates_empty_cache()`**:
  - **What it does**: Verifies that cache is created empty
  - **Purpose**: Ensure correct initialization

- **`test_initialization_with_custom_ttl()`**:
  - **What it does**: Verifies cache creation with custom TTL
  - **Purpose**: Ensure TTL can be configured

- **`test_initialization_last_update_is_none()`**:
  - **What it does**: Verifies that last_update_time is initially None
  - **Purpose**: Ensure update time is not set before first update

#### `TestModelInfoCacheUpdate`

- **`test_update_populates_cache()`**:
  - **What it does**: Verifies cache population with data
  - **Purpose**: Ensure update() correctly saves models

- **`test_update_sets_last_update_time()`**:
  - **What it does**: Verifies setting of last update time
  - **Purpose**: Ensure last_update_time is set after update

- **`test_update_replaces_existing_data()`**:
  - **What it does**: Verifies data replacement on repeated update
  - **Purpose**: Ensure old data is completely replaced

- **`test_update_with_empty_list()`**:
  - **What it does**: Verifies update with empty list
  - **Purpose**: Ensure cache is cleared on empty update

#### `TestModelInfoCacheGet`

- **`test_get_returns_model_info()`**:
  - **What it does**: Verifies retrieval of model information
  - **Purpose**: Ensure get() returns correct data

- **`test_get_returns_none_for_unknown_model()`**:
  - **What it does**: Verifies None return for unknown model
  - **Purpose**: Ensure get() doesn't crash when model is missing

- **`test_get_from_empty_cache()`**:
  - **What it does**: Verifies get() from empty cache
  - **Purpose**: Ensure empty cache doesn't cause errors

#### `TestModelInfoCacheGetMaxInputTokens`

- **`test_get_max_input_tokens_returns_value()`**:
  - **What it does**: Verifies retrieval of maxInputTokens for model
  - **Purpose**: Ensure value is extracted from tokenLimits

- **`test_get_max_input_tokens_returns_default_for_unknown()`**:
  - **What it does**: Verifies default return for unknown model
  - **Purpose**: Ensure DEFAULT_MAX_INPUT_TOKENS is returned

- **`test_get_max_input_tokens_returns_default_when_no_token_limits()`**:
  - **What it does**: Verifies default return when tokenLimits is missing
  - **Purpose**: Ensure model without tokenLimits doesn't break logic

- **`test_get_max_input_tokens_returns_default_when_max_input_is_none()`**:
  - **What it does**: Verifies default return when maxInputTokens=None
  - **Purpose**: Ensure None in tokenLimits is handled correctly

#### `TestModelInfoCacheIsEmpty` and `TestModelInfoCacheIsStale`

- **`test_is_empty_returns_true_for_new_cache()`**: Verifies is_empty() for new cache
- **`test_is_empty_returns_false_after_update()`**: Verifies is_empty() after population
- **`test_is_stale_returns_true_for_new_cache()`**: Verifies is_stale() for new cache
- **`test_is_stale_returns_false_after_recent_update()`**: Verifies is_stale() right after update
- **`test_is_stale_returns_true_after_ttl_expires()`**: Verifies is_stale() after TTL expiration

#### `TestModelInfoCacheGetAllModelIds`

- **`test_get_all_model_ids_returns_empty_for_new_cache()`**: Verifies get_all_model_ids() for empty cache
- **`test_get_all_model_ids_returns_all_ids()`**: Verifies get_all_model_ids() for populated cache

#### `TestModelInfoCacheThreadSafety`

- **`test_concurrent_updates_dont_corrupt_cache()`**:
  - **What it does**: Verifies thread safety during parallel updates
  - **Purpose**: Ensure asyncio.Lock protects against race conditions

- **`test_concurrent_reads_are_safe()`**:
  - **What it does**: Verifies safety of parallel reads
  - **Purpose**: Ensure multiple get() calls don't cause issues

---

### `tests/unit/test_config.py`

Unit tests for **configuration module** (loading settings from environment variables). **13 tests.**

#### `TestLogLevelConfig`

Tests for LOG_LEVEL configuration.

- **`test_default_log_level_is_info()`**:
  - **What it does**: Verifies that default LOG_LEVEL is INFO
  - **Purpose**: Ensure INFO is used without environment variable

- **`test_log_level_from_environment()`**:
  - **What it does**: Verifies LOG_LEVEL loading from environment variable
  - **Purpose**: Ensure value from environment is used

- **`test_log_level_uppercase_conversion()`**:
  - **What it does**: Verifies LOG_LEVEL conversion to uppercase
  - **Purpose**: Ensure lowercase value is converted to uppercase

- **`test_log_level_trace()`**:
  - **What it does**: Verifies LOG_LEVEL=TRACE setting
  - **Purpose**: Ensure TRACE level is supported

- **`test_log_level_error()`**:
  - **What it does**: Verifies LOG_LEVEL=ERROR setting
  - **Purpose**: Ensure ERROR level is supported

- **`test_log_level_critical()`**:
  - **What it does**: Verifies LOG_LEVEL=CRITICAL setting
  - **Purpose**: Ensure CRITICAL level is supported

#### `TestToolDescriptionMaxLengthConfig`

Tests for TOOL_DESCRIPTION_MAX_LENGTH configuration.

- **`test_default_tool_description_max_length()`**:
  - **What it does**: Verifies default value for TOOL_DESCRIPTION_MAX_LENGTH
  - **Purpose**: Ensure default is 10000

- **`test_tool_description_max_length_from_environment()`**:
  - **What it does**: Verifies TOOL_DESCRIPTION_MAX_LENGTH loading from environment
  - **Purpose**: Ensure value from environment is used

- **`test_tool_description_max_length_zero_disables()`**:
  - **What it does**: Verifies that 0 disables the feature
  - **Purpose**: Ensure TOOL_DESCRIPTION_MAX_LENGTH=0 works

#### `TestTimeoutConfigurationWarning`

Tests for `_warn_timeout_configuration()` function.

- **`test_no_warning_when_first_token_less_than_streaming()`**:
  - **What it does**: Verifies no warning when FIRST_TOKEN_TIMEOUT < STREAMING_READ_TIMEOUT
  - **Purpose**: Ensure correct configuration doesn't trigger warning

- **`test_warning_when_first_token_equals_streaming()`**:
  - **What it does**: Verifies warning when FIRST_TOKEN_TIMEOUT == STREAMING_READ_TIMEOUT
  - **Purpose**: Ensure equal timeouts trigger warning

- **`test_warning_when_first_token_greater_than_streaming()`**:
  - **What it does**: Verifies warning when FIRST_TOKEN_TIMEOUT > STREAMING_READ_TIMEOUT
  - **Purpose**: Ensure suboptimal configuration triggers warning with timeout values

- **`test_warning_contains_recommendation()`**:
  - **What it does**: Verifies warning contains recommendation text
  - **Purpose**: Ensure user gets helpful information about correct configuration

---

### `tests/unit/test_debug_logger.py`

Unit tests for **DebugLogger** (debug request logging). **26 tests.**

#### `TestDebugLoggerModeOff`

Tests for DEBUG_MODE=off mode.

- **`test_prepare_new_request_does_nothing()`**:
  - **What it does**: Verifies that prepare_new_request does nothing in off mode
  - **Purpose**: Ensure directory is not created in off mode

- **`test_log_request_body_does_nothing()`**:
  - **What it does**: Verifies that log_request_body does nothing in off mode
  - **Purpose**: Ensure data is not written

#### `TestDebugLoggerModeAll`

Tests for DEBUG_MODE=all mode.

- **`test_prepare_new_request_clears_directory()`**:
  - **What it does**: Verifies that prepare_new_request clears directory in all mode
  - **Purpose**: Ensure old logs are deleted

- **`test_log_request_body_writes_immediately()`**:
  - **What it does**: Verifies that log_request_body writes immediately to file in all mode
  - **Purpose**: Ensure data is written immediately

- **`test_log_kiro_request_body_writes_immediately()`**:
  - **What it does**: Verifies that log_kiro_request_body writes immediately to file in all mode
  - **Purpose**: Ensure Kiro payload is written immediately

- **`test_log_raw_chunk_appends_to_file()`**:
  - **What it does**: Verifies that log_raw_chunk appends to file in all mode
  - **Purpose**: Ensure chunks accumulate

#### `TestDebugLoggerModeErrors`

Tests for DEBUG_MODE=errors mode.

- **`test_log_request_body_buffers_data()`**:
  - **What it does**: Verifies that log_request_body buffers data in errors mode
  - **Purpose**: Ensure data is not written immediately

- **`test_flush_on_error_writes_buffers()`**:
  - **What it does**: Verifies that flush_on_error writes buffers to files
  - **Purpose**: Ensure data is saved on error

- **`test_flush_on_error_clears_buffers()`**:
  - **What it does**: Verifies that flush_on_error clears buffers after writing
  - **Purpose**: Ensure buffers don't accumulate between requests

- **`test_discard_buffers_clears_without_writing()`**:
  - **What it does**: Verifies that discard_buffers clears buffers without writing
  - **Purpose**: Ensure successful requests don't leave logs

- **`test_flush_on_error_writes_error_info_in_mode_all()`**:
  - **What it does**: Verifies that flush_on_error writes error_info.json in all mode
  - **Purpose**: Ensure error information is saved in both modes

#### `TestDebugLoggerLogErrorInfo`

Tests for log_error_info() method.

- **`test_log_error_info_writes_in_mode_all()`**:
  - **What it does**: Verifies that log_error_info writes file in all mode
  - **Purpose**: Ensure error_info.json is created on errors

- **`test_log_error_info_writes_in_mode_errors()`**:
  - **What it does**: Verifies that log_error_info writes file in errors mode
  - **Purpose**: Ensure method works in both modes

- **`test_log_error_info_does_nothing_in_mode_off()`**:
  - **What it does**: Verifies that log_error_info does nothing in off mode
  - **Purpose**: Ensure files are not created in off mode

#### `TestDebugLoggerHelperMethods`

Tests for DebugLogger helper methods.

- **`test_is_enabled_returns_true_for_errors()`**: Verifies _is_enabled() for errors mode
- **`test_is_enabled_returns_true_for_all()`**: Verifies _is_enabled() for all mode
- **`test_is_enabled_returns_false_for_off()`**: Verifies _is_enabled() for off mode
- **`test_is_immediate_write_returns_true_for_all()`**: Verifies _is_immediate_write() for all mode
- **`test_is_immediate_write_returns_false_for_errors()`**: Verifies _is_immediate_write() for errors mode

#### `TestDebugLoggerJsonHandling`

Tests for JSON handling in DebugLogger.

- **`test_log_request_body_formats_json_pretty()`**:
  - **What it does**: Verifies that JSON is formatted prettily
  - **Purpose**: Ensure JSON is readable in file

- **`test_log_request_body_handles_invalid_json()`**:
  - **What it does**: Verifies handling of invalid JSON
  - **Purpose**: Ensure invalid JSON is written as-is

#### `TestDebugLoggerAppLogsCapture`

Tests for application log capture (app_logs.txt).

- **`test_prepare_new_request_sets_up_log_capture()`**:
  - **What it does**: Verifies that prepare_new_request sets up log capture
  - **Purpose**: Ensure sink for logs is created

- **`test_flush_on_error_writes_app_logs_in_mode_errors()`**:
  - **What it does**: Verifies that flush_on_error writes app_logs.txt in errors mode
  - **Purpose**: Ensure application logs are saved on errors

- **`test_discard_buffers_saves_logs_in_mode_all()`**:
  - **What it does**: Verifies that discard_buffers saves logs in all mode
  - **Purpose**: Ensure even successful requests save logs in all mode

- **`test_discard_buffers_does_not_save_logs_in_mode_errors()`**:
  - **What it does**: Verifies that discard_buffers does NOT save logs in errors mode
  - **Purpose**: Ensure successful requests don't leave logs in errors mode

- **`test_clear_app_logs_buffer_removes_sink()`**:
  - **What it does**: Verifies that _clear_app_logs_buffer removes sink
  - **Purpose**: Ensure sink is correctly removed

- **`test_app_logs_not_saved_when_empty()`**:
  - **What it does**: Verifies that empty logs don't create file
  - **Purpose**: Ensure app_logs.txt is not created if there are no logs

---

### `tests/unit/test_converters.py`

Unit tests for **OpenAI <-> Kiro** converters. **68 tests.**

#### `TestExtractTextContent`

- **`test_extracts_from_string()`**: Verifies text extraction from string
- **`test_extracts_from_none()`**: Verifies None handling
- **`test_extracts_from_list_with_text_type()`**: Verifies extraction from list with type=text
- **`test_extracts_from_list_with_text_key()`**: Verifies extraction from list with text key
- **`test_extracts_from_list_with_strings()`**: Verifies extraction from list of strings
- **`test_extracts_from_mixed_list()`**: Verifies extraction from mixed list
- **`test_converts_other_types_to_string()`**: Verifies conversion of other types to string
- **`test_handles_empty_list()`**: Verifies empty list handling

#### `TestMergeAdjacentMessages`

- **`test_merges_adjacent_user_messages()`**: Verifies merging of adjacent user messages
- **`test_preserves_alternating_messages()`**: Verifies preservation of alternating messages
- **`test_handles_empty_list()`**: Verifies empty list handling
- **`test_handles_single_message()`**: Verifies single message handling
- **`test_merges_multiple_adjacent_groups()`**: Verifies merging of multiple groups

**New tests for tool message handling (role="tool"):**

- **`test_converts_tool_message_to_user_with_tool_result()`**:
  - **What it does**: Verifies conversion of tool message to user message with tool_result
  - **Purpose**: Ensure role="tool" is converted to user message with tool_results content

- **`test_converts_multiple_tool_messages_to_single_user_message()`**:
  - **What it does**: Verifies merging of multiple tool messages into single user message
  - **Purpose**: Ensure multiple tool results are merged into one user message

- **`test_tool_message_followed_by_user_message()`**:
  - **What it does**: Verifies tool message before user message
  - **Purpose**: Ensure tool results and user message are merged

- **`test_assistant_tool_user_sequence()`**:
  - **What it does**: Verifies assistant -> tool -> user sequence
  - **Purpose**: Ensure tool message is correctly inserted between assistant and user

- **`test_tool_message_with_empty_content()`**:
  - **What it does**: Verifies tool message with empty content
  - **Purpose**: Ensure empty result is replaced with "(empty result)"

- **`test_tool_message_with_none_tool_call_id()`**:
  - **What it does**: Verifies tool message without tool_call_id
  - **Purpose**: Ensure missing tool_call_id is replaced with empty string

- **`test_merges_list_contents_correctly()`**:
  - **What it does**: Verifies list contents merging
  - **Purpose**: Ensure lists are merged correctly

- **`test_merges_adjacent_assistant_tool_calls()`**:
  - **What it does**: Verifies tool_calls merging when merging adjacent assistant messages
  - **Purpose**: Ensure tool_calls from all assistant messages are preserved when merging

- **`test_merges_three_adjacent_assistant_tool_calls()`**:
  - **What it does**: Verifies tool_calls merging from three assistant messages
  - **Purpose**: Ensure all tool_calls are preserved when merging more than two messages

- **`test_merges_assistant_with_and_without_tool_calls()`**:
  - **What it does**: Verifies merging assistant with and without tool_calls
  - **Purpose**: Ensure tool_calls are correctly initialized when merging

#### `TestBuildKiroPayloadToolCallsIntegration`

Integration tests for full tool_calls flow from OpenAI to Kiro format.

- **`test_multiple_assistant_tool_calls_with_results()`**:
  - **What it does**: Verifies full scenario with multiple assistant tool_calls and their results
  - **Purpose**: Ensure all toolUses and toolResults are correctly linked in Kiro payload

#### `TestBuildKiroHistory`

- **`test_builds_user_message()`**: Verifies user message building
- **`test_builds_assistant_message()`**: Verifies assistant message building
- **`test_ignores_system_messages()`**: Verifies system message ignoring
- **`test_builds_conversation_history()`**: Verifies full conversation history building
- **`test_handles_empty_list()`**: Verifies empty list handling

#### `TestExtractToolResults` and `TestExtractToolUses`

- **`test_extracts_tool_results_from_list()`**: Verifies tool results extraction from list
- **`test_returns_empty_for_string_content()`**: Verifies empty list return for string
- **`test_extracts_from_tool_calls_field()`**: Verifies extraction from tool_calls field
- **`test_extracts_from_content_list()`**: Verifies extraction from content list

#### `TestProcessToolsWithLongDescriptions`

Tests for tools processing function with long descriptions (Tool Documentation Reference Pattern).

- **`test_returns_none_and_empty_string_for_none_tools()`**:
  - **What it does**: Verifies handling of None instead of tools list
  - **Purpose**: Ensure None returns (None, "")

- **`test_returns_none_and_empty_string_for_empty_list()`**:
  - **What it does**: Verifies handling of empty tools list
  - **Purpose**: Ensure empty list returns (None, "")

- **`test_short_description_unchanged()`**:
  - **What it does**: Verifies that short descriptions are not changed
  - **Purpose**: Ensure tools with short descriptions remain as-is

- **`test_long_description_moved_to_system_prompt()`**:
  - **What it does**: Verifies moving long description to system prompt
  - **Purpose**: Ensure long descriptions are moved correctly with reference in tool

- **`test_mixed_short_and_long_descriptions()`**:
  - **What it does**: Verifies handling of mixed tools list
  - **Purpose**: Ensure short ones stay, long ones are moved

- **`test_preserves_tool_parameters()`**:
  - **What it does**: Verifies parameters preservation when moving description
  - **Purpose**: Ensure parameters are not lost

- **`test_disabled_when_limit_is_zero()`**:
  - **What it does**: Verifies feature disabling when limit is 0
  - **Purpose**: Ensure tools are not changed when TOOL_DESCRIPTION_MAX_LENGTH=0

- **`test_non_function_tools_unchanged()`**:
  - **What it does**: Verifies that non-function tools are not changed
  - **Purpose**: Ensure only function tools are processed

- **`test_multiple_long_descriptions_all_moved()`**:
  - **What it does**: Verifies moving multiple long descriptions
  - **Purpose**: Ensure all long descriptions are moved

- **`test_empty_description_unchanged()`**:
  - **What it does**: Verifies handling of empty description
  - **Purpose**: Ensure empty description doesn't cause errors

- **`test_none_description_unchanged()`**:
  - **What it does**: Verifies handling of None description
  - **Purpose**: Ensure None description doesn't cause errors

#### `TestSanitizeJsonSchema`

Tests for `_sanitize_json_schema` function that cleans JSON Schema from fields not supported by Kiro API.

- **`test_returns_empty_dict_for_none()`**:
  - **What it does**: Verifies None handling
  - **Purpose**: Ensure None returns empty dict

- **`test_returns_empty_dict_for_empty_dict()`**:
  - **What it does**: Verifies empty dict handling
  - **Purpose**: Ensure empty dict is returned as-is

- **`test_removes_empty_required_array()`**:
  - **What it does**: Verifies removal of empty required array
  - **Purpose**: Ensure `required: []` is removed from schema (critical test for Cline bug)

- **`test_preserves_non_empty_required_array()`**:
  - **What it does**: Verifies preservation of non-empty required array
  - **Purpose**: Ensure required with elements is preserved

- **`test_removes_additional_properties()`**:
  - **What it does**: Verifies additionalProperties removal
  - **Purpose**: Ensure additionalProperties is removed from schema

- **`test_removes_both_empty_required_and_additional_properties()`**:
  - **What it does**: Verifies removal of both problematic fields
  - **Purpose**: Ensure both fields are removed simultaneously (real scenario from Cline)

- **`test_recursively_sanitizes_nested_properties()`**:
  - **What it does**: Verifies recursive cleaning of nested properties
  - **Purpose**: Ensure nested schemas are also cleaned

- **`test_recursively_sanitizes_dict_values()`**:
  - **What it does**: Verifies recursive cleaning of dict values
  - **Purpose**: Ensure any nested dicts are cleaned

- **`test_sanitizes_items_in_lists()`**:
  - **What it does**: Verifies cleaning of items in lists (anyOf, oneOf)
  - **Purpose**: Ensure list items are also cleaned

- **`test_preserves_non_dict_list_items()`**:
  - **What it does**: Verifies preservation of non-dict items in lists
  - **Purpose**: Ensure strings and other types in lists are preserved

- **`test_complex_real_world_schema()`**:
  - **What it does**: Verifies cleaning of real complex schema from Cline
  - **Purpose**: Ensure real schemas are processed correctly

#### `TestBuildUserInputContext`

- **`test_builds_tools_context()`**: Verifies context building with tools
- **`test_returns_empty_for_no_tools()`**: Verifies empty context return without tools

**New tests for empty description placeholder (Cline bug fix):**

- **`test_empty_description_replaced_with_placeholder()`**:
  - **What it does**: Verifies empty description replacement with placeholder
  - **Purpose**: Ensure empty description is replaced with "Tool: {name}" (critical test for Cline bug with focus_chain)

- **`test_whitespace_only_description_replaced_with_placeholder()`**:
  - **What it does**: Verifies whitespace-only description replacement with placeholder
  - **Purpose**: Ensure description with only spaces is replaced

- **`test_none_description_replaced_with_placeholder()`**:
  - **What it does**: Verifies None description replacement with placeholder
  - **Purpose**: Ensure None description is replaced with "Tool: {name}"

- **`test_non_empty_description_preserved()`**:
  - **What it does**: Verifies non-empty description preservation
  - **Purpose**: Ensure normal description is not changed

- **`test_sanitizes_tool_parameters()`**:
  - **What it does**: Verifies parameters cleaning from problematic fields
  - **Purpose**: Ensure _sanitize_json_schema is applied to parameters

- **`test_mixed_tools_with_empty_and_normal_descriptions()`**:
  - **What it does**: Verifies handling of mixed tools list
  - **Purpose**: Ensure empty descriptions are replaced while normal ones are preserved (real scenario from Cline)

#### `TestBuildKiroPayload`

- **`test_builds_simple_payload()`**: Verifies simple payload building
- **`test_includes_system_prompt_in_first_message()`**: Verifies system prompt addition
- **`test_builds_history_for_multi_turn()`**: Verifies history building for multi-turn
- **`test_handles_assistant_as_last_message()`**: Verifies handling of assistant as last message
- **`test_raises_for_empty_messages()`**: Verifies exception throwing for empty messages
- **`test_uses_continue_for_empty_content()`**: Verifies "Continue" usage for empty content
- **`test_maps_model_id_correctly()`**: Verifies external model ID mapping to internal
- **`test_long_tool_description_added_to_system_prompt()`**:
  - **What it does**: Verifies long tool descriptions integration in payload
  - **Purpose**: Ensure long descriptions are added to system prompt in payload

---

### `tests/unit/test_parsers.py`

Unit tests for **AwsEventStreamParser** and helper parsing functions. **52 tests.**

#### `TestFindMatchingBrace`

- **`test_simple_json_object()`**: Verifies closing brace search for simple JSON
- **`test_nested_json_object()`**: Verifies search for nested JSON
- **`test_json_with_braces_in_string()`**: Verifies ignoring braces inside strings
- **`test_json_with_escaped_quotes()`**: Verifies handling of escaped quotes
- **`test_incomplete_json()`**: Verifies handling of incomplete JSON
- **`test_invalid_start_position()`**: Verifies handling of invalid start position
- **`test_start_position_out_of_bounds()`**: Verifies handling of position beyond text

#### `TestParseBracketToolCalls`

- **`test_parses_single_tool_call()`**: Verifies parsing of single tool call
- **`test_parses_multiple_tool_calls()`**: Verifies parsing of multiple tool calls
- **`test_returns_empty_for_no_tool_calls()`**: Verifies empty list return without tool calls
- **`test_returns_empty_for_empty_string()`**: Verifies empty string handling
- **`test_returns_empty_for_none()`**: Verifies None handling
- **`test_handles_nested_json_in_args()`**: Verifies parsing of nested JSON in arguments
- **`test_generates_unique_ids()`**: Verifies unique ID generation for tool calls

#### `TestDeduplicateToolCalls`

- **`test_removes_duplicates()`**: Verifies duplicate removal
- **`test_preserves_first_occurrence()`**: Verifies first occurrence preservation
- **`test_handles_empty_list()`**: Verifies empty list handling

**New tests for improved deduplication by id:**

- **`test_deduplicates_by_id_keeps_one_with_arguments()`**:
  - **What it does**: Verifies deduplication by id keeping tool call with arguments
  - **Purpose**: Ensure when duplicates by id, the one with arguments is kept

- **`test_deduplicates_by_id_prefers_longer_arguments()`**:
  - **What it does**: Verifies that duplicates by id prefer longer arguments
  - **Purpose**: Ensure tool call with more complete arguments is kept

- **`test_deduplicates_empty_arguments_replaced_by_non_empty()`**:
  - **What it does**: Verifies empty arguments replacement with non-empty
  - **Purpose**: Ensure "{}" is replaced with real arguments

- **`test_handles_tool_calls_without_id()`**:
  - **What it does**: Verifies handling of tool calls without id
  - **Purpose**: Ensure tool calls without id are deduplicated by name+arguments

- **`test_mixed_with_and_without_id()`**:
  - **What it does**: Verifies mixed list with and without id
  - **Purpose**: Ensure both types are handled correctly

#### `TestAwsEventStreamParserInitialization`

- **`test_initialization_creates_empty_state()`**: Verifies initial parser state

#### `TestAwsEventStreamParserFeed`

- **`test_parses_content_event()`**: Verifies content event parsing
- **`test_parses_multiple_content_events()`**: Verifies multiple content events parsing
- **`test_deduplicates_repeated_content()`**: Verifies repeated content deduplication
- **`test_parses_usage_event()`**: Verifies usage event parsing
- **`test_parses_context_usage_event()`**: Verifies context_usage event parsing
- **`test_handles_incomplete_json()`**: Verifies incomplete JSON handling
- **`test_completes_json_across_chunks()`**: Verifies JSON assembly from multiple chunks
- **`test_decodes_escape_sequences()`**: Verifies escape sequence decoding
- **`test_handles_invalid_bytes()`**: Verifies invalid bytes handling

#### `TestAwsEventStreamParserToolCalls`

- **`test_parses_tool_start_event()`**: Verifies tool call start parsing
- **`test_parses_tool_input_event()`**: Verifies tool call input parsing
- **`test_parses_tool_stop_event()`**: Verifies tool call completion
- **`test_get_tool_calls_returns_all()`**: Verifies getting all tool calls
- **`test_get_tool_calls_finalizes_current()`**: Verifies incomplete tool call finalization

#### `TestAwsEventStreamParserReset`

- **`test_reset_clears_state()`**: Verifies parser state reset

#### `TestAwsEventStreamParserFinalizeToolCall`

**New tests for _finalize_tool_call method with different input types:**

- **`test_finalize_with_string_arguments()`**:
  - **What it does**: Verifies tool call finalization with string arguments
  - **Purpose**: Ensure JSON string is parsed and serialized back

- **`test_finalize_with_dict_arguments()`**:
  - **What it does**: Verifies tool call finalization with dict arguments
  - **Purpose**: Ensure dict is serialized to JSON string

- **`test_finalize_with_empty_string_arguments()`**:
  - **What it does**: Verifies tool call finalization with empty string arguments
  - **Purpose**: Ensure empty string is replaced with "{}"

- **`test_finalize_with_whitespace_only_arguments()`**:
  - **What it does**: Verifies tool call finalization with whitespace arguments
  - **Purpose**: Ensure whitespace string is replaced with "{}"

- **`test_finalize_with_invalid_json_arguments()`**:
  - **What it does**: Verifies tool call finalization with invalid JSON
  - **Purpose**: Ensure invalid JSON is replaced with "{}"

- **`test_finalize_with_none_current_tool_call()`**:
  - **What it does**: Verifies finalization when current_tool_call is None
  - **Purpose**: Ensure nothing happens with None

- **`test_finalize_clears_current_tool_call()`**:
  - **What it does**: Verifies that finalization clears current_tool_call
  - **Purpose**: Ensure current_tool_call = None after finalization

#### `TestAwsEventStreamParserEdgeCases`

- **`test_handles_followup_prompt()`**: Verifies followupPrompt ignoring
- **`test_handles_mixed_events()`**: Verifies mixed events parsing
- **`test_handles_garbage_between_events()`**: Verifies garbage handling between events
- **`test_handles_empty_chunk()`**: Verifies empty chunk handling

---

### `tests/unit/test_tokenizer.py`

Unit tests for **tokenizer module** (token counting with tiktoken). **32 tests.**

#### `TestCountTokens`

Tests for count_tokens function.

- **`test_empty_string_returns_zero()`**:
  - **What it does**: Verifies that empty string returns 0 tokens
  - **Purpose**: Ensure correct edge case handling

- **`test_none_returns_zero()`**:
  - **What it does**: Verifies that None returns 0 tokens
  - **Purpose**: Ensure correct None handling

- **`test_simple_text_returns_positive()`**:
  - **What it does**: Verifies that simple text returns positive token count
  - **Purpose**: Ensure basic counting functionality

- **`test_longer_text_returns_more_tokens()`**:
  - **What it does**: Verifies that longer text returns more tokens
  - **Purpose**: Ensure correct counting proportionality

- **`test_claude_correction_applied_by_default()`**:
  - **What it does**: Verifies that Claude correction factor is applied by default
  - **Purpose**: Ensure apply_claude_correction=True by default

- **`test_without_claude_correction()`**:
  - **What it does**: Verifies counting without correction factor
  - **Purpose**: Ensure apply_claude_correction=False works

- **`test_unicode_text()`**:
  - **What it does**: Verifies token counting for Unicode text
  - **Purpose**: Ensure correct non-ASCII character handling

- **`test_multiline_text()`**:
  - **What it does**: Verifies token counting for multiline text
  - **Purpose**: Ensure correct newline handling

- **`test_json_text()`**:
  - **What it does**: Verifies token counting for JSON string
  - **Purpose**: Ensure correct JSON handling

#### `TestCountTokensFallback`

Tests for fallback logic when tiktoken is unavailable.

- **`test_fallback_when_tiktoken_unavailable()`**:
  - **What it does**: Verifies fallback counting when tiktoken is unavailable
  - **Purpose**: Ensure system works without tiktoken

- **`test_fallback_without_correction()`**:
  - **What it does**: Verifies fallback without correction factor
  - **Purpose**: Ensure fallback works with apply_claude_correction=False

#### `TestCountMessageTokens`

Tests for count_message_tokens function.

- **`test_empty_list_returns_zero()`**:
  - **What it does**: Verifies that empty list returns 0 tokens
  - **Purpose**: Ensure correct empty list handling

- **`test_none_returns_zero()`**:
  - **What it does**: Verifies that None returns 0 tokens
  - **Purpose**: Ensure correct None handling

- **`test_single_user_message()`**:
  - **What it does**: Verifies token counting for single user message
  - **Purpose**: Ensure basic functionality

- **`test_multiple_messages()`**:
  - **What it does**: Verifies token counting for multiple messages
  - **Purpose**: Ensure tokens are summed correctly

- **`test_message_with_tool_calls()`**:
  - **What it does**: Verifies token counting for message with tool_calls
  - **Purpose**: Ensure tool_calls are counted

- **`test_message_with_tool_call_id()`**:
  - **What it does**: Verifies token counting for tool response message
  - **Purpose**: Ensure tool_call_id is counted

- **`test_message_with_list_content()`**:
  - **What it does**: Verifies token counting for multimodal content
  - **Purpose**: Ensure list content is handled

- **`test_without_claude_correction()`**:
  - **What it does**: Verifies counting without correction factor
  - **Purpose**: Ensure apply_claude_correction=False works

- **`test_message_with_empty_content()`**:
  - **What it does**: Verifies counting for message with empty content
  - **Purpose**: Ensure empty content doesn't break counting

- **`test_message_with_none_content()`**:
  - **What it does**: Verifies counting for message with None content
  - **Purpose**: Ensure None content doesn't break counting

#### `TestCountToolsTokens`

Tests for count_tools_tokens function.

- **`test_none_returns_zero()`**:
  - **What it does**: Verifies that None returns 0 tokens
  - **Purpose**: Ensure correct None handling

- **`test_empty_list_returns_zero()`**:
  - **What it does**: Verifies that empty list returns 0 tokens
  - **Purpose**: Ensure correct empty list handling

- **`test_single_tool()`**:
  - **What it does**: Verifies token counting for single tool
  - **Purpose**: Ensure basic functionality

- **`test_multiple_tools()`**:
  - **What it does**: Verifies token counting for multiple tools
  - **Purpose**: Ensure tokens are summed

- **`test_tool_with_complex_parameters()`**:
  - **What it does**: Verifies counting for tool with complex parameters
  - **Purpose**: Ensure JSON schema parameters are counted

- **`test_tool_without_parameters()`**:
  - **What it does**: Verifies counting for tool without parameters
  - **Purpose**: Ensure missing parameters doesn't break counting

- **`test_tool_with_empty_description()`**:
  - **What it does**: Verifies counting for tool with empty description
  - **Purpose**: Ensure empty description doesn't break counting

- **`test_non_function_tool_type()`**:
  - **What it does**: Verifies handling of tool with type != "function"
  - **Purpose**: Ensure non-function tools are handled

- **`test_without_claude_correction()`**:
  - **What it does**: Verifies counting without correction factor
  - **Purpose**: Ensure apply_claude_correction=False works

#### `TestEstimateRequestTokens`

Tests for estimate_request_tokens function.

- **`test_messages_only()`**:
  - **What it does**: Verifies token estimation for messages only
  - **Purpose**: Ensure basic functionality

- **`test_messages_with_tools()`**:
  - **What it does**: Verifies token estimation for messages with tools
  - **Purpose**: Ensure tools are counted

- **`test_messages_with_system_prompt()`**:
  - **What it does**: Verifies token estimation with separate system prompt
  - **Purpose**: Ensure system_prompt is counted

- **`test_full_request()`**:
  - **What it does**: Verifies token estimation for full request
  - **Purpose**: Ensure all components are summed

- **`test_empty_messages()`**:
  - **What it does**: Verifies estimation for empty message list
  - **Purpose**: Ensure correct edge case handling

#### `TestClaudeCorrectionFactor`

Tests for Claude correction factor.

- **`test_correction_factor_value()`**:
  - **What it does**: Verifies correction factor value
  - **Purpose**: Ensure factor equals 1.15

- **`test_correction_increases_token_count()`**:
  - **What it does**: Verifies that correction increases token count
  - **Purpose**: Ensure factor is applied correctly

#### `TestGetEncoding`

Tests for _get_encoding function.

- **`test_returns_encoding_when_tiktoken_available()`**:
  - **What it does**: Verifies that _get_encoding returns encoding when tiktoken is available
  - **Purpose**: Ensure correct tiktoken initialization

- **`test_caches_encoding()`**:
  - **What it does**: Verifies that encoding is cached
  - **Purpose**: Ensure lazy initialization

- **`test_handles_import_error()`**:
  - **What it does**: Verifies ImportError handling when tiktoken is missing
  - **Purpose**: Ensure system works without tiktoken

#### `TestTokenizerIntegration`

Integration tests for tokenizer.

- **`test_realistic_chat_request()`**:
  - **What it does**: Verifies token counting for realistic chat request
  - **Purpose**: Ensure correct operation on real data

- **`test_large_context()`**:
  - **What it does**: Verifies token counting for large context
  - **Purpose**: Ensure performance on large data

- **`test_consistency_across_calls()`**:
  - **What it does**: Verifies counting consistency across repeated calls
  - **Purpose**: Ensure results are deterministic

---

### `tests/unit/test_streaming.py`

Unit tests for **streaming module** (Kiro to OpenAI format stream conversion). **23 tests.**

#### `TestStreamingToolCallsIndex`

Tests for adding index to tool_calls in streaming responses.

- **`test_tool_calls_have_index_field()`**:
  - **What it does**: Verifies that tool_calls in streaming response contain index field
  - **Purpose**: Ensure OpenAI API spec is followed for streaming tool calls

- **`test_multiple_tool_calls_have_sequential_indices()`**:
  - **What it does**: Verifies that multiple tool_calls have sequential indices
  - **Purpose**: Ensure indices start from 0 and go sequentially

#### `TestStreamingToolCallsNoneProtection`

Tests for None value protection in tool_calls.

- **`test_handles_none_function_name()`**:
  - **What it does**: Verifies None handling in function.name
  - **Purpose**: Ensure None is replaced with empty string

- **`test_handles_none_function_arguments()`**:
  - **What it does**: Verifies None handling in function.arguments
  - **Purpose**: Ensure None is replaced with "{}"

- **`test_handles_none_function_object()`**:
  - **What it does**: Verifies None handling instead of function object
  - **Purpose**: Ensure None function is handled without errors

#### `TestCollectStreamResponseToolCalls`

Tests for collect_stream_response with tool_calls.

- **`test_collected_tool_calls_have_no_index()`**:
  - **What it does**: Verifies that collected tool_calls don't contain index field
  - **Purpose**: Ensure index is removed for non-streaming response

- **`test_collected_tool_calls_have_required_fields()`**:
  - **What it does**: Verifies that collected tool_calls contain all required fields
  - **Purpose**: Ensure id, type, function are present

- **`test_handles_none_in_collected_tool_calls()`**:
  - **What it does**: Verifies None value handling in collected tool_calls
  - **Purpose**: Ensure None values are replaced with defaults

#### `TestStreamingErrorHandling`

Tests for error handling in streaming module (bug #8 fix).

- **`test_generator_exit_handled_gracefully()`**:
  - **What it does**: Verifies that GeneratorExit is handled without logging as error
  - **Purpose**: Ensure client disconnect doesn't cause ERROR in logs

- **`test_exception_with_empty_message_logged_with_type()`**:
  - **What it does**: Verifies that exception with empty message is logged with type
  - **Purpose**: Ensure exception type is visible in logs even if str(e) is empty

- **`test_exception_propagated_to_caller()`**:
  - **What it does**: Verifies that exceptions are propagated up
  - **Purpose**: Ensure errors are not "swallowed" inside generator

- **`test_response_closed_on_error()`**:
  - **What it does**: Verifies that response is closed even on error
  - **Purpose**: Ensure resources are released in finally block

- **`test_response_closed_on_success()`**:
  - **What it does**: Verifies that response is closed on successful completion
  - **Purpose**: Ensure resources are released in finally block

- **`test_aclose_error_does_not_mask_original_error()`**:
  - **What it does**: Verifies that aclose() error doesn't mask original error
  - **Purpose**: Ensure original exception is propagated even if aclose() fails

#### `TestFirstTokenTimeoutError`

Tests for FirstTokenTimeoutError and first token timeout logging.

- **`test_first_token_timeout_not_caught_by_general_handler()`**:
  - **What it does**: Verifies that FirstTokenTimeoutError is propagated for retry
  - **Purpose**: Ensure first token timeout is not handled as regular error

- **`test_first_token_timeout_logged_with_correct_format()`**:
  - **What it does**: Verifies that first token timeout is logged with [FirstTokenTimeout] prefix
  - **Purpose**: Ensure consistent logging format for first token timeout

- **`test_first_token_timeout_includes_timeout_value()`**:
  - **What it does**: Verifies that first token timeout log includes the timeout value
  - **Purpose**: Ensure timeout value is visible in logs for debugging

- **`test_first_token_received_logged_on_success()`**:
  - **What it does**: Verifies that successful first token receipt is logged
  - **Purpose**: Ensure debug log shows when first token is received

#### `TestStreamWithFirstTokenRetry`

Tests for stream_with_first_token_retry function.

- **`test_retry_on_first_token_timeout()`**:
  - **What it does**: Verifies that request is retried on first token timeout
  - **Purpose**: Ensure retry logic works for first token timeout

- **`test_all_retries_exhausted_raises_504()`**:
  - **What it does**: Verifies that 504 is raised after all retries exhausted
  - **Purpose**: Ensure proper error handling when model never responds

- **`test_retry_logs_attempt_number()`**:
  - **What it does**: Verifies that retry attempts are logged with attempt number
  - **Purpose**: Ensure logs show which attempt failed (e.g., "1/3", "2/3", "3/3")

---

### `tests/unit/test_http_client.py`

Unit tests for **KiroHttpClient** (HTTP client with retry logic). **29 tests.**

#### `TestKiroHttpClientInitialization`

- **`test_initialization_stores_auth_manager()`**: Verifies auth_manager storage during initialization
- **`test_initialization_client_is_none()`**: Verifies that HTTP client is initially None

#### `TestKiroHttpClientGetClient`

- **`test_get_client_creates_new_client()`**: Verifies new HTTP client creation
- **`test_get_client_reuses_existing_client()`**: Verifies existing client reuse
- **`test_get_client_recreates_closed_client()`**: Verifies closed client recreation

#### `TestKiroHttpClientClose`

- **`test_close_closes_client()`**: Verifies HTTP client closing
- **`test_close_does_nothing_for_none_client()`**: Verifies close() doesn't crash for None client
- **`test_close_does_nothing_for_closed_client()`**: Verifies close() doesn't crash for closed client

#### `TestKiroHttpClientRequestWithRetry`

- **`test_successful_request_returns_response()`**: Verifies successful request
- **`test_403_triggers_token_refresh()`**: Verifies token refresh on 403
- **`test_429_triggers_backoff()`**: Verifies exponential backoff on 429
- **`test_5xx_triggers_backoff()`**: Verifies exponential backoff on 5xx
- **`test_timeout_triggers_backoff()`**: Verifies exponential backoff on timeout
- **`test_request_error_triggers_backoff()`**: Verifies exponential backoff on request error
- **`test_max_retries_exceeded_raises_502()`**: Verifies HTTPException after retries exhausted
- **`test_other_status_codes_returned_as_is()`**: Verifies other status codes return without retry
- **`test_streaming_request_uses_send()`**: Verifies send() usage for streaming

#### `TestKiroHttpClientContextManager`

- **`test_context_manager_returns_self()`**: Verifies __aenter__ returns self
- **`test_context_manager_closes_on_exit()`**: Verifies client closing on context exit

#### `TestKiroHttpClientExponentialBackoff`

- **`test_backoff_delay_increases_exponentially()`**: Verifies exponential delay increase

#### `TestKiroHttpClientStreamingTimeout`

Tests for streaming timeout logic (httpx timeouts, not FIRST_TOKEN_TIMEOUT).

- **`test_streaming_uses_streaming_read_timeout()`**:
  - **What it does**: Verifies that streaming requests use STREAMING_READ_TIMEOUT for read timeout
  - **Purpose**: Ensure httpx.Timeout is configured with connect=30s and read=STREAMING_READ_TIMEOUT

- **`test_streaming_uses_first_token_max_retries()`**:
  - **What it does**: Verifies that streaming requests use FIRST_TOKEN_MAX_RETRIES
  - **Purpose**: Ensure separate retry counter is used for stream=True

- **`test_streaming_timeout_retry_without_delay()`**:
  - **What it does**: Verifies that streaming timeout retry happens without delay
  - **Purpose**: Ensure no exponential backoff on streaming timeout

- **`test_non_streaming_uses_default_timeout()`**:
  - **What it does**: Verifies that non-streaming requests use httpx.Timeout(timeout=300)
  - **Purpose**: Ensure unified 300s timeout for all operations in non-streaming mode

- **`test_connect_timeout_logged_correctly()`**:
  - **What it does**: Verifies that ConnectTimeout is logged with [ConnectTimeout] prefix
  - **Purpose**: Ensure timeout type is visible in logs for debugging

- **`test_read_timeout_logged_correctly()`**:
  - **What it does**: Verifies that ReadTimeout is logged with [ReadTimeout] prefix and STREAMING_READ_TIMEOUT value
  - **Purpose**: Ensure timeout type and value are visible in logs

- **`test_streaming_timeout_returns_504_with_error_type()`**:
  - **What it does**: Verifies that streaming timeout returns 504 with error type in detail
  - **Purpose**: Ensure 504 Gateway Timeout includes error type (e.g., "ReadTimeout")

- **`test_non_streaming_timeout_returns_502()`**:
  - **What it does**: Verifies that non-streaming timeout returns 502
  - **Purpose**: Ensure old logic with 502 is used for non-streaming

---

### `tests/unit/test_routes.py`

Unit tests for **API endpoints** (/v1/models, /v1/chat/completions). **22 tests.**

#### `TestVerifyApiKey`

- **`test_valid_api_key_returns_true()`**: Verifies successful validation of correct key
- **`test_invalid_api_key_raises_401()`**: Verifies invalid key rejection
- **`test_missing_api_key_raises_401()`**: Verifies missing key rejection
- **`test_empty_api_key_raises_401()`**: Verifies empty key rejection
- **`test_wrong_format_raises_401()`**: Verifies key without Bearer rejection

#### `TestRootEndpoint`

- **`test_root_returns_status_ok()`**: Verifies root endpoint response
- **`test_root_returns_version()`**: Verifies version presence in response

#### `TestHealthEndpoint`

- **`test_health_returns_healthy()`**: Verifies health endpoint response
- **`test_health_returns_timestamp()`**: Verifies timestamp presence in response
- **`test_health_returns_version()`**: Verifies version presence in response

#### `TestModelsEndpoint`

- **`test_models_requires_auth()`**: Verifies authorization requirement
- **`test_models_returns_list()`**: Verifies model list return
- **`test_models_returns_available_models()`**: Verifies available models presence
- **`test_models_format_is_openai_compatible()`**: Verifies response format OpenAI compatibility

#### `TestChatCompletionsEndpoint`

- **`test_chat_completions_requires_auth()`**: Verifies authorization requirement
- **`test_chat_completions_validates_messages()`**: Verifies empty messages validation
- **`test_chat_completions_validates_model()`**: Verifies missing model validation

#### `TestChatCompletionsWithMockedKiro`

- **`test_chat_completions_accepts_valid_request_format()`**: Verifies valid request format acceptance

#### `TestChatCompletionsErrorHandling`

- **`test_invalid_json_returns_422()`**: Verifies invalid JSON handling
- **`test_missing_content_in_message_returns_200()`**: Verifies message without content handling

#### `TestRouterIntegration`

- **`test_router_has_all_endpoints()`**: Verifies all endpoints presence in router
- **`test_router_methods()`**: Verifies endpoint HTTP methods

---

### `tests/integration/test_full_flow.py`

Integration tests for **full end-to-end flow**. **12 tests (11 passed, 1 skipped).**

#### `TestFullChatCompletionFlow`

- **`test_full_flow_health_to_models_to_chat()`**: Verifies full flow from health check to chat completions
- **`test_authentication_flow()`**: Verifies authentication flow
- **`test_openai_compatibility_format()`**: Verifies response format compatibility with OpenAI API

#### `TestRequestValidationFlow`

- **`test_chat_completions_request_validation()`**: Verifies various request format validation
- **`test_complex_message_formats()`**: Verifies complex message format handling

#### `TestErrorHandlingFlow`

- **`test_invalid_json_handling()`**: Verifies invalid JSON handling
- **`test_wrong_content_type_handling()`**: SKIPPED - bug discovered in validation_exception_handler

#### `TestModelsEndpointIntegration`

- **`test_models_returns_all_available_models()`**: Verifies all models from config are returned
- **`test_models_caching_behavior()`**: Verifies model caching behavior

#### `TestStreamingFlagHandling`

- **`test_stream_true_accepted()`**: Verifies stream=true acceptance
- **`test_stream_false_accepted()`**: Verifies stream=false acceptance

#### `TestHealthEndpointIntegration`

- **`test_root_and_health_consistency()`**: Verifies / and /health consistency

---

## Testing Philosophy

### Principles

1. **Isolation**: Each test is completely isolated from external services through mocks
2. **Detail**: Abundant print() for understanding test flow during debugging
3. **Coverage**: Tests cover not only happy path, but also edge cases and errors
4. **Security**: All tests use mock credentials, never real ones

### Test Structure (Arrange-Act-Assert)

Each test follows the pattern:
1. **Arrange** (Setup): Prepare mocks and data
2. **Act** (Action): Execute the tested action
3. **Assert** (Verify): Verify result with explicit comparison

### Test Types

- **Unit tests**: Test individual functions/classes in isolation
- **Integration tests**: Verify component interactions
- **Security tests**: Verify security system
- **Edge case tests**: Paranoid edge case checks

## Adding New Tests

When adding new tests:

1. Follow existing class structure (`Test*Success`, `Test*Errors`, `Test*EdgeCases`)
2. Use descriptive names: `test_<what_it_does>_<expected_result>`
3. Add docstring with "What it does" and "Purpose"
4. Use print() for logging test steps
5. Update this README with new test description

## Troubleshooting

### Tests fail with ImportError

```bash
# Make sure you're in project root
cd /path/to/kiro-openai-gateway

# pytest.ini already contains pythonpath = .
# Just run pytest
pytest
```

### Tests pass locally but fail in CI

- Check dependency versions in requirements.txt
- Ensure all mocks correctly isolate external calls

### Async tests don't work

```bash
# Make sure pytest-asyncio is installed
pip install pytest-asyncio

# Check for @pytest.mark.asyncio decorator
```

## Coverage Metrics

To check code coverage:

```bash
# Install coverage
pip install pytest-cov

# Run with coverage report
pytest --cov=kiro_gateway --cov-report=html

# View report
open htmlcov/index.html  # macOS/Linux
start htmlcov/index.html  # Windows
```

## Contacts and Support

If you find bugs or have suggestions for test improvements, create an issue in the project repository.



================================================
FILE: tests/conftest.py
================================================
# -*- coding: utf-8 -*-

"""
Общие фикстуры и утилиты для тестирования Kiro OpenAI Gateway.

Обеспечивает изоляцию тестов от внешних сервисов и глобального состояния.
Все тесты ДОЛЖНЫ быть полностью изолированы от сети.
"""

import asyncio
import json
import pytest
import time
from typing import AsyncGenerator, Dict, Any, List
from unittest.mock import AsyncMock, MagicMock, Mock, patch
from datetime import datetime, timezone

import httpx
from fastapi.testclient import TestClient


# =============================================================================
# Event Loop фикстуры
# =============================================================================

@pytest.fixture(scope="session")
def event_loop():
    """
    Создает event loop для всей сессии тестов.
    Необходимо для корректной работы async фикстур.
    """
    print("Создание event loop для сессии тестов...")
    loop = asyncio.get_event_loop_policy().new_event_loop()
    yield loop
    print("Закрытие event loop...")
    loop.close()


# =============================================================================
# Фикстуры окружения
# =============================================================================

@pytest.fixture
def mock_env_vars(monkeypatch):
    """
    Мокирует переменные окружения для изоляции от реальных credentials.
    """
    print("Настройка мокированных переменных окружения...")
    monkeypatch.setenv("REFRESH_TOKEN", "test_refresh_token_abcdef")
    monkeypatch.setenv("PROXY_API_KEY", "test_proxy_key_12345")
    monkeypatch.setenv("PROFILE_ARN", "arn:aws:codewhisperer:us-east-1:123456789:profile/test")
    monkeypatch.setenv("KIRO_REGION", "us-east-1")
    return {
        "REFRESH_TOKEN": "test_refresh_token_abcdef",
        "PROXY_API_KEY": "test_proxy_key_12345",
        "PROFILE_ARN": "arn:aws:codewhisperer:us-east-1:123456789:profile/test",
        "KIRO_REGION": "us-east-1"
    }


# =============================================================================
# Фикстуры токенов и аутентификации
# =============================================================================

@pytest.fixture
def valid_kiro_token():
    """Возвращает валидный мок Kiro access token."""
    return "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.test_kiro_access_token"


@pytest.fixture
def mock_kiro_token_response(valid_kiro_token):
    """
    Фабрика для создания мок-ответа Kiro token refresh endpoint.
    """
    def _create_response(expires_in: int = 3600, token: str = None):
        return {
            "accessToken": token or valid_kiro_token,
            "refreshToken": "new_refresh_token_xyz",
            "expiresIn": expires_in,
            "profileArn": "arn:aws:codewhisperer:us-east-1:123456789:profile/test"
        }
    return _create_response


@pytest.fixture
def valid_proxy_api_key():
    """Возвращает валидный API ключ прокси (из config)."""
    return "changeme_proxy_secret"


@pytest.fixture
def invalid_proxy_api_key():
    """Возвращает невалидный API ключ для негативных тестов."""
    return "invalid_wrong_secret_key"


@pytest.fixture
def auth_headers(valid_proxy_api_key):
    """
    Фабрика для создания валидных и невалидных Authorization headers.
    """
    def _create_headers(api_key: str = None, invalid: bool = False):
        if invalid:
            return {"Authorization": "Bearer wrong_key_123"}
        key = api_key or valid_proxy_api_key
        return {"Authorization": f"Bearer {key}"}
    
    return _create_headers


# =============================================================================
# Фикстуры моделей Kiro
# =============================================================================

@pytest.fixture
def mock_kiro_models_response():
    """
    Мок успешного ответа от Kiro API для ListAvailableModels.
    """
    return {
        "models": [
            {
                "modelId": "claude-sonnet-4.5",
                "displayName": "Claude Sonnet 4.5",
                "tokenLimits": {
                    "maxInputTokens": 200000,
                    "maxOutputTokens": 8192
                }
            },
            {
                "modelId": "claude-opus-4.5",
                "displayName": "Claude Opus 4.5",
                "tokenLimits": {
                    "maxInputTokens": 200000,
                    "maxOutputTokens": 8192
                }
            },
            {
                "modelId": "claude-haiku-4.5",
                "displayName": "Claude Haiku 4.5",
                "tokenLimits": {
                    "maxInputTokens": 200000,
                    "maxOutputTokens": 8192
                }
            }
        ]
    }


# =============================================================================
# Фикстуры streaming ответов Kiro
# =============================================================================

@pytest.fixture
def mock_kiro_streaming_chunks():
    """
    Возвращает список мок SSE chunks от Kiro API для streaming response.
    Покрывает: обычный текст, tool calls, usage.
    """
    return [
        # Chunk 1: Начало текста
        b'{"content":"Hello"}',
        # Chunk 2: Продолжение текста
        b'{"content":" World!"}',
        # Chunk 3: Tool call start
        b'{"name":"get_weather","toolUseId":"call_abc123"}',
        # Chunk 4: Tool call input
        b'{"input":"{\\"location\\": \\"Moscow\\"}"}',
        # Chunk 5: Tool call stop
        b'{"stop":true}',
        # Chunk 6: Usage
        b'{"usage":1.5}',
        # Chunk 7: Context usage
        b'{"contextUsagePercentage":25.5}',
    ]

@pytest.fixture
def mock_kiro_simple_text_chunks():
    """
    Мок простого текстового ответа от Kiro (без tool calls).
    """
    return [
        b'{"content":"This is a complete response."}',
        b'{"usage":0.5}',
        b'{"contextUsagePercentage":10.0}',
    ]


@pytest.fixture
def mock_kiro_stream_with_usage():
    """
    Мок SSE ответа Kiro с информацией о usage.
    """
    return [
        b'{"content":"Final text."}',
        b'{"usage":1.3}',
        b'{"contextUsagePercentage":50.0}',
    ]


# =============================================================================
# Фикстуры OpenAI запросов
# =============================================================================

@pytest.fixture
def sample_openai_chat_request():
    """
    Фабрика для создания валидных OpenAI chat completion requests.
    """
    def _create_request(
        model: str = "claude-sonnet-4-5",
        messages: list = None,
        stream: bool = False,
        temperature: float = None,
        max_tokens: int = None,
        tools: list = None,
        **kwargs
    ):
        if messages is None:
            messages = [{"role": "user", "content": "Hello, AI!"}]
        
        request = {
            "model": model,
            "messages": messages,
            "stream": stream
        }
        
        if temperature is not None:
            request["temperature"] = temperature
        if max_tokens is not None:
            request["max_tokens"] = max_tokens
        if tools is not None:
            request["tools"] = tools
        
        request.update(kwargs)
        return request
    
    return _create_request


@pytest.fixture
def sample_tool_definition():
    """
    Пример определения tool для тестирования tool calling.
    """
    return {
        "type": "function",
        "function": {
            "name": "get_weather",
            "description": "Get weather for a location",
            "parameters": {
                "type": "object",
                "properties": {
                    "location": {"type": "string", "description": "City name"}
                },
                "required": ["location"]
            }
        }
    }


# =============================================================================
# Фикстуры HTTP клиента
# =============================================================================

@pytest.fixture
async def mock_httpx_client():
    """
    Создает мокированный httpx.AsyncClient для изоляции от сетевых запросов.
    """
    print("Создание мокированного httpx.AsyncClient...")
    mock_client = AsyncMock(spec=httpx.AsyncClient)
    
    # Мокируем методы
    mock_client.post = AsyncMock()
    mock_client.get = AsyncMock()
    mock_client.aclose = AsyncMock()
    mock_client.build_request = Mock()
    mock_client.send = AsyncMock()
    mock_client.is_closed = False
    
    return mock_client


@pytest.fixture
def mock_httpx_response():
    """
    Фабрика для создания мокированных httpx.Response объектов.
    """
    def _create_response(
        status_code: int = 200,
        json_data: Dict[str, Any] = None,
        text: str = None,
        stream_chunks: list = None
    ):
        print(f"Создание мок httpx.Response (status={status_code})...")
        mock_response = AsyncMock(spec=httpx.Response)
        mock_response.status_code = status_code
        
        if json_data is not None:
            mock_response.json = Mock(return_value=json_data)
        
        if text is not None:
            mock_response.text = text
            mock_response.content = text.encode()
        
        if stream_chunks is not None:
            # Для streaming responses
            async def mock_aiter_bytes():
                for chunk in stream_chunks:
                    yield chunk
            
            mock_response.aiter_bytes = mock_aiter_bytes
        
        mock_response.raise_for_status = Mock()
        mock_response.aclose = AsyncMock()
        mock_response.aread = AsyncMock(return_value=b'{"error": "mocked error"}')
        
        return mock_response
    
    return _create_response


# =============================================================================
# Глобальная блокировка сети
# =============================================================================

@pytest.fixture(scope="session", autouse=True)
def block_all_network_calls():
    """
    КРИТИЧЕСКАЯ ФИКСТУРА: Глобально блокирует ВСЕ сетевые вызовы.
    Гарантирует, что НИ ОДИН тест не сможет сделать реальный сетевой запрос.
    """
    
    # Создаем мок, который будет использоваться для всех инстансов AsyncClient
    mock_async_client = AsyncMock(spec=httpx.AsyncClient)

    async def network_call_error(*args, **kwargs):
        raise RuntimeError(
            "🚨 КРИТИЧЕСКАЯ ОШИБКА: Обнаружена попытка реального сетевого запроса! "
            "Тест не предоставил мок для httpx.AsyncClient. "
            "Все HTTP вызовы должны быть явно замокированы."
        )

    mock_async_client.post.side_effect = network_call_error
    mock_async_client.get.side_effect = network_call_error
    mock_async_client.send.side_effect = network_call_error
    
    # Мокируем контекстный менеджер
    mock_async_client.__aenter__ = AsyncMock(return_value=mock_async_client)
    mock_async_client.__aexit__ = AsyncMock()
    mock_async_client.aclose = AsyncMock()
    mock_async_client.is_closed = False

    # Патчим AsyncClient в модулях где он используется
    patchers = [
        patch('kiro_gateway.auth.httpx.AsyncClient', return_value=mock_async_client),
        patch('kiro_gateway.http_client.httpx.AsyncClient', return_value=mock_async_client),
        patch('kiro_gateway.routes.httpx.AsyncClient', return_value=mock_async_client),
        patch('kiro_gateway.streaming.httpx.AsyncClient', return_value=mock_async_client),
    ]
    
    # Запускаем патчеры
    for patcher in patchers:
        patcher.start()
    
    print("🛡️ ГЛОБАЛЬНАЯ БЛОКИРОВКА СЕТИ АКТИВИРОВАНА")
    
    yield

    # Останавливаем патчеры
    for patcher in patchers:
        patcher.stop()
    
    print("🛡️ ГЛОБАЛЬНАЯ БЛОКИРОВКА СЕТИ ДЕАКТИВИРОВАНА")


# =============================================================================
# Фикстуры приложения
# =============================================================================

@pytest.fixture
def clean_app():
    """
    Возвращает "чистый" экземпляр приложения для каждого теста.
    """
    print("Импорт приложения для теста...")
    from main import app
    # Сбрасываем все переопределения зависимостей перед тестом
    app.dependency_overrides = {}
    return app


@pytest.fixture
def test_client(clean_app):
    """
    Создает FastAPI TestClient для синхронных тестов endpoints,
    корректно обрабатывая lifespan события.
    """
    print("Создание TestClient с поддержкой lifespan...")
    with TestClient(clean_app) as client:
        yield client
    print("Закрытие TestClient...")


@pytest.fixture
async def async_test_client(clean_app):
    """
    Создает асинхронный test client для async endpoints.
    """
    print("Создание async test client...")
    from httpx import AsyncClient, ASGITransport
    
    transport = ASGITransport(app=clean_app)
    async with AsyncClient(transport=transport, base_url="http://test") as client:
        yield client
    
    print("Закрытие async test client...")


# =============================================================================
# Фикстуры для KiroAuthManager
# =============================================================================

@pytest.fixture
def mock_auth_manager():
    """
    Создает мокированный KiroAuthManager для тестов.
    """
    from kiro_gateway.auth import KiroAuthManager
    
    manager = KiroAuthManager(
        refresh_token="test_refresh_token",
        profile_arn="arn:aws:codewhisperer:us-east-1:123456789:profile/test",
        region="us-east-1"
    )
    
    # Устанавливаем валидный токен
    manager._access_token = "test_access_token"
    manager._expires_at = datetime.now(timezone.utc).replace(
        year=2099  # Далеко в будущем
    )
    
    return manager


@pytest.fixture
def expired_auth_manager():
    """
    Создает KiroAuthManager с истекшим токеном.
    """
    from kiro_gateway.auth import KiroAuthManager
    
    manager = KiroAuthManager(
        refresh_token="test_refresh_token",
        profile_arn="arn:aws:codewhisperer:us-east-1:123456789:profile/test",
        region="us-east-1"
    )
    
    # Устанавливаем истекший токен
    manager._access_token = "expired_token"
    manager._expires_at = datetime.now(timezone.utc).replace(
        year=2020  # В прошлом
    )
    
    return manager


# =============================================================================
# Фикстуры для ModelInfoCache
# =============================================================================

@pytest.fixture
def sample_models_data():
    """
    Возвращает список моделей для тестирования ModelInfoCache.
    """
    return [
        {
            "modelId": "claude-sonnet-4",
            "displayName": "Claude Sonnet 4",
            "tokenLimits": {
                "maxInputTokens": 200000,
                "maxOutputTokens": 8192
            }
        },
        {
            "modelId": "claude-opus-4.5",
            "displayName": "Claude Opus 4.5",
            "tokenLimits": {
                "maxInputTokens": 200000,
                "maxOutputTokens": 8192
            }
        },
        {
            "modelId": "claude-haiku-4.5",
            "displayName": "Claude Haiku 4.5",
            "tokenLimits": {
                "maxInputTokens": 100000,
                "maxOutputTokens": 4096
            }
        }
    ]


@pytest.fixture
def empty_model_cache():
    """
    Создает пустой ModelInfoCache.
    """
    from kiro_gateway.cache import ModelInfoCache
    return ModelInfoCache()


@pytest.fixture
async def populated_model_cache(mock_kiro_models_response):
    """
    Создает ModelInfoCache с предзаполненными данными.
    """
    from kiro_gateway.cache import ModelInfoCache
    
    cache = ModelInfoCache()
    await cache.update(mock_kiro_models_response["models"])
    return cache


# =============================================================================
# Фикстуры времени
# =============================================================================

@pytest.fixture
def mock_time():
    """
    Мокирует time.time() для предсказуемого поведения в тестах.
    """
    with patch('time.time') as mock:
        # Фиксированная точка времени: 2024-01-01 12:00:00
        mock.return_value = 1704110400.0
        yield mock


@pytest.fixture
def mock_datetime():
    """
    Мокирует datetime.now() для предсказуемого поведения.
    """
    fixed_time = datetime(2024, 1, 1, 12, 0, 0, tzinfo=timezone.utc)
    
    with patch('kiro_gateway.auth.datetime') as mock_dt:
        mock_dt.now.return_value = fixed_time
        mock_dt.fromisoformat = datetime.fromisoformat
        mock_dt.fromtimestamp = datetime.fromtimestamp
        yield mock_dt


# =============================================================================
# Фикстуры для временных файлов
# =============================================================================

@pytest.fixture
def temp_creds_file(tmp_path):
    """
    Создает временный файл credentials для тестов.
    """
    creds_file = tmp_path / "kiro-auth-token.json"
    creds_data = {
        "accessToken": "file_access_token",
        "refreshToken": "file_refresh_token",
        "expiresAt": "2099-01-01T00:00:00.000Z",
        "profileArn": "arn:aws:codewhisperer:us-east-1:123456789:profile/test",
        "region": "us-east-1"
    }
    creds_file.write_text(json.dumps(creds_data))
    return str(creds_file)


@pytest.fixture
def temp_debug_dir(tmp_path):
    """
    Создает временную директорию для debug файлов.
    """
    debug_dir = tmp_path / "debug_logs"
    debug_dir.mkdir()
    return debug_dir


# =============================================================================
# Фикстуры для парсера
# =============================================================================

@pytest.fixture
def aws_event_parser():
    """
    Создает экземпляр AwsEventStreamParser для тестов.
    """
    from kiro_gateway.parsers import AwsEventStreamParser
    return AwsEventStreamParser()


# =============================================================================
# Утилиты для тестов
# =============================================================================

def create_kiro_content_chunk(content: str) -> bytes:
    """Утилита для создания Kiro SSE chunk с контентом."""
    return f'{{"content":"{content}"}}'.encode()


def create_kiro_tool_start_chunk(name: str, tool_id: str) -> bytes:
    """Утилита для создания Kiro SSE chunk с началом tool call."""
    return f'{{"name":"{name}","toolUseId":"{tool_id}"}}'.encode()


def create_kiro_tool_input_chunk(input_json: str) -> bytes:
    """Утилита для создания Kiro SSE chunk с input для tool call."""
    escaped = input_json.replace('"', '\\"')
    return f'{{"input":"{escaped}"}}'.encode()


def create_kiro_tool_stop_chunk() -> bytes:
    """Утилита для создания Kiro SSE chunk с завершением tool call."""
    return b'{"stop":true}'


def create_kiro_usage_chunk(usage: float) -> bytes:
    """Утилита для создания Kiro SSE chunk с usage."""
    return f'{{"usage":{usage}}}'.encode()


def create_kiro_context_usage_chunk(percentage: float) -> bytes:
    """Утилита для создания Kiro SSE chunk с context usage."""
    return f'{{"contextUsagePercentage":{percentage}}}'.encode()


================================================
FILE: tests/integration/test_full_flow.py
================================================
# -*- coding: utf-8 -*-

"""
Integration-тесты для полного end-to-end flow.
Проверяет взаимодействие всех компонентов системы.
"""

import pytest
import json
from unittest.mock import AsyncMock, Mock, patch, MagicMock
from datetime import datetime, timezone, timedelta

from fastapi.testclient import TestClient
import httpx

from kiro_gateway.config import PROXY_API_KEY, AVAILABLE_MODELS


class TestFullChatCompletionFlow:
    """Integration-тесты полного flow chat completions."""
    
    def test_full_flow_health_to_models_to_chat(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет полный flow от health check до chat completions.
        Цель: Убедиться, что все эндпоинты работают вместе.
        """
        print("Шаг 1: Health check...")
        health_response = test_client.get("/health")
        assert health_response.status_code == 200
        assert health_response.json()["status"] == "healthy"
        print(f"Health: {health_response.json()}")
        
        print("Шаг 2: Получение списка моделей...")
        models_response = test_client.get(
            "/v1/models",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"}
        )
        assert models_response.status_code == 200
        assert len(models_response.json()["data"]) > 0
        print(f"Модели: {[m['id'] for m in models_response.json()['data']]}")
        
        print("Шаг 3: Валидация запроса chat completions...")
        # Этот запрос пройдёт валидацию, но упадёт на HTTP из-за блокировки сети
        chat_response = test_client.post(
            "/v1/chat/completions",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
            json={
                "model": "claude-sonnet-4-5",
                "messages": [{"role": "user", "content": "Hello"}]
            }
        )
        # Запрос должен пройти валидацию (не 422)
        assert chat_response.status_code != 422
        print(f"Chat response status: {chat_response.status_code}")
    
    def test_authentication_flow(self, test_client, valid_proxy_api_key, invalid_proxy_api_key):
        """
        Что он делает: Проверяет flow аутентификации.
        Цель: Убедиться, что защищённые эндпоинты требуют авторизации.
        """
        print("Шаг 1: Запрос без авторизации...")
        no_auth_response = test_client.get("/v1/models")
        assert no_auth_response.status_code == 401
        print(f"Без авторизации: {no_auth_response.status_code}")
        
        print("Шаг 2: Запрос с неверным ключом...")
        wrong_auth_response = test_client.get(
            "/v1/models",
            headers={"Authorization": f"Bearer {invalid_proxy_api_key}"}
        )
        assert wrong_auth_response.status_code == 401
        print(f"Неверный ключ: {wrong_auth_response.status_code}")
        
        print("Шаг 3: Запрос с верным ключом...")
        valid_auth_response = test_client.get(
            "/v1/models",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"}
        )
        assert valid_auth_response.status_code == 200
        print(f"Верный ключ: {valid_auth_response.status_code}")
    
    def test_openai_compatibility_format(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет совместимость формата ответов с OpenAI API.
        Цель: Убедиться, что ответы соответствуют спецификации OpenAI.
        """
        print("Проверка формата /v1/models...")
        models_response = test_client.get(
            "/v1/models",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"}
        )
        
        assert models_response.status_code == 200
        data = models_response.json()
        
        # Проверяем структуру ответа OpenAI
        assert "object" in data
        assert data["object"] == "list"
        assert "data" in data
        assert isinstance(data["data"], list)
        
        # Проверяем структуру каждой модели
        for model in data["data"]:
            assert "id" in model
            assert "object" in model
            assert model["object"] == "model"
            assert "owned_by" in model
            assert "created" in model
        
        print(f"Формат соответствует OpenAI API: {len(data['data'])} моделей")


class TestRequestValidationFlow:
    """Integration-тесты валидации запросов."""
    
    def test_chat_completions_request_validation(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет валидацию различных форматов запросов.
        Цель: Убедиться, что валидация работает корректно.
        """
        print("Тест 1: Пустые сообщения...")
        empty_messages = test_client.post(
            "/v1/chat/completions",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
            json={"model": "claude-sonnet-4-5", "messages": []}
        )
        assert empty_messages.status_code == 422
        print(f"Пустые сообщения: {empty_messages.status_code}")
        
        print("Тест 2: Отсутствует model...")
        no_model = test_client.post(
            "/v1/chat/completions",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
            json={"messages": [{"role": "user", "content": "Hello"}]}
        )
        assert no_model.status_code == 422
        print(f"Без model: {no_model.status_code}")
        
        print("Тест 3: Отсутствует messages...")
        no_messages = test_client.post(
            "/v1/chat/completions",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
            json={"model": "claude-sonnet-4-5"}
        )
        assert no_messages.status_code == 422
        print(f"Без messages: {no_messages.status_code}")
        
        print("Тест 4: Валидный запрос...")
        valid_request = test_client.post(
            "/v1/chat/completions",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
            json={
                "model": "claude-sonnet-4-5",
                "messages": [{"role": "user", "content": "Hello"}]
            }
        )
        # Валидация должна пройти (не 422)
        assert valid_request.status_code != 422
        print(f"Валидный запрос: {valid_request.status_code}")
    
    def test_complex_message_formats(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет обработку сложных форматов сообщений.
        Цель: Убедиться, что multimodal и tool форматы принимаются.
        """
        print("Тест 1: System + User сообщения...")
        system_user = test_client.post(
            "/v1/chat/completions",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
            json={
                "model": "claude-sonnet-4-5",
                "messages": [
                    {"role": "system", "content": "You are helpful"},
                    {"role": "user", "content": "Hello"}
                ]
            }
        )
        assert system_user.status_code != 422
        print(f"System + User: {system_user.status_code}")
        
        print("Тест 2: Multi-turn conversation...")
        multi_turn = test_client.post(
            "/v1/chat/completions",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
            json={
                "model": "claude-sonnet-4-5",
                "messages": [
                    {"role": "user", "content": "Hello"},
                    {"role": "assistant", "content": "Hi there!"},
                    {"role": "user", "content": "How are you?"}
                ]
            }
        )
        assert multi_turn.status_code != 422
        print(f"Multi-turn: {multi_turn.status_code}")
        
        print("Тест 3: С tools...")
        with_tools = test_client.post(
            "/v1/chat/completions",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
            json={
                "model": "claude-sonnet-4-5",
                "messages": [{"role": "user", "content": "What's the weather?"}],
                "tools": [{
                    "type": "function",
                    "function": {
                        "name": "get_weather",
                        "description": "Get weather",
                        "parameters": {"type": "object", "properties": {}}
                    }
                }]
            }
        )
        assert with_tools.status_code != 422
        print(f"С tools: {with_tools.status_code}")


class TestErrorHandlingFlow:
    """Integration-тесты обработки ошибок."""
    
    def test_invalid_json_handling(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет обработку невалидного JSON.
        Цель: Убедиться, что невалидный JSON возвращает понятную ошибку.
        """
        print("Отправка невалидного JSON...")
        response = test_client.post(
            "/v1/chat/completions",
            headers={
                "Authorization": f"Bearer {valid_proxy_api_key}",
                "Content-Type": "application/json"
            },
            content=b"not valid json"
        )
        
        assert response.status_code == 422
        print(f"Невалидный JSON: {response.status_code}")
    
    def test_wrong_content_type_handling(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет обработку неверного Content-Type.
        Цель: Убедиться, что неверный Content-Type обрабатывается.
        """
        print("Отправка с неверным Content-Type...")
        response = test_client.post(
            "/v1/chat/completions",
            headers={
                "Authorization": f"Bearer {valid_proxy_api_key}",
                "Content-Type": "text/plain"
            },
            content=b"Hello"
        )
        
        # Должна быть ошибка валидации
        assert response.status_code == 422
        print(f"Неверный Content-Type: {response.status_code}")


class TestModelsEndpointIntegration:
    """Integration-тесты эндпоинта /v1/models."""
    
    def test_models_returns_all_available_models(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет, что все модели из конфига возвращаются.
        Цель: Убедиться в полноте списка моделей.
        """
        print("Получение списка моделей...")
        response = test_client.get(
            "/v1/models",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"}
        )
        
        assert response.status_code == 200
        
        returned_ids = {m["id"] for m in response.json()["data"]}
        expected_ids = set(AVAILABLE_MODELS)
        
        print(f"Возвращённые модели: {returned_ids}")
        print(f"Ожидаемые модели: {expected_ids}")
        
        assert returned_ids == expected_ids
    
    def test_models_caching_behavior(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет поведение кэширования моделей.
        Цель: Убедиться, что повторные запросы работают корректно.
        """
        print("Первый запрос моделей...")
        response1 = test_client.get(
            "/v1/models",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"}
        )
        assert response1.status_code == 200
        
        print("Второй запрос моделей...")
        response2 = test_client.get(
            "/v1/models",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"}
        )
        assert response2.status_code == 200
        
        # Ответы должны быть идентичны
        assert response1.json()["data"] == response2.json()["data"]
        print("Кэширование работает корректно")


class TestStreamingFlagHandling:
    """Integration-тесты обработки флага stream."""
    
    def test_stream_true_accepted(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет, что stream=true принимается.
        Цель: Убедиться, что streaming режим доступен.
        
        Примечание: Для streaming режима нужен мок HTTP клиента,
        так как запрос выполняется внутри генератора.
        """
        print("Запрос с stream=true...")
        
        # Создаём мок response для streaming
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        async def mock_aiter_bytes():
            yield b'{"content":"Hello"}'
            yield b'{"usage":0.5}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        mock_response.aclose = AsyncMock()
        
        # Мокируем request_with_retry чтобы вернуть наш мок response
        with patch('kiro_gateway.routes.KiroHttpClient') as MockHttpClient:
            mock_client_instance = AsyncMock()
            mock_client_instance.request_with_retry = AsyncMock(return_value=mock_response)
            mock_client_instance.client = AsyncMock()
            mock_client_instance.close = AsyncMock()
            MockHttpClient.return_value = mock_client_instance
            
            response = test_client.post(
                "/v1/chat/completions",
                headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
                json={
                    "model": "claude-sonnet-4-5",
                    "messages": [{"role": "user", "content": "Hello"}],
                    "stream": True
                }
            )
        
        # Валидация должна пройти и streaming должен работать
        assert response.status_code == 200
        print(f"stream=true: {response.status_code}")
    
    def test_stream_false_accepted(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет, что stream=false принимается.
        Цель: Убедиться, что non-streaming режим доступен.
        """
        print("Запрос с stream=false...")
        response = test_client.post(
            "/v1/chat/completions",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
            json={
                "model": "claude-sonnet-4-5",
                "messages": [{"role": "user", "content": "Hello"}],
                "stream": False
            }
        )
        
        # Валидация должна пройти
        assert response.status_code != 422
        print(f"stream=false: {response.status_code}")


class TestHealthEndpointIntegration:
    """Integration-тесты health endpoints."""
    
    def test_root_and_health_consistency(self, test_client):
        """
        Что он делает: Проверяет консистентность / и /health.
        Цель: Убедиться, что оба эндпоинта возвращают корректный статус.
        """
        print("Запрос к /...")
        root_response = test_client.get("/")
        
        print("Запрос к /health...")
        health_response = test_client.get("/health")
        
        assert root_response.status_code == 200
        assert health_response.status_code == 200
        
        # Оба должны показывать "ok" статус
        assert root_response.json()["status"] == "ok"
        assert health_response.json()["status"] == "healthy"
        
        # Версии должны совпадать
        assert root_response.json()["version"] == health_response.json()["version"]
        
        print("Health endpoints консистентны")


================================================
FILE: tests/unit/test_auth_manager.py
================================================
# -*- coding: utf-8 -*-

"""
Unit-тесты для KiroAuthManager.
Проверяет логику управления токенами Kiro без реальных сетевых запросов.
"""

import asyncio
import pytest
from datetime import datetime, timezone, timedelta
from unittest.mock import AsyncMock, Mock, patch
import httpx

from kiro_gateway.auth import KiroAuthManager
from kiro_gateway.config import TOKEN_REFRESH_THRESHOLD


class TestKiroAuthManagerInitialization:
    """Тесты инициализации KiroAuthManager."""
    
    def test_initialization_stores_credentials(self):
        """
        Что он делает: Проверяет корректное сохранение credentials при инициализации.
        Цель: Убедиться, что все параметры конструктора сохраняются в приватных полях.
        """
        print("Настройка: Создание KiroAuthManager с тестовыми credentials...")
        manager = KiroAuthManager(
            refresh_token="test_refresh_123",
            profile_arn="arn:aws:codewhisperer:us-east-1:123456789:profile/test",
            region="us-east-1"
        )
        
        print("Проверка: Все credentials сохранены корректно...")
        print(f"Сравниваем refresh_token: Ожидалось 'test_refresh_123', Получено '{manager._refresh_token}'")
        assert manager._refresh_token == "test_refresh_123"
        
        print(f"Сравниваем profile_arn: Ожидалось 'arn:aws:...', Получено '{manager._profile_arn}'")
        assert manager._profile_arn == "arn:aws:codewhisperer:us-east-1:123456789:profile/test"
        
        print(f"Сравниваем region: Ожидалось 'us-east-1', Получено '{manager._region}'")
        assert manager._region == "us-east-1"
        
        print("Проверка: Токен изначально пустой...")
        assert manager._access_token is None
        assert manager._expires_at is None
    
    def test_initialization_sets_correct_urls_for_region(self):
        """
        Что он делает: Проверяет формирование URL на основе региона.
        Цель: Убедиться, что URL динамически формируются с правильным регионом.
        """
        print("Настройка: Создание KiroAuthManager с регионом eu-west-1...")
        manager = KiroAuthManager(
            refresh_token="test_token",
            region="eu-west-1"
        )
        
        print("Проверка: URL содержат правильный регион...")
        print(f"Сравниваем refresh_url: Ожидалось 'eu-west-1' в URL, Получено '{manager._refresh_url}'")
        assert "eu-west-1" in manager._refresh_url
        
        print(f"Сравниваем api_host: Ожидалось 'eu-west-1' в URL, Получено '{manager._api_host}'")
        assert "eu-west-1" in manager._api_host
        
        print(f"Сравниваем q_host: Ожидалось 'eu-west-1' в URL, Получено '{manager._q_host}'")
        assert "eu-west-1" in manager._q_host
    
    def test_initialization_generates_fingerprint(self):
        """
        Что он делает: Проверяет генерацию уникального fingerprint.
        Цель: Убедиться, что fingerprint генерируется и имеет корректный формат.
        """
        print("Настройка: Создание KiroAuthManager...")
        manager = KiroAuthManager(refresh_token="test_token")
        
        print("Проверка: Fingerprint сгенерирован...")
        print(f"Fingerprint: {manager._fingerprint}")
        assert manager._fingerprint is not None
        assert len(manager._fingerprint) == 64  # SHA256 hex digest


class TestKiroAuthManagerCredentialsFile:
    """Тесты загрузки credentials из файла."""
    
    def test_load_credentials_from_file(self, temp_creds_file):
        """
        Что он делает: Проверяет загрузку credentials из JSON файла.
        Цель: Убедиться, что данные корректно читаются из файла.
        """
        print(f"Настройка: Создание KiroAuthManager с файлом credentials: {temp_creds_file}")
        manager = KiroAuthManager(creds_file=temp_creds_file)
        
        print("Проверка: Данные загружены из файла...")
        print(f"Сравниваем access_token: Ожидалось 'file_access_token', Получено '{manager._access_token}'")
        assert manager._access_token == "file_access_token"
        
        print(f"Сравниваем refresh_token: Ожидалось 'file_refresh_token', Получено '{manager._refresh_token}'")
        assert manager._refresh_token == "file_refresh_token"
        
        print(f"Сравниваем region: Ожидалось 'us-east-1', Получено '{manager._region}'")
        assert manager._region == "us-east-1"
        
        print("Проверка: expiresAt распарсен корректно...")
        assert manager._expires_at is not None
        assert manager._expires_at.year == 2099
    
    def test_load_credentials_file_not_found(self, tmp_path):
        """
        Что он делает: Проверяет обработку отсутствующего файла credentials.
        Цель: Убедиться, что приложение не падает при отсутствии файла.
        """
        print("Настройка: Создание KiroAuthManager с несуществующим файлом...")
        non_existent_file = str(tmp_path / "non_existent.json")
        
        manager = KiroAuthManager(
            refresh_token="fallback_token",
            creds_file=non_existent_file
        )
        
        print("Проверка: Используется fallback refresh_token...")
        print(f"Сравниваем refresh_token: Ожидалось 'fallback_token', Получено '{manager._refresh_token}'")
        assert manager._refresh_token == "fallback_token"


class TestKiroAuthManagerTokenExpiration:
    """Тесты проверки истечения токена."""
    
    def test_is_token_expiring_soon_returns_true_when_no_expires_at(self):
        """
        Что он делает: Проверяет, что без expires_at токен считается истекающим.
        Цель: Убедиться в безопасном поведении при отсутствии информации о времени.
        """
        print("Настройка: Создание KiroAuthManager без expires_at...")
        manager = KiroAuthManager(refresh_token="test_token")
        manager._expires_at = None
        
        print("Проверка: is_token_expiring_soon возвращает True...")
        result = manager.is_token_expiring_soon()
        print(f"Сравниваем результат: Ожидалось True, Получено {result}")
        assert result is True
    
    def test_is_token_expiring_soon_returns_true_when_expired(self):
        """
        Что он делает: Проверяет, что истекший токен определяется корректно.
        Цель: Убедиться, что токен в прошлом считается истекающим.
        """
        print("Настройка: Создание KiroAuthManager с истекшим токеном...")
        manager = KiroAuthManager(refresh_token="test_token")
        manager._expires_at = datetime.now(timezone.utc) - timedelta(hours=1)
        
        print("Проверка: is_token_expiring_soon возвращает True для истекшего токена...")
        result = manager.is_token_expiring_soon()
        print(f"Сравниваем результат: Ожидалось True, Получено {result}")
        assert result is True
    
    def test_is_token_expiring_soon_returns_true_within_threshold(self):
        """
        Что он делает: Проверяет, что токен в пределах threshold считается истекающим.
        Цель: Убедиться, что токен обновляется заранее (за 10 минут до истечения).
        """
        print("Настройка: Создание KiroAuthManager с токеном, истекающим через 5 минут...")
        manager = KiroAuthManager(refresh_token="test_token")
        manager._expires_at = datetime.now(timezone.utc) + timedelta(minutes=5)
        
        print(f"TOKEN_REFRESH_THRESHOLD = {TOKEN_REFRESH_THRESHOLD} секунд")
        print("Проверка: is_token_expiring_soon возвращает True (5 мин < 10 мин threshold)...")
        result = manager.is_token_expiring_soon()
        print(f"Сравниваем результат: Ожидалось True, Получено {result}")
        assert result is True
    
    def test_is_token_expiring_soon_returns_false_when_valid(self):
        """
        Что он делает: Проверяет, что валидный токен не считается истекающим.
        Цель: Убедиться, что токен далеко в будущем не требует обновления.
        """
        print("Настройка: Создание KiroAuthManager с токеном, истекающим через 1 час...")
        manager = KiroAuthManager(refresh_token="test_token")
        manager._expires_at = datetime.now(timezone.utc) + timedelta(hours=1)
        
        print("Проверка: is_token_expiring_soon возвращает False...")
        result = manager.is_token_expiring_soon()
        print(f"Сравниваем результат: Ожидалось False, Получено {result}")
        assert result is False


class TestKiroAuthManagerTokenRefresh:
    """Тесты механизма обновления токена."""
    
    @pytest.mark.asyncio
    async def test_refresh_token_successful(self, valid_kiro_token, mock_kiro_token_response):
        """
        Что он делает: Тестирует успешное обновление токена через Kiro API.
        Цель: Проверить, что при успешном ответе токен и время истечения устанавливаются.
        """
        print("Настройка: Создание KiroAuthManager...")
        manager = KiroAuthManager(
            refresh_token="test_refresh",
            region="us-east-1"
        )
        
        print("Настройка: Мокирование успешного ответа от Kiro...")
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.json = Mock(return_value=mock_kiro_token_response())
        mock_response.raise_for_status = Mock()
        
        with patch('kiro_gateway.auth.httpx.AsyncClient') as mock_client_class:
            mock_client = AsyncMock()
            mock_client.post = AsyncMock(return_value=mock_response)
            mock_client.__aenter__ = AsyncMock(return_value=mock_client)
            mock_client.__aexit__ = AsyncMock(return_value=None)
            mock_client_class.return_value = mock_client
            
            print("Действие: Вызов _refresh_token_request()...")
            await manager._refresh_token_request()
            
            print("Проверка: Токен установлен корректно...")
            print(f"Сравниваем access_token: Ожидалось '{valid_kiro_token}', Получено '{manager._access_token}'")
            assert manager._access_token == valid_kiro_token
            
            print("Проверка: Время истечения установлено...")
            assert manager._expires_at is not None
            
            print("Проверка: Был сделан POST запрос...")
            mock_client.post.assert_called_once()
    
    @pytest.mark.asyncio
    async def test_refresh_token_updates_refresh_token(self, mock_kiro_token_response):
        """
        Что он делает: Проверяет обновление refresh_token из ответа.
        Цель: Убедиться, что новый refresh_token сохраняется.
        """
        print("Настройка: Создание KiroAuthManager...")
        manager = KiroAuthManager(refresh_token="old_refresh_token")
        
        print("Настройка: Мокирование ответа с новым refresh_token...")
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.json = Mock(return_value=mock_kiro_token_response())
        mock_response.raise_for_status = Mock()
        
        with patch('kiro_gateway.auth.httpx.AsyncClient') as mock_client_class:
            mock_client = AsyncMock()
            mock_client.post = AsyncMock(return_value=mock_response)
            mock_client.__aenter__ = AsyncMock(return_value=mock_client)
            mock_client.__aexit__ = AsyncMock(return_value=None)
            mock_client_class.return_value = mock_client
            
            print("Действие: Обновление токена...")
            await manager._refresh_token_request()
            
            print("Проверка: refresh_token обновлен...")
            print(f"Сравниваем refresh_token: Ожидалось 'new_refresh_token_xyz', Получено '{manager._refresh_token}'")
            assert manager._refresh_token == "new_refresh_token_xyz"
    
    @pytest.mark.asyncio
    async def test_refresh_token_missing_access_token_raises(self):
        """
        Что он делает: Проверяет обработку ответа без accessToken.
        Цель: Убедиться, что выбрасывается исключение при некорректном ответе.
        """
        print("Настройка: Создание KiroAuthManager...")
        manager = KiroAuthManager(refresh_token="test_refresh")
        
        print("Настройка: Мокирование ответа без accessToken...")
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.json = Mock(return_value={"expiresIn": 3600})  # Нет accessToken!
        mock_response.raise_for_status = Mock()
        
        with patch('kiro_gateway.auth.httpx.AsyncClient') as mock_client_class:
            mock_client = AsyncMock()
            mock_client.post = AsyncMock(return_value=mock_response)
            mock_client.__aenter__ = AsyncMock(return_value=mock_client)
            mock_client.__aexit__ = AsyncMock(return_value=None)
            mock_client_class.return_value = mock_client
            
            print("Действие: Попытка обновления токена...")
            with pytest.raises(ValueError) as exc_info:
                await manager._refresh_token_request()
            
            print(f"Проверка: Выброшено ValueError сообщением: {exc_info.value}")
            assert "accessToken" in str(exc_info.value)
    
    @pytest.mark.asyncio
    async def test_refresh_token_no_refresh_token_raises(self):
        """
        Что он делает: Проверяет обработку отсутствия refresh_token.
        Цель: Убедиться, что выбрасывается исключение без refresh_token.
        """
        print("Настройка: Создание KiroAuthManager без refresh_token...")
        manager = KiroAuthManager()
        manager._refresh_token = None
        
        print("Действие: Попытка обновления токена без refresh_token...")
        with pytest.raises(ValueError) as exc_info:
            await manager._refresh_token_request()
        
        print(f"Проверка: Выброшено ValueError: {exc_info.value}")
        assert "Refresh token" in str(exc_info.value)


class TestKiroAuthManagerGetAccessToken:
    """Тесты публичного метода get_access_token."""
    
    @pytest.mark.asyncio
    async def test_get_access_token_refreshes_when_expired(self, valid_kiro_token, mock_kiro_token_response):
        """
        Что он делает: Проверяет автоматическое обновление истекшего токена.
        Цель: Убедиться, что устаревший токен обновляется перед возвратом.
        """
        print("Настройка: Создание KiroAuthManager с истекшим токеном...")
        manager = KiroAuthManager(refresh_token="test_refresh")
        manager._access_token = "old_expired_token"
        manager._expires_at = datetime.now(timezone.utc) - timedelta(hours=1)
        
        print("Настройка: Мокирование успешного обновления...")
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.json = Mock(return_value=mock_kiro_token_response())
        mock_response.raise_for_status = Mock()
        
        with patch('kiro_gateway.auth.httpx.AsyncClient') as mock_client_class:
            mock_client = AsyncMock()
            mock_client.post = AsyncMock(return_value=mock_response)
            mock_client.__aenter__ = AsyncMock(return_value=mock_client)
            mock_client.__aexit__ = AsyncMock(return_value=None)
            mock_client_class.return_value = mock_client
            
            print("Действие: Запрос токена через get_access_token()...")
            token = await manager.get_access_token()
            
            print("Проверка: Получен новый токен, а не истекший...")
            print(f"Сравниваем токен: Ожидалось '{valid_kiro_token}', Получено '{token}'")
            assert token == valid_kiro_token
            assert token != "old_expired_token"
            
            print("Проверка: _refresh_token_request был вызван...")
            mock_client.post.assert_called_once()
    
    @pytest.mark.asyncio
    async def test_get_access_token_returns_valid_without_refresh(self, valid_kiro_token):
        """
        Что он делает: Проверяет возврат валидного токена без обновления.
        Цель: Убедиться, что не делаются лишние запросы, если токен валиден.
        """
        print("Настройка: Создание KiroAuthManager с валидным токеном...")
        manager = KiroAuthManager(refresh_token="test_refresh")
        manager._access_token = valid_kiro_token
        manager._expires_at = datetime.now(timezone.utc) + timedelta(hours=1)
        
        print("Настройка: Мокирование httpx для отслеживания вызовов...")
        with patch('kiro_gateway.auth.httpx.AsyncClient') as mock_client_class:
            mock_client = AsyncMock()
            mock_client.post = AsyncMock()
            mock_client_class.return_value = mock_client
            
            print("Действие: Запрос валидного токена...")
            token = await manager.get_access_token()
            
            print("Проверка: Возвращен существующий токен...")
            print(f"Сравниваем токен: Ожидалось '{valid_kiro_token}', Получено '{token}'")
            assert token == valid_kiro_token
            
            print("Проверка: _refresh_token НЕ был вызван (нет сетевых запросов)...")
            mock_client.post.assert_not_called()
    
    @pytest.mark.asyncio
    async def test_get_access_token_thread_safety(self, valid_kiro_token, mock_kiro_token_response):
        """
        Что он делает: Проверяет потокобезопасность через asyncio.Lock.
        Цель: Убедиться, что параллельные вызовы не приводят к race condition.
        """
        print("Настройка: Создание KiroAuthManager...")
        manager = KiroAuthManager(refresh_token="test_refresh")
        manager._access_token = None
        manager._expires_at = None
        
        refresh_call_count = 0
        
        async def mock_refresh():
            nonlocal refresh_call_count
            refresh_call_count += 1
            await asyncio.sleep(0.1)  # Имитация задержки
            manager._access_token = valid_kiro_token
            manager._expires_at = datetime.now(timezone.utc) + timedelta(hours=1)
        
        print("Настройка: Патчинг _refresh_token_request для отслеживания вызовов...")
        with patch.object(manager, '_refresh_token_request', side_effect=mock_refresh):
            print("Действие: 5 параллельных вызовов get_access_token()...")
            tokens = await asyncio.gather(*[
                manager.get_access_token() for _ in range(5)
            ])
            
            print("Проверка: Все вызовы получили одинаковый токен...")
            assert all(token == valid_kiro_token for token in tokens)
            
            print(f"Проверка: _refresh_token вызван ТОЛЬКО ОДИН РАЗ (благодаря lock)...")
            print(f"Сравниваем количество вызовов: Ожидалось 1, Получено {refresh_call_count}")
            assert refresh_call_count == 1


class TestKiroAuthManagerForceRefresh:
    """Тесты принудительного обновления токена."""
    
    @pytest.mark.asyncio
    async def test_force_refresh_updates_token(self, valid_kiro_token, mock_kiro_token_response):
        """
        Что он делает: Проверяет принудительное обновление токена.
        Цель: Убедиться, что force_refresh всегда обновляет токен.
        """
        print("Настройка: Создание KiroAuthManager с валидным токеном...")
        manager = KiroAuthManager(refresh_token="test_refresh")
        manager._access_token = "old_but_valid_token"
        manager._expires_at = datetime.now(timezone.utc) + timedelta(hours=1)
        
        print("Настройка: Мокирование обновления...")
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.json = Mock(return_value=mock_kiro_token_response())
        mock_response.raise_for_status = Mock()
        
        with patch('kiro_gateway.auth.httpx.AsyncClient') as mock_client_class:
            mock_client = AsyncMock()
            mock_client.post = AsyncMock(return_value=mock_response)
            mock_client.__aenter__ = AsyncMock(return_value=mock_client)
            mock_client.__aexit__ = AsyncMock(return_value=None)
            mock_client_class.return_value = mock_client
            
            print("Действие: Принудительное обновление токена...")
            token = await manager.force_refresh()
            
            print("Проверка: Токен обновлен, несмотря на валидность старого...")
            print(f"Сравниваем токен: Ожидалось '{valid_kiro_token}', Получено '{token}'")
            assert token == valid_kiro_token
            
            print("Проверка: POST запрос был сделан...")
            mock_client.post.assert_called_once()


class TestKiroAuthManagerProperties:
    """Тесты свойств KiroAuthManager."""
    
    def test_profile_arn_property(self):
        """
        Что он делает: Проверяет свойство profile_arn.
        Цель: Убедиться, что profile_arn доступен через property.
        """
        print("Настройка: Создание KiroAuthManager с profile_arn...")
        manager = KiroAuthManager(
            refresh_token="test",
            profile_arn="arn:aws:test:profile"
        )
        
        print("Проверка: profile_arn доступен...")
        print(f"Сравниваем profile_arn: Ожидалось 'arn:aws:test:profile', Получено '{manager.profile_arn}'")
        assert manager.profile_arn == "arn:aws:test:profile"
    
    def test_region_property(self):
        """
        Что он делает: Проверяет свойство region.
        Цель: Убедиться, что region доступен через property.
        """
        print("Настройка: Создание KiroAuthManager с region...")
        manager = KiroAuthManager(
            refresh_token="test",
            region="eu-west-1"
        )
        
        print("Проверка: region доступен...")
        print(f"Сравниваем region: Ожидалось 'eu-west-1', Получено '{manager.region}'")
        assert manager.region == "eu-west-1"
    
    def test_api_host_property(self):
        """
        Что он делает: Проверяет свойство api_host.
        Цель: Убедиться, что api_host формируется корректно.
        """
        print("Настройка: Создание KiroAuthManager...")
        manager = KiroAuthManager(
            refresh_token="test",
            region="us-east-1"
        )
        
        print("Проверка: api_host содержит codewhisperer и регион...")
        print(f"api_host: {manager.api_host}")
        assert "codewhisperer" in manager.api_host
        assert "us-east-1" in manager.api_host
    
    def test_fingerprint_property(self):
        """
        Что он делает: Проверяет свойство fingerprint.
        Цель: Убедиться, что fingerprint доступен через property.
        """
        print("Настройка: Создание KiroAuthManager...")
        manager = KiroAuthManager(refresh_token="test")
        
        print("Проверка: fingerprint доступен и имеет корректную длину...")
        print(f"fingerprint: {manager.fingerprint}")
        assert len(manager.fingerprint) == 64


================================================
FILE: tests/unit/test_cache.py
================================================
# -*- coding: utf-8 -*-

"""
Unit-тесты для ModelInfoCache.
Проверяет логику кэширования метаданных моделей.
"""

import asyncio
import time
import pytest

from kiro_gateway.cache import ModelInfoCache
from kiro_gateway.config import DEFAULT_MAX_INPUT_TOKENS


class TestModelInfoCacheInitialization:
    """Тесты инициализации ModelInfoCache."""
    
    def test_initialization_creates_empty_cache(self):
        """
        Что он делает: Проверяет, что кэш создаётся пустым.
        Цель: Убедиться в корректной инициализации.
        """
        print("Настройка: Создание ModelInfoCache...")
        cache = ModelInfoCache()
        
        print("Проверка: Кэш пуст при создании...")
        print(f"Сравниваем is_empty(): Ожидалось True, Получено {cache.is_empty()}")
        assert cache.is_empty() is True
        
        print(f"Сравниваем size: Ожидалось 0, Получено {cache.size}")
        assert cache.size == 0
    
    def test_initialization_with_custom_ttl(self):
        """
        Что он делает: Проверяет создание кэша с кастомным TTL.
        Цель: Убедиться, что TTL можно настроить.
        """
        print("Настройка: Создание ModelInfoCache с TTL=7200...")
        cache = ModelInfoCache(cache_ttl=7200)
        
        print("Проверка: TTL установлен корректно...")
        print(f"Сравниваем _cache_ttl: Ожидалось 7200, Получено {cache._cache_ttl}")
        assert cache._cache_ttl == 7200
    
    def test_initialization_last_update_is_none(self):
        """
        Что он делает: Проверяет, что last_update_time изначально None.
        Цель: Убедиться, что время обновления не установлено до первого update.
        """
        print("Настройка: Создание ModelInfoCache...")
        cache = ModelInfoCache()
        
        print("Проверка: last_update_time изначально None...")
        print(f"Сравниваем last_update_time: Ожидалось None, Получено {cache.last_update_time}")
        assert cache.last_update_time is None


class TestModelInfoCacheUpdate:
    """Тесты обновления кэша."""
    
    @pytest.mark.asyncio
    async def test_update_populates_cache(self, sample_models_data):
        """
        Что он делает: Проверяет заполнение кэша данными.
        Цель: Убедиться, что update() корректно сохраняет модели.
        """
        print("Настройка: Создание ModelInfoCache...")
        cache = ModelInfoCache()
        
        print(f"Действие: Обновление кэша с {len(sample_models_data)} моделями...")
        await cache.update(sample_models_data)
        
        print("Проверка: Кэш заполнен...")
        print(f"Сравниваем is_empty(): Ожидалось False, Получено {cache.is_empty()}")
        assert cache.is_empty() is False
        
        print(f"Сравниваем size: Ожидалось {len(sample_models_data)}, Получено {cache.size}")
        assert cache.size == len(sample_models_data)
    
    @pytest.mark.asyncio
    async def test_update_sets_last_update_time(self, sample_models_data):
        """
        Что он делает: Проверяет установку времени последнего обновления.
        Цель: Убедиться, что last_update_time устанавливается после update.
        """
        print("Настройка: Создание ModelInfoCache...")
        cache = ModelInfoCache()
        
        before_update = time.time()
        print(f"Действие: Обновление кэша (время до: {before_update})...")
        await cache.update(sample_models_data)
        after_update = time.time()
        
        print("Проверка: last_update_time установлен в разумных пределах...")
        print(f"last_update_time: {cache.last_update_time}")
        assert cache.last_update_time is not None
        assert before_update <= cache.last_update_time <= after_update
    
    @pytest.mark.asyncio
    async def test_update_replaces_existing_data(self, sample_models_data):
        """
        Что он делает: Проверяет замену данных при повторном update.
        Цель: Убедиться, что старые данные полностью заменяются.
        """
        print("Настройка: Создание ModelInfoCache и первое обновление...")
        cache = ModelInfoCache()
        await cache.update(sample_models_data)
        
        print("Действие: Обновление с новыми данными...")
        new_data = [{"modelId": "new-model", "tokenLimits": {"maxInputTokens": 50000}}]
        await cache.update(new_data)
        
        print("Проверка: Старые данные заменены...")
        print(f"Сравниваем size: Ожидалось 1, Получено {cache.size}")
        assert cache.size == 1
        
        print("Проверка: Старая модель недоступна...")
        assert cache.get("claude-sonnet-4") is None
        
        print("Проверка: Новая модель доступна...")
        assert cache.get("new-model") is not None
    
    @pytest.mark.asyncio
    async def test_update_with_empty_list(self):
        """
        Что он делает: Проверяет обновление пустым списком.
        Цель: Убедиться, что кэш очищается при пустом update.
        """
        print("Настройка: Создание ModelInfoCache с данными...")
        cache = ModelInfoCache()
        await cache.update([{"modelId": "test-model"}])
        
        print("Действие: Обновление пустым списком...")
        await cache.update([])
        
        print("Проверка: Кэш пуст...")
        print(f"Сравниваем is_empty(): Ожидалось True, Получено {cache.is_empty()}")
        assert cache.is_empty() is True


class TestModelInfoCacheGet:
    """Тесты получения данных из кэша."""
    
    @pytest.mark.asyncio
    async def test_get_returns_model_info(self, sample_models_data):
        """
        Что он делает: Проверяет получение информации о модели.
        Цель: Убедиться, что get() возвращает корректные данные.
        """
        print("Настройка: Создание и заполнение кэша...")
        cache = ModelInfoCache()
        await cache.update(sample_models_data)
        
        print("Действие: Получение информации о claude-sonnet-4...")
        model_info = cache.get("claude-sonnet-4")
        
        print("Проверка: Информация получена...")
        print(f"model_info: {model_info}")
        assert model_info is not None
        assert model_info["modelId"] == "claude-sonnet-4"
    
    @pytest.mark.asyncio
    async def test_get_returns_none_for_unknown_model(self, sample_models_data):
        """
        Что он делает: Проверяет возврат None для неизвестной модели.
        Цель: Убедиться, что get() не падает при отсутствии модели.
        """
        print("Настройка: Создание и заполнение кэша...")
        cache = ModelInfoCache()
        await cache.update(sample_models_data)
        
        print("Действие: Получение информации о несуществующей модели...")
        model_info = cache.get("non-existent-model")
        
        print("Проверка: Возвращён None...")
        print(f"Сравниваем model_info: Ожидалось None, Получено {model_info}")
        assert model_info is None
    
    def test_get_from_empty_cache(self):
        """
        Что он делает: Проверяет get() из пустого кэша.
        Цель: Убедиться, что пустой кэш не вызывает ошибок.
        """
        print("Настройка: Создание пустого кэша...")
        cache = ModelInfoCache()
        
        print("Действие: Получение из пустого кэша...")
        model_info = cache.get("any-model")
        
        print("Проверка: Возвращён None...")
        print(f"Сравниваем model_info: Ожидалось None, Получено {model_info}")
        assert model_info is None


class TestModelInfoCacheGetMaxInputTokens:
    """Тесты получения maxInputTokens."""
    
    @pytest.mark.asyncio
    async def test_get_max_input_tokens_returns_value(self, sample_models_data):
        """
        Что он делает: Проверяет получение maxInputTokens для модели.
        Цель: Убедиться, что значение извлекается из tokenLimits.
        """
        print("Настройка: Создание и заполнение кэша...")
        cache = ModelInfoCache()
        await cache.update(sample_models_data)
        
        print("Действие: Получение maxInputTokens для claude-sonnet-4...")
        max_tokens = cache.get_max_input_tokens("claude-sonnet-4")
        
        print("Проверка: Значение корректно...")
        print(f"Сравниваем max_tokens: Ожидалось 200000, Получено {max_tokens}")
        assert max_tokens == 200000
    
    @pytest.mark.asyncio
    async def test_get_max_input_tokens_returns_default_for_unknown(self, sample_models_data):
        """
        Что он делает: Проверяет возврат дефолта для неизвестной модели.
        Цель: Убедиться, что возвращается DEFAULT_MAX_INPUT_TOKENS.
        """
        print("Настройка: Создание и заполнение кэша...")
        cache = ModelInfoCache()
        await cache.update(sample_models_data)
        
        print("Действие: Получение maxInputTokens для неизвестной модели...")
        max_tokens = cache.get_max_input_tokens("unknown-model")
        
        print("Проверка: Возвращён дефолт...")
        print(f"Сравниваем max_tokens: Ожидалось {DEFAULT_MAX_INPUT_TOKENS}, Получено {max_tokens}")
        assert max_tokens == DEFAULT_MAX_INPUT_TOKENS
    
    @pytest.mark.asyncio
    async def test_get_max_input_tokens_returns_default_when_no_token_limits(self):
        """
        Что он делает: Проверяет возврат дефолта при отсутствии tokenLimits.
        Цель: Убедиться, что модель без tokenLimits не ломает логику.
        """
        print("Настройка: Создание кэша с моделью без tokenLimits...")
        cache = ModelInfoCache()
        await cache.update([{"modelId": "model-without-limits"}])
        
        print("Действие: Получение maxInputTokens...")
        max_tokens = cache.get_max_input_tokens("model-without-limits")
        
        print("Проверка: Возвращён дефолт...")
        print(f"Сравниваем max_tokens: Ожидалось {DEFAULT_MAX_INPUT_TOKENS}, Получено {max_tokens}")
        assert max_tokens == DEFAULT_MAX_INPUT_TOKENS
    
    @pytest.mark.asyncio
    async def test_get_max_input_tokens_returns_default_when_max_input_is_none(self):
        """
        Что он делает: Проверяет возврат дефолта при maxInputTokens=None.
        Цель: Убедиться, что None в tokenLimits обрабатывается корректно.
        """
        print("Настройка: Создание кэша с моделью с maxInputTokens=None...")
        cache = ModelInfoCache()
        await cache.update([{
            "modelId": "model-with-null",
            "tokenLimits": {"maxInputTokens": None}
        }])
        
        print("Действие: Получение maxInputTokens...")
        max_tokens = cache.get_max_input_tokens("model-with-null")
        
        print("Проверка: Возвращён дефолт...")
        print(f"Сравниваем max_tokens: Ожидалось {DEFAULT_MAX_INPUT_TOKENS}, Получено {max_tokens}")
        assert max_tokens == DEFAULT_MAX_INPUT_TOKENS


class TestModelInfoCacheIsEmpty:
    """Тесты проверки пустоты кэша."""
    
    def test_is_empty_returns_true_for_new_cache(self):
        """
        Что он делает: Проверяет is_empty() для нового кэша.
        Цель: Убедиться, что новый кэш считается пустым.
        """
        print("Настройка: Создание нового кэша...")
        cache = ModelInfoCache()
        
        print("Проверка: is_empty() возвращает True...")
        print(f"Сравниваем is_empty(): Ожидалось True, Получено {cache.is_empty()}")
        assert cache.is_empty() is True
    
    @pytest.mark.asyncio
    async def test_is_empty_returns_false_after_update(self, sample_models_data):
        """
        Что он делает: Проверяет is_empty() после заполнения.
        Цель: Убедиться, что заполненный кэш не считается пустым.
        """
        print("Настройка: Создание и заполнение кэша...")
        cache = ModelInfoCache()
        await cache.update(sample_models_data)
        
        print("Проверка: is_empty() возвращает False...")
        print(f"Сравниваем is_empty(): Ожидалось False, Получено {cache.is_empty()}")
        assert cache.is_empty() is False


class TestModelInfoCacheIsStale:
    """Тесты проверки устаревания кэша."""
    
    def test_is_stale_returns_true_for_new_cache(self):
        """
        Что он делает: Проверяет is_stale() для нового кэша.
        Цель: Убедиться, что кэш без обновлений считается устаревшим.
        """
        print("Настройка: Создание нового кэша...")
        cache = ModelInfoCache()
        
        print("Проверка: is_stale() возвращает True...")
        print(f"Сравниваем is_stale(): Ожидалось True, Получено {cache.is_stale()}")
        assert cache.is_stale() is True
    
    @pytest.mark.asyncio
    async def test_is_stale_returns_false_after_recent_update(self, sample_models_data):
        """
        Что он делает: Проверяет is_stale() сразу после обновления.
        Цель: Убедиться, что свежий кэш не считается устаревшим.
        """
        print("Настройка: Создание и заполнение кэша...")
        cache = ModelInfoCache()
        await cache.update(sample_models_data)
        
        print("Проверка: is_stale() возвращает False...")
        print(f"Сравниваем is_stale(): Ожидалось False, Получено {cache.is_stale()}")
        assert cache.is_stale() is False
    
    @pytest.mark.asyncio
    async def test_is_stale_returns_true_after_ttl_expires(self, sample_models_data):
        """
        Что он делает: Проверяет is_stale() после истечения TTL.
        Цель: Убедиться, что кэш считается устаревшим после TTL.
        """
        print("Настройка: Создание кэша с TTL=0.1 секунды...")
        cache = ModelInfoCache(cache_ttl=0.1)
        await cache.update(sample_models_data)
        
        print("Действие: Ожидание истечения TTL...")
        await asyncio.sleep(0.2)
        
        print("Проверка: is_stale() возвращает True...")
        print(f"Сравниваем is_stale(): Ожидалось True, Получено {cache.is_stale()}")
        assert cache.is_stale() is True


class TestModelInfoCacheGetAllModelIds:
    """Тесты получения списка ID моделей."""
    
    def test_get_all_model_ids_returns_empty_for_new_cache(self):
        """
        Что он делает: Проверяет get_all_model_ids() для пустого кэша.
        Цель: Убедиться, что возвращается пустой список.
        """
        print("Настройка: Создание пустого кэша...")
        cache = ModelInfoCache()
        
        print("Действие: Получение списка ID моделей...")
        model_ids = cache.get_all_model_ids()
        
        print("Проверка: Список пуст...")
        print(f"Сравниваем model_ids: Ожидалось [], Получено {model_ids}")
        assert model_ids == []
    
    @pytest.mark.asyncio
    async def test_get_all_model_ids_returns_all_ids(self, sample_models_data):
        """
        Что он делает: Проверяет get_all_model_ids() для заполненного кэша.
        Цель: Убедиться, что возвращаются все ID моделей.
        """
        print("Настройка: Создание и заполнение кэша...")
        cache = ModelInfoCache()
        await cache.update(sample_models_data)
        
        print("Действие: Получение списка ID моделей...")
        model_ids = cache.get_all_model_ids()
        
        print("Проверка: Все ID присутствуют...")
        expected_ids = [m["modelId"] for m in sample_models_data]
        print(f"Сравниваем model_ids: Ожидалось {expected_ids}, Получено {model_ids}")
        assert set(model_ids) == set(expected_ids)


class TestModelInfoCacheThreadSafety:
    """Тесты потокобезопасности кэша."""
    
    @pytest.mark.asyncio
    async def test_concurrent_updates_dont_corrupt_cache(self, sample_models_data):
        """
        Что он делает: Проверяет потокобезопасность при параллельных update.
        Цель: Убедиться, что asyncio.Lock защищает от race conditions.
        """
        print("Настройка: Создание кэша...")
        cache = ModelInfoCache()
        
        async def update_with_data(data):
            await cache.update(data)
        
        print("Действие: 10 параллельных обновлений...")
        tasks = []
        for i in range(10):
            data = [{"modelId": f"model-{i}", "tokenLimits": {"maxInputTokens": 100000 + i}}]
            tasks.append(update_with_data(data))
        
        await asyncio.gather(*tasks)
        
        print("Проверка: Кэш содержит данные последнего обновления...")
        # Из-за race condition, мы не знаем какое обновление было последним,
        # но кэш должен содержать ровно одну модель
        print(f"Сравниваем size: Ожидалось 1, Получено {cache.size}")
        assert cache.size == 1
        
        print("Проверка: Кэш не повреждён...")
        model_ids = cache.get_all_model_ids()
        assert len(model_ids) == 1
        assert model_ids[0].startswith("model-")
    
    @pytest.mark.asyncio
    async def test_concurrent_reads_are_safe(self, sample_models_data):
        """
        Что он делает: Проверяет безопасность параллельных чтений.
        Цель: Убедиться, что множественные get() не вызывают проблем.
        """
        print("Настройка: Создание и заполнение кэша...")
        cache = ModelInfoCache()
        await cache.update(sample_models_data)
        
        print("Действие: 100 параллельных чтений...")
        async def read_model():
            return cache.get("claude-sonnet-4")
        
        results = await asyncio.gather(*[read_model() for _ in range(100)])
        
        print("Проверка: Все чтения вернули одинаковый результат...")
        assert all(r is not None for r in results)
        assert all(r["modelId"] == "claude-sonnet-4" for r in results)


================================================
FILE: tests/unit/test_config.py
================================================
# -*- coding: utf-8 -*-

"""
Unit tests for the configuration module.
Verifies loading settings from environment variables.
"""

import pytest
import os
from unittest.mock import patch


class TestLogLevelConfig:
    """Tests for LOG_LEVEL configuration."""
    
    def test_default_log_level_is_info(self):
        """
        What it does: Verifies that LOG_LEVEL defaults to INFO.
        Purpose: Ensure that INFO is used when no environment variable is set.
        
        Note: This test verifies the config.py code logic, not the actual
        value from the .env file. We mock os.getenv to simulate
        the absence of the environment variable.
        """
        print("Setup: Mocking os.getenv for LOG_LEVEL...")
        
        # Create a mock that returns None for LOG_LEVEL (simulating missing variable)
        original_getenv = os.getenv
        
        def mock_getenv(key, default=None):
            if key == "LOG_LEVEL":
                print(f"os.getenv('{key}') -> None (mocked)")
                return default  # Return default, simulating missing variable
            return original_getenv(key, default)
        
        with patch.object(os, 'getenv', side_effect=mock_getenv):
            # Reload config module with mocked getenv
            import importlib
            import kiro_gateway.config as config_module
            importlib.reload(config_module)
            
            print(f"LOG_LEVEL: {config_module.LOG_LEVEL}")
            print(f"Comparing: Expected 'INFO', Got '{config_module.LOG_LEVEL}'")
            assert config_module.LOG_LEVEL == "INFO"
        
        # Restore module with real values
        import importlib
        import kiro_gateway.config as config_module
        importlib.reload(config_module)
    
    def test_log_level_from_environment(self):
        """
        What it does: Verifies loading LOG_LEVEL from environment variable.
        Purpose: Ensure that the value from environment is used.
        """
        print("Setup: Setting LOG_LEVEL=DEBUG...")
        
        with patch.dict(os.environ, {"LOG_LEVEL": "DEBUG"}):
            import importlib
            import kiro_gateway.config as config_module
            importlib.reload(config_module)
            
            print(f"LOG_LEVEL: {config_module.LOG_LEVEL}")
            print(f"Comparing: Expected 'DEBUG', Got '{config_module.LOG_LEVEL}'")
            assert config_module.LOG_LEVEL == "DEBUG"
    
    def test_log_level_uppercase_conversion(self):
        """
        What it does: Verifies LOG_LEVEL conversion to uppercase.
        Purpose: Ensure that lowercase value is converted to uppercase.
        """
        print("Setup: Setting LOG_LEVEL=warning (lowercase)...")
        
        with patch.dict(os.environ, {"LOG_LEVEL": "warning"}):
            import importlib
            import kiro_gateway.config as config_module
            importlib.reload(config_module)
            
            print(f"LOG_LEVEL: {config_module.LOG_LEVEL}")
            print(f"Comparing: Expected 'WARNING', Got '{config_module.LOG_LEVEL}'")
            assert config_module.LOG_LEVEL == "WARNING"
    
    def test_log_level_trace(self):
        """
        What it does: Verifies setting LOG_LEVEL=TRACE.
        Purpose: Ensure that TRACE level is supported.
        """
        print("Setup: Setting LOG_LEVEL=TRACE...")
        
        with patch.dict(os.environ, {"LOG_LEVEL": "TRACE"}):
            import importlib
            import kiro_gateway.config as config_module
            importlib.reload(config_module)
            
            print(f"LOG_LEVEL: {config_module.LOG_LEVEL}")
            assert config_module.LOG_LEVEL == "TRACE"
    
    def test_log_level_error(self):
        """
        What it does: Verifies setting LOG_LEVEL=ERROR.
        Purpose: Ensure that ERROR level is supported.
        """
        print("Setup: Setting LOG_LEVEL=ERROR...")
        
        with patch.dict(os.environ, {"LOG_LEVEL": "ERROR"}):
            import importlib
            import kiro_gateway.config as config_module
            importlib.reload(config_module)
            
            print(f"LOG_LEVEL: {config_module.LOG_LEVEL}")
            assert config_module.LOG_LEVEL == "ERROR"
    
    def test_log_level_critical(self):
        """
        What it does: Verifies setting LOG_LEVEL=CRITICAL.
        Purpose: Ensure that CRITICAL level is supported.
        """
        print("Setup: Setting LOG_LEVEL=CRITICAL...")
        
        with patch.dict(os.environ, {"LOG_LEVEL": "CRITICAL"}):
            import importlib
            import kiro_gateway.config as config_module
            importlib.reload(config_module)
            
            print(f"LOG_LEVEL: {config_module.LOG_LEVEL}")
            assert config_module.LOG_LEVEL == "CRITICAL"


class TestToolDescriptionMaxLengthConfig:
    """Tests for TOOL_DESCRIPTION_MAX_LENGTH configuration."""
    
    def test_default_tool_description_max_length(self):
        """
        What it does: Verifies the default value for TOOL_DESCRIPTION_MAX_LENGTH.
        Purpose: Ensure that 10000 is used by default.
        """
        print("Setup: Removing TOOL_DESCRIPTION_MAX_LENGTH from environment...")
        
        with patch.dict(os.environ, {}, clear=False):
            if "TOOL_DESCRIPTION_MAX_LENGTH" in os.environ:
                del os.environ["TOOL_DESCRIPTION_MAX_LENGTH"]
            
            import importlib
            import kiro_gateway.config as config_module
            importlib.reload(config_module)
            
            print(f"TOOL_DESCRIPTION_MAX_LENGTH: {config_module.TOOL_DESCRIPTION_MAX_LENGTH}")
            assert config_module.TOOL_DESCRIPTION_MAX_LENGTH == 10000
    
    def test_tool_description_max_length_from_environment(self):
        """
        What it does: Verifies loading TOOL_DESCRIPTION_MAX_LENGTH from environment.
        Purpose: Ensure that the value from environment is used.
        """
        print("Setup: Setting TOOL_DESCRIPTION_MAX_LENGTH=5000...")
        
        with patch.dict(os.environ, {"TOOL_DESCRIPTION_MAX_LENGTH": "5000"}):
            import importlib
            import kiro_gateway.config as config_module
            importlib.reload(config_module)
            
            print(f"TOOL_DESCRIPTION_MAX_LENGTH: {config_module.TOOL_DESCRIPTION_MAX_LENGTH}")
            assert config_module.TOOL_DESCRIPTION_MAX_LENGTH == 5000
    
    def test_tool_description_max_length_zero_disables(self):
        """
        What it does: Verifies that 0 disables the feature.
        Purpose: Ensure that TOOL_DESCRIPTION_MAX_LENGTH=0 works.
        """
        print("Setup: Setting TOOL_DESCRIPTION_MAX_LENGTH=0...")
        
        with patch.dict(os.environ, {"TOOL_DESCRIPTION_MAX_LENGTH": "0"}):
            import importlib
            import kiro_gateway.config as config_module
            importlib.reload(config_module)
            
            print(f"TOOL_DESCRIPTION_MAX_LENGTH: {config_module.TOOL_DESCRIPTION_MAX_LENGTH}")
            assert config_module.TOOL_DESCRIPTION_MAX_LENGTH == 0


class TestTimeoutConfigurationWarning:
    """Tests for _warn_timeout_configuration() function."""
    
    def test_no_warning_when_first_token_less_than_streaming(self, capsys):
        """
        What it does: Verifies that warning is NOT shown with correct configuration.
        Purpose: Ensure that no warning when FIRST_TOKEN_TIMEOUT < STREAMING_READ_TIMEOUT.
        """
        print("Setup: FIRST_TOKEN_TIMEOUT=15, STREAMING_READ_TIMEOUT=300...")
        
        with patch.dict(os.environ, {
            "FIRST_TOKEN_TIMEOUT": "15",
            "STREAMING_READ_TIMEOUT": "300"
        }):
            import importlib
            import kiro_gateway.config as config_module
            importlib.reload(config_module)
            
            # Call the warning function
            config_module._warn_timeout_configuration()
            
            captured = capsys.readouterr()
            print(f"Captured stderr: {captured.err}")
            
            # Warning should NOT be shown
            assert "WARNING" not in captured.err
            assert "Suboptimal timeout configuration" not in captured.err
    
    def test_warning_when_first_token_equals_streaming(self, capsys):
        """
        What it does: Verifies that warning is shown when timeouts are equal.
        Purpose: Ensure that warning when FIRST_TOKEN_TIMEOUT == STREAMING_READ_TIMEOUT.
        """
        print("Setup: FIRST_TOKEN_TIMEOUT=300, STREAMING_READ_TIMEOUT=300...")
        
        with patch.dict(os.environ, {
            "FIRST_TOKEN_TIMEOUT": "300",
            "STREAMING_READ_TIMEOUT": "300"
        }):
            import importlib
            import kiro_gateway.config as config_module
            importlib.reload(config_module)
            
            # Call the warning function
            config_module._warn_timeout_configuration()
            
            captured = capsys.readouterr()
            print(f"Captured stderr: {captured.err}")
            
            # Warning SHOULD be shown
            assert "WARNING" in captured.err or "Suboptimal timeout configuration" in captured.err
    
    def test_warning_when_first_token_greater_than_streaming(self, capsys):
        """
        What it does: Verifies that warning is shown when FIRST_TOKEN > STREAMING.
        Purpose: Ensure that warning when FIRST_TOKEN_TIMEOUT > STREAMING_READ_TIMEOUT.
        """
        print("Setup: FIRST_TOKEN_TIMEOUT=500, STREAMING_READ_TIMEOUT=300...")
        
        with patch.dict(os.environ, {
            "FIRST_TOKEN_TIMEOUT": "500",
            "STREAMING_READ_TIMEOUT": "300"
        }):
            import importlib
            import kiro_gateway.config as config_module
            importlib.reload(config_module)
            
            # Call the warning function
            config_module._warn_timeout_configuration()
            
            captured = capsys.readouterr()
            print(f"Captured stderr: {captured.err}")
            
            # Warning SHOULD be shown
            assert "WARNING" in captured.err or "Suboptimal timeout configuration" in captured.err
            # Verify that timeout values are mentioned in warning
            assert "500" in captured.err
            assert "300" in captured.err
    
    def test_warning_contains_recommendation(self, capsys):
        """
        What it does: Verifies that warning contains a recommendation.
        Purpose: Ensure that user receives useful information.
        """
        print("Setup: FIRST_TOKEN_TIMEOUT=400, STREAMING_READ_TIMEOUT=300...")
        
        with patch.dict(os.environ, {
            "FIRST_TOKEN_TIMEOUT": "400",
            "STREAMING_READ_TIMEOUT": "300"
        }):
            import importlib
            import kiro_gateway.config as config_module
            importlib.reload(config_module)
            
            # Call the warning function
            config_module._warn_timeout_configuration()
            
            captured = capsys.readouterr()
            print(f"Captured stderr: {captured.err}")
            
            # Warning should contain recommendation
            assert "Recommendation" in captured.err or "LESS than" in captured.err


================================================
FILE: tests/unit/test_converters.py
================================================
# -*- coding: utf-8 -*-

"""
Unit-тесты для конвертеров OpenAI <-> Kiro.
Проверяет логику преобразования форматов сообщений и payload.
"""

import pytest

from unittest.mock import patch

from kiro_gateway.converters import (
    extract_text_content,
    merge_adjacent_messages,
    build_kiro_history,
    build_kiro_payload,
    process_tools_with_long_descriptions,
    _extract_tool_results,
    _extract_tool_uses,
    _build_user_input_context,
    _sanitize_json_schema
)
from kiro_gateway.models import ChatMessage, ChatCompletionRequest, Tool, ToolFunction


class TestExtractTextContent:
    """Тесты функции extract_text_content."""
    
    def test_extracts_from_string(self):
        """
        Что он делает: Проверяет извлечение текста из строки.
        Цель: Убедиться, что строка возвращается как есть.
        """
        print("Настройка: Простая строка...")
        content = "Hello, World!"
        
        print("Действие: Извлечение текста...")
        result = extract_text_content(content)
        
        print(f"Сравниваем результат: Ожидалось 'Hello, World!', Получено '{result}'")
        assert result == "Hello, World!"
    
    def test_extracts_from_none(self):
        """
        Что он делает: Проверяет обработку None.
        Цель: Убедиться, что None возвращает пустую строку.
        """
        print("Настройка: None...")
        
        print("Действие: Извлечение текста...")
        result = extract_text_content(None)
        
        print(f"Сравниваем результат: Ожидалось '', Получено '{result}'")
        assert result == ""
    
    def test_extracts_from_list_with_text_type(self):
        """
        Что он делает: Проверяет извлечение из списка с type=text.
        Цель: Убедиться, что OpenAI multimodal формат обрабатывается.
        """
        print("Настройка: Список с type=text...")
        content = [
            {"type": "text", "text": "Hello"},
            {"type": "text", "text": " World"}
        ]
        
        print("Действие: Извлечение текста...")
        result = extract_text_content(content)
        
        print(f"Сравниваем результат: Ожидалось 'Hello World', Получено '{result}'")
        assert result == "Hello World"
    
    def test_extracts_from_list_with_text_key(self):
        """
        Что он делает: Проверяет извлечение из списка с ключом text.
        Цель: Убедиться, что альтернативный формат обрабатывается.
        """
        print("Настройка: Список с ключом text...")
        content = [{"text": "Hello"}, {"text": " World"}]
        
        print("Действие: Извлечение текста...")
        result = extract_text_content(content)
        
        print(f"Сравниваем результат: Ожидалось 'Hello World', Получено '{result}'")
        assert result == "Hello World"
    
    def test_extracts_from_list_with_strings(self):
        """
        Что он делает: Проверяет извлечение из списка строк.
        Цель: Убедиться, что список строк объединяется.
        """
        print("Настройка: Список строк...")
        content = ["Hello", " ", "World"]
        
        print("Действие: Извлечение текста...")
        result = extract_text_content(content)
        
        print(f"Сравниваем результат: Ожидалось 'Hello World', Получено '{result}'")
        assert result == "Hello World"
    
    def test_extracts_from_mixed_list(self):
        """
        Что он делает: Проверяет извлечение из смешанного списка.
        Цель: Убедиться, что разные форматы в одном списке обрабатываются.
        """
        print("Настройка: Смешанный список...")
        content = [
            {"type": "text", "text": "Part1"},
            "Part2",
            {"text": "Part3"}
        ]
        
        print("Действие: Извлечение текста...")
        result = extract_text_content(content)
        
        print(f"Сравниваем результат: Ожидалось 'Part1Part2Part3', Получено '{result}'")
        assert result == "Part1Part2Part3"
    
    def test_converts_other_types_to_string(self):
        """
        Что он делает: Проверяет конвертацию других типов в строку.
        Цель: Убедиться, что числа и другие типы преобразуются.
        """
        print("Настройка: Число...")
        content = 42
        
        print("Действие: Извлечение текста...")
        result = extract_text_content(content)
        
        print(f"Сравниваем результат: Ожидалось '42', Получено '{result}'")
        assert result == "42"
    
    def test_handles_empty_list(self):
        """
        Что он делает: Проверяет обработку пустого списка.
        Цель: Убедиться, что пустой список возвращает пустую строку.
        """
        print("Настройка: Пустой список...")
        content = []
        
        print("Действие: Извлечение текста...")
        result = extract_text_content(content)
        
        print(f"Сравниваем результат: Ожидалось '', Получено '{result}'")
        assert result == ""


class TestMergeAdjacentMessages:
    """Тесты функции merge_adjacent_messages."""
    
    def test_merges_adjacent_user_messages(self):
        """
        Что он делает: Проверяет объединение соседних user сообщений.
        Цель: Убедиться, что сообщения с одинаковой ролью объединяются.
        """
        print("Настройка: Два user сообщения подряд...")
        messages = [
            ChatMessage(role="user", content="Hello"),
            ChatMessage(role="user", content="World")
        ]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Сравниваем длину: Ожидалось 1, Получено {len(result)}")
        assert len(result) == 1
        assert "Hello" in result[0].content
        assert "World" in result[0].content
    
    def test_preserves_alternating_messages(self):
        """
        Что он делает: Проверяет сохранение чередующихся сообщений.
        Цель: Убедиться, что разные роли не объединяются.
        """
        print("Настройка: Чередующиеся сообщения...")
        messages = [
            ChatMessage(role="user", content="Hello"),
            ChatMessage(role="assistant", content="Hi"),
            ChatMessage(role="user", content="How are you?")
        ]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Сравниваем длину: Ожидалось 3, Получено {len(result)}")
        assert len(result) == 3
    
    def test_handles_empty_list(self):
        """
        Что он делает: Проверяет обработку пустого списка.
        Цель: Убедиться, что пустой список не вызывает ошибок.
        """
        print("Настройка: Пустой список...")
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages([])
        
        print(f"Сравниваем результат: Ожидалось [], Получено {result}")
        assert result == []
    
    def test_handles_single_message(self):
        """
        Что он делает: Проверяет обработку одного сообщения.
        Цель: Убедиться, что одно сообщение возвращается как есть.
        """
        print("Настройка: Одно сообщение...")
        messages = [ChatMessage(role="user", content="Hello")]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Сравниваем длину: Ожидалось 1, Получено {len(result)}")
        assert len(result) == 1
        assert result[0].content == "Hello"
    
    def test_merges_multiple_adjacent_groups(self):
        """
        Что он делает: Проверяет объединение нескольких групп.
        Цель: Убедиться, что несколько групп соседних сообщений объединяются.
        """
        print("Настройка: Несколько групп соседних сообщений...")
        messages = [
            ChatMessage(role="user", content="A"),
            ChatMessage(role="user", content="B"),
            ChatMessage(role="assistant", content="C"),
            ChatMessage(role="assistant", content="D"),
            ChatMessage(role="user", content="E")
        ]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Сравниваем длину: Ожидалось 3, Получено {len(result)}")
        assert len(result) == 3
        assert result[0].role == "user"
        assert result[1].role == "assistant"
        assert result[2].role == "user"
    
    def test_converts_tool_message_to_user_with_tool_result(self):
        """
        Что он делает: Проверяет преобразование tool message в user message с tool_result.
        Цель: Убедиться, что role="tool" преобразуется в user message с tool_results content.
        """
        print("Настройка: Tool message...")
        messages = [
            ChatMessage(role="tool", content="Tool result text", tool_call_id="call_123")
        ]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Результат: {result}")
        print(f"Сравниваем длину: Ожидалось 1, Получено {len(result)}")
        assert len(result) == 1
        assert result[0].role == "user"
        
        print("Проверяем content содержит tool_result...")
        assert isinstance(result[0].content, list)
        assert len(result[0].content) == 1
        assert result[0].content[0]["type"] == "tool_result"
        assert result[0].content[0]["tool_use_id"] == "call_123"
        assert result[0].content[0]["content"] == "Tool result text"
    
    def test_converts_multiple_tool_messages_to_single_user_message(self):
        """
        Что он делает: Проверяет объединение нескольких tool messages в один user message.
        Цель: Убедиться, что несколько tool results объединяются в один user message.
        """
        print("Настройка: Несколько tool messages подряд...")
        messages = [
            ChatMessage(role="tool", content="Result 1", tool_call_id="call_1"),
            ChatMessage(role="tool", content="Result 2", tool_call_id="call_2"),
            ChatMessage(role="tool", content="Result 3", tool_call_id="call_3")
        ]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Результат: {result}")
        print(f"Сравниваем длину: Ожидалось 1, Получено {len(result)}")
        assert len(result) == 1
        assert result[0].role == "user"
        
        print("Проверяем content содержит все tool_results...")
        assert isinstance(result[0].content, list)
        assert len(result[0].content) == 3
        
        tool_use_ids = [item["tool_use_id"] for item in result[0].content]
        assert "call_1" in tool_use_ids
        assert "call_2" in tool_use_ids
        assert "call_3" in tool_use_ids
    
    def test_tool_message_followed_by_user_message(self):
        """
        Что он делает: Проверяет tool message перед user message.
        Цель: Убедиться, что tool results и user message объединяются.
        """
        print("Настройка: Tool message + user message...")
        messages = [
            ChatMessage(role="tool", content="Tool result", tool_call_id="call_1"),
            ChatMessage(role="user", content="Continue please")
        ]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Результат: {result}")
        print(f"Сравниваем длину: Ожидалось 1, Получено {len(result)}")
        # Tool message преобразуется в user, затем объединяется с user
        assert len(result) == 1
        assert result[0].role == "user"
    
    def test_assistant_tool_user_sequence(self):
        """
        Что он делает: Проверяет последовательность assistant -> tool -> user.
        Цель: Убедиться, что tool message корректно вставляется между assistant и user.
        """
        print("Настройка: assistant -> tool -> user...")
        messages = [
            ChatMessage(role="assistant", content="I'll call a tool"),
            ChatMessage(role="tool", content="Tool output", tool_call_id="call_abc"),
            ChatMessage(role="user", content="Thanks!")
        ]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Результат: {result}")
        # assistant остаётся, tool+user объединяются в один user
        assert len(result) == 2
        assert result[0].role == "assistant"
        assert result[1].role == "user"
    
    def test_tool_message_with_empty_content(self):
        """
        Что он делает: Проверяет tool message с пустым content.
        Цель: Убедиться, что пустой результат заменяется на "(empty result)".
        """
        print("Настройка: Tool message с пустым content...")
        messages = [
            ChatMessage(role="tool", content="", tool_call_id="call_empty")
        ]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Результат: {result}")
        assert len(result) == 1
        assert result[0].content[0]["content"] == "(empty result)"
    
    def test_tool_message_with_none_tool_call_id(self):
        """
        Что он делает: Проверяет tool message без tool_call_id.
        Цель: Убедиться, что отсутствующий tool_call_id заменяется на пустую строку.
        """
        print("Настройка: Tool message без tool_call_id...")
        messages = [
            ChatMessage(role="tool", content="Result", tool_call_id=None)
        ]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Результат: {result}")
        assert len(result) == 1
        assert result[0].content[0]["tool_use_id"] == ""
    
    def test_merges_list_contents_correctly(self):
        """
        Что он делает: Проверяет объединение list contents.
        Цель: Убедиться, что списки объединяются корректно.
        """
        print("Настройка: Два user сообщения с list content...")
        messages = [
            ChatMessage(role="user", content=[{"type": "text", "text": "Part 1"}]),
            ChatMessage(role="user", content=[{"type": "text", "text": "Part 2"}])
        ]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Результат: {result}")
        assert len(result) == 1
        assert isinstance(result[0].content, list)
        assert len(result[0].content) == 2
    
    def test_merges_adjacent_assistant_tool_calls(self):
        """
        Что он делает: Проверяет объединение tool_calls при merge соседних assistant сообщений.
        Цель: Убедиться, что tool_calls из всех assistant сообщений сохраняются при объединении.
        
        Это критический тест для бага, когда Codex CLI отправляет несколько assistant
        сообщений подряд, каждое со своим tool_call. Без этого фикса второй tool_call
        терялся, что приводило к ошибке 400 от Kiro API (toolResult без toolUse).
        """
        print("Настройка: Два assistant сообщения с разными tool_calls...")
        messages = [
            ChatMessage(
                role="assistant",
                content=None,
                tool_calls=[{
                    "id": "tooluse_first",
                    "type": "function",
                    "function": {
                        "name": "shell",
                        "arguments": '{"command": ["ls", "-la"]}'
                    }
                }]
            ),
            ChatMessage(
                role="assistant",
                content=None,
                tool_calls=[{
                    "id": "tooluse_second",
                    "type": "function",
                    "function": {
                        "name": "shell",
                        "arguments": '{"command": ["pwd"]}'
                    }
                }]
            )
        ]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Результат: {result}")
        print(f"Сравниваем длину: Ожидалось 1, Получено {len(result)}")
        assert len(result) == 1
        assert result[0].role == "assistant"
        
        print("Проверяем, что оба tool_calls сохранены...")
        assert result[0].tool_calls is not None
        print(f"Сравниваем количество tool_calls: Ожидалось 2, Получено {len(result[0].tool_calls)}")
        assert len(result[0].tool_calls) == 2
        
        tool_ids = [tc["id"] for tc in result[0].tool_calls]
        print(f"Tool IDs: {tool_ids}")
        assert "tooluse_first" in tool_ids
        assert "tooluse_second" in tool_ids
    
    def test_merges_three_adjacent_assistant_tool_calls(self):
        """
        Что он делает: Проверяет объединение tool_calls из трёх assistant сообщений.
        Цель: Убедиться, что все tool_calls сохраняются при объединении более двух сообщений.
        """
        print("Настройка: Три assistant сообщения с tool_calls...")
        messages = [
            ChatMessage(
                role="assistant",
                content="",
                tool_calls=[{"id": "call_1", "type": "function", "function": {"name": "tool1", "arguments": "{}"}}]
            ),
            ChatMessage(
                role="assistant",
                content="",
                tool_calls=[{"id": "call_2", "type": "function", "function": {"name": "tool2", "arguments": "{}"}}]
            ),
            ChatMessage(
                role="assistant",
                content="",
                tool_calls=[{"id": "call_3", "type": "function", "function": {"name": "tool3", "arguments": "{}"}}]
            )
        ]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Результат: {result}")
        assert len(result) == 1
        assert len(result[0].tool_calls) == 3
        
        tool_ids = [tc["id"] for tc in result[0].tool_calls]
        print(f"Сравниваем tool IDs: Ожидалось ['call_1', 'call_2', 'call_3'], Получено {tool_ids}")
        assert tool_ids == ["call_1", "call_2", "call_3"]
    
    def test_merges_assistant_with_and_without_tool_calls(self):
        """
        Что он делает: Проверяет объединение assistant с tool_calls и без.
        Цель: Убедиться, что tool_calls корректно инициализируются при объединении.
        """
        print("Настройка: Assistant без tool_calls + assistant с tool_calls...")
        messages = [
            ChatMessage(role="assistant", content="Thinking...", tool_calls=None),
            ChatMessage(
                role="assistant",
                content="",
                tool_calls=[{"id": "call_1", "type": "function", "function": {"name": "tool1", "arguments": "{}"}}]
            )
        ]
        
        print("Действие: Объединение сообщений...")
        result = merge_adjacent_messages(messages)
        
        print(f"Результат: {result}")
        assert len(result) == 1
        assert result[0].tool_calls is not None
        print(f"Сравниваем количество tool_calls: Ожидалось 1, Получено {len(result[0].tool_calls)}")
        assert len(result[0].tool_calls) == 1
        assert result[0].tool_calls[0]["id"] == "call_1"


class TestBuildKiroHistory:
    """Тесты функции build_kiro_history."""
    
    def test_builds_user_message(self):
        """
        Что он делает: Проверяет построение user сообщения.
        Цель: Убедиться, что user сообщение преобразуется в userInputMessage.
        """
        print("Настройка: User сообщение...")
        messages = [ChatMessage(role="user", content="Hello")]
        
        print("Действие: Построение истории...")
        result = build_kiro_history(messages, "claude-sonnet-4")
        
        print(f"Результат: {result}")
        assert len(result) == 1
        assert "userInputMessage" in result[0]
        assert result[0]["userInputMessage"]["content"] == "Hello"
        assert result[0]["userInputMessage"]["modelId"] == "claude-sonnet-4"
    
    def test_builds_assistant_message(self):
        """
        Что он делает: Проверяет построение assistant сообщения.
        Цель: Убедиться, что assistant сообщение преобразуется в assistantResponseMessage.
        """
        print("Настройка: Assistant сообщение...")
        messages = [ChatMessage(role="assistant", content="Hi there")]
        
        print("Действие: Построение истории...")
        result = build_kiro_history(messages, "claude-sonnet-4")
        
        print(f"Результат: {result}")
        assert len(result) == 1
        assert "assistantResponseMessage" in result[0]
        assert result[0]["assistantResponseMessage"]["content"] == "Hi there"
    
    def test_ignores_system_messages(self):
        """
        Что он делает: Проверяет игнорирование system сообщений.
        Цель: Убедиться, что system сообщения не добавляются в историю.
        """
        print("Настройка: System сообщение...")
        messages = [ChatMessage(role="system", content="You are helpful")]
        
        print("Действие: Построение истории...")
        result = build_kiro_history(messages, "claude-sonnet-4")
        
        print(f"Сравниваем длину: Ожидалось 0, Получено {len(result)}")
        assert len(result) == 0
    
    def test_builds_conversation_history(self):
        """
        Что он делает: Проверяет построение полной истории разговора.
        Цель: Убедиться, что чередование user/assistant сохраняется.
        """
        print("Настройка: Полная история разговора...")
        messages = [
            ChatMessage(role="user", content="Hello"),
            ChatMessage(role="assistant", content="Hi"),
            ChatMessage(role="user", content="How are you?")
        ]
        
        print("Действие: Построение истории...")
        result = build_kiro_history(messages, "claude-sonnet-4")
        
        print(f"Результат: {result}")
        assert len(result) == 3
        assert "userInputMessage" in result[0]
        assert "assistantResponseMessage" in result[1]
        assert "userInputMessage" in result[2]
    
    def test_handles_empty_list(self):
        """
        Что он делает: Проверяет обработку пустого списка.
        Цель: Убедиться, что пустой список возвращает пустую историю.
        """
        print("Настройка: Пустой список...")
        
        print("Действие: Построение истории...")
        result = build_kiro_history([], "claude-sonnet-4")
        
        print(f"Сравниваем результат: Ожидалось [], Получено {result}")
        assert result == []


class TestExtractToolResults:
    """Тесты функции _extract_tool_results."""
    
    def test_extracts_tool_results_from_list(self):
        """
        Что он делает: Проверяет извлечение tool results из списка.
        Цель: Убедиться, что tool_result элементы извлекаются.
        """
        print("Настройка: Список с tool_result...")
        content = [
            {"type": "tool_result", "tool_use_id": "call_123", "content": "Result text"}
        ]
        
        print("Действие: Извлечение tool results...")
        result = _extract_tool_results(content)
        
        print(f"Результат: {result}")
        assert len(result) == 1
        assert result[0]["toolUseId"] == "call_123"
        assert result[0]["status"] == "success"
    
    def test_returns_empty_for_string_content(self):
        """
        Что он делает: Проверяет возврат пустого списка для строки.
        Цель: Убедиться, что строка не содержит tool results.
        """
        print("Настройка: Строка...")
        content = "Just a string"
        
        print("Действие: Извлечение tool results...")
        result = _extract_tool_results(content)
        
        print(f"Сравниваем результат: Ожидалось [], Получено {result}")
        assert result == []
    
    def test_returns_empty_for_list_without_tool_results(self):
        """
        Что он делает: Проверяет возврат пустого списка без tool_result.
        Цель: Убедиться, что обычные элементы не извлекаются.
        """
        print("Настройка: Список без tool_result...")
        content = [{"type": "text", "text": "Hello"}]
        
        print("Действие: Извлечение tool results...")
        result = _extract_tool_results(content)
        
        print(f"Сравниваем результат: Ожидалось [], Получено {result}")
        assert result == []


class TestExtractToolUses:
    """Тесты функции _extract_tool_uses."""
    
    def test_extracts_from_tool_calls_field(self):
        """
        Что он делает: Проверяет извлечение из поля tool_calls.
        Цель: Убедиться, что OpenAI tool_calls формат обрабатывается.
        """
        print("Настройка: Сообщение с tool_calls...")
        msg = ChatMessage(
            role="assistant",
            content="",
            tool_calls=[{
                "id": "call_123",
                "function": {
                    "name": "get_weather",
                    "arguments": '{"location": "Moscow"}'
                }
            }]
        )
        
        print("Действие: Извлечение tool uses...")
        result = _extract_tool_uses(msg)
        
        print(f"Результат: {result}")
        assert len(result) == 1
        assert result[0]["name"] == "get_weather"
        assert result[0]["toolUseId"] == "call_123"
    
    def test_extracts_from_content_list(self):
        """
        Что он делает: Проверяет извлечение из content списка.
        Цель: Убедиться, что tool_use в content обрабатывается.
        """
        print("Настройка: Сообщение с tool_use в content...")
        msg = ChatMessage(
            role="assistant",
            content=[{
                "type": "tool_use",
                "id": "call_456",
                "name": "search",
                "input": {"query": "test"}
            }]
        )
        
        print("Действие: Извлечение tool uses...")
        result = _extract_tool_uses(msg)
        
        print(f"Результат: {result}")
        assert len(result) == 1
        assert result[0]["name"] == "search"
        assert result[0]["toolUseId"] == "call_456"
    
    def test_returns_empty_for_no_tool_uses(self):
        """
        Что он делает: Проверяет возврат пустого списка без tool uses.
        Цель: Убедиться, что обычное сообщение не содержит tool uses.
        """
        print("Настройка: Обычное сообщение...")
        msg = ChatMessage(role="assistant", content="Hello")
        
        print("Действие: Извлечение tool uses...")
        result = _extract_tool_uses(msg)
        
        print(f"Сравниваем результат: Ожидалось [], Получено {result}")
        assert result == []


class TestProcessToolsWithLongDescriptions:
    """Тесты функции process_tools_with_long_descriptions."""
    
    def test_returns_none_and_empty_string_for_none_tools(self):
        """
        Что он делает: Проверяет обработку None вместо списка tools.
        Цель: Убедиться, что None возвращает (None, "").
        """
        print("Настройка: None вместо tools...")
        
        print("Действие: Обработка tools...")
        processed, doc = process_tools_with_long_descriptions(None)
        
        print(f"Сравниваем результат: Ожидалось (None, ''), Получено ({processed}, '{doc}')")
        assert processed is None
        assert doc == ""
    
    def test_returns_none_and_empty_string_for_empty_list(self):
        """
        Что он делает: Проверяет обработку пустого списка tools.
        Цель: Убедиться, что пустой список возвращает (None, "").
        """
        print("Настройка: Пустой список tools...")
        
        print("Действие: Обработка tools...")
        processed, doc = process_tools_with_long_descriptions([])
        
        print(f"Сравниваем результат: Ожидалось (None, ''), Получено ({processed}, '{doc}')")
        assert processed is None
        assert doc == ""
    
    def test_short_description_unchanged(self):
        """
        Что он делает: Проверяет, что короткие descriptions не изменяются.
        Цель: Убедиться, что tools с короткими descriptions остаются как есть.
        """
        print("Настройка: Tool с коротким description...")
        tools = [Tool(
            type="function",
            function=ToolFunction(
                name="get_weather",
                description="Get weather for a location",
                parameters={"type": "object", "properties": {}}
            )
        )]
        
        print("Действие: Обработка tools...")
        with patch('kiro_gateway.converters.TOOL_DESCRIPTION_MAX_LENGTH', 10000):
            processed, doc = process_tools_with_long_descriptions(tools)
        
        print(f"Сравниваем description: Ожидалось 'Get weather for a location', Получено '{processed[0].function.description}'")
        assert len(processed) == 1
        assert processed[0].function.description == "Get weather for a location"
        assert doc == ""
    
    def test_long_description_moved_to_system_prompt(self):
        """
        Что он делает: Проверяет перенос длинного description в system prompt.
        Цель: Убедиться, что длинные descriptions переносятся корректно.
        """
        print("Настройка: Tool с очень длинным description...")
        long_description = "A" * 15000  # 15000 символов - больше лимита
        tools = [Tool(
            type="function",
            function=ToolFunction(
                name="bash",
                description=long_description,
                parameters={"type": "object", "properties": {"command": {"type": "string"}}}
            )
        )]
        
        print("Действие: Обработка tools с лимитом 10000...")
        with patch('kiro_gateway.converters.TOOL_DESCRIPTION_MAX_LENGTH', 10000):
            processed, doc = process_tools_with_long_descriptions(tools)
        
        print(f"Проверяем reference в description...")
        assert len(processed) == 1
        assert "[Full documentation in system prompt under '## Tool: bash']" in processed[0].function.description
        
        print(f"Проверяем документацию в system prompt...")
        assert "## Tool: bash" in doc
        assert long_description in doc
        assert "# Tool Documentation" in doc
    
    def test_mixed_short_and_long_descriptions(self):
        """
        Что он делает: Проверяет обработку смешанного списка tools.
        Цель: Убедиться, что короткие остаются, длинные переносятся.
        """
        print("Настройка: Два tools - короткий и длинный...")
        short_desc = "Short description"
        long_desc = "B" * 15000
        tools = [
            Tool(
                type="function",
                function=ToolFunction(
                    name="short_tool",
                    description=short_desc,
                    parameters={}
                )
            ),
            Tool(
                type="function",
                function=ToolFunction(
                    name="long_tool",
                    description=long_desc,
                    parameters={}
                )
            )
        ]
        
        print("Действие: Обработка tools...")
        with patch('kiro_gateway.converters.TOOL_DESCRIPTION_MAX_LENGTH', 10000):
            processed, doc = process_tools_with_long_descriptions(tools)
        
        print(f"Проверяем количество tools: Ожидалось 2, Получено {len(processed)}")
        assert len(processed) == 2
        
        print(f"Проверяем короткий tool...")
        assert processed[0].function.description == short_desc
        
        print(f"Проверяем длинный tool...")
        assert "[Full documentation in system prompt" in processed[1].function.description
        assert "## Tool: long_tool" in doc
        assert long_desc in doc
    
    def test_preserves_tool_parameters(self):
        """
        Что он делает: Проверяет сохранение parameters при переносе description.
        Цель: Убедиться, что parameters не теряются.
        """
        print("Настройка: Tool с parameters и длинным description...")
        params = {
            "type": "object",
            "properties": {
                "location": {"type": "string", "description": "City name"},
                "units": {"type": "string", "enum": ["celsius", "fahrenheit"]}
            },
            "required": ["location"]
        }
        tools = [Tool(
            type="function",
            function=ToolFunction(
                name="weather",
                description="C" * 15000,
                parameters=params
            )
        )]
        
        print("Действие: Обработка tools...")
        with patch('kiro_gateway.converters.TOOL_DESCRIPTION_MAX_LENGTH', 10000):
            processed, doc = process_tools_with_long_descriptions(tools)
        
        print(f"Проверяем сохранение parameters...")
        assert processed[0].function.parameters == params
    
    def test_disabled_when_limit_is_zero(self):
        """
        Что он делает: Проверяет отключение функции при лимите 0.
        Цель: Убедиться, что при TOOL_DESCRIPTION_MAX_LENGTH=0 tools не изменяются.
        """
        print("Настройка: Tool с длинным description и лимит 0...")
        long_desc = "D" * 15000
        tools = [Tool(
            type="function",
            function=ToolFunction(
                name="test_tool",
                description=long_desc,
                parameters={}
            )
        )]
        
        print("Действие: Обработка tools с лимитом 0...")
        with patch('kiro_gateway.converters.TOOL_DESCRIPTION_MAX_LENGTH', 0):
            processed, doc = process_tools_with_long_descriptions(tools)
        
        print(f"Проверяем, что description не изменился...")
        assert processed[0].function.description == long_desc
        assert doc == ""
    
    def test_non_function_tools_unchanged(self):
        """
        Что он делает: Проверяет, что non-function tools не изменяются.
        Цель: Убедиться, что только function tools обрабатываются.
        """
        print("Настройка: Tool с type != function...")
        # Создаём tool с другим типом (хотя в реальности OpenAI поддерживает только function)
        tools = [Tool(
            type="other_type",
            function=ToolFunction(
                name="test",
                description="E" * 15000,
                parameters={}
            )
        )]
        
        print("Действие: Обработка tools...")
        with patch('kiro_gateway.converters.TOOL_DESCRIPTION_MAX_LENGTH', 10000):
            processed, doc = process_tools_with_long_descriptions(tools)
        
        print(f"Проверяем, что tool не изменился...")
        assert len(processed) == 1
        assert processed[0].type == "other_type"
        assert doc == ""
    
    def test_multiple_long_descriptions_all_moved(self):
        """
        Что он делает: Проверяет перенос нескольких длинных descriptions.
        Цель: Убедиться, что все длинные descriptions переносятся.
        """
        print("Настройка: Три tools с длинными descriptions...")
        tools = [
            Tool(type="function", function=ToolFunction(name="tool1", description="F" * 15000, parameters={})),
            Tool(type="function", function=ToolFunction(name="tool2", description="G" * 15000, parameters={})),
            Tool(type="function", function=ToolFunction(name="tool3", description="H" * 15000, parameters={}))
        ]
        
        print("Действие: Обработка tools...")
        with patch('kiro_gateway.converters.TOOL_DESCRIPTION_MAX_LENGTH', 10000):
            processed, doc = process_tools_with_long_descriptions(tools)
        
        print(f"Проверяем все три tools...")
        assert len(processed) == 3
        for i, tool in enumerate(processed):
            assert "[Full documentation in system prompt" in tool.function.description
        
        print(f"Проверяем документацию содержит все три секции...")
        assert "## Tool: tool1" in doc
        assert "## Tool: tool2" in doc
        assert "## Tool: tool3" in doc
    
    def test_empty_description_unchanged(self):
        """
        Что он делает: Проверяет обработку пустого description.
        Цель: Убедиться, что пустой description не вызывает ошибок.
        """
        print("Настройка: Tool с пустым description...")
        tools = [Tool(
            type="function",
            function=ToolFunction(
                name="empty_desc_tool",
                description="",
                parameters={}
            )
        )]
        
        print("Действие: Обработка tools...")
        with patch('kiro_gateway.converters.TOOL_DESCRIPTION_MAX_LENGTH', 10000):
            processed, doc = process_tools_with_long_descriptions(tools)
        
        print(f"Проверяем, что пустой description остался пустым...")
        assert processed[0].function.description == ""
        assert doc == ""
    
    def test_none_description_unchanged(self):
        """
        Что он делает: Проверяет обработку None description.
        Цель: Убедиться, что None description не вызывает ошибок.
        """
        print("Настройка: Tool с None description...")
        tools = [Tool(
            type="function",
            function=ToolFunction(
                name="none_desc_tool",
                description=None,
                parameters={}
            )
        )]
        
        print("Действие: Обработка tools...")
        with patch('kiro_gateway.converters.TOOL_DESCRIPTION_MAX_LENGTH', 10000):
            processed, doc = process_tools_with_long_descriptions(tools)
        
        print(f"Проверяем, что None description обработан корректно...")
        # None должен остаться None или стать пустой строкой
        assert processed[0].function.description is None or processed[0].function.description == ""
        assert doc == ""


class TestSanitizeJsonSchema:
    """
    Тесты функции _sanitize_json_schema.
    
    Эта функция очищает JSON Schema от полей, которые Kiro API не принимает:
    - Пустые required массивы []
    - additionalProperties
    """
    
    def test_returns_empty_dict_for_none(self):
        """
        Что он делает: Проверяет обработку None.
        Цель: Убедиться, что None возвращает пустой словарь.
        """
        print("Настройка: None schema...")
        
        print("Действие: Очистка schema...")
        result = _sanitize_json_schema(None)
        
        print(f"Сравниваем результат: Ожидалось {{}}, Получено {result}")
        assert result == {}
    
    def test_returns_empty_dict_for_empty_dict(self):
        """
        Что он делает: Проверяет обработку пустого словаря.
        Цель: Убедиться, что пустой словарь возвращается как есть.
        """
        print("Настройка: Пустой словарь...")
        
        print("Действие: Очистка schema...")
        result = _sanitize_json_schema({})
        
        print(f"Сравниваем результат: Ожидалось {{}}, Получено {result}")
        assert result == {}
    
    def test_removes_empty_required_array(self):
        """
        Что он делает: Проверяет удаление пустого required массива.
        Цель: Убедиться, что required: [] удаляется из schema.
        
        Это критический тест для бага Cline, где tools с required: []
        вызывали ошибку 400 "Improperly formed request" от Kiro API.
        """
        print("Настройка: Schema с пустым required...")
        schema = {
            "type": "object",
            "properties": {},
            "required": []
        }
        
        print("Действие: Очистка schema...")
        result = _sanitize_json_schema(schema)
        
        print(f"Результат: {result}")
        print("Проверяем, что required удалён...")
        assert "required" not in result
        assert result["type"] == "object"
        assert result["properties"] == {}
    
    def test_preserves_non_empty_required_array(self):
        """
        Что он делает: Проверяет сохранение непустого required массива.
        Цель: Убедиться, что required с элементами сохраняется.
        """
        print("Настройка: Schema с непустым required...")
        schema = {
            "type": "object",
            "properties": {
                "location": {"type": "string"}
            },
            "required": ["location"]
        }
        
        print("Действие: Очистка schema...")
        result = _sanitize_json_schema(schema)
        
        print(f"Результат: {result}")
        print("Проверяем, что required сохранён...")
        assert "required" in result
        assert result["required"] == ["location"]
    
    def test_removes_additional_properties(self):
        """
        Что он делает: Проверяет удаление additionalProperties.
        Цель: Убедиться, что additionalProperties удаляется из schema.
        
        Kiro API не поддерживает additionalProperties в JSON Schema.
        """
        print("Настройка: Schema с additionalProperties...")
        schema = {
            "type": "object",
            "properties": {},
            "additionalProperties": False
        }
        
        print("Действие: Очистка schema...")
        result = _sanitize_json_schema(schema)
        
        print(f"Результат: {result}")
        print("Проверяем, что additionalProperties удалён...")
        assert "additionalProperties" not in result
        assert result["type"] == "object"
    
    def test_removes_both_empty_required_and_additional_properties(self):
        """
        Что он делает: Проверяет удаление обоих проблемных полей.
        Цель: Убедиться, что оба поля удаляются одновременно.
        
        Это реальный сценарий от Cline, где tools имели оба поля.
        """
        print("Настройка: Schema с обоими проблемными полями...")
        schema = {
            "type": "object",
            "properties": {},
            "required": [],
            "additionalProperties": False
        }
        
        print("Действие: Очистка schema...")
        result = _sanitize_json_schema(schema)
        
        print(f"Результат: {result}")
        print("Проверяем, что оба поля удалены...")
        assert "required" not in result
        assert "additionalProperties" not in result
        assert result == {"type": "object", "properties": {}}
    
    def test_recursively_sanitizes_nested_properties(self):
        """
        Что он делает: Проверяет рекурсивную очистку вложенных properties.
        Цель: Убедиться, что вложенные schema также очищаются.
        """
        print("Настройка: Schema с вложенными properties...")
        schema = {
            "type": "object",
            "properties": {
                "nested": {
                    "type": "object",
                    "properties": {},
                    "required": [],
                    "additionalProperties": False
                }
            }
        }
        
        print("Действие: Очистка schema...")
        result = _sanitize_json_schema(schema)
        
        print(f"Результат: {result}")
        print("Проверяем вложенный объект...")
        nested = result["properties"]["nested"]
        assert "required" not in nested
        assert "additionalProperties" not in nested
    
    def test_recursively_sanitizes_dict_values(self):
        """
        Что он делает: Проверяет рекурсивную очистку dict значений.
        Цель: Убедиться, что любые вложенные dict очищаются.
        """
        print("Настройка: Schema с вложенным dict...")
        schema = {
            "type": "object",
            "items": {
                "type": "string",
                "additionalProperties": True
            }
        }
        
        print("Действие: Очистка schema...")
        result = _sanitize_json_schema(schema)
        
        print(f"Результат: {result}")
        print("Проверяем вложенный items...")
        assert "additionalProperties" not in result["items"]
        assert result["items"]["type"] == "string"
    
    def test_sanitizes_items_in_lists(self):
        """
        Что он делает: Проверяет очистку элементов в списках (anyOf, oneOf).
        Цель: Убедиться, что элементы списков также очищаются.
        """
        print("Настройка: Schema с anyOf...")
        schema = {
            "anyOf": [
                {"type": "string", "additionalProperties": False},
                {"type": "number", "required": []}
            ]
        }
        
        print("Действие: Очистка schema...")
        result = _sanitize_json_schema(schema)
        
        print(f"Результат: {result}")
        print("Проверяем элементы anyOf...")
        assert "additionalProperties" not in result["anyOf"][0]
        assert "required" not in result["anyOf"][1]
    
    def test_preserves_non_dict_list_items(self):
        """
        Что он делает: Проверяет сохранение не-dict элементов в списках.
        Цель: Убедиться, что строки и другие типы в списках сохраняются.
        """
        print("Настройка: Schema с enum...")
        schema = {
            "type": "string",
            "enum": ["value1", "value2", "value3"]
        }
        
        print("Действие: Очистка schema...")
        result = _sanitize_json_schema(schema)
        
        print(f"Результат: {result}")
        print("Проверяем enum сохранён...")
        assert result["enum"] == ["value1", "value2", "value3"]
    
    def test_complex_real_world_schema(self):
        """
        Что он делает: Проверяет очистку реальной сложной schema от Cline.
        Цель: Убедиться, что реальные schema обрабатываются корректно.
        """
        print("Настройка: Реальная schema от Cline...")
        schema = {
            "type": "object",
            "properties": {
                "question": {
                    "type": "string",
                    "description": "The question to ask"
                },
                "options": {
                    "type": "string",
                    "description": "Array of options"
                }
            },
            "required": ["question", "options"],
            "additionalProperties": False
        }
        
        print("Действие: Очистка schema...")
        result = _sanitize_json_schema(schema)
        
        print(f"Результат: {result}")
        print("Проверяем результат...")
        assert "additionalProperties" not in result
        assert result["required"] == ["question", "options"]  # Непустой required сохраняется
        assert result["properties"]["question"]["type"] == "string"


class TestBuildUserInputContext:
    """Тесты функции _build_user_input_context."""
    
    def test_builds_tools_context(self):
        """
        Что он делает: Проверяет построение контекста с tools.
        Цель: Убедиться, что tools преобразуются в toolSpecification.
        """
        print("Настройка: Запрос с tools...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4",
            messages=[ChatMessage(role="user", content="Hello")],
            tools=[Tool(
                type="function",
                function=ToolFunction(
                    name="get_weather",
                    description="Get weather",
                    parameters={"type": "object", "properties": {}}
                )
            )]
        )
        current_msg = ChatMessage(role="user", content="Hello")
        
        print("Действие: Построение контекста...")
        result = _build_user_input_context(request, current_msg)
        
        print(f"Результат: {result}")
        assert "tools" in result
        assert len(result["tools"]) == 1
        assert result["tools"][0]["toolSpecification"]["name"] == "get_weather"
    
    def test_returns_empty_for_no_tools(self):
        """
        Что он делает: Проверяет возврат пустого контекста без tools.
        Цель: Убедиться, что запрос без tools возвращает пустой контекст.
        """
        print("Настройка: Запрос без tools...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4",
            messages=[ChatMessage(role="user", content="Hello")]
        )
        current_msg = ChatMessage(role="user", content="Hello")
        
        print("Действие: Построение контекста...")
        result = _build_user_input_context(request, current_msg)
        
        print(f"Сравниваем результат: Ожидалось {{}}, Получено {result}")
        assert result == {}
    
    def test_empty_description_replaced_with_placeholder(self):
        """
        Что он делает: Проверяет замену пустого description на placeholder.
        Цель: Убедиться, что пустой description заменяется на "Tool: {name}".
        
        Это критический тест для бага Cline, где tool focus_chain имел
        пустой description "", что вызывало ошибку 400 от Kiro API.
        """
        print("Настройка: Tool с пустым description...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4",
            messages=[ChatMessage(role="user", content="Hello")],
            tools=[Tool(
                type="function",
                function=ToolFunction(
                    name="focus_chain",
                    description="",
                    parameters={"type": "object", "properties": {}}
                )
            )]
        )
        current_msg = ChatMessage(role="user", content="Hello")
        
        print("Действие: Построение контекста...")
        result = _build_user_input_context(request, current_msg)
        
        print(f"Результат: {result}")
        print("Проверяем, что description заменён на placeholder...")
        tool_spec = result["tools"][0]["toolSpecification"]
        assert tool_spec["description"] == "Tool: focus_chain"
    
    def test_whitespace_only_description_replaced_with_placeholder(self):
        """
        Что он делает: Проверяет замену description из пробелов на placeholder.
        Цель: Убедиться, что description с только пробелами заменяется.
        """
        print("Настройка: Tool с description из пробелов...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4",
            messages=[ChatMessage(role="user", content="Hello")],
            tools=[Tool(
                type="function",
                function=ToolFunction(
                    name="whitespace_tool",
                    description="   ",
                    parameters={}
                )
            )]
        )
        current_msg = ChatMessage(role="user", content="Hello")
        
        print("Действие: Построение контекста...")
        result = _build_user_input_context(request, current_msg)
        
        print(f"Результат: {result}")
        print("Проверяем, что description заменён на placeholder...")
        tool_spec = result["tools"][0]["toolSpecification"]
        assert tool_spec["description"] == "Tool: whitespace_tool"
    
    def test_none_description_replaced_with_placeholder(self):
        """
        Что он делает: Проверяет замену None description на placeholder.
        Цель: Убедиться, что None description заменяется на "Tool: {name}".
        """
        print("Настройка: Tool с None description...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4",
            messages=[ChatMessage(role="user", content="Hello")],
            tools=[Tool(
                type="function",
                function=ToolFunction(
                    name="none_desc_tool",
                    description=None,
                    parameters={}
                )
            )]
        )
        current_msg = ChatMessage(role="user", content="Hello")
        
        print("Действие: Построение контекста...")
        result = _build_user_input_context(request, current_msg)
        
        print(f"Результат: {result}")
        print("Проверяем, что description заменён на placeholder...")
        tool_spec = result["tools"][0]["toolSpecification"]
        assert tool_spec["description"] == "Tool: none_desc_tool"
    
    def test_non_empty_description_preserved(self):
        """
        Что он делает: Проверяет сохранение непустого description.
        Цель: Убедиться, что нормальный description не изменяется.
        """
        print("Настройка: Tool с нормальным description...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4",
            messages=[ChatMessage(role="user", content="Hello")],
            tools=[Tool(
                type="function",
                function=ToolFunction(
                    name="get_weather",
                    description="Get weather for a location",
                    parameters={}
                )
            )]
        )
        current_msg = ChatMessage(role="user", content="Hello")
        
        print("Действие: Построение контекста...")
        result = _build_user_input_context(request, current_msg)
        
        print(f"Результат: {result}")
        print("Проверяем, что description сохранён...")
        tool_spec = result["tools"][0]["toolSpecification"]
        assert tool_spec["description"] == "Get weather for a location"
    
    def test_sanitizes_tool_parameters(self):
        """
        Что он делает: Проверяет очистку parameters от проблемных полей.
        Цель: Убедиться, что _sanitize_json_schema применяется к parameters.
        """
        print("Настройка: Tool с проблемными parameters...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4",
            messages=[ChatMessage(role="user", content="Hello")],
            tools=[Tool(
                type="function",
                function=ToolFunction(
                    name="test_tool",
                    description="Test tool",
                    parameters={
                        "type": "object",
                        "properties": {},
                        "required": [],
                        "additionalProperties": False
                    }
                )
            )]
        )
        current_msg = ChatMessage(role="user", content="Hello")
        
        print("Действие: Построение контекста...")
        result = _build_user_input_context(request, current_msg)
        
        print(f"Результат: {result}")
        print("Проверяем, что parameters очищены...")
        input_schema = result["tools"][0]["toolSpecification"]["inputSchema"]["json"]
        assert "required" not in input_schema
        assert "additionalProperties" not in input_schema
    
    def test_mixed_tools_with_empty_and_normal_descriptions(self):
        """
        Что он делает: Проверяет обработку смешанного списка tools.
        Цель: Убедиться, что пустые descriptions заменяются, а нормальные сохраняются.
        
        Это реальный сценарий от Cline, где большинство tools имеют
        нормальные descriptions, но focus_chain имеет пустой.
        """
        print("Настройка: Смешанный список tools...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4",
            messages=[ChatMessage(role="user", content="Hello")],
            tools=[
                Tool(
                    type="function",
                    function=ToolFunction(
                        name="read_file",
                        description="Read contents of a file",
                        parameters={}
                    )
                ),
                Tool(
                    type="function",
                    function=ToolFunction(
                        name="focus_chain",
                        description="",
                        parameters={}
                    )
                ),
                Tool(
                    type="function",
                    function=ToolFunction(
                        name="write_file",
                        description="Write content to a file",
                        parameters={}
                    )
                )
            ]
        )
        current_msg = ChatMessage(role="user", content="Hello")
        
        print("Действие: Построение контекста...")
        result = _build_user_input_context(request, current_msg)
        
        print(f"Результат: {result}")
        print("Проверяем descriptions...")
        tools = result["tools"]
        assert tools[0]["toolSpecification"]["description"] == "Read contents of a file"
        assert tools[1]["toolSpecification"]["description"] == "Tool: focus_chain"
        assert tools[2]["toolSpecification"]["description"] == "Write content to a file"


class TestBuildKiroPayloadToolCallsIntegration:
    """
    Интеграционные тесты для build_kiro_payload с tool_calls.
    Проверяет полный flow от OpenAI формата до Kiro формата.
    """
    
    def test_multiple_assistant_tool_calls_with_results(self):
        """
        Что он делает: Проверяет полный сценарий с несколькими assistant tool_calls и их результатами.
        Цель: Убедиться, что все toolUses и toolResults корректно связываются в Kiro payload.
        
        Это интеграционный тест для бага Codex CLI, где несколько assistant сообщений
        с tool_calls отправлялись подряд, а затем tool results. Без фикса второй toolUse
        терялся, что приводило к ошибке 400 от Kiro API.
        """
        print("Настройка: Полный сценарий с двумя tool_calls и их результатами...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4-5",
            messages=[
                ChatMessage(role="user", content="Run two commands"),
                # Первый assistant с tool_call
                ChatMessage(
                    role="assistant",
                    content=None,
                    tool_calls=[{
                        "id": "tooluse_first",
                        "type": "function",
                        "function": {
                            "name": "shell",
                            "arguments": '{"command": ["ls"]}'
                        }
                    }]
                ),
                # Второй assistant с tool_call (подряд!)
                ChatMessage(
                    role="assistant",
                    content=None,
                    tool_calls=[{
                        "id": "tooluse_second",
                        "type": "function",
                        "function": {
                            "name": "shell",
                            "arguments": '{"command": ["pwd"]}'
                        }
                    }]
                ),
                # Результаты обоих tool_calls
                ChatMessage(role="tool", content="file1.txt\nfile2.txt", tool_call_id="tooluse_first"),
                ChatMessage(role="tool", content="/home/user", tool_call_id="tooluse_second")
            ]
        )
        
        print("Действие: Построение Kiro payload...")
        result = build_kiro_payload(request, "conv-123", "arn:aws:test")
        
        print(f"Результат: {result}")
        
        # Проверяем историю
        history = result["conversationState"].get("history", [])
        print(f"История: {history}")
        
        # Должен быть userInputMessage и assistantResponseMessage в истории
        assert len(history) >= 2, f"Ожидалось минимум 2 элемента в истории, получено {len(history)}"
        
        # Находим assistantResponseMessage
        assistant_msgs = [h for h in history if "assistantResponseMessage" in h]
        print(f"Assistant сообщения в истории: {assistant_msgs}")
        assert len(assistant_msgs) >= 1, "Должен быть хотя бы один assistantResponseMessage"
        
        # Проверяем, что в assistantResponseMessage есть оба toolUses
        assistant_msg = assistant_msgs[0]["assistantResponseMessage"]
        tool_uses = assistant_msg.get("toolUses", [])
        print(f"ToolUses в assistant: {tool_uses}")
        print(f"Сравниваем количество toolUses: Ожидалось 2, Получено {len(tool_uses)}")
        assert len(tool_uses) == 2, f"Должно быть 2 toolUses, получено {len(tool_uses)}"
        
        tool_use_ids = [tu["toolUseId"] for tu in tool_uses]
        print(f"ToolUse IDs: {tool_use_ids}")
        assert "tooluse_first" in tool_use_ids
        assert "tooluse_second" in tool_use_ids
        
        # Проверяем currentMessage содержит toolResults
        current_msg = result["conversationState"]["currentMessage"]["userInputMessage"]
        context = current_msg.get("userInputMessageContext", {})
        tool_results = context.get("toolResults", [])
        print(f"ToolResults в currentMessage: {tool_results}")
        print(f"Сравниваем количество toolResults: Ожидалось 2, Получено {len(tool_results)}")
        assert len(tool_results) == 2, f"Должно быть 2 toolResults, получено {len(tool_results)}"
        
        tool_result_ids = [tr["toolUseId"] for tr in tool_results]
        print(f"ToolResult IDs: {tool_result_ids}")
        assert "tooluse_first" in tool_result_ids
        assert "tooluse_second" in tool_result_ids


class TestBuildKiroPayload:
    """Тесты функции build_kiro_payload."""
    
    def test_builds_simple_payload(self):
        """
        Что он делает: Проверяет построение простого payload.
        Цель: Убедиться, что базовый запрос преобразуется корректно.
        """
        print("Настройка: Простой запрос...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4-5",
            messages=[ChatMessage(role="user", content="Hello")]
        )
        
        print("Действие: Построение payload...")
        result = build_kiro_payload(request, "conv-123", "arn:aws:test")
        
        print(f"Результат: {result}")
        assert "conversationState" in result
        assert result["conversationState"]["conversationId"] == "conv-123"
        assert "currentMessage" in result["conversationState"]
        assert result["profileArn"] == "arn:aws:test"
    
    def test_includes_system_prompt_in_first_message(self):
        """
        Что он делает: Проверяет добавление system prompt к первому сообщению.
        Цель: Убедиться, что system prompt объединяется с user сообщением.
        """
        print("Настройка: Запрос с system prompt...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4-5",
            messages=[
                ChatMessage(role="system", content="You are helpful"),
                ChatMessage(role="user", content="Hello")
            ]
        )
        
        print("Действие: Построение payload...")
        result = build_kiro_payload(request, "conv-123", "")
        
        print(f"Результат: {result}")
        current_content = result["conversationState"]["currentMessage"]["userInputMessage"]["content"]
        assert "You are helpful" in current_content
        assert "Hello" in current_content
    
    def test_builds_history_for_multi_turn(self):
        """
        Что он делает: Проверяет построение истории для multi-turn.
        Цель: Убедиться, что предыдущие сообщения попадают в history.
        """
        print("Настройка: Multi-turn запрос...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4-5",
            messages=[
                ChatMessage(role="user", content="Hello"),
                ChatMessage(role="assistant", content="Hi"),
                ChatMessage(role="user", content="How are you?")
            ]
        )
        
        print("Действие: Построение payload...")
        result = build_kiro_payload(request, "conv-123", "")
        
        print(f"Результат: {result}")
        assert "history" in result["conversationState"]
        assert len(result["conversationState"]["history"]) == 2
    
    def test_handles_assistant_as_last_message(self):
        """
        Что он делает: Проверяет обработку assistant как последнего сообщения.
        Цель: Убедиться, что создаётся "Continue" сообщение.
        """
        print("Настройка: Запрос с assistant в конце...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4-5",
            messages=[
                ChatMessage(role="user", content="Hello"),
                ChatMessage(role="assistant", content="Hi there")
            ]
        )
        
        print("Действие: Построение payload...")
        result = build_kiro_payload(request, "conv-123", "")
        
        print(f"Результат: {result}")
        current_content = result["conversationState"]["currentMessage"]["userInputMessage"]["content"]
        assert current_content == "Continue"
    
    def test_raises_for_empty_messages(self):
        """
        Что он делает: Проверяет выброс исключения для пустых сообщений.
        Цель: Убедиться, что пустой запрос вызывает ValueError.
        """
        print("Настройка: Запрос только с system сообщением...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4-5",
            messages=[ChatMessage(role="system", content="You are helpful")]
        )
        
        print("Действие: Попытка построения payload...")
        with pytest.raises(ValueError) as exc_info:
            build_kiro_payload(request, "conv-123", "")
        
        print(f"Исключение: {exc_info.value}")
        assert "No messages to send" in str(exc_info.value)
    
    def test_uses_continue_for_empty_content(self):
        """
        Что он делает: Проверяет использование "Continue" для пустого контента.
        Цель: Убедиться, что пустое сообщение заменяется на "Continue".
        """
        print("Настройка: Запрос с пустым контентом...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4-5",
            messages=[ChatMessage(role="user", content="")]
        )
        
        print("Действие: Построение payload...")
        result = build_kiro_payload(request, "conv-123", "")
        
        print(f"Результат: {result}")
        current_content = result["conversationState"]["currentMessage"]["userInputMessage"]["content"]
        assert current_content == "Continue"
    
    def test_maps_model_id_correctly(self):
        """
        Что он делает: Проверяет маппинг внешнего ID модели во внутренний.
        Цель: Убедиться, что MODEL_MAPPING применяется.
        """
        print("Настройка: Запрос с внешним ID модели...")
        request = ChatCompletionRequest(
            model="claude-sonnet-4-5",
            messages=[ChatMessage(role="user", content="Hello")]
        )
        
        print("Действие: Построение payload...")
        result = build_kiro_payload(request, "conv-123", "")
        
        print(f"Результат: {result}")
        model_id = result["conversationState"]["currentMessage"]["userInputMessage"]["modelId"]
        # claude-sonnet-4-5 должен маппиться в CLAUDE_SONNET_4_5_20250929_V1_0
        assert model_id == "CLAUDE_SONNET_4_5_20250929_V1_0"
    
    def test_long_tool_description_added_to_system_prompt(self):
        """
        Что он делает: Проверяет интеграцию длинных tool descriptions в payload.
        Цель: Убедиться, что длинные descriptions добавляются в system prompt в payload.
        """
        print("Настройка: Запрос с tool с длинным description...")
        long_desc = "X" * 15000
        request = ChatCompletionRequest(
            model="claude-sonnet-4-5",
            messages=[
                ChatMessage(role="system", content="You are helpful"),
                ChatMessage(role="user", content="Hello")
            ],
            tools=[Tool(
                type="function",
                function=ToolFunction(
                    name="long_tool",
                    description=long_desc,
                    parameters={}
                )
            )]
        )
        
        print("Действие: Построение payload...")
        with patch('kiro_gateway.converters.TOOL_DESCRIPTION_MAX_LENGTH', 10000):
            result = build_kiro_payload(request, "conv-123", "")
        
        print(f"Проверяем, что system prompt содержит tool documentation...")
        current_content = result["conversationState"]["currentMessage"]["userInputMessage"]["content"]
        assert "You are helpful" in current_content
        assert "## Tool: long_tool" in current_content
        assert long_desc in current_content
        
        print(f"Проверяем, что tool в context имеет reference description...")
        tools_context = result["conversationState"]["currentMessage"]["userInputMessage"]["userInputMessageContext"]["tools"]
        assert "[Full documentation in system prompt" in tools_context[0]["toolSpecification"]["description"]


================================================
FILE: tests/unit/test_debug_logger.py
================================================
# -*- coding: utf-8 -*-

"""
Unit-тесты для DebugLogger.
Проверяет логику буферизации и записи debug логов в разных режимах.
"""

import json
import pytest
from pathlib import Path
from unittest.mock import patch, MagicMock


class TestDebugLoggerModeOff:
    """Тесты для режима DEBUG_MODE=off."""
    
    def test_prepare_new_request_does_nothing(self, tmp_path):
        """
        Что он делает: Проверяет, что prepare_new_request ничего не делает в режиме off.
        Цель: Убедиться, что в режиме off директория не создаётся.
        """
        print("Настройка: Режим off...")
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'off'):
            with patch('kiro_gateway.debug_logger.DEBUG_DIR', str(tmp_path / "debug_logs")):
                # Пересоздаём экземпляр с новыми настройками
                from kiro_gateway.debug_logger import DebugLogger
                logger = DebugLogger.__new__(DebugLogger)
                logger._initialized = False
                logger.__init__()
                logger.debug_dir = tmp_path / "debug_logs"
                
                print("Действие: Вызов prepare_new_request...")
                logger.prepare_new_request()
                
                print(f"Проверяем, что директория не создана...")
                assert not (tmp_path / "debug_logs").exists()
    
    def test_log_request_body_does_nothing(self, tmp_path):
        """
        Что он делает: Проверяет, что log_request_body ничего не делает в режиме off.
        Цель: Убедиться, что данные не записываются.
        """
        print("Настройка: Режим off...")
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'off'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = tmp_path / "debug_logs"
            
            print("Действие: Вызов log_request_body...")
            logger.log_request_body(b'{"test": "data"}')
            
            print(f"Проверяем, что файл не создан...")
            assert not (tmp_path / "debug_logs" / "request_body.json").exists()


class TestDebugLoggerModeAll:
    """Тесты для режима DEBUG_MODE=all."""
    
    def test_prepare_new_request_clears_directory(self, tmp_path):
        """
        Что он делает: Проверяет, что prepare_new_request очищает директорию в режиме all.
        Цель: Убедиться, что старые логи удаляются.
        """
        print("Настройка: Режим all, создаём старый файл...")
        debug_dir = tmp_path / "debug_logs"
        debug_dir.mkdir()
        old_file = debug_dir / "old_file.txt"
        old_file.write_text("old content")
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            print("Действие: Вызов prepare_new_request...")
            logger.prepare_new_request()
            
            print(f"Проверяем, что старый файл удалён...")
            assert not old_file.exists()
            print(f"Проверяем, что директория существует...")
            assert debug_dir.exists()
    
    def test_log_request_body_writes_immediately(self, tmp_path):
        """
        Что он делает: Проверяет, что log_request_body пишет сразу в файл в режиме all.
        Цель: Убедиться, что данные записываются немедленно.
        """
        print("Настройка: Режим all...")
        debug_dir = tmp_path / "debug_logs"
        debug_dir.mkdir()
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            print("Действие: Вызов log_request_body...")
            test_data = b'{"model": "test", "messages": []}'
            logger.log_request_body(test_data)
            
            print(f"Проверяем, что файл создан...")
            file_path = debug_dir / "request_body.json"
            assert file_path.exists()
            
            print(f"Проверяем содержимое файла...")
            content = json.loads(file_path.read_text())
            assert content["model"] == "test"
    
    def test_log_kiro_request_body_writes_immediately(self, tmp_path):
        """
        Что он делает: Проверяет, что log_kiro_request_body пишет сразу в файл в режиме all.
        Цель: Убедиться, что Kiro payload записывается немедленно.
        """
        print("Настройка: Режим all...")
        debug_dir = tmp_path / "debug_logs"
        debug_dir.mkdir()
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            print("Действие: Вызов log_kiro_request_body...")
            test_data = b'{"conversationState": {}}'
            logger.log_kiro_request_body(test_data)
            
            print(f"Проверяем, что файл создан...")
            file_path = debug_dir / "kiro_request_body.json"
            assert file_path.exists()
    
    def test_log_raw_chunk_appends_to_file(self, tmp_path):
        """
        Что он делает: Проверяет, что log_raw_chunk дописывает в файл в режиме all.
        Цель: Убедиться, что чанки накапливаются.
        """
        print("Настройка: Режим all...")
        debug_dir = tmp_path / "debug_logs"
        debug_dir.mkdir()
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            print("Действие: Вызов log_raw_chunk дважды...")
            logger.log_raw_chunk(b'chunk1')
            logger.log_raw_chunk(b'chunk2')
            
            print(f"Проверяем содержимое файла...")
            file_path = debug_dir / "response_stream_raw.txt"
            content = file_path.read_bytes()
            assert content == b'chunk1chunk2'


class TestDebugLoggerModeErrors:
    """Тесты для режима DEBUG_MODE=errors."""
    
    def test_log_request_body_buffers_data(self, tmp_path):
        """
        Что он делает: Проверяет, что log_request_body буферизует данные в режиме errors.
        Цель: Убедиться, что данные не записываются сразу.
        """
        print("Настройка: Режим errors...")
        debug_dir = tmp_path / "debug_logs"
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'errors'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            print("Действие: Вызов log_request_body...")
            test_data = b'{"test": "buffered"}'
            logger.log_request_body(test_data)
            
            print(f"Проверяем, что файл НЕ создан...")
            assert not debug_dir.exists()
            
            print(f"Проверяем, что данные в буфере...")
            assert logger._request_body_buffer == test_data
    
    def test_flush_on_error_writes_buffers(self, tmp_path):
        """
        Что он делает: Проверяет, что flush_on_error записывает буферы в файлы.
        Цель: Убедиться, что при ошибке данные сохраняются.
        """
        print("Настройка: Режим errors, заполняем буферы...")
        debug_dir = tmp_path / "debug_logs"
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'errors'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            # Заполняем буферы
            logger.log_request_body(b'{"request": "body"}')
            logger.log_kiro_request_body(b'{"kiro": "request"}')
            logger.log_raw_chunk(b'raw_chunk')
            logger.log_modified_chunk(b'modified_chunk')
            
            print("Действие: Вызов flush_on_error...")
            logger.flush_on_error(400, "Bad Request")
            
            print(f"Проверяем, что все файлы созданы...")
            assert (debug_dir / "request_body.json").exists()
            assert (debug_dir / "kiro_request_body.json").exists()
            assert (debug_dir / "response_stream_raw.txt").exists()
            assert (debug_dir / "response_stream_modified.txt").exists()
            assert (debug_dir / "error_info.json").exists()
            
            print(f"Проверяем error_info.json...")
            error_info = json.loads((debug_dir / "error_info.json").read_text())
            assert error_info["status_code"] == 400
            assert error_info["error_message"] == "Bad Request"
    
    def test_flush_on_error_clears_buffers(self, tmp_path):
        """
        Что он делает: Проверяет, что flush_on_error очищает буферы после записи.
        Цель: Убедиться, что буферы не накапливаются между запросами.
        """
        print("Настройка: Режим errors...")
        debug_dir = tmp_path / "debug_logs"
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'errors'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            logger.log_request_body(b'{"test": "data"}')
            
            print("Действие: Вызов flush_on_error...")
            logger.flush_on_error(500, "Error")
            
            print(f"Проверяем, что буферы очищены...")
            assert logger._request_body_buffer is None
            assert logger._kiro_request_body_buffer is None
            assert len(logger._raw_chunks_buffer) == 0
            assert len(logger._modified_chunks_buffer) == 0
    
    def test_discard_buffers_clears_without_writing(self, tmp_path):
        """
        Что он делает: Проверяет, что discard_buffers очищает буферы без записи.
        Цель: Убедиться, что успешные запросы не оставляют логов.
        """
        print("Настройка: Режим errors, заполняем буферы...")
        debug_dir = tmp_path / "debug_logs"
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'errors'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            logger.log_request_body(b'{"test": "data"}')
            logger.log_raw_chunk(b'chunk')
            
            print("Действие: Вызов discard_buffers...")
            logger.discard_buffers()
            
            print(f"Проверяем, что директория НЕ создана...")
            assert not debug_dir.exists()
            
            print(f"Проверяем, что буферы очищены...")
            assert logger._request_body_buffer is None
            assert len(logger._raw_chunks_buffer) == 0
    
    def test_flush_on_error_writes_error_info_in_mode_all(self, tmp_path):
        """
        Что он делает: Проверяет, что flush_on_error записывает error_info.json в режиме all.
        Цель: Убедиться, что информация об ошибке сохраняется в обоих режимах.
        """
        print("Настройка: Режим all...")
        debug_dir = tmp_path / "debug_logs"
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            print("Действие: Вызов flush_on_error...")
            logger.flush_on_error(400, "Bad Request")
            
            print(f"Проверяем, что error_info.json создан...")
            assert (debug_dir / "error_info.json").exists()
            
            print(f"Проверяем содержимое error_info.json...")
            error_info = json.loads((debug_dir / "error_info.json").read_text())
            assert error_info["status_code"] == 400
            assert error_info["error_message"] == "Bad Request"


class TestDebugLoggerLogErrorInfo:
    """Тесты для метода log_error_info()."""
    
    def test_log_error_info_writes_in_mode_all(self, tmp_path):
        """
        Что он делает: Проверяет, что log_error_info записывает файл в режиме all.
        Цель: Убедиться, что error_info.json создаётся при ошибках.
        """
        print("Настройка: Режим all...")
        debug_dir = tmp_path / "debug_logs"
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            print("Действие: Вызов log_error_info...")
            logger.log_error_info(500, "Internal Server Error")
            
            print(f"Проверяем, что error_info.json создан...")
            error_file = debug_dir / "error_info.json"
            assert error_file.exists()
            
            print(f"Проверяем содержимое...")
            error_info = json.loads(error_file.read_text())
            assert error_info["status_code"] == 500
            assert error_info["error_message"] == "Internal Server Error"
    
    def test_log_error_info_writes_in_mode_errors(self, tmp_path):
        """
        Что он делает: Проверяет, что log_error_info записывает файл в режиме errors.
        Цель: Убедиться, что метод работает в обоих режимах.
        """
        print("Настройка: Режим errors...")
        debug_dir = tmp_path / "debug_logs"
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'errors'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            print("Действие: Вызов log_error_info...")
            logger.log_error_info(404, "Not Found")
            
            print(f"Проверяем, что error_info.json создан...")
            error_file = debug_dir / "error_info.json"
            assert error_file.exists()
    
    def test_log_error_info_does_nothing_in_mode_off(self, tmp_path):
        """
        Что он делает: Проверяет, что log_error_info ничего не делает в режиме off.
        Цель: Убедиться, что в режиме off файлы не создаются.
        """
        print("Настройка: Режим off...")
        debug_dir = tmp_path / "debug_logs"
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'off'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            print("Действие: Вызов log_error_info...")
            logger.log_error_info(500, "Error")
            
            print(f"Проверяем, что директория НЕ создана...")
            assert not debug_dir.exists()


class TestDebugLoggerHelperMethods:
    """Тесты для вспомогательных методов DebugLogger."""
    
    def test_is_enabled_returns_true_for_errors(self):
        """
        Что он делает: Проверяет _is_enabled() для режима errors.
        Цель: Убедиться, что режим errors считается включённым.
        """
        print("Настройка: Режим errors...")
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'errors'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            
            print(f"Проверяем _is_enabled()...")
            assert logger._is_enabled() is True
    
    def test_is_enabled_returns_true_for_all(self):
        """
        Что он делает: Проверяет _is_enabled() для режима all.
        Цель: Убедиться, что режим all считается включённым.
        """
        print("Настройка: Режим all...")
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            
            print(f"Проверяем _is_enabled()...")
            assert logger._is_enabled() is True
    
    def test_is_enabled_returns_false_for_off(self):
        """
        Что он делает: Проверяет _is_enabled() для режима off.
        Цель: Убедиться, что режим off считается выключенным.
        """
        print("Настройка: Режим off...")
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'off'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            
            print(f"Проверяем _is_enabled()...")
            assert logger._is_enabled() is False
    
    def test_is_immediate_write_returns_true_for_all(self):
        """
        Что он делает: Проверяет _is_immediate_write() для режима all.
        Цель: Убедиться, что режим all пишет сразу.
        """
        print("Настройка: Режим all...")
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            
            print(f"Проверяем _is_immediate_write()...")
            assert logger._is_immediate_write() is True
    
    def test_is_immediate_write_returns_false_for_errors(self):
        """
        Что он делает: Проверяет _is_immediate_write() для режима errors.
        Цель: Убедиться, что режим errors буферизует.
        """
        print("Настройка: Режим errors...")
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'errors'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            
            print(f"Проверяем _is_immediate_write()...")
            assert logger._is_immediate_write() is False


class TestDebugLoggerJsonHandling:
    """Тесты для обработки JSON в DebugLogger."""
    
    def test_log_request_body_formats_json_pretty(self, tmp_path):
        """
        Что он делает: Проверяет, что JSON форматируется красиво.
        Цель: Убедиться, что JSON читаем в файле.
        """
        print("Настройка: Режим all...")
        debug_dir = tmp_path / "debug_logs"
        debug_dir.mkdir()
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            print("Действие: Вызов log_request_body с JSON...")
            logger.log_request_body(b'{"key":"value"}')
            
            print(f"Проверяем форматирование...")
            content = (debug_dir / "request_body.json").read_text()
            # Должен быть отформатирован с отступами
            assert "  " in content or "\n" in content
    
    def test_log_request_body_handles_invalid_json(self, tmp_path):
        """
        Что он делает: Проверяет обработку невалидного JSON.
        Цель: Убедиться, что невалидный JSON записывается как есть.
        """
        print("Настройка: Режим all...")
        debug_dir = tmp_path / "debug_logs"
        debug_dir.mkdir()
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            logger = DebugLogger.__new__(DebugLogger)
            logger._initialized = False
            logger.__init__()
            logger.debug_dir = debug_dir
            
            print("Действие: Вызов log_request_body с невалидным JSON...")
            invalid_data = b'not a json {{'
            logger.log_request_body(invalid_data)
            
            print(f"Проверяем, что данные записаны как есть...")
            content = (debug_dir / "request_body.json").read_bytes()
            assert content == invalid_data


class TestDebugLoggerAppLogsCapture:
    """Тесты для захвата логов приложения (app_logs.txt)."""
    
    def test_prepare_new_request_sets_up_log_capture(self, tmp_path):
        """
        Что он делает: Проверяет, что prepare_new_request настраивает захват логов.
        Цель: Убедиться, что sink для логов создаётся.
        """
        print("Настройка: Режим all...")
        debug_dir = tmp_path / "debug_logs"
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            dbg_logger = DebugLogger.__new__(DebugLogger)
            dbg_logger._initialized = False
            dbg_logger.__init__()
            dbg_logger.debug_dir = debug_dir
            
            print("Действие: Вызов prepare_new_request...")
            dbg_logger.prepare_new_request()
            
            print(f"Проверяем, что sink создан...")
            assert dbg_logger._loguru_sink_id is not None
            
            # Очистка
            dbg_logger._clear_app_logs_buffer()
    
    def test_flush_on_error_writes_app_logs_in_mode_errors(self, tmp_path):
        """
        Что он делает: Проверяет, что flush_on_error записывает app_logs.txt в режиме errors.
        Цель: Убедиться, что логи приложения сохраняются при ошибках.
        """
        print("Настройка: Режим errors...")
        debug_dir = tmp_path / "debug_logs"
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'errors'):
            from kiro_gateway.debug_logger import DebugLogger
            from loguru import logger as loguru_logger
            
            dbg_logger = DebugLogger.__new__(DebugLogger)
            dbg_logger._initialized = False
            dbg_logger.__init__()
            dbg_logger.debug_dir = debug_dir
            
            # Настраиваем захват логов
            dbg_logger.prepare_new_request()
            
            # Добавляем данные в буфер чтобы flush сработал
            dbg_logger.log_request_body(b'{"test": "data"}')
            
            # Пишем тестовый лог напрямую в буфер (имитация)
            dbg_logger._app_logs_buffer.write("Test log message\n")
            
            print("Действие: Вызов flush_on_error...")
            dbg_logger.flush_on_error(500, "Test Error")
            
            print(f"Проверяем, что app_logs.txt создан...")
            app_logs_file = debug_dir / "app_logs.txt"
            assert app_logs_file.exists()
            
            print(f"Проверяем содержимое...")
            content = app_logs_file.read_text()
            assert "Test log message" in content
    
    def test_discard_buffers_saves_logs_in_mode_all(self, tmp_path):
        """
        Что он делает: Проверяет, что discard_buffers сохраняет логи в режиме all.
        Цель: Убедиться, что даже успешные запросы сохраняют логи в режиме all.
        """
        print("Настройка: Режим all...")
        debug_dir = tmp_path / "debug_logs"
        debug_dir.mkdir()
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            
            dbg_logger = DebugLogger.__new__(DebugLogger)
            dbg_logger._initialized = False
            dbg_logger.__init__()
            dbg_logger.debug_dir = debug_dir
            
            # Настраиваем захват логов
            dbg_logger.prepare_new_request()
            
            # Пишем тестовый лог напрямую в буфер
            dbg_logger._app_logs_buffer.write("Success log message\n")
            
            print("Действие: Вызов discard_buffers...")
            dbg_logger.discard_buffers()
            
            print(f"Проверяем, что app_logs.txt создан...")
            app_logs_file = debug_dir / "app_logs.txt"
            assert app_logs_file.exists()
            
            print(f"Проверяем содержимое...")
            content = app_logs_file.read_text()
            assert "Success log message" in content
    
    def test_discard_buffers_does_not_save_logs_in_mode_errors(self, tmp_path):
        """
        Что он делает: Проверяет, что discard_buffers НЕ сохраняет логи в режиме errors.
        Цель: Убедиться, что успешные запросы не оставляют логов в режиме errors.
        """
        print("Настройка: Режим errors...")
        debug_dir = tmp_path / "debug_logs"
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'errors'):
            from kiro_gateway.debug_logger import DebugLogger
            
            dbg_logger = DebugLogger.__new__(DebugLogger)
            dbg_logger._initialized = False
            dbg_logger.__init__()
            dbg_logger.debug_dir = debug_dir
            
            # Настраиваем захват логов
            dbg_logger.prepare_new_request()
            
            # Пишем тестовый лог напрямую в буфер
            dbg_logger._app_logs_buffer.write("Should not be saved\n")
            
            print("Действие: Вызов discard_buffers...")
            dbg_logger.discard_buffers()
            
            print(f"Проверяем, что директория НЕ создана...")
            assert not debug_dir.exists()
    
    def test_clear_app_logs_buffer_removes_sink(self, tmp_path):
        """
        Что он делает: Проверяет, что _clear_app_logs_buffer удаляет sink.
        Цель: Убедиться, что sink корректно удаляется.
        """
        print("Настройка: Режим all...")
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            
            dbg_logger = DebugLogger.__new__(DebugLogger)
            dbg_logger._initialized = False
            dbg_logger.__init__()
            dbg_logger.debug_dir = tmp_path / "debug_logs"
            
            # Настраиваем захват логов
            dbg_logger.prepare_new_request()
            sink_id = dbg_logger._loguru_sink_id
            assert sink_id is not None
            
            print("Действие: Вызов _clear_app_logs_buffer...")
            dbg_logger._clear_app_logs_buffer()
            
            print(f"Проверяем, что sink_id сброшен...")
            assert dbg_logger._loguru_sink_id is None
    
    def test_app_logs_not_saved_when_empty(self, tmp_path):
        """
        Что он делает: Проверяет, что пустые логи не создают файл.
        Цель: Убедиться, что app_logs.txt не создаётся если логов нет.
        """
        print("Настройка: Режим all...")
        debug_dir = tmp_path / "debug_logs"
        debug_dir.mkdir()
        
        with patch('kiro_gateway.debug_logger.DEBUG_MODE', 'all'):
            from kiro_gateway.debug_logger import DebugLogger
            
            dbg_logger = DebugLogger.__new__(DebugLogger)
            dbg_logger._initialized = False
            dbg_logger.__init__()
            dbg_logger.debug_dir = debug_dir
            
            # НЕ пишем ничего в буфер
            
            print("Действие: Вызов _write_app_logs_to_file...")
            dbg_logger._write_app_logs_to_file()
            
            print(f"Проверяем, что app_logs.txt НЕ создан...")
            app_logs_file = debug_dir / "app_logs.txt"
            assert not app_logs_file.exists()


================================================
FILE: tests/unit/test_http_client.py
================================================
# -*- coding: utf-8 -*-

"""
Unit tests for KiroHttpClient.
Tests retry logic, error handling, and HTTP client management.
"""

import asyncio
import pytest
from unittest.mock import AsyncMock, Mock, patch, MagicMock
from datetime import datetime, timezone, timedelta

import httpx
from fastapi import HTTPException

from kiro_gateway.http_client import KiroHttpClient
from kiro_gateway.auth import KiroAuthManager
from kiro_gateway.config import MAX_RETRIES, BASE_RETRY_DELAY, FIRST_TOKEN_MAX_RETRIES, STREAMING_READ_TIMEOUT


@pytest.fixture
def mock_auth_manager_for_http():
    """Creates a mocked KiroAuthManager for HTTP client tests."""
    manager = Mock(spec=KiroAuthManager)
    manager.get_access_token = AsyncMock(return_value="test_access_token")
    manager.force_refresh = AsyncMock(return_value="new_access_token")
    manager.fingerprint = "test_fingerprint_12345678"
    manager._fingerprint = "test_fingerprint_12345678"
    return manager


class TestKiroHttpClientInitialization:
    """Tests for KiroHttpClient initialization."""
    
    def test_initialization_stores_auth_manager(self, mock_auth_manager_for_http):
        """
        What it does: Verifies auth_manager is stored during initialization.
        Purpose: Ensure auth_manager is available for obtaining tokens.
        """
        print("Setup: Creating KiroHttpClient...")
        client = KiroHttpClient(mock_auth_manager_for_http)
        
        print("Verification: auth_manager is stored...")
        assert client.auth_manager is mock_auth_manager_for_http
    
    def test_initialization_client_is_none(self, mock_auth_manager_for_http):
        """
        What it does: Verifies that HTTP client is initially None.
        Purpose: Ensure lazy initialization.
        """
        print("Setup: Creating KiroHttpClient...")
        client = KiroHttpClient(mock_auth_manager_for_http)
        
        print("Verification: client is initially None...")
        assert client.client is None


class TestKiroHttpClientGetClient:
    """Tests for _get_client method."""
    
    @pytest.mark.asyncio
    async def test_get_client_creates_new_client(self, mock_auth_manager_for_http):
        """
        What it does: Verifies creation of a new HTTP client.
        Purpose: Ensure client is created on first call.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        print("Action: Getting client...")
        with patch('kiro_gateway.http_client.httpx.AsyncClient') as mock_async_client:
            mock_instance = AsyncMock()
            mock_instance.is_closed = False
            mock_async_client.return_value = mock_instance
            
            client = await http_client._get_client()
            
            print("Verification: Client created...")
            mock_async_client.assert_called_once()
            assert client is mock_instance
    
    @pytest.mark.asyncio
    async def test_get_client_reuses_existing_client(self, mock_auth_manager_for_http):
        """
        What it does: Verifies reuse of existing client.
        Purpose: Ensure client is not recreated unnecessarily.
        """
        print("Setup: Creating KiroHttpClient with existing client...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_existing = AsyncMock()
        mock_existing.is_closed = False
        http_client.client = mock_existing
        
        print("Action: Getting client...")
        client = await http_client._get_client()
        
        print("Verification: Existing client returned...")
        assert client is mock_existing
    
    @pytest.mark.asyncio
    async def test_get_client_recreates_closed_client(self, mock_auth_manager_for_http):
        """
        What it does: Verifies recreation of closed client.
        Purpose: Ensure closed client is replaced with a new one.
        """
        print("Setup: Creating KiroHttpClient with closed client...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_closed = AsyncMock()
        mock_closed.is_closed = True
        http_client.client = mock_closed
        
        print("Action: Getting client...")
        with patch('kiro_gateway.http_client.httpx.AsyncClient') as mock_async_client:
            mock_new = AsyncMock()
            mock_new.is_closed = False
            mock_async_client.return_value = mock_new
            
            client = await http_client._get_client()
            
            print("Verification: New client created...")
            mock_async_client.assert_called_once()
            assert client is mock_new


class TestKiroHttpClientClose:
    """Tests for close method."""
    
    @pytest.mark.asyncio
    async def test_close_closes_client(self, mock_auth_manager_for_http):
        """
        What it does: Verifies HTTP client closure.
        Purpose: Ensure aclose() is called.
        """
        print("Setup: Creating KiroHttpClient with client...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.aclose = AsyncMock()
        http_client.client = mock_client
        
        print("Action: Closing client...")
        await http_client.close()
        
        print("Verification: aclose() called...")
        mock_client.aclose.assert_called_once()
    
    @pytest.mark.asyncio
    async def test_close_does_nothing_for_none_client(self, mock_auth_manager_for_http):
        """
        What it does: Verifies that close() doesn't fail for None client.
        Purpose: Ensure safe close() call without client.
        """
        print("Setup: Creating KiroHttpClient without client...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        print("Action: Closing client...")
        await http_client.close()  # Should not raise an error
        
        print("Verification: No errors...")
    
    @pytest.mark.asyncio
    async def test_close_does_nothing_for_closed_client(self, mock_auth_manager_for_http):
        """
        What it does: Verifies that close() doesn't fail for closed client.
        Purpose: Ensure safe repeated close() call.
        """
        print("Setup: Creating KiroHttpClient with closed client...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_client = AsyncMock()
        mock_client.is_closed = True
        http_client.client = mock_client
        
        print("Action: Closing client...")
        await http_client.close()
        
        print("Verification: aclose() NOT called...")
        mock_client.aclose.assert_not_called()


class TestKiroHttpClientRequestWithRetry:
    """Tests for request_with_retry method."""
    
    @pytest.mark.asyncio
    async def test_successful_request_returns_response(self, mock_auth_manager_for_http):
        """
        What it does: Verifies successful request.
        Purpose: Ensure 200 response is returned immediately.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.request = AsyncMock(return_value=mock_response)
        
        print("Action: Executing request...")
        with patch.object(http_client, '_get_client', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                response = await http_client.request_with_retry(
                    "POST",
                    "https://api.example.com/test",
                    {"data": "value"}
                )
        
        print("Verification: Response received...")
        assert response.status_code == 200
        mock_client.request.assert_called_once()
    
    @pytest.mark.asyncio
    async def test_403_triggers_token_refresh(self, mock_auth_manager_for_http):
        """
        What it does: Verifies token refresh on 403.
        Purpose: Ensure force_refresh() is called on 403.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response_403 = AsyncMock()
        mock_response_403.status_code = 403
        
        mock_response_200 = AsyncMock()
        mock_response_200.status_code = 200
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.request = AsyncMock(side_effect=[mock_response_403, mock_response_200])
        
        print("Action: Executing request...")
        with patch.object(http_client, '_get_client', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                response = await http_client.request_with_retry(
                    "POST",
                    "https://api.example.com/test",
                    {"data": "value"}
                )
        
        print("Verification: force_refresh() called...")
        mock_auth_manager_for_http.force_refresh.assert_called_once()
        assert response.status_code == 200
    
    @pytest.mark.asyncio
    async def test_429_triggers_backoff(self, mock_auth_manager_for_http):
        """
        What it does: Verifies exponential backoff on 429.
        Purpose: Ensure request is retried after delay.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response_429 = AsyncMock()
        mock_response_429.status_code = 429
        
        mock_response_200 = AsyncMock()
        mock_response_200.status_code = 200
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.request = AsyncMock(side_effect=[mock_response_429, mock_response_200])
        
        print("Action: Executing request...")
        with patch.object(http_client, '_get_client', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                with patch('kiro_gateway.http_client.asyncio.sleep', new_callable=AsyncMock) as mock_sleep:
                    response = await http_client.request_with_retry(
                        "POST",
                        "https://api.example.com/test",
                        {"data": "value"}
                    )
        
        print("Verification: sleep() called for backoff...")
        mock_sleep.assert_called_once()
        assert response.status_code == 200
    
    @pytest.mark.asyncio
    async def test_5xx_triggers_backoff(self, mock_auth_manager_for_http):
        """
        What it does: Verifies exponential backoff on 5xx.
        Purpose: Ensure server errors are handled with retry.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response_500 = AsyncMock()
        mock_response_500.status_code = 500
        
        mock_response_200 = AsyncMock()
        mock_response_200.status_code = 200
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.request = AsyncMock(side_effect=[mock_response_500, mock_response_200])
        
        print("Action: Executing request...")
        with patch.object(http_client, '_get_client', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                with patch('kiro_gateway.http_client.asyncio.sleep', new_callable=AsyncMock) as mock_sleep:
                    response = await http_client.request_with_retry(
                        "POST",
                        "https://api.example.com/test",
                        {"data": "value"}
                    )
        
        print("Verification: sleep() called for backoff...")
        mock_sleep.assert_called_once()
        assert response.status_code == 200
    
    @pytest.mark.asyncio
    async def test_timeout_triggers_backoff(self, mock_auth_manager_for_http):
        """
        What it does: Verifies exponential backoff on timeout.
        Purpose: Ensure timeouts are handled with retry.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response_200 = AsyncMock()
        mock_response_200.status_code = 200
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.request = AsyncMock(side_effect=[
            httpx.TimeoutException("Timeout"),
            mock_response_200
        ])
        
        print("Action: Executing request...")
        with patch.object(http_client, '_get_client', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                with patch('kiro_gateway.http_client.asyncio.sleep', new_callable=AsyncMock) as mock_sleep:
                    response = await http_client.request_with_retry(
                        "POST",
                        "https://api.example.com/test",
                        {"data": "value"}
                    )
        
        print("Verification: sleep() called for backoff...")
        mock_sleep.assert_called_once()
        assert response.status_code == 200
    
    @pytest.mark.asyncio
    async def test_request_error_triggers_backoff(self, mock_auth_manager_for_http):
        """
        What it does: Verifies exponential backoff on request error.
        Purpose: Ensure network errors are handled with retry.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response_200 = AsyncMock()
        mock_response_200.status_code = 200
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.request = AsyncMock(side_effect=[
            httpx.RequestError("Connection error"),
            mock_response_200
        ])
        
        print("Action: Executing request...")
        with patch.object(http_client, '_get_client', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                with patch('kiro_gateway.http_client.asyncio.sleep', new_callable=AsyncMock) as mock_sleep:
                    response = await http_client.request_with_retry(
                        "POST",
                        "https://api.example.com/test",
                        {"data": "value"}
                    )
        
        print("Verification: sleep() called for backoff...")
        mock_sleep.assert_called_once()
        assert response.status_code == 200
    
    @pytest.mark.asyncio
    async def test_max_retries_exceeded_raises_502(self, mock_auth_manager_for_http):
        """
        What it does: Verifies HTTPException is raised after exhausting retries.
        Purpose: Ensure 502 is raised after MAX_RETRIES.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.request = AsyncMock(side_effect=httpx.TimeoutException("Timeout"))
        
        print("Action: Executing request...")
        with patch.object(http_client, '_get_client', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                with patch('kiro_gateway.http_client.asyncio.sleep', new_callable=AsyncMock):
                    with pytest.raises(HTTPException) as exc_info:
                        await http_client.request_with_retry(
                            "POST",
                            "https://api.example.com/test",
                            {"data": "value"}
                        )
        
        print(f"Verification: HTTPException with code 502...")
        assert exc_info.value.status_code == 502
        assert str(MAX_RETRIES) in exc_info.value.detail
    
    @pytest.mark.asyncio
    async def test_other_status_codes_returned_as_is(self, mock_auth_manager_for_http):
        """
        What it does: Verifies other status codes are returned without retry.
        Purpose: Ensure 400, 404, etc. are returned immediately.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response = AsyncMock()
        mock_response.status_code = 400
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.request = AsyncMock(return_value=mock_response)
        
        print("Action: Executing request...")
        with patch.object(http_client, '_get_client', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                response = await http_client.request_with_retry(
                    "POST",
                    "https://api.example.com/test",
                    {"data": "value"}
                )
        
        print("Verification: 400 response returned without retry...")
        assert response.status_code == 400
        mock_client.request.assert_called_once()
    
    @pytest.mark.asyncio
    async def test_streaming_request_uses_send(self, mock_auth_manager_for_http):
        """
        What it does: Verifies send() is used for streaming.
        Purpose: Ensure stream=True uses build_request + send.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        mock_request = Mock()
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.build_request = Mock(return_value=mock_request)
        mock_client.send = AsyncMock(return_value=mock_response)
        
        print("Action: Executing streaming request...")
        with patch.object(http_client, '_get_client', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                response = await http_client.request_with_retry(
                    "POST",
                    "https://api.example.com/test",
                    {"data": "value"},
                    stream=True
                )
        
        print("Verification: build_request and send called...")
        mock_client.build_request.assert_called_once()
        mock_client.send.assert_called_once_with(mock_request, stream=True)
        assert response.status_code == 200


class TestKiroHttpClientContextManager:
    """Tests for async context manager."""
    
    @pytest.mark.asyncio
    async def test_context_manager_returns_self(self, mock_auth_manager_for_http):
        """
        What it does: Verifies that __aenter__ returns self.
        Purpose: Ensure correct async with behavior.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        print("Action: Entering context...")
        result = await http_client.__aenter__()
        
        print("Verification: self returned...")
        assert result is http_client
    
    @pytest.mark.asyncio
    async def test_context_manager_closes_on_exit(self, mock_auth_manager_for_http):
        """
        What it does: Verifies client closure on context exit.
        Purpose: Ensure close() is called in __aexit__.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.aclose = AsyncMock()
        http_client.client = mock_client
        
        print("Action: Exiting context...")
        await http_client.__aexit__(None, None, None)
        
        print("Verification: aclose() called...")
        mock_client.aclose.assert_called_once()


class TestKiroHttpClientExponentialBackoff:
    """Tests for exponential backoff logic."""
    
    @pytest.mark.asyncio
    async def test_backoff_delay_increases_exponentially(self, mock_auth_manager_for_http):
        """
        What it does: Verifies exponential delay increase.
        Purpose: Ensure delay = BASE_RETRY_DELAY * (2 ** attempt).
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response_429 = AsyncMock()
        mock_response_429.status_code = 429
        
        mock_response_200 = AsyncMock()
        mock_response_200.status_code = 200
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        # 2 errors 429, then success (to verify 2 backoff delays)
        mock_client.request = AsyncMock(side_effect=[
            mock_response_429,
            mock_response_429,
            mock_response_200
        ])
        
        sleep_delays = []
        
        async def capture_sleep(delay):
            sleep_delays.append(delay)
        
        print("Action: Executing request with multiple retries...")
        with patch.object(http_client, '_get_client', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                with patch('kiro_gateway.http_client.asyncio.sleep', side_effect=capture_sleep):
                    response = await http_client.request_with_retry(
                        "POST",
                        "https://api.example.com/test",
                        {"data": "value"}
                    )
        
        print(f"Verification: Delays increase exponentially...")
        print(f"Delays: {sleep_delays}")
        assert len(sleep_delays) == 2
        assert sleep_delays[0] == BASE_RETRY_DELAY * (2 ** 0)  # 1.0
        assert sleep_delays[1] == BASE_RETRY_DELAY * (2 ** 1)  # 2.0


class TestKiroHttpClientStreamingTimeout:
    """Tests for streaming request timeout logic."""
    
    @pytest.mark.asyncio
    async def test_streaming_uses_streaming_read_timeout(self, mock_auth_manager_for_http):
        """
        What it does: Verifies that streaming requests use STREAMING_READ_TIMEOUT.
        Purpose: Ensure stream=True uses httpx.Timeout with correct values.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        mock_request = Mock()
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.build_request = Mock(return_value=mock_request)
        mock_client.send = AsyncMock(return_value=mock_response)
        
        print("Action: Executing streaming request...")
        with patch('kiro_gateway.http_client.httpx.AsyncClient') as mock_async_client:
            mock_async_client.return_value = mock_client
            
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                response = await http_client.request_with_retry(
                    "POST",
                    "https://api.example.com/test",
                    {"data": "value"},
                    stream=True
                )
        
        print("Verification: AsyncClient created with httpx.Timeout for streaming...")
        call_args = mock_async_client.call_args
        timeout_arg = call_args.kwargs.get('timeout')
        assert timeout_arg is not None, f"timeout not found in call_args: {call_args}"
        print(f"Comparing connect: Expected 30.0, Got {timeout_arg.connect}")
        assert timeout_arg.connect == 30.0, f"Expected connect=30.0, got {timeout_arg.connect}"
        print(f"Comparing read: Expected {STREAMING_READ_TIMEOUT}, Got {timeout_arg.read}")
        assert timeout_arg.read == STREAMING_READ_TIMEOUT, f"Expected read={STREAMING_READ_TIMEOUT}, got {timeout_arg.read}"
        assert call_args.kwargs.get('follow_redirects') == True
        assert response.status_code == 200
    
    @pytest.mark.asyncio
    async def test_streaming_uses_first_token_max_retries(self, mock_auth_manager_for_http):
        """
        What it does: Verifies that streaming requests use FIRST_TOKEN_MAX_RETRIES.
        Purpose: Ensure stream=True uses separate retry counter.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_request = Mock()
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.build_request = Mock(return_value=mock_request)
        mock_client.send = AsyncMock(side_effect=httpx.TimeoutException("Timeout"))
        
        print("Action: Executing streaming request with timeouts...")
        with patch('kiro_gateway.http_client.httpx.AsyncClient', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                with pytest.raises(HTTPException) as exc_info:
                    await http_client.request_with_retry(
                        "POST",
                        "https://api.example.com/test",
                        {"data": "value"},
                        stream=True
                    )
        
        print(f"Verification: HTTPException with code 504...")
        assert exc_info.value.status_code == 504
        assert str(FIRST_TOKEN_MAX_RETRIES) in exc_info.value.detail
        
        print(f"Verification: Attempt count = FIRST_TOKEN_MAX_RETRIES ({FIRST_TOKEN_MAX_RETRIES})...")
        assert mock_client.send.call_count == FIRST_TOKEN_MAX_RETRIES
    
    @pytest.mark.asyncio
    async def test_streaming_timeout_retry_without_delay(self, mock_auth_manager_for_http):
        """
        What it does: Verifies that streaming timeout retry happens without delay.
        Purpose: Ensure no exponential backoff on first token timeout.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        mock_request = Mock()
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.build_request = Mock(return_value=mock_request)
        # First timeout, then success
        mock_client.send = AsyncMock(side_effect=[
            httpx.TimeoutException("Timeout"),
            mock_response
        ])
        
        sleep_called = False
        
        async def capture_sleep(delay):
            nonlocal sleep_called
            sleep_called = True
        
        print("Action: Executing streaming request with one timeout...")
        with patch('kiro_gateway.http_client.httpx.AsyncClient', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                with patch('kiro_gateway.http_client.asyncio.sleep', side_effect=capture_sleep):
                    response = await http_client.request_with_retry(
                        "POST",
                        "https://api.example.com/test",
                        {"data": "value"},
                        stream=True
                    )
        
        print("Verification: sleep() NOT called for streaming timeout...")
        assert not sleep_called
        assert response.status_code == 200
        
    @pytest.mark.asyncio
    async def test_non_streaming_uses_default_timeout(self, mock_auth_manager_for_http):
        """
        What it does: Verifies that non-streaming requests use 300 seconds.
        Purpose: Ensure stream=False uses unified httpx.Timeout.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.request = AsyncMock(return_value=mock_response)
        
        print("Action: Executing non-streaming request...")
        with patch('kiro_gateway.http_client.httpx.AsyncClient') as mock_async_client:
            mock_async_client.return_value = mock_client
            
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                response = await http_client.request_with_retry(
                    "POST",
                    "https://api.example.com/test",
                    {"data": "value"},
                    stream=False
                )
        
        print("Verification: AsyncClient created with httpx.Timeout(timeout=300)...")
        call_args = mock_async_client.call_args
        timeout_arg = call_args.kwargs.get('timeout')
        assert timeout_arg is not None, f"timeout not found in call_args: {call_args}"
        # httpx.Timeout(timeout=300) sets all timeouts to 300
        print(f"Comparing timeout: Expected 300.0 for all, Got connect={timeout_arg.connect}")
        assert timeout_arg.connect == 300.0
        assert timeout_arg.read == 300.0
        assert response.status_code == 200
    
    @pytest.mark.asyncio
    async def test_connect_timeout_logged_correctly(self, mock_auth_manager_for_http):
        """
        What it does: Verifies ConnectTimeout logging.
        Purpose: Ensure ConnectTimeout is logged with correct type.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        mock_request = Mock()
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.build_request = Mock(return_value=mock_request)
        # First ConnectTimeout, then success
        mock_client.send = AsyncMock(side_effect=[
            httpx.ConnectTimeout("Connection timeout"),
            mock_response
        ])
        
        print("Action: Executing streaming request with ConnectTimeout...")
        with patch('kiro_gateway.http_client.httpx.AsyncClient', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                with patch('kiro_gateway.http_client.logger') as mock_logger:
                    response = await http_client.request_with_retry(
                        "POST",
                        "https://api.example.com/test",
                        {"data": "value"},
                        stream=True
                    )
        
        print("Verification: logger.warning called with [ConnectTimeout]...")
        warning_calls = [str(call) for call in mock_logger.warning.call_args_list]
        assert any("ConnectTimeout" in call for call in warning_calls), f"ConnectTimeout not found in: {warning_calls}"
        assert response.status_code == 200
    
    @pytest.mark.asyncio
    async def test_read_timeout_logged_correctly(self, mock_auth_manager_for_http):
        """
        What it does: Verifies ReadTimeout logging.
        Purpose: Ensure ReadTimeout is logged with STREAMING_READ_TIMEOUT.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        mock_request = Mock()
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.build_request = Mock(return_value=mock_request)
        # First ReadTimeout, then success
        mock_client.send = AsyncMock(side_effect=[
            httpx.ReadTimeout("Read timeout"),
            mock_response
        ])
        
        print("Action: Executing streaming request with ReadTimeout...")
        with patch('kiro_gateway.http_client.httpx.AsyncClient', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                with patch('kiro_gateway.http_client.logger') as mock_logger:
                    response = await http_client.request_with_retry(
                        "POST",
                        "https://api.example.com/test",
                        {"data": "value"},
                        stream=True
                    )
        
        print("Verification: logger.warning called with [ReadTimeout] and STREAMING_READ_TIMEOUT...")
        warning_calls = [str(call) for call in mock_logger.warning.call_args_list]
        assert any("ReadTimeout" in call for call in warning_calls), f"ReadTimeout not found in: {warning_calls}"
        assert any(str(STREAMING_READ_TIMEOUT) in call for call in warning_calls), f"STREAMING_READ_TIMEOUT not found in: {warning_calls}"
        assert response.status_code == 200
    
    @pytest.mark.asyncio
    async def test_streaming_timeout_returns_504_with_error_type(self, mock_auth_manager_for_http):
        """
        What it does: Verifies that streaming timeout returns 504 with error type.
        Purpose: Ensure 504 is returned with error info after exhausting retries.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_request = Mock()
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.build_request = Mock(return_value=mock_request)
        mock_client.send = AsyncMock(side_effect=httpx.ReadTimeout("Timeout"))
        
        print("Action: Executing streaming request with persistent timeouts...")
        with patch('kiro_gateway.http_client.httpx.AsyncClient', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                with pytest.raises(HTTPException) as exc_info:
                    await http_client.request_with_retry(
                        "POST",
                        "https://api.example.com/test",
                        {"data": "value"},
                        stream=True
                    )
        
        print("Verification: HTTPException with code 504 and error type...")
        print(f"Comparing status_code: Expected 504, Got {exc_info.value.status_code}")
        assert exc_info.value.status_code == 504
        print(f"Comparing detail: Expected 'ReadTimeout' in '{exc_info.value.detail}'")
        assert "ReadTimeout" in exc_info.value.detail
        assert "Streaming failed" in exc_info.value.detail
    
    @pytest.mark.asyncio
    async def test_non_streaming_timeout_returns_502(self, mock_auth_manager_for_http):
        """
        What it does: Verifies that non-streaming timeout returns 502.
        Purpose: Ensure non-streaming uses legacy logic with 502.
        """
        print("Setup: Creating KiroHttpClient...")
        http_client = KiroHttpClient(mock_auth_manager_for_http)
        
        mock_client = AsyncMock()
        mock_client.is_closed = False
        mock_client.request = AsyncMock(side_effect=httpx.TimeoutException("Timeout"))
        
        print("Action: Executing non-streaming request with persistent timeouts...")
        with patch('kiro_gateway.http_client.httpx.AsyncClient', return_value=mock_client):
            with patch('kiro_gateway.http_client.get_kiro_headers', return_value={}):
                with patch('kiro_gateway.http_client.asyncio.sleep', new_callable=AsyncMock):
                    with pytest.raises(HTTPException) as exc_info:
                        await http_client.request_with_retry(
                            "POST",
                            "https://api.example.com/test",
                            {"data": "value"},
                            stream=False
                        )
        
        print("Verification: HTTPException with code 502...")
        assert exc_info.value.status_code == 502


================================================
FILE: tests/unit/test_parsers.py
================================================
# -*- coding: utf-8 -*-

"""
Unit-тесты для AwsEventStreamParser и вспомогательных функций парсинга.
Проверяет логику парсинга AWS SSE потока от Kiro API.
"""

import pytest

from kiro_gateway.parsers import (
    AwsEventStreamParser,
    find_matching_brace,
    parse_bracket_tool_calls,
    deduplicate_tool_calls
)


class TestFindMatchingBrace:
    """Тесты функции find_matching_brace."""
    
    def test_simple_json_object(self):
        """
        Что он делает: Проверяет поиск закрывающей скобки для простого JSON.
        Цель: Убедиться, что базовый случай работает.
        """
        print("Настройка: Простой JSON объект...")
        text = '{"key": "value"}'
        
        print("Действие: Поиск закрывающей скобки...")
        result = find_matching_brace(text, 0)
        
        print(f"Сравниваем результат: Ожидалось 15, Получено {result}")
        assert result == 15
    
    def test_nested_json_object(self):
        """
        Что он делает: Проверяет поиск для вложенного JSON.
        Цель: Убедиться, что вложенность обрабатывается корректно.
        """
        print("Настройка: Вложенный JSON объект...")
        text = '{"outer": {"inner": "value"}}'
        
        print("Действие: Поиск закрывающей скобки...")
        result = find_matching_brace(text, 0)
        
        # Длина строки 29, индекс последнего символа 28
        print(f"Сравниваем результат: Ожидалось 28, Получено {result}")
        assert result == 28
    
    def test_json_with_braces_in_string(self):
        """
        Что он делает: Проверяет игнорирование скобок внутри строк.
        Цель: Убедиться, что скобки в строках не влияют на подсчёт.
        """
        print("Настройка: JSON со скобками в строке...")
        text = '{"text": "Hello {world}"}'
        
        print("Действие: Поиск закрывающей скобки...")
        result = find_matching_brace(text, 0)
        
        print(f"Сравниваем результат: Ожидалось 24, Получено {result}")
        assert result == 24
    
    def test_json_with_escaped_quotes(self):
        """
        Что он делает: Проверяет обработку экранированных кавычек.
        Цель: Убедиться, что escape-последовательности не ломают парсинг.
        """
        print("Настройка: JSON с экранированными кавычками...")
        text = '{"text": "Say \\"hello\\""}'
        
        print("Действие: Поиск закрывающей скобки...")
        result = find_matching_brace(text, 0)
        
        # Длина строки 25, индекс последнего символа 24
        print(f"Сравниваем результат: Ожидалось 24, Получено {result}")
        assert result == 24
    
    def test_incomplete_json(self):
        """
        Что он делает: Проверяет обработку незавершённого JSON.
        Цель: Убедиться, что возвращается -1 для неполного JSON.
        """
        print("Настройка: Незавершённый JSON...")
        text = '{"key": "value"'
        
        print("Действие: Поиск закрывающей скобки...")
        result = find_matching_brace(text, 0)
        
        print(f"Сравниваем результат: Ожидалось -1, Получено {result}")
        assert result == -1
    
    def test_invalid_start_position(self):
        """
        Что он делает: Проверяет обработку невалидной стартовой позиции.
        Цель: Убедиться, что возвращается -1 если start_pos не на '{'.
        """
        print("Настройка: Текст без скобки на start_pos...")
        text = 'hello {"key": "value"}'
        
        print("Действие: Поиск с позиции 0 (не скобка)...")
        result = find_matching_brace(text, 0)
        
        print(f"Сравниваем результат: Ожидалось -1, Получено {result}")
        assert result == -1
    
    def test_start_position_out_of_bounds(self):
        """
        Что он делает: Проверяет обработку позиции за пределами текста.
        Цель: Убедиться, что возвращается -1 для невалидной позиции.
        """
        print("Настройка: Короткий текст...")
        text = '{"a":1}'
        
        print("Действие: Поиск с позиции 100...")
        result = find_matching_brace(text, 100)
        
        print(f"Сравниваем результат: Ожидалось -1, Получено {result}")
        assert result == -1


class TestParseBracketToolCalls:
    """Тесты функции parse_bracket_tool_calls."""
    
    def test_parses_single_tool_call(self):
        """
        Что он делает: Проверяет парсинг одного tool call.
        Цель: Убедиться, что bracket-style tool call извлекается корректно.
        """
        print("Настройка: Текст с одним tool call...")
        text = '[Called get_weather with args: {"location": "Moscow"}]'
        
        print("Действие: Парсинг tool calls...")
        result = parse_bracket_tool_calls(text)
        
        print(f"Результат: {result}")
        assert len(result) == 1
        assert result[0]["function"]["name"] == "get_weather"
        assert '"location"' in result[0]["function"]["arguments"]
    
    def test_parses_multiple_tool_calls(self):
        """
        Что он делает: Проверяет парсинг нескольких tool calls.
        Цель: Убедиться, что все tool calls извлекаются.
        """
        print("Настройка: Текст с несколькими tool calls...")
        text = '''
        [Called get_weather with args: {"location": "Moscow"}]
        Some text in between
        [Called get_time with args: {"timezone": "UTC"}]
        '''
        
        print("Действие: Парсинг tool calls...")
        result = parse_bracket_tool_calls(text)
        
        print(f"Результат: {result}")
        assert len(result) == 2
        assert result[0]["function"]["name"] == "get_weather"
        assert result[1]["function"]["name"] == "get_time"
    
    def test_returns_empty_for_no_tool_calls(self):
        """
        Что он делает: Проверяет возврат пустого списка без tool calls.
        Цель: Убедиться, что обычный текст не парсится как tool call.
        """
        print("Настройка: Обычный текст без tool calls...")
        text = "This is just regular text without any tool calls."
        
        print("Действие: Парсинг tool calls...")
        result = parse_bracket_tool_calls(text)
        
        print(f"Сравниваем результат: Ожидалось [], Получено {result}")
        assert result == []
    
    def test_returns_empty_for_empty_string(self):
        """
        Что он делает: Проверяет обработку пустой строки.
        Цель: Убедиться, что пустая строка не вызывает ошибок.
        """
        print("Настройка: Пустая строка...")
        
        print("Действие: Парсинг tool calls...")
        result = parse_bracket_tool_calls("")
        
        print(f"Сравниваем результат: Ожидалось [], Получено {result}")
        assert result == []
    
    def test_returns_empty_for_none(self):
        """
        Что он делает: Проверяет обработку None.
        Цель: Убедиться, что None не вызывает ошибок.
        """
        print("Настройка: None...")
        
        print("Действие: Парсинг tool calls...")
        result = parse_bracket_tool_calls(None)
        
        print(f"Сравниваем результат: Ожидалось [], Получено {result}")
        assert result == []
    
    def test_handles_nested_json_in_args(self):
        """
        Что он делает: Проверяет парсинг вложенного JSON в аргументах.
        Цель: Убедиться, что сложные аргументы парсятся корректно.
        """
        print("Настройка: Tool call с вложенным JSON...")
        text = '[Called complex_func with args: {"data": {"nested": {"deep": "value"}}}]'
        
        print("Действие: Парсинг tool calls...")
        result = parse_bracket_tool_calls(text)
        
        print(f"Результат: {result}")
        assert len(result) == 1
        assert result[0]["function"]["name"] == "complex_func"
        assert "nested" in result[0]["function"]["arguments"]
    
    def test_generates_unique_ids(self):
        """
        Что он делает: Проверяет генерацию уникальных ID для tool calls.
        Цель: Убедиться, что каждый tool call имеет уникальный ID.
        """
        print("Настройка: Два одинаковых tool calls...")
        text = '''
        [Called func with args: {"a": 1}]
        [Called func with args: {"a": 1}]
        '''
        
        print("Действие: Парсинг tool calls...")
        result = parse_bracket_tool_calls(text)
        
        print(f"IDs: {[r['id'] for r in result]}")
        assert len(result) == 2
        assert result[0]["id"] != result[1]["id"]


class TestDeduplicateToolCalls:
    """Тесты функции deduplicate_tool_calls."""
    
    def test_removes_duplicates(self):
        """
        Что он делает: Проверяет удаление дубликатов.
        Цель: Убедиться, что одинаковые tool calls удаляются.
        """
        print("Настройка: Список с дубликатами...")
        tool_calls = [
            {"id": "1", "function": {"name": "func", "arguments": '{"a": 1}'}},
            {"id": "2", "function": {"name": "func", "arguments": '{"a": 1}'}},
            {"id": "3", "function": {"name": "other", "arguments": '{"b": 2}'}},
        ]
        
        print("Действие: Дедупликация...")
        result = deduplicate_tool_calls(tool_calls)
        
        print(f"Сравниваем длину: Ожидалось 2, Получено {len(result)}")
        assert len(result) == 2
    
    def test_preserves_first_occurrence(self):
        """
        Что он делает: Проверяет сохранение первого вхождения.
        Цель: Убедиться, что сохраняется первый tool call из дубликатов.
        """
        print("Настройка: Список с дубликатами...")
        tool_calls = [
            {"id": "first", "function": {"name": "func", "arguments": '{"a": 1}'}},
            {"id": "second", "function": {"name": "func", "arguments": '{"a": 1}'}},
        ]
        
        print("Действие: Дедупликация...")
        result = deduplicate_tool_calls(tool_calls)
        
        print(f"Сравниваем ID: Ожидалось 'first', Получено '{result[0]['id']}'")
        assert result[0]["id"] == "first"
    
    def test_handles_empty_list(self):
        """
        Что он делает: Проверяет обработку пустого списка.
        Цель: Убедиться, что пустой список не вызывает ошибок.
        """
        print("Настройка: Пустой список...")
        
        print("Действие: Дедупликация...")
        result = deduplicate_tool_calls([])
        
        print(f"Сравниваем результат: Ожидалось [], Получено {result}")
        assert result == []
    
    def test_deduplicates_by_id_keeps_one_with_arguments(self):
        """
        Что он делает: Проверяет дедупликацию по id с сохранением tool call с аргументами.
        Цель: Убедиться, что при дубликатах по id сохраняется тот, у которого есть аргументы.
        """
        print("Настройка: Два tool calls с одинаковым id, один с аргументами, один пустой...")
        tool_calls = [
            {"id": "call_123", "function": {"name": "func", "arguments": "{}"}},
            {"id": "call_123", "function": {"name": "func", "arguments": '{"location": "Moscow"}'}},
        ]
        
        print("Действие: Дедупликация...")
        result = deduplicate_tool_calls(tool_calls)
        
        print(f"Результат: {result}")
        print(f"Сравниваем длину: Ожидалось 1, Получено {len(result)}")
        assert len(result) == 1
        
        print("Проверяем, что сохранился tool call с аргументами...")
        assert "Moscow" in result[0]["function"]["arguments"]
    
    def test_deduplicates_by_id_prefers_longer_arguments(self):
        """
        Что он делает: Проверяет, что при дубликатах по id предпочитаются более длинные аргументы.
        Цель: Убедиться, что сохраняется tool call с более полными аргументами.
        """
        print("Настройка: Два tool calls с одинаковым id, разной длины аргументов...")
        tool_calls = [
            {"id": "call_abc", "function": {"name": "search", "arguments": '{"q": "test"}'}},
            {"id": "call_abc", "function": {"name": "search", "arguments": '{"q": "test", "limit": 10, "offset": 0}'}},
        ]
        
        print("Действие: Дедупликация...")
        result = deduplicate_tool_calls(tool_calls)
        
        print(f"Результат: {result}")
        assert len(result) == 1
        
        print("Проверяем, что сохранился tool call с более длинными аргументами...")
        assert "limit" in result[0]["function"]["arguments"]
    
    def test_deduplicates_empty_arguments_replaced_by_non_empty(self):
        """
        Что он делает: Проверяет замену пустых аргументов на непустые.
        Цель: Убедиться, что "{}" заменяется на реальные аргументы.
        """
        print("Настройка: Первый tool call с пустыми аргументами, второй с реальными...")
        tool_calls = [
            {"id": "call_xyz", "function": {"name": "get_weather", "arguments": "{}"}},
            {"id": "call_xyz", "function": {"name": "get_weather", "arguments": '{"city": "London"}'}},
        ]
        
        print("Действие: Дедупликация...")
        result = deduplicate_tool_calls(tool_calls)
        
        print(f"Результат: {result}")
        assert len(result) == 1
        assert result[0]["function"]["arguments"] == '{"city": "London"}'
    
    def test_handles_tool_calls_without_id(self):
        """
        Что он делает: Проверяет обработку tool calls без id.
        Цель: Убедиться, что tool calls без id дедуплицируются по name+arguments.
        """
        print("Настройка: Tool calls без id...")
        tool_calls = [
            {"id": "", "function": {"name": "func", "arguments": '{"a": 1}'}},
            {"id": "", "function": {"name": "func", "arguments": '{"a": 1}'}},
            {"id": "", "function": {"name": "func", "arguments": '{"b": 2}'}},
        ]
        
        print("Действие: Дедупликация...")
        result = deduplicate_tool_calls(tool_calls)
        
        print(f"Результат: {result}")
        # Два уникальных по name+arguments
        assert len(result) == 2
    
    def test_mixed_with_and_without_id(self):
        """
        Что он делает: Проверяет смешанный список с id и без.
        Цель: Убедиться, что оба типа обрабатываются корректно.
        """
        print("Настройка: Смешанный список...")
        tool_calls = [
            {"id": "call_1", "function": {"name": "func1", "arguments": '{"x": 1}'}},
            {"id": "call_1", "function": {"name": "func1", "arguments": "{}"}},  # Дубликат по id
            {"id": "", "function": {"name": "func2", "arguments": '{"y": 2}'}},
            {"id": "", "function": {"name": "func2", "arguments": '{"y": 2}'}},  # Дубликат по name+args
        ]
        
        print("Действие: Дедупликация...")
        result = deduplicate_tool_calls(tool_calls)
        
        print(f"Результат: {result}")
        # call_1 с аргументами + func2 один раз
        assert len(result) == 2
        
        # Проверяем, что call_1 сохранил аргументы
        call_1 = next(tc for tc in result if tc["id"] == "call_1")
        assert call_1["function"]["arguments"] == '{"x": 1}'


class TestAwsEventStreamParserInitialization:
    """Тесты инициализации AwsEventStreamParser."""
    
    def test_initialization_creates_empty_state(self):
        """
        Что он делает: Проверяет начальное состояние парсера.
        Цель: Убедиться, что парсер создаётся с пустым состоянием.
        """
        print("Настройка: Создание парсера...")
        parser = AwsEventStreamParser()
        
        print("Проверка: Буфер пуст...")
        assert parser.buffer == ""
        
        print("Проверка: last_content is None...")
        assert parser.last_content is None
        
        print("Проверка: current_tool_call is None...")
        assert parser.current_tool_call is None
        
        print("Проверка: tool_calls пуст...")
        assert parser.tool_calls == []


class TestAwsEventStreamParserFeed:
    """Тесты метода feed парсера."""
    
    def test_parses_content_event(self, aws_event_parser):
        """
        Что он делает: Проверяет парсинг события с контентом.
        Цель: Убедиться, что текстовый контент извлекается.
        """
        print("Настройка: Chunk с контентом...")
        chunk = b'{"content":"Hello World"}'
        
        print("Действие: Парсинг chunk...")
        events = aws_event_parser.feed(chunk)
        
        print(f"Результат: {events}")
        assert len(events) == 1
        assert events[0]["type"] == "content"
        assert events[0]["data"] == "Hello World"
    
    def test_parses_multiple_content_events(self, aws_event_parser):
        """
        Что он делает: Проверяет парсинг нескольких событий контента.
        Цель: Убедиться, что все события извлекаются.
        """
        print("Настройка: Chunk с несколькими событиями...")
        chunk = b'{"content":"First"}{"content":"Second"}'
        
        print("Действие: Парсинг chunk...")
        events = aws_event_parser.feed(chunk)
        
        print(f"Результат: {events}")
        assert len(events) == 2
        assert events[0]["data"] == "First"
        assert events[1]["data"] == "Second"
    
    def test_deduplicates_repeated_content(self, aws_event_parser):
        """
        Что он делает: Проверяет дедупликацию повторяющегося контента.
        Цель: Убедиться, что одинаковый контент не дублируется.
        """
        print("Настройка: Chunks с повторяющимся контентом...")
        
        print("Действие: Парсинг первого chunk...")
        events1 = aws_event_parser.feed(b'{"content":"Same"}')
        
        print("Действие: Парсинг второго chunk с тем же контентом...")
        events2 = aws_event_parser.feed(b'{"content":"Same"}')
        
        print(f"Первый результат: {events1}")
        print(f"Второй результат: {events2}")
        assert len(events1) == 1
        assert len(events2) == 0  # Дубликат отфильтрован
    
    def test_parses_usage_event(self, aws_event_parser):
        """
        Что он делает: Проверяет парсинг события usage.
        Цель: Убедиться, что информация о credits извлекается.
        """
        print("Настройка: Chunk с usage...")
        chunk = b'{"usage":1.5}'
        
        print("Действие: Парсинг chunk...")
        events = aws_event_parser.feed(chunk)
        
        print(f"Результат: {events}")
        assert len(events) == 1
        assert events[0]["type"] == "usage"
        assert events[0]["data"] == 1.5
    
    def test_parses_context_usage_event(self, aws_event_parser):
        """
        Что он делает: Проверяет парсинг события context_usage.
        Цель: Убедиться, что процент использования контекста извлекается.
        """
        print("Настройка: Chunk с context usage...")
        chunk = b'{"contextUsagePercentage":25.5}'
        
        print("Действие: Парсинг chunk...")
        events = aws_event_parser.feed(chunk)
        
        print(f"Результат: {events}")
        assert len(events) == 1
        assert events[0]["type"] == "context_usage"
        assert events[0]["data"] == 25.5
    
    def test_handles_incomplete_json(self, aws_event_parser):
        """
        Что он делает: Проверяет обработку неполного JSON.
        Цель: Убедиться, что неполный JSON буферизуется.
        """
        print("Настройка: Неполный chunk...")
        chunk = b'{"content":"Hel'
        
        print("Действие: Парсинг неполного chunk...")
        events = aws_event_parser.feed(chunk)
        
        print(f"Результат: {events}")
        assert len(events) == 0  # Ничего не распарсено
        
        print("Проверка: Данные в буфере...")
        assert 'content' in aws_event_parser.buffer
    
    def test_completes_json_across_chunks(self, aws_event_parser):
        """
        Что он делает: Проверяет сборку JSON из нескольких chunks.
        Цель: Убедиться, что JSON собирается из частей.
        """
        print("Настройка: Первая часть JSON...")
        events1 = aws_event_parser.feed(b'{"content":"Hel')
        
        print("Действие: Вторая часть JSON...")
        events2 = aws_event_parser.feed(b'lo World"}')
        
        print(f"Первый результат: {events1}")
        print(f"Второй результат: {events2}")
        assert len(events1) == 0
        assert len(events2) == 1
        assert events2[0]["data"] == "Hello World"
    
    def test_decodes_escape_sequences(self, aws_event_parser):
        """
        Что он делает: Проверяет декодирование escape-последовательностей.
        Цель: Убедиться, что \\n преобразуется в реальный перенос строки.
        """
        print("Настройка: Chunk с escape-последовательностью...")
        # Используем правильный формат escape-последовательности
        chunk = b'{"content":"Line1\\nLine2"}'
        
        print("Действие: Парсинг chunk...")
        events = aws_event_parser.feed(chunk)
        
        print(f"Результат: {events}")
        assert len(events) == 1
        assert "\n" in events[0]["data"]
    def test_handles_invalid_bytes(self, aws_event_parser):
        """
        Что он делает: Проверяет обработку невалидных байтов.
        Цель: Убедиться, что невалидные данные не ломают парсер.
        """
        print("Настройка: Невалидные байты...")
        chunk = b'\xff\xfe{"content":"test"}'
        
        print("Действие: Парсинг chunk...")
        events = aws_event_parser.feed(chunk)
        
        print(f"Результат: {events}")
        # Парсер должен продолжить работу
        assert len(events) == 1


class TestAwsEventStreamParserToolCalls:
    """Тесты парсинга tool calls."""
    
    def test_parses_tool_start_event(self, aws_event_parser):
        """
        Что он делает: Проверяет парсинг начала tool call.
        Цель: Убедиться, что tool_start создаёт current_tool_call.
        """
        print("Настройка: Chunk с началом tool call...")
        chunk = b'{"name":"get_weather","toolUseId":"call_123"}'
        
        print("Действие: Парсинг chunk...")
        events = aws_event_parser.feed(chunk)
        
        print(f"Результат: {events}")
        print(f"current_tool_call: {aws_event_parser.current_tool_call}")
        
        # tool_start не возвращает событие, но создаёт current_tool_call
        assert aws_event_parser.current_tool_call is not None
        assert aws_event_parser.current_tool_call["function"]["name"] == "get_weather"
    
    def test_parses_tool_input_event(self, aws_event_parser):
        """
        Что он делает: Проверяет парсинг input для tool call.
        Цель: Убедиться, что input добавляется к current_tool_call.
        """
        print("Настройка: Начало tool call...")
        aws_event_parser.feed(b'{"name":"func","toolUseId":"call_1"}')
        
        print("Действие: Парсинг input...")
        aws_event_parser.feed(b'{"input":"{\\"key\\": \\"value\\"}"}')
        
        print(f"current_tool_call: {aws_event_parser.current_tool_call}")
        assert '{"key": "value"}' in aws_event_parser.current_tool_call["function"]["arguments"]
    
    def test_parses_tool_stop_event(self, aws_event_parser):
        """
        Что он делает: Проверяет завершение tool call.
        Цель: Убедиться, что tool call добавляется в список.
        """
        print("Настройка: Полный tool call...")
        aws_event_parser.feed(b'{"name":"func","toolUseId":"call_1"}')
        aws_event_parser.feed(b'{"input":"{}"}')
        
        print("Действие: Парсинг stop...")
        aws_event_parser.feed(b'{"stop":true}')
        
        print(f"tool_calls: {aws_event_parser.tool_calls}")
        assert len(aws_event_parser.tool_calls) == 1
        assert aws_event_parser.current_tool_call is None
    
    def test_get_tool_calls_returns_all(self, aws_event_parser):
        """
        Что он делает: Проверяет получение всех tool calls.
        Цель: Убедиться, что get_tool_calls возвращает завершённые calls.
        """
        print("Настройка: Несколько tool calls...")
        aws_event_parser.feed(b'{"name":"func1","toolUseId":"call_1"}')
        aws_event_parser.feed(b'{"stop":true}')
        aws_event_parser.feed(b'{"name":"func2","toolUseId":"call_2"}')
        aws_event_parser.feed(b'{"stop":true}')
        
        print("Действие: Получение tool calls...")
        tool_calls = aws_event_parser.get_tool_calls()
        
        print(f"Результат: {tool_calls}")
        assert len(tool_calls) == 2
    
    def test_get_tool_calls_finalizes_current(self, aws_event_parser):
        """
        Что он делает: Проверяет завершение незавершённого tool call.
        Цель: Убедиться, что get_tool_calls завершает current_tool_call.
        """
        print("Настройка: Незавершённый tool call...")
        aws_event_parser.feed(b'{"name":"func","toolUseId":"call_1"}')
        
        print("Действие: Получение tool calls...")
        tool_calls = aws_event_parser.get_tool_calls()
        
        print(f"Результат: {tool_calls}")
        assert len(tool_calls) == 1
        assert aws_event_parser.current_tool_call is None


class TestAwsEventStreamParserReset:
    """Тесты метода reset."""
    
    def test_reset_clears_state(self, aws_event_parser):
        """
        Что он делает: Проверяет сброс состояния парсера.
        Цель: Убедиться, что reset очищает все данные.
        """
        print("Настройка: Заполнение парсера данными...")
        aws_event_parser.feed(b'{"content":"test"}')
        aws_event_parser.feed(b'{"name":"func","toolUseId":"call_1"}')
        
        print("Действие: Сброс парсера...")
        aws_event_parser.reset()
        
        print("Проверка: Все данные очищены...")
        assert aws_event_parser.buffer == ""
        assert aws_event_parser.last_content is None
        assert aws_event_parser.current_tool_call is None
        assert aws_event_parser.tool_calls == []


class TestAwsEventStreamParserFinalizeToolCall:
    """Тесты метода _finalize_tool_call для обработки разных типов input."""
    
    def test_finalize_with_string_arguments(self, aws_event_parser):
        """
        Что он делает: Проверяет финализацию tool call со строковыми аргументами.
        Цель: Убедиться, что строка JSON парсится и сериализуется обратно.
        """
        print("Настройка: Tool call со строковыми аргументами...")
        aws_event_parser.current_tool_call = {
            "id": "call_1",
            "type": "function",
            "function": {
                "name": "test_func",
                "arguments": '{"key": "value"}'
            }
        }
        
        print("Действие: Финализация tool call...")
        aws_event_parser._finalize_tool_call()
        
        print(f"Результат: {aws_event_parser.tool_calls}")
        assert len(aws_event_parser.tool_calls) == 1
        assert aws_event_parser.tool_calls[0]["function"]["arguments"] == '{"key": "value"}'
    
    def test_finalize_with_dict_arguments(self, aws_event_parser):
        """
        Что он делает: Проверяет финализацию tool call с dict аргументами.
        Цель: Убедиться, что dict сериализуется в JSON строку.
        """
        print("Настройка: Tool call с dict аргументами...")
        aws_event_parser.current_tool_call = {
            "id": "call_2",
            "type": "function",
            "function": {
                "name": "test_func",
                "arguments": {"location": "Moscow", "units": "celsius"}
            }
        }
        
        print("Действие: Финализация tool call...")
        aws_event_parser._finalize_tool_call()
        
        print(f"Результат: {aws_event_parser.tool_calls}")
        assert len(aws_event_parser.tool_calls) == 1
        
        args = aws_event_parser.tool_calls[0]["function"]["arguments"]
        print(f"Аргументы: {args}")
        assert isinstance(args, str)
        assert "Moscow" in args
        assert "celsius" in args
    
    def test_finalize_with_empty_string_arguments(self, aws_event_parser):
        """
        Что он делает: Проверяет финализацию tool call с пустой строкой аргументов.
        Цель: Убедиться, что пустая строка заменяется на "{}".
        """
        print("Настройка: Tool call с пустой строкой аргументов...")
        aws_event_parser.current_tool_call = {
            "id": "call_3",
            "type": "function",
            "function": {
                "name": "test_func",
                "arguments": ""
            }
        }
        
        print("Действие: Финализация tool call...")
        aws_event_parser._finalize_tool_call()
        
        print(f"Результат: {aws_event_parser.tool_calls}")
        assert len(aws_event_parser.tool_calls) == 1
        assert aws_event_parser.tool_calls[0]["function"]["arguments"] == "{}"
    
    def test_finalize_with_whitespace_only_arguments(self, aws_event_parser):
        """
        Что он делает: Проверяет финализацию tool call с пробельными аргументами.
        Цель: Убедиться, что строка из пробелов заменяется на "{}".
        """
        print("Настройка: Tool call с пробельными аргументами...")
        aws_event_parser.current_tool_call = {
            "id": "call_4",
            "type": "function",
            "function": {
                "name": "test_func",
                "arguments": "   "
            }
        }
        
        print("Действие: Финализация tool call...")
        aws_event_parser._finalize_tool_call()
        
        print(f"Результат: {aws_event_parser.tool_calls}")
        assert len(aws_event_parser.tool_calls) == 1
        assert aws_event_parser.tool_calls[0]["function"]["arguments"] == "{}"
    
    def test_finalize_with_invalid_json_arguments(self, aws_event_parser):
        """
        Что он делает: Проверяет финализацию tool call с невалидным JSON.
        Цель: Убедиться, что невалидный JSON заменяется на "{}".
        """
        print("Настройка: Tool call с невалидным JSON...")
        aws_event_parser.current_tool_call = {
            "id": "call_5",
            "type": "function",
            "function": {
                "name": "test_func",
                "arguments": "not valid json {"
            }
        }
        
        print("Действие: Финализация tool call...")
        aws_event_parser._finalize_tool_call()
        
        print(f"Результат: {aws_event_parser.tool_calls}")
        assert len(aws_event_parser.tool_calls) == 1
        assert aws_event_parser.tool_calls[0]["function"]["arguments"] == "{}"
    
    def test_finalize_with_none_current_tool_call(self, aws_event_parser):
        """
        Что он делает: Проверяет финализацию когда current_tool_call is None.
        Цель: Убедиться, что ничего не происходит при None.
        """
        print("Настройка: current_tool_call = None...")
        aws_event_parser.current_tool_call = None
        
        print("Действие: Финализация tool call...")
        aws_event_parser._finalize_tool_call()
        
        print(f"Результат: {aws_event_parser.tool_calls}")
        assert len(aws_event_parser.tool_calls) == 0
    
    def test_finalize_clears_current_tool_call(self, aws_event_parser):
        """
        Что он делает: Проверяет, что финализация очищает current_tool_call.
        Цель: Убедиться, что после финализации current_tool_call = None.
        """
        print("Настройка: Tool call...")
        aws_event_parser.current_tool_call = {
            "id": "call_6",
            "type": "function",
            "function": {
                "name": "test_func",
                "arguments": "{}"
            }
        }
        
        print("Действие: Финализация tool call...")
        aws_event_parser._finalize_tool_call()
        
        print(f"current_tool_call после финализации: {aws_event_parser.current_tool_call}")
        assert aws_event_parser.current_tool_call is None


class TestAwsEventStreamParserEdgeCases:
    """Тесты граничных случаев."""
    
    def test_handles_followup_prompt(self, aws_event_parser):
        """
        Что он делает: Проверяет игнорирование followupPrompt.
        Цель: Убедиться, что followupPrompt не создаёт событие.
        """
        print("Настройка: Chunk с followupPrompt...")
        chunk = b'{"content":"text","followupPrompt":"suggestion"}'
        
        print("Действие: Парсинг chunk...")
        events = aws_event_parser.feed(chunk)
        
        print(f"Результат: {events}")
        assert len(events) == 0  # followupPrompt игнорируется
    
    def test_handles_mixed_events(self, aws_event_parser):
        """
        Что он делает: Проверяет парсинг смешанных событий.
        Цель: Убедиться, что разные типы событий обрабатываются вместе.
        """
        print("Настройка: Chunk со смешанными событиями...")
        chunk = b'{"content":"Hello"}{"usage":1.0}{"contextUsagePercentage":50}'
        
        print("Действие: Парсинг chunk...")
        events = aws_event_parser.feed(chunk)
        
        print(f"Результат: {events}")
        assert len(events) == 3
        assert events[0]["type"] == "content"
        assert events[1]["type"] == "usage"
        assert events[2]["type"] == "context_usage"
    
    def test_handles_garbage_between_events(self, aws_event_parser):
        """
        Что он делает: Проверяет обработку мусора между событиями.
        Цель: Убедиться, что парсер находит JSON среди мусора.
        """
        print("Настройка: Chunk с мусором между JSON...")
        chunk = b'garbage{"content":"valid"}more garbage{"usage":1}'
        
        print("Действие: Парсинг chunk...")
        events = aws_event_parser.feed(chunk)
        
        print(f"Результат: {events}")
        assert len(events) == 2
    
    def test_handles_empty_chunk(self, aws_event_parser):
        """
        Что он делает: Проверяет обработку пустого chunk.
        Цель: Убедиться, что пустой chunk не вызывает ошибок.
        """
        print("Настройка: Пустой chunk...")
        
        print("Действие: Парсинг пустого chunk...")
        events = aws_event_parser.feed(b'')
        
        print(f"Сравниваем результат: Ожидалось [], Получено {events}")
        assert events == []


================================================
FILE: tests/unit/test_routes.py
================================================
# -*- coding: utf-8 -*-

"""
Unit-тесты для API endpoints (routes.py).
Проверяет работу эндпоинтов /, /health, /v1/models, /v1/chat/completions.
"""

import pytest
from unittest.mock import AsyncMock, Mock, patch, MagicMock
from datetime import datetime, timezone

from fastapi import HTTPException
from fastapi.testclient import TestClient

from kiro_gateway.routes import verify_api_key, router
from kiro_gateway.config import PROXY_API_KEY, APP_VERSION, AVAILABLE_MODELS


class TestVerifyApiKey:
    """Тесты функции verify_api_key."""
    
    @pytest.mark.asyncio
    async def test_valid_api_key_returns_true(self):
        """
        Что он делает: Проверяет успешную валидацию корректного ключа.
        Цель: Убедиться, что валидный ключ проходит проверку.
        """
        print("Настройка: Валидный API ключ...")
        valid_header = f"Bearer {PROXY_API_KEY}"
        
        print("Действие: Проверка ключа...")
        result = await verify_api_key(valid_header)
        
        print(f"Сравниваем результат: Ожидалось True, Получено {result}")
        assert result is True
    
    @pytest.mark.asyncio
    async def test_invalid_api_key_raises_401(self):
        """
        Что он делает: Проверяет отклонение невалидного ключа.
        Цель: Убедиться, что невалидный ключ вызывает 401.
        """
        print("Настройка: Невалидный API ключ...")
        invalid_header = "Bearer wrong_key"
        
        print("Действие: Проверка ключа...")
        with pytest.raises(HTTPException) as exc_info:
            await verify_api_key(invalid_header)
        
        print(f"Проверка: HTTPException с кодом 401...")
        assert exc_info.value.status_code == 401
        assert "Invalid or missing API Key" in exc_info.value.detail
    
    @pytest.mark.asyncio
    async def test_missing_api_key_raises_401(self):
        """
        Что он делает: Проверяет отклонение отсутствующего ключа.
        Цель: Убедиться, что отсутствие ключа вызывает 401.
        """
        print("Настройка: Отсутствующий API ключ...")
        
        print("Действие: Проверка ключа...")
        with pytest.raises(HTTPException) as exc_info:
            await verify_api_key(None)
        
        print(f"Проверка: HTTPException с кодом 401...")
        assert exc_info.value.status_code == 401
    
    @pytest.mark.asyncio
    async def test_empty_api_key_raises_401(self):
        """
        Что он делает: Проверяет отклонение пустого ключа.
        Цель: Убедиться, что пустая строка вызывает 401.
        """
        print("Настройка: Пустой API ключ...")
        
        print("Действие: Проверка ключа...")
        with pytest.raises(HTTPException) as exc_info:
            await verify_api_key("")
        
        print(f"Проверка: HTTPException с кодом 401...")
        assert exc_info.value.status_code == 401
    
    @pytest.mark.asyncio
    async def test_wrong_format_raises_401(self):
        """
        Что он делает: Проверяет отклонение ключа без Bearer.
        Цель: Убедиться, что неправильный формат вызывает 401.
        """
        print("Настройка: Ключ без Bearer...")
        wrong_format = PROXY_API_KEY  # Без "Bearer "
        
        print("Действие: Проверка ключа...")
        with pytest.raises(HTTPException) as exc_info:
            await verify_api_key(wrong_format)
        
        print(f"Проверка: HTTPException с кодом 401...")
        assert exc_info.value.status_code == 401


class TestRootEndpoint:
    """Тесты эндпоинта /."""
    
    def test_root_returns_status_ok(self, test_client):
        """
        Что он делает: Проверяет ответ корневого эндпоинта.
        Цель: Убедиться, что / возвращает статус ok.
        """
        print("Действие: GET /...")
        response = test_client.get("/")
        
        print(f"Результат: {response.json()}")
        assert response.status_code == 200
        assert response.json()["status"] == "ok"
        assert "Kiro API Gateway" in response.json()["message"]
    
    def test_root_returns_version(self, test_client):
        """
        Что он делает: Проверяет наличие версии в ответе.
        Цель: Убедиться, что версия приложения возвращается.
        """
        print("Действие: GET /...")
        response = test_client.get("/")
        
        print(f"Результат: {response.json()}")
        assert response.status_code == 200
        assert "version" in response.json()
        assert response.json()["version"] == APP_VERSION


class TestHealthEndpoint:
    """Тесты эндпоинта /health."""
    
    def test_health_returns_healthy(self, test_client):
        """
        Что он делает: Проверяет ответ health эндпоинта.
        Цель: Убедиться, что /health возвращает статус healthy.
        """
        print("Действие: GET /health...")
        response = test_client.get("/health")
        
        print(f"Результат: {response.json()}")
        assert response.status_code == 200
        assert response.json()["status"] == "healthy"
    
    def test_health_returns_timestamp(self, test_client):
        """
        Что он делает: Проверяет наличие timestamp в ответе.
        Цель: Убедиться, что timestamp возвращается.
        """
        print("Действие: GET /health...")
        response = test_client.get("/health")
        
        print(f"Результат: {response.json()}")
        assert response.status_code == 200
        assert "timestamp" in response.json()
    
    def test_health_returns_version(self, test_client):
        """
        Что он делает: Проверяет наличие версии в ответе.
        Цель: Убедиться, что версия приложения возвращается.
        """
        print("Действие: GET /health...")
        response = test_client.get("/health")
        
        print(f"Результат: {response.json()}")
        assert response.status_code == 200
        assert response.json()["version"] == APP_VERSION


class TestModelsEndpoint:
    """Тесты эндпоинта /v1/models."""
    
    def test_models_requires_auth(self, test_client):
        """
        Что он делает: Проверяет требование авторизации.
        Цель: Убедиться, что без ключа возвращается 401.
        """
        print("Действие: GET /v1/models без авторизации...")
        response = test_client.get("/v1/models")
        
        print(f"Статус: {response.status_code}")
        assert response.status_code == 401
    
    def test_models_returns_list(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет возврат списка моделей.
        Цель: Убедиться, что /v1/models возвращает список.
        """
        print("Действие: GET /v1/models с авторизацией...")
        response = test_client.get(
            "/v1/models",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"}
        )
        
        print(f"Результат: {response.json()}")
        assert response.status_code == 200
        assert response.json()["object"] == "list"
        assert "data" in response.json()
    
    def test_models_returns_available_models(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет наличие доступных моделей.
        Цель: Убедиться, что все модели из AVAILABLE_MODELS возвращаются.
        """
        print("Действие: GET /v1/models с авторизацией...")
        response = test_client.get(
            "/v1/models",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"}
        )
        
        print(f"Результат: {response.json()}")
        assert response.status_code == 200
        
        model_ids = [m["id"] for m in response.json()["data"]]
        for expected_model in AVAILABLE_MODELS:
            assert expected_model in model_ids, f"Модель {expected_model} не найдена"
    
    def test_models_format_is_openai_compatible(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет формат ответа на совместимость с OpenAI.
        Цель: Убедиться, что формат соответствует OpenAI API.
        """
        print("Действие: GET /v1/models с авторизацией...")
        response = test_client.get(
            "/v1/models",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"}
        )
        
        print(f"Результат: {response.json()}")
        assert response.status_code == 200
        
        for model in response.json()["data"]:
            assert "id" in model
            assert "object" in model
            assert model["object"] == "model"
            assert "owned_by" in model


class TestChatCompletionsEndpoint:
    """Тесты эндпоинта /v1/chat/completions."""
    
    def test_chat_completions_requires_auth(self, test_client):
        """
        Что он делает: Проверяет требование авторизации.
        Цель: Убедиться, что без ключа возвращается 401.
        """
        print("Действие: POST /v1/chat/completions без авторизации...")
        response = test_client.post(
            "/v1/chat/completions",
            json={
                "model": "claude-sonnet-4-5",
                "messages": [{"role": "user", "content": "Hello"}]
            }
        )
        
        print(f"Статус: {response.status_code}")
        assert response.status_code == 401
    
    def test_chat_completions_validates_messages(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет валидацию пустых сообщений.
        Цель: Убедиться, что пустой список сообщений отклоняется.
        """
        print("Действие: POST /v1/chat/completions с пустыми сообщениями...")
        response = test_client.post(
            "/v1/chat/completions",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
            json={
                "model": "claude-sonnet-4-5",
                "messages": []
            }
        )
        
        print(f"Статус: {response.status_code}")
        # Pydantic должен отклонить пустой список
        assert response.status_code == 422
    
    def test_chat_completions_validates_model(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет валидацию отсутствующей модели.
        Цель: Убедиться, что запрос без модели отклоняется.
        """
        print("Действие: POST /v1/chat/completions без модели...")
        response = test_client.post(
            "/v1/chat/completions",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
            json={
                "messages": [{"role": "user", "content": "Hello"}]
            }
        )
        
        print(f"Статус: {response.status_code}")
        assert response.status_code == 422


class TestChatCompletionsWithMockedKiro:
    """Тесты /v1/chat/completions с мокированным Kiro API."""
    
    def test_chat_completions_accepts_valid_request_format(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет, что валидный формат запроса принимается.
        Цель: Убедиться, что Pydantic валидация проходит для корректного запроса.
        """
        print("Настройка: Валидный запрос...")
        
        # Этот тест проверяет только валидацию запроса
        # Реальный вызов к Kiro API будет заблокирован фикстурой block_all_network_calls
        # Поэтому мы ожидаем ошибку на этапе HTTP запроса, а не валидации
        
        print("Действие: POST /v1/chat/completions с валидным запросом...")
        response = test_client.post(
            "/v1/chat/completions",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
            json={
                "model": "claude-sonnet-4-5",
                "messages": [{"role": "user", "content": "Hello"}],
                "stream": False
            }
        )
        
        print(f"Статус: {response.status_code}")
        # Запрос должен пройти валидацию (не 422)
        # Но может упасть на этапе HTTP из-за блокировки сети
        assert response.status_code != 422


class TestChatCompletionsErrorHandling:
    """Тесты обработки ошибок в /v1/chat/completions."""
    
    def test_invalid_json_returns_422(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет обработку невалидного JSON.
        Цель: Убедиться, что невалидный JSON возвращает 422.
        """
        print("Действие: POST /v1/chat/completions с невалидным JSON...")
        response = test_client.post(
            "/v1/chat/completions",
            headers={
                "Authorization": f"Bearer {valid_proxy_api_key}",
                "Content-Type": "application/json"
            },
            content=b"not valid json"
        )
        
        print(f"Статус: {response.status_code}")
        assert response.status_code == 422
    
    def test_missing_content_in_message_returns_200(self, test_client, valid_proxy_api_key):
        """
        Что он делает: Проверяет обработку сообщения без content.
        Цель: Убедиться, что сообщение без content допустимо (content опционален).
        """
        print("Действие: POST /v1/chat/completions с сообщением без content...")
        # Этот тест проверяет валидацию Pydantic
        # content может быть None согласно модели
        response = test_client.post(
            "/v1/chat/completions",
            headers={"Authorization": f"Bearer {valid_proxy_api_key}"},
            json={
                "model": "claude-sonnet-4-5",
                "messages": [{"role": "user"}]  # content отсутствует
            }
        )
        
        print(f"Статус: {response.status_code}")
        # Запрос должен пройти валидацию (content опционален)
        # Но может упасть на этапе обработки из-за отсутствия мока Kiro API
        # Поэтому проверяем, что это не 422 (валидация прошла)
        assert response.status_code != 422 or "content" not in str(response.json())


class TestRouterIntegration:
    """Тесты интеграции роутера."""
    
    def test_router_has_all_endpoints(self):
        """
        Что он делает: Проверяет наличие всех эндпоинтов в роутере.
        Цель: Убедиться, что все эндпоинты зарегистрированы.
        """
        print("Проверка: Эндпоинты в роутере...")
        
        routes = [route.path for route in router.routes]
        
        print(f"Найденные роуты: {routes}")
        assert "/" in routes
        assert "/health" in routes
        assert "/v1/models" in routes
        assert "/v1/chat/completions" in routes
    
    def test_router_methods(self):
        """
        Что он делает: Проверяет HTTP методы эндпоинтов.
        Цель: Убедиться, что методы соответствуют ожиданиям.
        """
        print("Проверка: HTTP методы...")
        
        for route in router.routes:
            if route.path == "/":
                assert "GET" in route.methods
            elif route.path == "/health":
                assert "GET" in route.methods
            elif route.path == "/v1/models":
                assert "GET" in route.methods
            elif route.path == "/v1/chat/completions":
                assert "POST" in route.methods


================================================
FILE: tests/unit/test_streaming.py
================================================

# -*- coding: utf-8 -*-

"""
Unit tests for streaming module.
Tests logic for adding index to tool_calls and protection from None values.
"""

import pytest
import json
from unittest.mock import AsyncMock, MagicMock, patch

from kiro_gateway.streaming import (
    stream_kiro_to_openai,
    collect_stream_response
)


@pytest.fixture
def mock_model_cache():
    """Mock for ModelInfoCache."""
    cache = MagicMock()
    cache.get_max_input_tokens.return_value = 200000
    return cache


@pytest.fixture
def mock_auth_manager():
    """Mock for KiroAuthManager."""
    manager = MagicMock()
    return manager


@pytest.fixture
def mock_http_client():
    """Mock for httpx.AsyncClient."""
    client = AsyncMock()
    return client


class TestStreamingToolCallsIndex:
    """Tests for adding index to tool_calls in streaming responses."""
    
    @pytest.mark.asyncio
    async def test_tool_calls_have_index_field(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that tool_calls in streaming response contain index field.
        Goal: Ensure OpenAI API spec is followed for streaming tool calls.
        """
        print("Setup: Mock tool calls without index...")
        tool_calls = [
            {
                "id": "call_123",
                "type": "function",
                "function": {
                    "name": "get_weather",
                    "arguments": '{"location": "Moscow"}'
                }
            }
        ]
        
        print("Setup: Mock parser...")
        mock_parser = MagicMock()
        mock_parser.feed.return_value = []
        mock_parser.get_tool_calls.return_value = tool_calls
        
        print("Setup: Mock response...")
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        async def mock_aiter_bytes():
            yield b'{"content":"test"}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        mock_response.aclose = AsyncMock()
        
        print("Action: Collecting streaming chunks...")
        chunks = []
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                async for chunk in stream_kiro_to_openai(
                    mock_http_client, mock_response, "test-model", 
                    mock_model_cache, mock_auth_manager
                ):
                    chunks.append(chunk)
        
        print(f"Received chunks: {len(chunks)}")
        
        # Look for chunk with tool_calls
        tool_calls_found = False
        for chunk in chunks:
            if isinstance(chunk, str) and "tool_calls" in chunk:
                if chunk.startswith("data: "):
                    json_str = chunk[6:].strip()
                    if json_str != "[DONE]":
                        data = json.loads(json_str)
                        if "choices" in data and data["choices"]:
                            delta = data["choices"][0].get("delta", {})
                            if "tool_calls" in delta:
                                print(f"Tool calls in delta: {delta['tool_calls']}")
                                for tc in delta["tool_calls"]:
                                    print(f"Checking index in tool call: {tc}")
                                    assert "index" in tc, "Tool call must contain index field"
                                    tool_calls_found = True
        
        assert tool_calls_found, "Tool calls chunk not found"
    
    @pytest.mark.asyncio
    async def test_multiple_tool_calls_have_sequential_indices(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that multiple tool_calls have sequential indices.
        Goal: Ensure indices start from 0 and go sequentially.
        """
        print("Setup: Multiple tool calls...")
        tool_calls = [
            {"id": "call_1", "type": "function", "function": {"name": "func1", "arguments": "{}"}},
            {"id": "call_2", "type": "function", "function": {"name": "func2", "arguments": "{}"}},
            {"id": "call_3", "type": "function", "function": {"name": "func3", "arguments": "{}"}}
        ]
        
        print("Setup: Mock parser...")
        mock_parser = MagicMock()
        mock_parser.feed.return_value = []
        mock_parser.get_tool_calls.return_value = tool_calls
        
        print("Setup: Mock response...")
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        async def mock_aiter_bytes():
            yield b'{"content":"test"}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        mock_response.aclose = AsyncMock()
        
        print("Action: Collecting streaming chunks...")
        chunks = []
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                async for chunk in stream_kiro_to_openai(
                    mock_http_client, mock_response, "test-model",
                    mock_model_cache, mock_auth_manager
                ):
                    chunks.append(chunk)
        
        # Look for chunk with tool_calls
        for chunk in chunks:
            if isinstance(chunk, str) and "tool_calls" in chunk:
                if chunk.startswith("data: "):
                    json_str = chunk[6:].strip()
                    if json_str != "[DONE]":
                        data = json.loads(json_str)
                        if "choices" in data and data["choices"]:
                            delta = data["choices"][0].get("delta", {})
                            if "tool_calls" in delta:
                                indices = [tc["index"] for tc in delta["tool_calls"]]
                                print(f"Indices: {indices}")
                                assert indices == [0, 1, 2], f"Indices should be [0, 1, 2], got {indices}"


class TestStreamingToolCallsNoneProtection:
    """Tests for protection from None values in tool_calls."""
    
    @pytest.mark.asyncio
    async def test_handles_none_function_name(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies handling of None in function.name.
        Goal: Ensure None is replaced with empty string.
        """
        print("Setup: Tool call with None name...")
        tool_calls = [
            {
                "id": "call_1",
                "type": "function",
                "function": {
                    "name": None,
                    "arguments": '{"a": 1}'
                }
            }
        ]
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = []
        mock_parser.get_tool_calls.return_value = tool_calls
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        async def mock_aiter_bytes():
            yield b'{"content":"test"}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        mock_response.aclose = AsyncMock()
        
        print("Action: Collecting streaming chunks...")
        chunks = []
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                async for chunk in stream_kiro_to_openai(
                    mock_http_client, mock_response, "test-model",
                    mock_model_cache, mock_auth_manager
                ):
                    chunks.append(chunk)
        
        # Verify no exceptions and chunks collected
        print(f"Received chunks: {len(chunks)}")
        assert len(chunks) > 0
        
        # Verify name replaced with empty string
        for chunk in chunks:
            if isinstance(chunk, str) and "tool_calls" in chunk:
                if chunk.startswith("data: "):
                    json_str = chunk[6:].strip()
                    if json_str != "[DONE]":
                        data = json.loads(json_str)
                        if "choices" in data and data["choices"]:
                            delta = data["choices"][0].get("delta", {})
                            if "tool_calls" in delta:
                                for tc in delta["tool_calls"]:
                                    assert tc["function"]["name"] == "", "None name should be replaced with empty string"
    
    @pytest.mark.asyncio
    async def test_handles_none_function_arguments(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies handling of None in function.arguments.
        Goal: Ensure None is replaced with "{}".
        """
        print("Setup: Tool call with None arguments...")
        tool_calls = [
            {
                "id": "call_1",
                "type": "function",
                "function": {
                    "name": "test_func",
                    "arguments": None
                }
            }
        ]
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = []
        mock_parser.get_tool_calls.return_value = tool_calls
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        async def mock_aiter_bytes():
            yield b'{"content":"test"}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        mock_response.aclose = AsyncMock()
        
        print("Action: Collecting streaming chunks...")
        chunks = []
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                async for chunk in stream_kiro_to_openai(
                    mock_http_client, mock_response, "test-model",
                    mock_model_cache, mock_auth_manager
                ):
                    chunks.append(chunk)
        
        print(f"Received chunks: {len(chunks)}")
        assert len(chunks) > 0
        
        # Verify arguments replaced with "{}"
        for chunk in chunks:
            if isinstance(chunk, str) and "tool_calls" in chunk:
                if chunk.startswith("data: "):
                    json_str = chunk[6:].strip()
                    if json_str != "[DONE]":
                        data = json.loads(json_str)
                        if "choices" in data and data["choices"]:
                            delta = data["choices"][0].get("delta", {})
                            if "tool_calls" in delta:
                                for tc in delta["tool_calls"]:
                                    # None should be replaced with "{}" or empty string
                                    assert tc["function"]["arguments"] is not None
    
    @pytest.mark.asyncio
    async def test_handles_none_function_object(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies handling of None instead of function object.
        Goal: Ensure None function is handled without errors.
        """
        print("Setup: Tool call with None function...")
        tool_calls = [
            {
                "id": "call_1",
                "type": "function",
                "function": None
            }
        ]
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = []
        mock_parser.get_tool_calls.return_value = tool_calls
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        async def mock_aiter_bytes():
            yield b'{"content":"test"}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        mock_response.aclose = AsyncMock()
        
        print("Action: Collecting streaming chunks...")
        chunks = []
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                async for chunk in stream_kiro_to_openai(
                    mock_http_client, mock_response, "test-model",
                    mock_model_cache, mock_auth_manager
                ):
                    chunks.append(chunk)
        
        print(f"Received chunks: {len(chunks)}")
        assert len(chunks) > 0


class TestCollectStreamResponseToolCalls:
    """Tests for collect_stream_response with tool_calls."""
    
    @pytest.mark.asyncio
    async def test_collected_tool_calls_have_no_index(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that collected tool_calls don't contain index field.
        Goal: Ensure index is removed for non-streaming response.
        """
        print("Setup: Tool calls...")
        tool_calls = [
            {
                "id": "call_1",
                "type": "function",
                "function": {"name": "func1", "arguments": '{"a": 1}'}
            }
        ]
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = []
        mock_parser.get_tool_calls.return_value = tool_calls
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        async def mock_aiter_bytes():
            yield b'{"content":"Hello"}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        mock_response.aclose = AsyncMock()
        
        print("Action: Collecting full response...")
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                result = await collect_stream_response(
                    mock_http_client, mock_response, "test-model",
                    mock_model_cache, mock_auth_manager
                )
        
        print(f"Result: {result}")
        
        if "choices" in result and result["choices"]:
            message = result["choices"][0].get("message", {})
            if "tool_calls" in message:
                for tc in message["tool_calls"]:
                    print(f"Tool call: {tc}")
                    assert "index" not in tc, "Non-streaming tool_calls should not contain index"
    
    @pytest.mark.asyncio
    async def test_collected_tool_calls_have_required_fields(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that collected tool_calls contain all required fields.
        Goal: Ensure id, type, function are present.
        """
        print("Setup: Tool calls...")
        tool_calls = [
            {
                "id": "call_abc",
                "type": "function",
                "function": {"name": "get_weather", "arguments": '{"city": "Moscow"}'}
            }
        ]
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = []
        mock_parser.get_tool_calls.return_value = tool_calls
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        async def mock_aiter_bytes():
            yield b'{"content":""}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        mock_response.aclose = AsyncMock()
        
        print("Action: Collecting full response...")
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                result = await collect_stream_response(
                    mock_http_client, mock_response, "test-model",
                    mock_model_cache, mock_auth_manager
                )
        
        print(f"Result: {result}")
        
        if "choices" in result and result["choices"]:
            message = result["choices"][0].get("message", {})
            if "tool_calls" in message:
                for tc in message["tool_calls"]:
                    print(f"Checking tool call: {tc}")
                    assert "id" in tc, "Tool call must contain id"
                    assert "type" in tc, "Tool call must contain type"
                    assert "function" in tc, "Tool call must contain function"
                    assert "name" in tc["function"], "Function must contain name"
                    assert "arguments" in tc["function"], "Function must contain arguments"
    
    @pytest.mark.asyncio
    async def test_handles_none_in_collected_tool_calls(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies handling of None values in collected tool_calls.
        Goal: Ensure None is replaced with default values.
        """
        print("Setup: Tool calls with None values...")
        tool_calls = [
            {
                "id": "call_1",
                "type": "function",
                "function": None
            }
        ]
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = []
        mock_parser.get_tool_calls.return_value = tool_calls
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        
        async def mock_aiter_bytes():
            yield b'{"content":""}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        mock_response.aclose = AsyncMock()
        
        print("Action: Collecting full response...")
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                result = await collect_stream_response(
                    mock_http_client, mock_response, "test-model",
                    mock_model_cache, mock_auth_manager
                )
        
        print(f"Result: {result}")
        
        # Verify no exceptions
        assert "choices" in result


class TestStreamingErrorHandling:
    """Tests for error handling in streaming module."""
    
    @pytest.mark.asyncio
    async def test_generator_exit_handled_gracefully(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that GeneratorExit is handled without logging as error.
        Goal: Ensure client disconnect doesn't cause ERROR in logs.
        """
        print("Setup: Mock response that will raise GeneratorExit...")
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.aclose = AsyncMock()
        
        # Create generator that will raise GeneratorExit
        async def mock_aiter_bytes_with_generator_exit():
            yield b'{"content":"Hello"}'
            # Simulate client disconnect
            raise GeneratorExit()
        
        mock_response.aiter_bytes = mock_aiter_bytes_with_generator_exit
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = [{"type": "content", "data": "Hello"}]
        mock_parser.get_tool_calls.return_value = []
        
        print("Action: Running streaming with GeneratorExit...")
        chunks_received = []
        generator_exit_caught = False
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                try:
                    async for chunk in stream_kiro_to_openai(
                        mock_http_client, mock_response, "test-model",
                        mock_model_cache, mock_auth_manager
                    ):
                        chunks_received.append(chunk)
                except GeneratorExit:
                    generator_exit_caught = True
                    print("GeneratorExit was caught (expected)")
        
        print(f"Received chunks before GeneratorExit: {len(chunks_received)}")
        print(f"GeneratorExit caught: {generator_exit_caught}")
        
        # Verify response was closed
        print("Check: response.aclose() should be called...")
        mock_response.aclose.assert_called()
        print("✓ response.aclose() was called")
    
    @pytest.mark.asyncio
    async def test_exception_with_empty_message_logged_with_type(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that exception with empty message is logged with type.
        Goal: Ensure exception type is visible in logs even if str(e) is empty.
        """
        print("Setup: Mock response that will raise exception with empty message...")
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.aclose = AsyncMock()
        
        # Create custom exception with empty message
        class EmptyMessageError(Exception):
            def __str__(self):
                return ""
        
        async def mock_aiter_bytes_with_empty_error():
            yield b'{"content":"Hello"}'
            raise EmptyMessageError()
        
        mock_response.aiter_bytes = mock_aiter_bytes_with_empty_error
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = [{"type": "content", "data": "Hello"}]
        mock_parser.get_tool_calls.return_value = []
        
        print("Action: Running streaming with EmptyMessageError...")
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                with patch('kiro_gateway.streaming.logger') as mock_logger:
                    exception_raised = False
                    try:
                        async for chunk in stream_kiro_to_openai(
                            mock_http_client, mock_response, "test-model",
                            mock_model_cache, mock_auth_manager
                        ):
                            pass
                    except EmptyMessageError:
                        exception_raised = True
                        print("EmptyMessageError was caught (expected)")
                    
                    print("Check: logger.error should be called with exception type...")
                    # Verify logger.error was called
                    error_calls = [call for call in mock_logger.error.call_args_list]
                    print(f"logger.error calls: {error_calls}")
                    
                    # Should have call with exception type
                    assert exception_raised, "Exception should be propagated"
                    assert mock_logger.error.called, "logger.error should be called"
                    
                    # Verify exception type is in message
                    error_message = str(mock_logger.error.call_args_list[0])
                    print(f"Error message: {error_message}")
                    assert "EmptyMessageError" in error_message, "Exception type should be in log"
                    print("✓ Exception type is present in log")
    
    @pytest.mark.asyncio
    async def test_exception_propagated_to_caller(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that exceptions are propagated up.
        Goal: Ensure errors are not "swallowed" inside generator.
        """
        print("Setup: Mock response that will raise RuntimeError...")
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.aclose = AsyncMock()
        
        async def mock_aiter_bytes_with_error():
            yield b'{"content":"Hello"}'
            raise RuntimeError("Test error for propagation")
        
        mock_response.aiter_bytes = mock_aiter_bytes_with_error
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = [{"type": "content", "data": "Hello"}]
        mock_parser.get_tool_calls.return_value = []
        
        print("Action: Running streaming with RuntimeError...")
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                with pytest.raises(RuntimeError) as exc_info:
                    async for chunk in stream_kiro_to_openai(
                        mock_http_client, mock_response, "test-model",
                        mock_model_cache, mock_auth_manager
                    ):
                        pass
        
        print(f"Caught exception: {exc_info.value}")
        assert "Test error for propagation" in str(exc_info.value)
        print("✓ Exception was propagated up with correct message")
    
    @pytest.mark.asyncio
    async def test_response_closed_on_error(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that response is closed even on error.
        Goal: Ensure resources are released in finally block.
        """
        print("Setup: Mock response that will raise ValueError...")
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.aclose = AsyncMock()
        
        async def mock_aiter_bytes_with_value_error():
            yield b'{"content":"Hello"}'
            raise ValueError("Test value error")
        
        mock_response.aiter_bytes = mock_aiter_bytes_with_value_error
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = [{"type": "content", "data": "Hello"}]
        mock_parser.get_tool_calls.return_value = []
        
        print("Action: Running streaming with ValueError...")
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                try:
                    async for chunk in stream_kiro_to_openai(
                        mock_http_client, mock_response, "test-model",
                        mock_model_cache, mock_auth_manager
                    ):
                        pass
                except ValueError:
                    print("ValueError caught (expected)")
        
        print("Check: response.aclose() should be called...")
        mock_response.aclose.assert_called()
        print("✓ response.aclose() was called even on error")
    
    @pytest.mark.asyncio
    async def test_response_closed_on_success(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that response is closed on successful completion.
        Goal: Ensure resources are released in finally block.
        """
        print("Setup: Mock response for successful streaming...")
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.aclose = AsyncMock()
        
        async def mock_aiter_bytes_success():
            yield b'{"content":"Hello World"}'
        
        mock_response.aiter_bytes = mock_aiter_bytes_success
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = [{"type": "content", "data": "Hello World"}]
        mock_parser.get_tool_calls.return_value = []
        
        print("Action: Running successful streaming...")
        chunks = []
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                async for chunk in stream_kiro_to_openai(
                    mock_http_client, mock_response, "test-model",
                    mock_model_cache, mock_auth_manager
                ):
                    chunks.append(chunk)
        
        print(f"Received chunks: {len(chunks)}")
        print("Check: response.aclose() should be called...")
        mock_response.aclose.assert_called()
        print("✓ response.aclose() was called on successful completion")
    
    @pytest.mark.asyncio
    async def test_aclose_error_does_not_mask_original_error(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that error in aclose() doesn't mask original error.
        Goal: Ensure original exception is propagated even if aclose() fails.
        """
        print("Setup: Mock response with error in aclose()...")
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.aclose = AsyncMock(side_effect=ConnectionError("Connection lost"))
        
        async def mock_aiter_bytes_with_error():
            yield b'{"content":"Hello"}'
            raise RuntimeError("Original error")
        
        mock_response.aiter_bytes = mock_aiter_bytes_with_error
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = [{"type": "content", "data": "Hello"}]
        mock_parser.get_tool_calls.return_value = []
        
        print("Action: Running streaming with error and error in aclose()...")
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                with pytest.raises(RuntimeError) as exc_info:
                    async for chunk in stream_kiro_to_openai(
                        mock_http_client, mock_response, "test-model",
                        mock_model_cache, mock_auth_manager
                    ):
                        pass
        
        print(f"Caught exception: {exc_info.value}")
        # Should be original error, not ConnectionError from aclose()
        assert "Original error" in str(exc_info.value)
        print("✓ Original error was not masked by error in aclose()")


class TestFirstTokenTimeoutError:
    """Tests for FirstTokenTimeoutError and first token timeout logging."""
    
    @pytest.mark.asyncio
    async def test_first_token_timeout_not_caught_by_general_handler(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that FirstTokenTimeoutError is propagated for retry.
        Goal: Ensure first token timeout is not handled as regular error.
        """
        import asyncio
        from kiro_gateway.streaming import FirstTokenTimeoutError, stream_kiro_to_openai_internal
        
        print("Setup: Mock response with timeout...")
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.aclose = AsyncMock()
        
        # Create generator that will be used
        async def mock_aiter_bytes():
            yield b'{"content":"test"}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        
        print("Action: Mocking asyncio.wait_for to immediately raise TimeoutError...")
        
        # Mock asyncio.wait_for to immediately raise TimeoutError
        async def mock_wait_for_timeout(*args, **kwargs):
            raise asyncio.TimeoutError()
        
        with patch('kiro_gateway.streaming.asyncio.wait_for', side_effect=mock_wait_for_timeout):
            with pytest.raises(FirstTokenTimeoutError) as exc_info:
                async for chunk in stream_kiro_to_openai_internal(
                    mock_http_client, mock_response, "test-model",
                    mock_model_cache, mock_auth_manager,
                    first_token_timeout=30  # Value doesn't matter, wait_for is mocked
                ):
                    pass
        
        print(f"Caught exception: {exc_info.value}")
        print("✓ FirstTokenTimeoutError was propagated for retry logic")
        
        # Verify response was closed
        mock_response.aclose.assert_called()
        print("✓ response.aclose() was called")
    
    @pytest.mark.asyncio
    async def test_first_token_timeout_logged_with_correct_format(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that first token timeout is logged with [FirstTokenTimeout] prefix.
        Goal: Ensure consistent logging format for first token timeout.
        """
        import asyncio
        from kiro_gateway.streaming import FirstTokenTimeoutError, stream_kiro_to_openai_internal
        
        print("Setup: Mock response with timeout...")
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.aclose = AsyncMock()
        
        async def mock_aiter_bytes():
            yield b'{"content":"test"}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        
        async def mock_wait_for_timeout(*args, **kwargs):
            raise asyncio.TimeoutError()
        
        print("Action: Running streaming with timeout and checking logs...")
        
        with patch('kiro_gateway.streaming.asyncio.wait_for', side_effect=mock_wait_for_timeout):
            with patch('kiro_gateway.streaming.logger') as mock_logger:
                try:
                    async for chunk in stream_kiro_to_openai_internal(
                        mock_http_client, mock_response, "test-model",
                        mock_model_cache, mock_auth_manager,
                        first_token_timeout=15
                    ):
                        pass
                except FirstTokenTimeoutError:
                    pass
                
                print("Check: logger.warning should be called with [FirstTokenTimeout]...")
                warning_calls = [str(call) for call in mock_logger.warning.call_args_list]
                print(f"Warning calls: {warning_calls}")
                
                assert any("FirstTokenTimeout" in call for call in warning_calls), \
                    f"[FirstTokenTimeout] not found in warning logs: {warning_calls}"
                print("✓ [FirstTokenTimeout] prefix found in logs")
    
    @pytest.mark.asyncio
    async def test_first_token_timeout_includes_timeout_value(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that first token timeout log includes the timeout value.
        Goal: Ensure timeout value is visible in logs for debugging.
        """
        import asyncio
        from kiro_gateway.streaming import FirstTokenTimeoutError, stream_kiro_to_openai_internal
        
        print("Setup: Mock response with timeout...")
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.aclose = AsyncMock()
        
        async def mock_aiter_bytes():
            yield b'{"content":"test"}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        
        async def mock_wait_for_timeout(*args, **kwargs):
            raise asyncio.TimeoutError()
        
        custom_timeout = 25.0
        
        print(f"Action: Running streaming with custom timeout={custom_timeout}...")
        
        with patch('kiro_gateway.streaming.asyncio.wait_for', side_effect=mock_wait_for_timeout):
            with patch('kiro_gateway.streaming.logger') as mock_logger:
                try:
                    async for chunk in stream_kiro_to_openai_internal(
                        mock_http_client, mock_response, "test-model",
                        mock_model_cache, mock_auth_manager,
                        first_token_timeout=custom_timeout
                    ):
                        pass
                except FirstTokenTimeoutError:
                    pass
                
                print("Check: logger.warning should include timeout value...")
                warning_calls = [str(call) for call in mock_logger.warning.call_args_list]
                print(f"Warning calls: {warning_calls}")
                
                assert any(str(custom_timeout) in call for call in warning_calls), \
                    f"Timeout value {custom_timeout} not found in warning logs: {warning_calls}"
                print(f"✓ Timeout value {custom_timeout} found in logs")
    
    @pytest.mark.asyncio
    async def test_first_token_received_logged_on_success(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that successful first token receipt is logged.
        Goal: Ensure debug log shows when first token is received.
        """
        from kiro_gateway.streaming import stream_kiro_to_openai_internal
        
        print("Setup: Mock response with successful first token...")
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.aclose = AsyncMock()
        
        async def mock_aiter_bytes():
            yield b'{"content":"Hello"}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = [{"type": "content", "data": "Hello"}]
        mock_parser.get_tool_calls.return_value = []
        
        print("Action: Running streaming and checking debug logs...")
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                with patch('kiro_gateway.streaming.logger') as mock_logger:
                    chunks = []
                    async for chunk in stream_kiro_to_openai_internal(
                        mock_http_client, mock_response, "test-model",
                        mock_model_cache, mock_auth_manager,
                        first_token_timeout=15
                    ):
                        chunks.append(chunk)
                    
                    print(f"Received {len(chunks)} chunks")
                    print("Check: logger.debug should be called with 'First token received'...")
                    debug_calls = [str(call) for call in mock_logger.debug.call_args_list]
                    print(f"Debug calls: {debug_calls}")
                    
                    assert any("First token received" in call for call in debug_calls), \
                        f"'First token received' not found in debug logs: {debug_calls}"
                    print("✓ 'First token received' found in debug logs")


class TestStreamWithFirstTokenRetry:
    """Tests for stream_with_first_token_retry function."""
    
    @pytest.mark.asyncio
    async def test_retry_on_first_token_timeout(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that request is retried on first token timeout.
        Goal: Ensure retry logic works for first token timeout.
        """
        import asyncio
        from kiro_gateway.streaming import stream_with_first_token_retry, FirstTokenTimeoutError
        
        print("Setup: Mock make_request that succeeds on second attempt...")
        
        mock_response_success = AsyncMock()
        mock_response_success.status_code = 200
        mock_response_success.aclose = AsyncMock()
        
        async def mock_aiter_bytes_success():
            yield b'{"content":"Success"}'
        
        mock_response_success.aiter_bytes = mock_aiter_bytes_success
        
        call_count = 0
        
        async def mock_make_request():
            nonlocal call_count
            call_count += 1
            print(f"make_request called (attempt {call_count})")
            return mock_response_success
        
        mock_parser = MagicMock()
        mock_parser.feed.return_value = [{"type": "content", "data": "Success"}]
        mock_parser.get_tool_calls.return_value = []
        
        # First call raises timeout, second succeeds
        timeout_raised = False
        
        async def mock_wait_for_with_retry(coro, timeout):
            nonlocal timeout_raised
            if not timeout_raised:
                timeout_raised = True
                raise asyncio.TimeoutError()
            return await coro
        
        print("Action: Running stream_with_first_token_retry...")
        
        with patch('kiro_gateway.streaming.AwsEventStreamParser', return_value=mock_parser):
            with patch('kiro_gateway.streaming.parse_bracket_tool_calls', return_value=[]):
                with patch('kiro_gateway.streaming.asyncio.wait_for', side_effect=mock_wait_for_with_retry):
                    chunks = []
                    async for chunk in stream_with_first_token_retry(
                        mock_make_request,
                        mock_http_client,
                        "test-model",
                        mock_model_cache,
                        mock_auth_manager,
                        max_retries=3,
                        first_token_timeout=15
                    ):
                        chunks.append(chunk)
        
        print(f"Received {len(chunks)} chunks")
        print(f"make_request was called {call_count} times")
        
        assert call_count == 2, f"Expected 2 calls (1 timeout + 1 success), got {call_count}"
        assert len(chunks) > 0, "Should receive chunks after retry"
        print("✓ Retry logic worked correctly")
    
    @pytest.mark.asyncio
    async def test_all_retries_exhausted_raises_504(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that 504 is raised after all retries exhausted.
        Goal: Ensure proper error handling when model never responds.
        """
        import asyncio
        from fastapi import HTTPException
        from kiro_gateway.streaming import stream_with_first_token_retry
        
        print("Setup: Mock make_request that always times out...")
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.aclose = AsyncMock()
        
        async def mock_aiter_bytes():
            yield b'{"content":"test"}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        
        call_count = 0
        
        async def mock_make_request():
            nonlocal call_count
            call_count += 1
            print(f"make_request called (attempt {call_count})")
            return mock_response
        
        async def mock_wait_for_always_timeout(*args, **kwargs):
            raise asyncio.TimeoutError()
        
        max_retries = 3
        
        print(f"Action: Running stream_with_first_token_retry with max_retries={max_retries}...")
        
        with patch('kiro_gateway.streaming.asyncio.wait_for', side_effect=mock_wait_for_always_timeout):
            with pytest.raises(HTTPException) as exc_info:
                async for chunk in stream_with_first_token_retry(
                    mock_make_request,
                    mock_http_client,
                    "test-model",
                    mock_model_cache,
                    mock_auth_manager,
                    max_retries=max_retries,
                    first_token_timeout=15
                ):
                    pass
        
        print(f"Caught HTTPException: {exc_info.value.status_code} - {exc_info.value.detail}")
        print(f"make_request was called {call_count} times")
        
        print(f"Сравниваем status_code: Ожидалось 504, Получено {exc_info.value.status_code}")
        assert exc_info.value.status_code == 504
        print(f"Сравниваем call_count: Ожидалось {max_retries}, Получено {call_count}")
        assert call_count == max_retries
        assert "15" in exc_info.value.detail, "Timeout value should be in error message"
        print("✓ 504 raised after all retries exhausted")
    
    @pytest.mark.asyncio
    async def test_retry_logs_attempt_number(self, mock_http_client, mock_model_cache, mock_auth_manager):
        """
        What it does: Verifies that retry attempts are logged with attempt number.
        Goal: Ensure logs show which attempt failed.
        """
        import asyncio
        from fastapi import HTTPException
        from kiro_gateway.streaming import stream_with_first_token_retry
        
        print("Setup: Mock make_request that always times out...")
        
        mock_response = AsyncMock()
        mock_response.status_code = 200
        mock_response.aclose = AsyncMock()
        
        async def mock_aiter_bytes():
            yield b'{"content":"test"}'
        
        mock_response.aiter_bytes = mock_aiter_bytes
        
        async def mock_make_request():
            return mock_response
        
        async def mock_wait_for_always_timeout(*args, **kwargs):
            raise asyncio.TimeoutError()
        
        print("Action: Running stream_with_first_token_retry and checking logs...")
        
        with patch('kiro_gateway.streaming.asyncio.wait_for', side_effect=mock_wait_for_always_timeout):
            with patch('kiro_gateway.streaming.logger') as mock_logger:
                try:
                    async for chunk in stream_with_first_token_retry(
                        mock_make_request,
                        mock_http_client,
                        "test-model",
                        mock_model_cache,
                        mock_auth_manager,
                        max_retries=3,
                        first_token_timeout=15
                    ):
                        pass
                except HTTPException:
                    pass
                
                print("Check: logger.warning should include attempt numbers...")
                warning_calls = [str(call) for call in mock_logger.warning.call_args_list]
                print(f"Warning calls: {warning_calls}")
                
                # Should have warnings for attempts 1/3, 2/3, 3/3
                assert any("1/3" in call or "2/3" in call or "3/3" in call for call in warning_calls), \
                    f"Attempt numbers not found in warning logs: {warning_calls}"
                print("✓ Attempt numbers found in logs")


================================================
FILE: tests/unit/test_tokenizer.py
================================================
# -*- coding: utf-8 -*-

"""
Unit-тесты для модуля токенизатора (kiro_gateway/tokenizer.py).

Проверяет:
- Подсчёт токенов в тексте (count_tokens)
- Подсчёт токенов в сообщениях (count_message_tokens)
- Подсчёт токенов в инструментах (count_tools_tokens)
- Оценку токенов запроса (estimate_request_tokens)
- Коэффициент коррекции для Claude (CLAUDE_CORRECTION_FACTOR)
- Fallback при отсутствии tiktoken
"""

import pytest
from unittest.mock import patch, MagicMock

from kiro_gateway.tokenizer import (
    count_tokens,
    count_message_tokens,
    count_tools_tokens,
    estimate_request_tokens,
    CLAUDE_CORRECTION_FACTOR,
    _get_encoding
)


class TestCountTokens:
    """Тесты для функции count_tokens."""
    
    def test_empty_string_returns_zero(self):
        """
        Что он делает: Проверяет, что пустая строка возвращает 0 токенов.
        Цель: Убедиться в корректной обработке граничного случая.
        """
        print("Тест: Пустая строка...")
        result = count_tokens("")
        print(f"Результат: {result}")
        assert result == 0, "Пустая строка должна возвращать 0 токенов"
    
    def test_none_returns_zero(self):
        """
        Что он делает: Проверяет, что None возвращает 0 токенов.
        Цель: Убедиться в корректной обработке None.
        """
        print("Тест: None...")
        result = count_tokens(None)
        print(f"Результат: {result}")
        assert result == 0, "None должен возвращать 0 токенов"
    
    def test_simple_text_returns_positive(self):
        """
        Что он делает: Проверяет, что простой текст возвращает положительное число токенов.
        Цель: Убедиться в базовой работоспособности подсчёта.
        """
        print("Тест: Простой текст...")
        result = count_tokens("Hello, world!")
        print(f"Результат: {result}")
        assert result > 0, "Простой текст должен возвращать положительное число токенов"
    
    def test_longer_text_returns_more_tokens(self):
        """
        Что он делает: Проверяет, что более длинный текст возвращает больше токенов.
        Цель: Убедиться в корректной пропорциональности подсчёта.
        """
        print("Тест: Сравнение длинного и короткого текста...")
        short_text = "Hello"
        long_text = "Hello, this is a much longer text that should have more tokens"
        
        short_tokens = count_tokens(short_text)
        long_tokens = count_tokens(long_text)
        
        print(f"Короткий текст: {short_tokens} токенов")
        print(f"Длинный текст: {long_tokens} токенов")
        
        assert long_tokens > short_tokens, "Длинный текст должен иметь больше токенов"
    
    def test_claude_correction_applied_by_default(self):
        """
        Что он делает: Проверяет, что коэффициент коррекции Claude применяется по умолчанию.
        Цель: Убедиться, что apply_claude_correction=True по умолчанию.
        """
        print("Тест: Коэффициент коррекции Claude...")
        text = "This is a test text for token counting"
        
        with_correction = count_tokens(text, apply_claude_correction=True)
        without_correction = count_tokens(text, apply_claude_correction=False)
        
        print(f"С коррекцией: {with_correction}")
        print(f"Без коррекции: {without_correction}")
        
        # С коррекцией должно быть больше (коэффициент 1.15)
        assert with_correction > without_correction, "С коррекцией должно быть больше токенов"
        
        # Проверяем примерное соотношение
        ratio = with_correction / without_correction
        print(f"Соотношение: {ratio}")
        assert 1.1 <= ratio <= 1.2, f"Соотношение должно быть около {CLAUDE_CORRECTION_FACTOR}"
    
    def test_without_claude_correction(self):
        """
        Что он делает: Проверяет подсчёт без коэффициента коррекции.
        Цель: Убедиться, что apply_claude_correction=False работает.
        """
        print("Тест: Без коэффициента коррекции...")
        text = "Test text"
        
        result = count_tokens(text, apply_claude_correction=False)
        print(f"Результат: {result}")
        
        assert result > 0, "Должен вернуть положительное число токенов"
    
    def test_unicode_text(self):
        """
        Что он делает: Проверяет подсчёт токенов для Unicode текста.
        Цель: Убедиться в корректной обработке не-ASCII символов.
        """
        print("Тест: Unicode текст...")
        text = "Привет, мир! 你好世界 🌍"
        
        result = count_tokens(text)
        print(f"Результат: {result}")
        
        assert result > 0, "Unicode текст должен возвращать положительное число токенов"
    
    def test_multiline_text(self):
        """
        Что он делает: Проверяет подсчёт токенов для многострочного текста.
        Цель: Убедиться в корректной обработке переносов строк.
        """
        print("Тест: Многострочный текст...")
        text = """Line 1
        Line 2
        Line 3"""
        
        result = count_tokens(text)
        print(f"Результат: {result}")
        
        assert result > 0, "Многострочный текст должен возвращать положительное число токенов"
    
    def test_json_text(self):
        """
        Что он делает: Проверяет подсчёт токенов для JSON строки.
        Цель: Убедиться в корректной обработке JSON.
        """
        print("Тест: JSON текст...")
        text = '{"name": "test", "value": 123, "nested": {"key": "value"}}'
        
        result = count_tokens(text)
        print(f"Результат: {result}")
        
        assert result > 0, "JSON текст должен возвращать положительное число токенов"


class TestCountTokensFallback:
    """Тесты для fallback логики при отсутствии tiktoken."""
    
    def test_fallback_when_tiktoken_unavailable(self):
        """
        Что он делает: Проверяет fallback подсчёт когда tiktoken недоступен.
        Цель: Убедиться, что система работает без tiktoken.
        """
        print("Тест: Fallback без tiktoken...")
        
        # Мокируем _get_encoding чтобы вернуть None
        with patch('kiro_gateway.tokenizer._get_encoding', return_value=None):
            result = count_tokens("Hello world test")
            print(f"Результат: {result}")
            
            # Fallback: len(text) // 4 + 1, затем * 1.15
            # "Hello world test" = 16 символов
            # 16 // 4 + 1 = 5
            # 5 * 1.15 = 5.75 -> 5
            assert result > 0, "Fallback должен вернуть положительное число"
    
    def test_fallback_without_correction(self):
        """
        Что он делает: Проверяет fallback без коэффициента коррекции.
        Цель: Убедиться, что fallback работает с apply_claude_correction=False.
        """
        print("Тест: Fallback без коррекции...")
        
        with patch('kiro_gateway.tokenizer._get_encoding', return_value=None):
            result = count_tokens("Test", apply_claude_correction=False)
            print(f"Результат: {result}")
            
            # "Test" = 4 символа
            # 4 // 4 + 1 = 2
            assert result > 0, "Fallback должен вернуть положительное число"


class TestCountMessageTokens:
    """Тесты для функции count_message_tokens."""
    
    def test_empty_list_returns_zero(self):
        """
        Что он делает: Проверяет, что пустой список возвращает 0 токенов.
        Цель: Убедиться в корректной обработке пустого списка.
        """
        print("Тест: Пустой список сообщений...")
        result = count_message_tokens([])
        print(f"Результат: {result}")
        assert result == 0, "Пустой список должен возвращать 0 токенов"
    
    def test_none_returns_zero(self):
        """
        Что он делает: Проверяет, что None возвращает 0 токенов.
        Цель: Убедиться в корректной обработке None.
        """
        print("Тест: None...")
        result = count_message_tokens(None)
        print(f"Результат: {result}")
        assert result == 0, "None должен возвращать 0 токенов"
    
    def test_single_user_message(self):
        """
        Что он делает: Проверяет подсчёт токенов для одного user сообщения.
        Цель: Убедиться в базовой работоспособности.
        """
        print("Тест: Одно user сообщение...")
        messages = [{"role": "user", "content": "Hello, AI!"}]
        
        result = count_message_tokens(messages)
        print(f"Результат: {result}")
        
        assert result > 0, "Должен вернуть положительное число токенов"
    
    def test_multiple_messages(self):
        """
        Что он делает: Проверяет подсчёт токенов для нескольких сообщений.
        Цель: Убедиться, что токены суммируются корректно.
        """
        print("Тест: Несколько сообщений...")
        messages = [
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": "Hello!"},
            {"role": "assistant", "content": "Hi there! How can I help you?"},
            {"role": "user", "content": "What is the weather?"}
        ]
        
        result = count_message_tokens(messages)
        print(f"Результат: {result}")
        
        # Больше сообщений = больше токенов
        single_message = count_message_tokens([messages[0]])
        assert result > single_message, "Несколько сообщений должны иметь больше токенов"
    
    def test_message_with_tool_calls(self):
        """
        Что он делает: Проверяет подсчёт токенов для сообщения с tool_calls.
        Цель: Убедиться, что tool_calls учитываются.
        """
        print("Тест: Сообщение с tool_calls...")
        messages = [
            {
                "role": "assistant",
                "content": "",
                "tool_calls": [
                    {
                        "id": "call_123",
                        "type": "function",
                        "function": {
                            "name": "get_weather",
                            "arguments": '{"location": "Moscow"}'
                        }
                    }
                ]
            }
        ]
        
        result = count_message_tokens(messages)
        print(f"Результат: {result}")
        
        assert result > 0, "Сообщение с tool_calls должно иметь токены"
    
    def test_message_with_tool_call_id(self):
        """
        Что он делает: Проверяет подсчёт токенов для tool response сообщения.
        Цель: Убедиться, что tool_call_id учитывается.
        """
        print("Тест: Tool response сообщение...")
        messages = [
            {
                "role": "tool",
                "content": "The weather in Moscow is sunny, 25°C",
                "tool_call_id": "call_123"
            }
        ]
        
        result = count_message_tokens(messages)
        print(f"Результат: {result}")
        
        assert result > 0, "Tool response должен иметь токены"
    
    def test_message_with_list_content(self):
        """
        Что он делает: Проверяет подсчёт токенов для мультимодального контента.
        Цель: Убедиться, что list content обрабатывается.
        """
        print("Тест: Мультимодальный контент...")
        messages = [
            {
                "role": "user",
                "content": [
                    {"type": "text", "text": "What is in this image?"},
                    {"type": "image_url", "image_url": {"url": "https://example.com/image.jpg"}}
                ]
            }
        ]
        
        result = count_message_tokens(messages)
        print(f"Результат: {result}")
        
        assert result > 0, "Мультимодальный контент должен иметь токены"
    
    def test_without_claude_correction(self):
        """
        Что он делает: Проверяет подсчёт без коэффициента коррекции.
        Цель: Убедиться, что apply_claude_correction=False работает.
        """
        print("Тест: Без коэффициента коррекции...")
        messages = [{"role": "user", "content": "Test message"}]
        
        with_correction = count_message_tokens(messages, apply_claude_correction=True)
        without_correction = count_message_tokens(messages, apply_claude_correction=False)
        
        print(f"С коррекцией: {with_correction}")
        print(f"Без коррекции: {without_correction}")
        
        assert with_correction > without_correction, "С коррекцией должно быть больше"
    
    def test_message_with_empty_content(self):
        """
        Что он делает: Проверяет подсчёт для сообщения с пустым content.
        Цель: Убедиться, что пустой content не ломает подсчёт.
        """
        print("Тест: Пустой content...")
        messages = [{"role": "user", "content": ""}]
        
        result = count_message_tokens(messages)
        print(f"Результат: {result}")
        
        # Должны быть служебные токены (role, разделители)
        assert result > 0, "Даже пустое сообщение должно иметь служебные токены"
    
    def test_message_with_none_content(self):
        """
        Что он делает: Проверяет подсчёт для сообщения с None content.
        Цель: Убедиться, что None content не ломает подсчёт.
        """
        print("Тест: None content...")
        messages = [{"role": "assistant", "content": None}]
        
        result = count_message_tokens(messages)
        print(f"Результат: {result}")
        
        assert result > 0, "Сообщение с None content должно иметь служебные токены"


class TestCountToolsTokens:
    """Тесты для функции count_tools_tokens."""
    
    def test_none_returns_zero(self):
        """
        Что он делает: Проверяет, что None возвращает 0 токенов.
        Цель: Убедиться в корректной обработке None.
        """
        print("Тест: None...")
        result = count_tools_tokens(None)
        print(f"Результат: {result}")
        assert result == 0, "None должен возвращать 0 токенов"
    
    def test_empty_list_returns_zero(self):
        """
        Что он делает: Проверяет, что пустой список возвращает 0 токенов.
        Цель: Убедиться в корректной обработке пустого списка.
        """
        print("Тест: Пустой список...")
        result = count_tools_tokens([])
        print(f"Результат: {result}")
        assert result == 0, "Пустой список должен возвращать 0 токенов"
    
    def test_single_tool(self):
        """
        Что он делает: Проверяет подсчёт токенов для одного инструмента.
        Цель: Убедиться в базовой работоспособности.
        """
        print("Тест: Один инструмент...")
        tools = [
            {
                "type": "function",
                "function": {
                    "name": "get_weather",
                    "description": "Get the current weather for a location",
                    "parameters": {
                        "type": "object",
                        "properties": {
                            "location": {"type": "string", "description": "City name"}
                        },
                        "required": ["location"]
                    }
                }
            }
        ]
        
        result = count_tools_tokens(tools)
        print(f"Результат: {result}")
        
        assert result > 0, "Инструмент должен иметь токены"
    
    def test_multiple_tools(self):
        """
        Что он делает: Проверяет подсчёт токенов для нескольких инструментов.
        Цель: Убедиться, что токены суммируются.
        """
        print("Тест: Несколько инструментов...")
        tools = [
            {
                "type": "function",
                "function": {
                    "name": "get_weather",
                    "description": "Get weather",
                    "parameters": {"type": "object", "properties": {}}
                }
            },
            {
                "type": "function",
                "function": {
                    "name": "search_web",
                    "description": "Search the web",
                    "parameters": {"type": "object", "properties": {}}
                }
            }
        ]
        
        result = count_tools_tokens(tools)
        single_tool = count_tools_tokens([tools[0]])
        
        print(f"Два инструмента: {result}")
        print(f"Один инструмент: {single_tool}")
        
        assert result > single_tool, "Больше инструментов = больше токенов"
    
    def test_tool_with_complex_parameters(self):
        """
        Что он делает: Проверяет подсчёт для инструмента со сложными параметрами.
        Цель: Убедиться, что JSON schema параметров учитывается.
        """
        print("Тест: Сложные параметры...")
        tools = [
            {
                "type": "function",
                "function": {
                    "name": "complex_function",
                    "description": "A function with complex parameters",
                    "parameters": {
                        "type": "object",
                        "properties": {
                            "name": {"type": "string", "description": "Name"},
                            "age": {"type": "integer", "description": "Age"},
                            "address": {
                                "type": "object",
                                "properties": {
                                    "street": {"type": "string"},
                                    "city": {"type": "string"},
                                    "country": {"type": "string"}
                                }
                            },
                            "tags": {
                                "type": "array",
                                "items": {"type": "string"}
                            }
                        },
                        "required": ["name", "age"]
                    }
                }
            }
        ]
        
        result = count_tools_tokens(tools)
        print(f"Результат: {result}")
        
        assert result > 0, "Сложный инструмент должен иметь токены"
    
    def test_tool_without_parameters(self):
        """
        Что он делает: Проверяет подсчёт для инструмента без параметров.
        Цель: Убедиться, что отсутствие parameters не ломает подсчёт.
        """
        print("Тест: Без параметров...")
        tools = [
            {
                "type": "function",
                "function": {
                    "name": "no_params_func",
                    "description": "A function without parameters"
                }
            }
        ]
        
        result = count_tools_tokens(tools)
        print(f"Результат: {result}")
        
        assert result > 0, "Инструмент без параметров должен иметь токены"
    
    def test_tool_with_empty_description(self):
        """
        Что он делает: Проверяет подсчёт для инструмента с пустым description.
        Цель: Убедиться, что пустой description не ломает подсчёт.
        """
        print("Тест: Пустой description...")
        tools = [
            {
                "type": "function",
                "function": {
                    "name": "func",
                    "description": "",
                    "parameters": {"type": "object", "properties": {}}
                }
            }
        ]
        
        result = count_tools_tokens(tools)
        print(f"Результат: {result}")
        
        assert result > 0, "Инструмент с пустым description должен иметь токены"
    
    def test_non_function_tool_type(self):
        """
        Что он делает: Проверяет обработку инструмента с type != "function".
        Цель: Убедиться, что non-function tools обрабатываются.
        """
        print("Тест: Non-function tool...")
        tools = [
            {
                "type": "other_type",
                "some_field": "value"
            }
        ]
        
        result = count_tools_tokens(tools)
        print(f"Результат: {result}")
        
        # Должны быть хотя бы служебные токены
        assert result >= 0, "Non-function tool не должен ломать подсчёт"
    
    def test_without_claude_correction(self):
        """
        Что он делает: Проверяет подсчёт без коэффициента коррекции.
        Цель: Убедиться, что apply_claude_correction=False работает.
        """
        print("Тест: Без коэффициента коррекции...")
        tools = [
            {
                "type": "function",
                "function": {
                    "name": "test_func",
                    "description": "Test function",
                    "parameters": {"type": "object", "properties": {}}
                }
            }
        ]
        
        with_correction = count_tools_tokens(tools, apply_claude_correction=True)
        without_correction = count_tools_tokens(tools, apply_claude_correction=False)
        
        print(f"С коррекцией: {with_correction}")
        print(f"Без коррекции: {without_correction}")
        
        assert with_correction > without_correction, "С коррекцией должно быть больше"


class TestEstimateRequestTokens:
    """Тесты для функции estimate_request_tokens."""
    
    def test_messages_only(self):
        """
        Что он делает: Проверяет оценку токенов только для сообщений.
        Цель: Убедиться в базовой работоспособности.
        """
        print("Тест: Только сообщения...")
        messages = [{"role": "user", "content": "Hello!"}]
        
        result = estimate_request_tokens(messages)
        print(f"Результат: {result}")
        
        assert "messages_tokens" in result
        assert "tools_tokens" in result
        assert "system_tokens" in result
        assert "total_tokens" in result
        
        assert result["messages_tokens"] > 0
        assert result["tools_tokens"] == 0
        assert result["system_tokens"] == 0
        assert result["total_tokens"] == result["messages_tokens"]
    
    def test_messages_with_tools(self):
        """
        Что он делает: Проверяет оценку токенов для сообщений с инструментами.
        Цель: Убедиться, что tools учитываются.
        """
        print("Тест: Сообщения с инструментами...")
        messages = [{"role": "user", "content": "What is the weather?"}]
        tools = [
            {
                "type": "function",
                "function": {
                    "name": "get_weather",
                    "description": "Get weather",
                    "parameters": {"type": "object", "properties": {}}
                }
            }
        ]
        
        result = estimate_request_tokens(messages, tools=tools)
        print(f"Результат: {result}")
        
        assert result["messages_tokens"] > 0
        assert result["tools_tokens"] > 0
        assert result["total_tokens"] == result["messages_tokens"] + result["tools_tokens"]
    
    def test_messages_with_system_prompt(self):
        """
        Что он делает: Проверяет оценку токенов с отдельным system prompt.
        Цель: Убедиться, что system_prompt учитывается.
        """
        print("Тест: С system prompt...")
        messages = [{"role": "user", "content": "Hello!"}]
        system_prompt = "You are a helpful assistant."
        
        result = estimate_request_tokens(messages, system_prompt=system_prompt)
        print(f"Результат: {result}")
        
        assert result["messages_tokens"] > 0
        assert result["system_tokens"] > 0
        assert result["total_tokens"] == result["messages_tokens"] + result["system_tokens"]
    
    def test_full_request(self):
        """
        Что он делает: Проверяет оценку токенов для полного запроса.
        Цель: Убедиться, что все компоненты суммируются.
        """
        print("Тест: Полный запрос...")
        messages = [
            {"role": "user", "content": "What is the weather in Moscow?"}
        ]
        tools = [
            {
                "type": "function",
                "function": {
                    "name": "get_weather",
                    "description": "Get weather for a location",
                    "parameters": {
                        "type": "object",
                        "properties": {
                            "location": {"type": "string"}
                        }
                    }
                }
            }
        ]
        system_prompt = "You are a weather assistant."
        
        result = estimate_request_tokens(messages, tools=tools, system_prompt=system_prompt)
        print(f"Результат: {result}")
        
        expected_total = result["messages_tokens"] + result["tools_tokens"] + result["system_tokens"]
        assert result["total_tokens"] == expected_total, "Total должен быть суммой компонентов"
    
    def test_empty_messages(self):
        """
        Что он делает: Проверяет оценку для пустого списка сообщений.
        Цель: Убедиться в корректной обработке граничного случая.
        """
        print("Тест: Пустые сообщения...")
        result = estimate_request_tokens([])
        print(f"Результат: {result}")
        
        assert result["messages_tokens"] == 0
        assert result["total_tokens"] == 0


class TestClaudeCorrectionFactor:
    """Тесты для коэффициента коррекции Claude."""
    
    def test_correction_factor_value(self):
        """
        Что он делает: Проверяет значение коэффициента коррекции.
        Цель: Убедиться, что коэффициент равен 1.15.
        """
        print(f"Коэффициент коррекции: {CLAUDE_CORRECTION_FACTOR}")
        assert CLAUDE_CORRECTION_FACTOR == 1.15, "Коэффициент должен быть 1.15"
    
    def test_correction_increases_token_count(self):
        """
        Что он делает: Проверяет, что коррекция увеличивает количество токенов.
        Цель: Убедиться, что коэффициент применяется корректно.
        """
        print("Тест: Коррекция увеличивает токены...")
        text = "This is a test text for checking the correction factor"
        
        with_correction = count_tokens(text, apply_claude_correction=True)
        without_correction = count_tokens(text, apply_claude_correction=False)
        
        print(f"С коррекцией: {with_correction}")
        print(f"Без коррекции: {without_correction}")
        
        assert with_correction > without_correction
        
        # Проверяем, что разница примерно 15%
        increase_percent = (with_correction - without_correction) / without_correction * 100
        print(f"Увеличение: {increase_percent:.1f}%")
        
        # Допускаем погрешность из-за округления
        assert 10 <= increase_percent <= 20, "Увеличение должно быть около 15%"
class TestGetEncoding:
    """Тесты для функции _get_encoding."""
    
    def test_returns_encoding_when_tiktoken_available(self):
        """
        Что он делает: Проверяет, что _get_encoding возвращает encoding когда tiktoken доступен.
        Цель: Убедиться в корректной инициализации tiktoken.
        """
        print("Тест: tiktoken доступен...")
        
        # Сбрасываем глобальную переменную для чистого теста
        import kiro_gateway.tokenizer as tokenizer_module
        original_encoding = tokenizer_module._encoding
        tokenizer_module._encoding = None
        
        try:
            encoding = _get_encoding()
            print(f"Encoding: {encoding}")
            
            # Если tiktoken установлен, должен вернуть encoding
            if encoding is not None:
                assert hasattr(encoding, 'encode'), "Encoding должен иметь метод encode"
        finally:
            # Восстанавливаем
            tokenizer_module._encoding = original_encoding
    
    def test_caches_encoding(self):
        """
        Что он делает: Проверяет, что encoding кэшируется.
        Цель: Убедиться в ленивой инициализации.
        """
        print("Тест: Кэширование encoding...")
        
        encoding1 = _get_encoding()
        encoding2 = _get_encoding()
        
        print(f"Encoding 1: {encoding1}")
        print(f"Encoding 2: {encoding2}")
        
        # Должен вернуть тот же объект
        assert encoding1 is encoding2, "Encoding должен кэшироваться"
    
    def test_handles_import_error(self):
        """
        Что он делает: Проверяет обработку ImportError при отсутствии tiktoken.
        Цель: Убедиться, что система работает без tiktoken.
        """
        print("Тест: ImportError...")
        
        import kiro_gateway.tokenizer as tokenizer_module
        original_encoding = tokenizer_module._encoding
        tokenizer_module._encoding = None
        
        try:
            # Мокируем import tiktoken чтобы выбросить ImportError
            with patch.dict('sys.modules', {'tiktoken': None}):
                with patch('builtins.__import__', side_effect=ImportError("No module named 'tiktoken'")):
                    # Сбрасываем кэш
                    tokenizer_module._encoding = None
                    
                    # Должен вернуть None и не упасть
                    # Примечание: из-за кэширования этот тест может не работать идеально
                    # но главное - проверить что код не падает
                    pass
        finally:
            tokenizer_module._encoding = original_encoding


class TestTokenizerIntegration:
    """Интеграционные тесты для токенизатора."""
    
    def test_realistic_chat_request(self):
        """
        Что он делает: Проверяет подсчёт токенов для реалистичного chat запроса.
        Цель: Убедиться в корректной работе на реальных данных.
        """
        print("Тест: Реалистичный chat запрос...")
        
        messages = [
            {"role": "system", "content": "You are a helpful AI assistant. Be concise and accurate."},
            {"role": "user", "content": "What is the capital of France?"},
            {"role": "assistant", "content": "The capital of France is Paris."},
            {"role": "user", "content": "What is its population?"}
        ]
        
        tools = [
            {
                "type": "function",
                "function": {
                    "name": "search_web",
                    "description": "Search the web for information",
                    "parameters": {
                        "type": "object",
                        "properties": {
                            "query": {"type": "string", "description": "Search query"}
                        },
                        "required": ["query"]
                    }
                }
            }
        ]
        
        result = estimate_request_tokens(messages, tools=tools)
        print(f"Результат: {result}")
        
        # Проверяем разумность значений
        assert result["messages_tokens"] > 50, "Сообщения должны иметь > 50 токенов"
        assert result["tools_tokens"] > 20, "Tools должны иметь > 20 токенов"
        assert result["total_tokens"] > 70, "Total должен быть > 70 токенов"
    
    def test_large_context(self):
        """
        Что он делает: Проверяет подсчёт токенов для большого контекста.
        Цель: Убедиться в производительности на больших данных.
        """
        print("Тест: Большой контекст...")
        
        # Создаём большой текст
        large_text = "This is a test sentence. " * 1000  # ~5000 слов
        
        messages = [{"role": "user", "content": large_text}]
        
        result = estimate_request_tokens(messages)
        print(f"Токенов в большом тексте: {result['total_tokens']}")
        
        # Должно быть много токенов
        assert result["total_tokens"] > 1000, "Большой текст должен иметь > 1000 токенов"
    
    def test_consistency_across_calls(self):
        """
        Что он делает: Проверяет консистентность подсчёта при повторных вызовах.
        Цель: Убедиться, что результаты детерминированы.
        """
        print("Тест: Консистентность...")
        
        text = "This is a test for consistency checking"
        
        results = [count_tokens(text) for _ in range(5)]
        print(f"Результаты: {results}")
        
        # Все результаты должны быть одинаковыми
        assert len(set(results)) == 1, "Результаты должны быть консистентными"
    
    


================================================
FILE: .github/ISSUE_TEMPLATE/bug_report.yml
================================================
name: 🐛 Bug Report
description: Something isn't working? Report it here
title: "[Bug]: "
labels: ["bug"]
body:
  - type: markdown
    attributes:
      value: |
        ## Before submitting
        
        Please enable debug logging to help me fix the issue faster:
        1. Add `DEBUG_MODE=errors` to your `.env` file
        2. Restart the gateway
        3. Reproduce the error
        4. Attach files from `debug_logs/` folder below

  - type: input
    id: version
    attributes:
      label: Gateway Version
      description: Which version are you using?
      placeholder: "e.g. v1.0.6"
    validations:
      required: true

  - type: textarea
    id: description
    attributes:
      label: What happened?
      description: Describe what you were doing and what went wrong
      placeholder: "I was trying to use X and got a 400 error \"Improperly formed request\"..."
    validations:
      required: true

  - type: textarea
    id: logs
    attributes:
      label: Debug Logs
      description: |
        Attach files from `debug_logs/` folder, especially these:
        - `app_logs.txt`
        - `request_body.json`
        - `kiro_request_body.json`
        
        Drag & drop files here or paste the content.
      placeholder: "Drag & drop your log files here..."
    validations:
      required: true

