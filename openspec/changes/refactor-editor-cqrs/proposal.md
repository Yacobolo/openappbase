# Change: Refactor Editor to CQRS with NATS KV State Management

## Why
The current editor implementation mixes query and command operations, lacks reactive state management, and doesn't leverage the existing NATS JetStream infrastructure for session state. This refactoring aligns the editor with CQRS principles and provides a foundation for real-time collaborative editing.

The key architectural insight is to separate concerns:
- **Postgres** = Source of truth for actual table data
- **NATS KV** = Ephemeral session state (current page, selected table, filters) with auto-cleanup
- **SSE** = Push updates to browser when state changes

This approach ensures UI state doesn't pollute the database while providing reactive, real-time updates through Datastar.

## What Changes
- Separate query routes (`/editor`) from command routes (`/api/editor/*`)
- Introduce `EditorState` struct to track session-specific UI state (table name, current page, filters)
- Use NATS JetStream KV bucket for storing and watching editor session state
- Add SSE endpoint (`/editor/sse`) that watches NATS KV for state changes and re-renders components
- Refactor command handlers to update state in NATS KV (triggering SSE re-renders)
- Implement CQRS flow: User Action → Command API → Update NATS State → SSE Watches → Re-render Component
- Add `StateStore[EditorState]` integration for session management
- Add watcher functionality to `EditorService` for real-time state updates

## Impact
- Affected specs: `editor-state-management` (new capability)
- Affected code:
  - `internal/features/editor/handlers.go` - split into query and command handlers, add SSE handler
  - `internal/features/editor/routes.go` - separate query and command routes
  - `internal/features/editor/services/editor_service.go` - add state management methods and watcher
  - `internal/features/editor/pages/editor.templ` - simplify to shell page
  - `internal/features/editor/components/` - may need error/empty state components
  - `internal/session/store.go` - potentially add Watch method for generic state watching
- Dependencies:
  - Builds on existing `add-dynamic-table-editor` change
  - Requires NATS JetStream KV (already available in project)
  - Uses existing `StateStore[T]` pattern from `internal/session/store.go`
