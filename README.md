# User Registration + OTP via SMS — App Scaffolding

A clean, minimal starter for **user registration with OTP (one-time password) delivered via SMS**.
This scaffold includes configuration loading (YAML + env), structured logging (zap), pluggable SMS notifier adapters (Kavenegar and SMS.ir), Redis-backed OTP storage, an in-memory user repository, validators, and the main wiring to run a registration flow.

---

## Highlights / Features

* Config loader using **koanf** (YAML + ENV precedence, prefix/divider handling).

* Structured logging with **zap** (dev / prod configs).

* SMS notifier abstraction + provider adapters:
  
  * **Kavenegar** adapter implemented.
  * **SMS.ir** adapter stubbed (easy to finish).

* Redis-backed OTP repository (stores OTP with TTL via `Set(key,value,ttl)`).

* In-memory user repository for quick bootstrapping.

* Validation for registration inputs (name/family/phone).

* Fluent response builder (responser) and organized request/response param types.

* Simple main wiring that demonstrates the registration flow.

---

| Cohort / File(s)                                                                                                    | Summary                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| **Project environment & ignores**<br>`.env`, `.gitignore`, `.idea/.gitignore`                                       | Adds env placeholders for Kavenegar/Redis; updates gitignore to ignore `*.xml`, `*.iml`; IDE-specific ignore rules.    |
| **Module & dependencies**<br>`go.mod`                                                                               | Declares module, Go version, and dependencies (zap, koanf, redis, kavenegar, etc.).                                    |
| **App configuration**<br>`config/app.yaml`                                                                          | Adds `sms-provider: "kavenegar"`.                                                                                      |
| **Entity & params**<br>`entity/user.go`, `param/register_param.go`                                                  | Adds User entity and RegisterParam DTO.                                                                                |
| **Config loader**<br>`pkg/configloader/config.go`                                                                   | Adds fluent config loader using koanf for YAML and env with prefix/divider handling.                                   |
| **Logger**<br>`pkg/logger/config.yaml`, `pkg/logger/logger.go`                                                      | Adds zap configs for dev/prod and a builder that selects mode and returns a logger.                                    |
| **Responder**<br>`pkg/responser/responser.go`                                                                       | Adds Response struct with fluent setters and `Build`.                                                                  |
| **Validation**<br>`validator/user_validator.go`                                                                     | Adds user registration validator (name/family/phone checks).                                                           |
| **SMS notifier (core)**<br>`pkg/notifier/sms/sms_provider.go`, `sms_message.go`, `sms_templates.go`                 | Defines SMSProvider interface, message builder, and a Register template.                                               |
| **SMS adapters & factory**<br>`pkg/notifier/sms/notifier.go`, `kavenegar_adapter.go`, `smsir_adapter.go`            | Adds notifier selecting provider via config; Kavenegar adapter implemented; SMS.ir stubbed.                            |
| **Repositories (users & OTP)**<br>`repository/userrepo/*`, `repository/otprepo/*`, `repository/redis_connection.go` | Adds in-memory user repo; Redis connection factory; Redis OTP repo with `Set(key,value,ttl)`.                          |
| **App service**<br>`service/appservice/app.go`                                                                      | Adds AppService to read sms-provider and API key from config.                                                          |
| **User service**<br>`service/userservice/user_service.go`                                                           | Implements registration flow: validate, user lookup, OTP generate/store (Redis), SMS send, logging, response building. |
| **Main wiring**<br>`main.go`                                                                                        | Wires validator, repos, responser, notifier, user service; invokes Register.                                           |

---

## Quick intro — how the registration flow works

1. **Input**: a `RegisterParam` DTO is provided (name, family, phone).
2. **Validate**: `user_validator` validates inputs (format, required fields).
3. **Lookup**: user repository checks whether the user exists (in-memory here).
4. **OTP generation**: a numeric OTP is generated by the service.
5. **Store OTP**: OTP is stored in Redis with TTL via the Redis OTP repo (`Set(key,value,ttl)`).
6. **Send SMS**: SMS is sent using the configured provider (Kavenegar or SMS.ir) via the `SMSProvider` interface.
7. **Respond**: A structured `Response` is built and returned. Logs are written using zap.

---

## Prerequisites

* Go (version declared in `go.mod`)
* Redis (for OTP storage)
* API key(s) for the SMS provider you plan to use (e.g. Kavenegar or SMS.ir)

---

## Setup & Run (quick)

```bash
# 1. clone
git clone <repo-url>
cd <repo>

# 2. populate env (file names from scaffold)
cp .env.example .env      # edit values inside .env
# edit config/app.yaml to set sms-provider: "kavenegar" or "smsir"

# 3. fetch deps
go mod download

# 4. run
go run main.go
```

> `main.go` wires validators, repositories, the responder, notifier and the user service and demonstrates the Register flow.

---

## Example config snippets

### `config/app.yaml` (example)

```yaml
# primary app configuration
sms-provider: "kavenegar"   # other option: "smsir"
# other app-level options can be added here
```

### `.env` (example)

```env
# Redis
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# Kavenegar (if using Kavenegar)
KAVENEGAR_API_KEY=your_kavenegar_api_key_here

# SMS.ir (if using SMS.ir)
SMSIR_API_KEY=your_smsir_api_key_here

# App environment
APP_ENV=dev  # used by logger builder to pick dev/prod
```

---

## Project structure (important files added)

```
.
├─ .env, .gitignore, .idea/...
├─ go.mod
├─ config/
│  └─ app.yaml
├─ entity/
│  └─ user.go                    # User entity
├─ param/
│  └─ register_param.go          # RegisterParam DTO
├─ pkg/
│  ├─ configloader/config.go     # koanf-based config loader
│  ├─ logger/
│  │  ├─ config.yaml             # zap dev/prod configs
│  │  └─ logger.go               # zap builder
│  ├─ responser/
│  │  └─ responser.go            # Response builder
│  └─ notifier/
│     └─ sms/
│         ├─ sms_provider.go     # SMSProvider interface
│         ├─ sms_message.go      # SMS message builder
│         ├─ sms_templates.go    # Register template
│         ├─ notifier.go         # factory / provider selector
│         ├─ kavenegar_adapter.go# implemented adapter
│         └─ smsir_adapter.go    # stubbed adapter
├─ repository/
│  ├─ userrepo/                  # in-memory user repo
│  ├─ otprepo/                   # Redis-based OTP repo
│  └─ redis_connection.go        # Redis connection factory
├─ service/
│  ├─ appservice/app.go          # reads provider/config keys
│  └─ userservice/user_service.go# registration flow
└─ main.go
```

---

## How to change SMS provider

1. Edit `config/app.yaml` and set:

```yaml
sms-provider: "smsir"   # or "kavenegar"
```

2. Make sure to set the required API key in `.env` for the provider selected (e.g. `SMSIR_API_KEY` or `KAVENEGAR_API_KEY`).
3. Restart the app.

The notifier factory (`pkg/notifier/sms/notifier.go`) picks the adapter based on `sms-provider`.
