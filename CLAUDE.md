# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Start the database
docker compose -f ./scripts/docker-compose.yml -p budget-app up -d

# Run migrations (requires .env to be loaded)
go run ./cmd/migrations up
go run ./cmd/migrations down        # one step back
go run ./cmd/migrations down all    # revert all
go run ./cmd/migrations version     # check current version

# Start the API (listens on :8080)
go run ./cmd/api

# Run all tests
go test ./...

# Run a single test
go test ./service/... -run TestExpenses
go test ./service/... -run TestGetUserByID
```

Required `.env` (copy values from `.env`, which has working local defaults):
```
DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, SSL_MODE, ALLOWED_ORIGINS
```

## Architecture

Request flow:
```
HTTP Request → Echo Router → auth.Middleware (JWT) → HandlerWebRoute (permission check) → Handler func → Service func → sql.DB
```

### Layer responsibilities

**`contexts/`** — the central glue layer. `Context` is passed to every handler and carries lazy factory functions for `Database`, `API` (response builder), and `Logs`. `NewContext()` is called once in `router.Init`. Handlers receive a `*contexts.Context` (not Echo's context directly).

**Route registration pattern** — Routes are declared as package-level `*contexts.WebRoute` vars in `handler/`. Each `WebRoute` holds the path, method, handler func, and `PermissionLevel`. `router.Init` collects them into `noAuthRoutes` / `authRoutes` slices and registers them via `contexts.HandlerWebRoute`, which wraps each handler with permission checking before it reaches the service layer.

**`contexts.WebRoute` vs `contexts.EchoHandler`** — `WebRoute` is the declarative definition (written in `handler/`). `HandlerWebRoute()` converts it to an `EchoHandler` (what Echo actually registers). `Routes` in `router/router.go` is a type alias for `*EchoHandler`.

**Permission levels** — defined as `AdminLevel`, `MediumLevel`, `BasicLevel` in `contexts/web-route.go`. Each route declares which levels are allowed. Permission is checked inside `HandlerWebRoute` against `ctx.API().Session().UserLevel` — note this is currently a stub returning an empty string; real session data comes from the JWT claims stored in the request context via `auth/context.go`.

**`service/`** — business logic only; no Echo/HTTP imports. All service functions accept a `DB` interface (`service/services_repository.go`) so they can be tested with `go-sqlmock` without a real database. SQL queries are in separate `*_query.go` files.

**`auth/`** — `Middleware` validates the Bearer JWT and injects `user_id` into the request context. `ConfigCORS()` reads `ALLOWED_ORIGINS` from env. The `jwtSecret` is currently hardcoded in both `auth/middleware.go` and `handler/login.go` — these need to be aligned and moved to env.

**`common/`** — `GenerateCustomGuideID()` produces a 15-char lowercase random string used as all primary keys (no UUIDs). `helper.go` has `GetEnv`/`GetEnvArray`.

**`db/`** — `Database.Connect()` opens and pings a new `*sql.DB` on every call; callers must `defer db.Close()`. Write operations use `db.Begin()` / `tx.Commit()` / `tx.Rollback()` in the handler layer.

## Key known issues

- `jwtSecret` is duplicated and hardcoded — `auth/middleware.go` uses `"sua-chave-secreta-super-segura"` and `handler/login.go` uses `"secure-secret-key"`. These must match or auth will always fail after login.
- `APIClient.Session()` returns a stub `UserLevel: ""` — permission checks for routes that require a level above anonymous will always fail until this is wired to real JWT claims.
- `Database.Connect()` opens a new connection per request; no connection pooling is configured.
