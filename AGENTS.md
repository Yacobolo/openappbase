# Agent Guidelines for Northstar PostgreSQL UI

## Build & Test Commands

- **Dev server**: `task live` (hot reload) or `task setup && task live` (first time)
- **Production build**: `task build` (generates to `bin/main`)
- **Run production**: `task run` (builds first)
- **Database migrations**: `task db:migrate:up` | `task db:migrate:down` | `task db:migrate:status`
- **Create migration**: `task db:migrate:create -- migration_name`
- **Reset database**: `task db:reset` (WARNING: deletes all data)
- **Context search**: `task context:search -- "query"` (token-efficient search using OpenCode CLI subagent)
- **Tests**: No test files currently exist; use manual testing via the dev server

### Context Documentation Search

**IMPORTANT**: When searching context documentation, use the Taskfile command via bash instead of the `task` tool:

```bash
task context:search -- "your search query"
```

This runs the `context-searcher` subagent externally via OpenCode CLI, reducing token usage by ~97% compared to using the `task` tool directly. The subagent's exploration happens outside your token budget - you only receive the final summary.

**Examples:**

```bash
task context:search -- "find daisyui modal component documentation"
task context:search -- "show datastar SSE server examples"
task context:search -- "how to use NATS pub/sub patterns"
```

## Code Style & Conventions

### Project Structure

- **Backend**: Go 1.25+ with chi router, templ templates, SQLite (app state) + PostgreSQL (target databases)
- **Frontend**: Datastar (reactive SSE), Templ (Go templates), Tailwind v4 + DaisyUI, iconify-icon web components
- **Real-time**: Embedded NATS for pub/sub messaging between services
- **Database queries**: sqlc (see sqlc.yml) - never write raw SQL in Go, always use sqlc-generated queries

### Go Style

- **Imports**: Standard library first, then third-party, then local packages (northstar/...). Group with blank lines.
- **Error handling**: Always wrap errors with context: `fmt.Errorf("failed to X: %w", err)`. Use slog for logging: `slog.Error("message", "key", value, "error", err)`
- **Naming**: Follow Go conventions - PascalCase for exports, camelCase for private. Services should be named `*Service`, handlers `*Handlers`
- **Contexts**: Always use `ctx` from request context. For async callbacks (e.g., NATS), use `context.Background()` not request context
- **Null handling**: Use `sql.NullString`, `sql.NullInt64` for nullable DB fields. Helper: `sqlNullString(*string) sql.NullString`
- **Encryption**: Use AES-256-GCM for sensitive data. Key must be exactly 32 bytes (see connection_service.go)
- **Dependencies**: Handlers receive dependencies via constructor (NewHandlers), never use globals

### Frontend Style (Templ + Datastar)

- **Templates**: Use templ (_.templ files, generates _\_templ.go). Run `task build:templ` or use `task live` for watch mode
- **Styling**: Use DaisyUI semantic tokens (bg-base-100, btn-primary, text-base-content) NOT raw Tailwind colors. See DESIGN.md for full guide
- **Icons**: Use `<iconify-icon icon="mdi:name">` web component, NOT SVG sprites
- **SSE/Datastar**: Use `datastar.NewSSE(w, r)` for real-time updates. Patch fragments with `sse.PatchElementTempl(component)`
- **NATS topics**: Subscribe to topics like `admin.connections.update` in SSE handlers for real-time push

### Database Conventions

- **Migrations**: SQL files in `internal/store/migrations/`, use Goose. Multi-statement blocks need `-- +goose StatementBegin/End`
- **Queries**: Write in `internal/store/queries/*.sql`, generate with `sqlc generate` (auto-run by build tasks)
- **IDs**: Use int64 for all IDs (SQLite INTEGER PRIMARY KEY). Timestamps are TEXT in ISO8601 format
- **Encryption**: Store encrypted passwords in `encrypted_password` column, decrypt in service layer, never in queries

### Error Handling & Logging

- **HTTP errors**: Use `http.Error(w, msg, statusCode)` or return JSON errors with proper status codes
- **Logging**: Use structured logging with slog: `slog.Info("msg", "key", value)`. Levels: DEBUG, INFO, WARN, ERROR
- **Context**: Log errors with context before returning: `slog.Error("Failed to X", "id", id, "error", err)`

### Architecture Patterns

- **Features**: Organized by feature in `internal/features/*` with handlers.go, routes.go, services/, components/, pages/
- **Services**: Business logic lives in services (e.g., ConnectionService), handlers orchestrate HTTP + services
- **Real-time**: PostgreSQL LISTEN/NOTIFY → NATS → SSE to clients. Publish to NATS on data changes, subscribers re-query DB
- **Session state**: Stored in SQLite, not PostgreSQL target databases. Keep target databases clean

## Key Files to Reference

- `Taskfile.yml` - All build/dev commands
- `DESIGN.md` - Complete UI/styling guidelines (DaisyUI semantic tokens, component patterns)
- `context/README.md` - Overview of up-to-date reference docs for Datastar, NATS, sqlc, Templ with grep search examples
- `internal/features/admin/handlers.go` - Example SSE + NATS pattern
- `internal/features/admin/services/connection_service.go` - Example service with encryption
