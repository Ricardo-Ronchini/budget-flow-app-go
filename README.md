# Budget Flow API

> A RESTful API for personal expense management, built in Go with JWT authentication, role-based access control, and versioned database migrations.

## Overview

Budget Flow API is a backend service for tracking and managing personal expenses. It exposes a set of authenticated endpoints that allow users to create, read, update, and delete expense records tied to their accounts.

The project was built as a learning exercise to explore advanced Go patterns — specifically around dependency injection, interface-based testing, and layered HTTP architecture — while producing something close enough to a production-grade service to serve as a portfolio reference.

## Architecture

```
Request → Echo Router → auth.Middleware (JWT) → HandlerWebRoute (permission check) → Handler func → Service func → sql.DB
```

| Layer | Package | Responsibility |
|-------|---------|----------------|
| Entry points | `cmd/` | API server and migration CLI binaries |
| HTTP | `handler/` | Declares routes as package-level vars; binds request data |
| Business logic | `service/` | All domain logic; no HTTP imports |
| Database | `db/` | Connection factory and versioned SQL migrations |
| Authentication | `auth/` | JWT middleware and CORS configuration |
| Routing | `router/` | Echo setup and route registration |
| Shared context | `contexts/` | Dependency container, response builder, and permission system |

### Declarative route pattern

Routes are defined in `handler/` as `*contexts.WebRoute` structs — each one declares its HTTP method, path, handler function, and the permission levels allowed to call it. `contexts.HandlerWebRoute()` wraps each route with a permission check before registering it with Echo. This keeps access control co-located with the route definition rather than scattered across middleware.

### Context as a dependency container

`contexts.Context` is initialized once in `router.Init` and passed to every handler. It carries factory functions for the database connection, the structured logger, and the API response builder. Handlers receive this context and return `(int, any)` — they never import Echo directly.

## Permission System

The API defines three access levels: `AdminLevel`, `MediumLevel`, and `BasicLevel`. Each route declares which levels are permitted in its `PermissionLevel` field. `HandlerWebRoute` enforces this before the handler runs. Routes that don't require authentication set `Authenticate: false` and bypass the level check entirely.

## Tech Stack

| Technology | Role |
|-----------|------|
| **Go 1.23** | Primary language — strong standard library, fast compile times |
| **Echo v4** | HTTP framework — middleware chaining, route grouping, minimal overhead |
| **PostgreSQL** | Relational database — structured schema for financial records |
| **golang-migrate** | Versioned, reversible SQL migrations |
| **golang-jwt/jwt** | JWT generation and validation for stateless authentication |
| **go-sqlmock** | SQL mock for unit-testing service functions without a database |
| **logrus** | Structured logging with configurable levels |
| **godotenv** | `.env` file loading for local development |

## API Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|:---:|
| `GET` | `/test` | Health check | No |
| `POST` | `/v1/login` | Authenticate and receive a JWT token | No |
| `GET` | `/v1/api/user` | Get authenticated user data | Yes |
| `POST` | `/v1/api/user` | Create a new user | Yes |
| `GET` | `/v1/api/expenses` | List all expenses | Yes |
| `GET` | `/v1/api/expenses/:expense_id` | Get a single expense by ID | Yes |
| `POST` | `/v1/api/expenses` | Create a new expense | Yes |
| `PUT` | `/v1/api/expenses/:expense_id` | Update an expense | Yes |
| `DELETE` | `/v1/api/expenses/:expense_id` | Delete an expense | Yes |

Authenticated routes require a `Authorization: Bearer <token>` header. The token is obtained from `POST /v1/login`.

## Getting Started

### Prerequisites

- [Go 1.23+](https://go.dev/dl/)
- [Docker](https://www.docker.com/)

### Running locally

**1. Clone the repository**
```bash
git clone https://github.com/ricardo-ronchini/budget-flow-app-go.git
cd budget-flow-app-go
```

**2. Configure environment**
```bash
cp .env.example .env
# Edit .env with your database credentials
```

**3. Start the database**
```bash
docker compose -f ./scripts/docker-compose.yml -p budget-app up -d
```

**4. Run migrations**
```bash
go run ./cmd/migrations up
```

**5. Start the API**
```bash
go run ./cmd/api
```

The API will be available at `http://localhost:8080`. Confirm with:
```bash
curl http://localhost:8080/test
```

### Running tests

```bash
go test ./...
```

To run a specific test:
```bash
go test ./service/... -run TestExpenses
```

## Database

Two tables are created by the migrations:

**`users`** — stores user accounts.

| Column | Type | Notes |
|--------|------|-------|
| `user_id` | text | Primary key (custom 15-char random ID) |
| `name` | text | Required |
| `email` | text | Unique, required |
| `username` | text | Used for login |
| `password_hash` | text | bcrypt hash of the user's password |
| `created_at` | timestamp | |
| `updated_at` | timestamp | |

**`expenses`** — stores expense records, linked to a user.

| Column | Type | Notes |
|--------|------|-------|
| `expense_id` | text | Primary key (custom 15-char random ID) |
| `user_id` | text | Foreign reference to `users.user_id` |
| `name` | text | Expense label |
| `value` | numeric | Amount |
| `date` | timestamp | Date of the expense |
| `created_at` | timestamp | |
| `updated_at` | timestamp | |

## Project Highlights

**1. Declarative route + permission system**
Each route in `handler/` is a `*contexts.WebRoute` struct that declares its HTTP method, path, handler function, and allowed permission levels. `contexts.HandlerWebRoute()` enforces those levels before the handler runs — access control is defined at the route declaration site, not inside the handler or in a separate middleware file.

**2. Testable service layer via interface injection**
The `service` package defines a minimal `DB` interface with `Query`, `QueryRow`, and `Exec`. Service functions accept this interface instead of a concrete `*sql.DB`, which allows tests to use `go-sqlmock` to assert exact SQL queries and arguments without a running database.

**3. Single Context as dependency container**
`contexts.Context` is initialized once at startup and passed through the entire request lifecycle. It exposes factory functions for the database connection, structured logger, and API response builder — so handlers remain decoupled from infrastructure concerns and the dependency graph is explicit.

**4. JWT authentication with request-scoped user injection**
`auth.Middleware` validates the Bearer token, extracts `user_id` from the claims, and injects it into the request context via `context.WithValue`. Downstream handlers retrieve the user ID without re-parsing the token, keeping the authentication logic contained in a single place.

**5. Migration runner as a standalone CLI**
`cmd/migrations` is a separate binary that accepts `up`, `down`, `down all`, and `version` commands. Before applying any migration it checks for and resolves dirty state automatically, so interrupted migration runs don't leave the database schema in a broken state.
