## 1. Foundation Setup
- [ ] 1.1 Add `EditorState` struct to `internal/features/editor/services/editor_service.go`
- [ ] 1.2 Add `Watch(ctx, sessionID)` method to `internal/session/store.go` for generic state watching
- [ ] 1.3 Create NATS KV bucket "editor_sessions" in application initialization (likely `cmd/web/main.go`)
- [ ] 1.4 Add helper components: `EmptyTableView()` and `ErrorView(message)` in `internal/features/editor/components/`

## 2. Service Layer Enhancement
- [ ] 2.1 Add `GetSessionState(w, r)` method to EditorService that returns (sessionID, *EditorState, error)
- [ ] 2.2 Add `SaveState(ctx, sessionID, state)` method to EditorService
- [ ] 2.3 Add `WatchUpdates(ctx, sessionID)` method to EditorService that returns state change channel
- [ ] 2.4 Integrate `StateStore[EditorState]` in EditorService initialization

## 3. Query Routes Implementation
- [ ] 3.1 Refactor `EditorPage` handler to render minimal shell page (remove data loading logic)
- [ ] 3.2 Add `EditorSSE` handler that:
  - [ ] 3.2.1 Establishes SSE connection
  - [ ] 3.2.2 Gets session state from NATS KV
  - [ ] 3.2.3 Watches for state changes
  - [ ] 3.2.4 Renders initial view based on state
  - [ ] 3.2.5 Re-renders on state changes
- [ ] 3.3 Add `renderTableFromState(ctx, sse, state)` helper method that:
  - [ ] 3.3.1 Handles empty state (no table selected)
  - [ ] 3.3.2 Handles invalid table names (error view)
  - [ ] 3.3.3 Queries Postgres and renders DataTable component
  - [ ] 3.3.4 Pushes updates via SSE with correct selector

## 4. Command Routes Implementation
- [ ] 4.1 Add `LoadTable` handler (POST /api/editor/load):
  - [ ] 4.1.1 Validate table parameter
  - [ ] 4.1.2 Get session state
  - [ ] 4.1.3 Update state.TableName and reset state.Page to 1
  - [ ] 4.1.4 Save state to NATS KV
  - [ ] 4.1.5 Return HTTP 200 OK (no rendering)
- [ ] 4.2 Add `ChangePage` handler (POST /api/editor/page/{page}):
  - [ ] 4.2.1 Parse and validate page parameter
  - [ ] 4.2.2 Get session state
  - [ ] 4.2.3 Update state.Page
  - [ ] 4.2.4 Save state to NATS KV
  - [ ] 4.2.5 Return HTTP 200 OK
- [ ] 4.3 Remove or comment out old `LoadTableData` and `ChangePage` query handlers

## 5. Routes Configuration
- [ ] 5.1 Update `internal/features/editor/routes.go`:
  - [ ] 5.1.1 Keep GET /editor for page load
  - [ ] 5.1.2 Add GET /editor/sse for SSE connection
  - [ ] 5.1.3 Add POST /api/editor/load for loading tables
  - [ ] 5.1.4 Add POST /api/editor/page/{page} for pagination
  - [ ] 5.1.5 Remove old /editor/load and /editor/change-page query routes

## 6. Template Updates
- [ ] 6.1 Update `internal/features/editor/pages/editor.templ` to render shell with SSE connection
- [ ] 6.2 Add `EmptyTableView` component in `internal/features/editor/components/`
- [ ] 6.3 Add `ErrorView` component in `internal/features/editor/components/`
- [ ] 6.4 Update DataTable component to use command API endpoints (data-post attributes)
- [ ] 6.5 Regenerate templ files: `templ generate`

## 7. Initialization & Wiring
- [ ] 7.1 Initialize "editor_sessions" KV bucket on app startup
- [ ] 7.2 Pass StateStore[EditorState] to editor handlers constructor
- [ ] 7.3 Verify session store integration in routes setup

## 8. Testing & Validation
- [ ] 8.1 Test initial page load with empty state
- [ ] 8.2 Test loading a table via command API
- [ ] 8.3 Test page navigation via command API
- [ ] 8.4 Test SSE reconnection after network interruption
- [ ] 8.5 Test multiple concurrent users (session isolation)
- [ ] 8.6 Test invalid table names (error handling)
- [ ] 8.7 Test NATS KV TTL cleanup (verify after 24h or manually expire)
- [ ] 8.8 Performance test: measure SSE latency on state changes

## 9. Documentation & Cleanup
- [ ] 9.1 Add code comments explaining CQRS flow
- [ ] 9.2 Update any feature documentation if needed
- [ ] 9.3 Remove commented-out old code after validation
- [ ] 9.4 Run `go mod tidy` to ensure clean dependencies

## Dependencies
- Requires: `add-dynamic-table-editor` change to be completed (provides base editor service and data loading)
- Requires: NATS JetStream enabled (already in project)
- Requires: Datastar integration (already in project)

## Validation Criteria
- SSE connection established on page load
- Command API calls trigger SSE updates automatically
- No direct data queries in command handlers
- Each user session maintains independent state
- NATS KV bucket visible with `nats kv ls` (if NATS CLI available)
- State changes reflect in browser within <200ms
