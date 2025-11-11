# Design: CQRS Editor with NATS KV State Management

## Context
The PostgresUI editor needs to evolve from a simple query-response model to a reactive, event-driven architecture that can support real-time collaboration and complex UI state management. The project already has:
- Embedded NATS Server with JetStream enabled
- Generic `StateStore[T]` for session state management
- Datastar SSE for reactive updates
- Existing editor with pagination and table data loading

This design leverages these existing capabilities to implement a clean CQRS pattern.

### Datastar Documentation
For up-to-date information on Datastar SSE patterns, capabilities, and best practices, use the Context7 tool to fetch the latest documentation:
```
context7_resolve-library-id: "datastar"
context7_get-library-docs: Use the resolved library ID
```
Datastar provides the reactive SSE foundation for this CQRS implementation, enabling efficient DOM updates via `PatchElementTempl()` and other SSE events.

## Goals / Non-Goals

### Goals
- Implement CQRS pattern: separate query and command routes
- Use NATS KV for ephemeral UI state (not data)
- Enable reactive UI updates via SSE watching NATS KV
- Keep Postgres as the source of truth for table data
- Maintain session isolation (each user has their own state)
- Enable future real-time collaboration features
- Auto-cleanup stale sessions via NATS KV TTL

### Non-Goals
- Multi-user collaboration in this change (foundation only)
- Complex filtering or sorting (can be added later)
- Cell editing/updates (future capability)
- Optimistic UI updates (server-side rendering only for now)
- WebSocket support (SSE is sufficient)

## Decisions

### Decision 1: NATS KV for Session State
**Choice**: Store editor session state (table name, page, filters) in NATS JetStream KV bucket.

**Rationale**:
- Ephemeral UI state doesn't belong in Postgres
- NATS KV provides built-in TTL for auto-cleanup
- Enables real-time watching with minimal overhead
- Project already has embedded NATS running
- Consistent with existing `StateStore[T]` pattern used in explorer feature

**Alternatives considered**:
- Postgres LISTEN/NOTIFY: Limited to single instance, no persistence
- In-memory channels: Lost on restart, single instance only
- Redis: Requires additional infrastructure

### Decision 2: Query/Command Route Separation
**Choice**: 
- Query routes: `/editor` (page), `/editor/sse` (SSE stream)
- Command routes: `/api/editor/load`, `/api/editor/page/{page}`, etc.

**Rationale**:
- Clear CQRS separation improves code organization
- Commands use POST/PUT/DELETE (RESTful semantics)
- Queries use GET (cacheable, idempotent)
- API prefix makes it obvious which endpoints modify state
- Follows HTTP best practices

**Alternatives considered**:
- All routes under `/editor/*`: Less clear separation
- RPC-style single endpoint: Harder to reason about side effects

### Decision 3: SSE-Based State Watching
**Choice**: SSE endpoint watches NATS KV for state changes and pushes component updates via Datastar.

**Rationale**:
- Datastar already integrated for SSE-based updates
- Browser automatically reconnects on connection loss
- More efficient than polling
- Works across all browsers (unlike WebSocket)
- Simple mental model: state change → automatic re-render

**Flow**:
```
User clicks "Next Page"
  ↓
POST /api/editor/page/2
  ↓
Update EditorState.Page = 2
  ↓
Save to NATS KV (triggers watch notification)
  ↓
SSE handler receives watch event
  ↓
Query Postgres for page 2 data
  ↓
Render DataTable component
  ↓
Push to browser via SSE PatchElement
  ↓
Datastar swaps out #table-container
```

### Decision 4: EditorState Structure
**Choice**:
```go
type EditorState struct {
    TableName string `json:"tableName"` // e.g., "public.users"
    Page      int    `json:"page"`      // Current page number
    Version   int    `json:"version"`   // For manual refresh triggers
}
```

**Rationale**:
- Minimal state needed for current features
- Version field enables force-refresh without changing other state
- Easy to extend with SortColumn, FilterText, etc.
- JSON serializable for NATS KV storage

### Decision 5: Generic StateStore Enhancement
**Choice**: Add `Watch()` method to existing `StateStore[T]` to support any state type.

**Rationale**:
- Keeps state management logic centralized
- Reusable across features (explorer, editor, future features)
- Consistent API with existing Get/Save methods
- Type-safe with Go generics

## Architecture

### Component Interaction
```
┌─────────────────────────────────────────────────────────────┐
│ Browser                                                      │
│  ┌──────────────┐          ┌─────────────────────────────┐ │
│  │ Editor Page  │◄─────────│ SSE Connection              │ │
│  │ (Datastar)   │          │ (/editor/sse)               │ │
│  └──────┬───────┘          └─────────────────────────────┘ │
│         │ POST                                              │
│         ▼                                                   │
│  ┌──────────────────────────────────────────────────────┐  │
│  │ Command API (/api/editor/*)                          │  │
│  └──────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
         │                           ▲
         │ Update State              │ Watch & Push Updates
         ▼                           │
┌─────────────────────────────────────────────────────────────┐
│ NATS JetStream KV Bucket: "editor_sessions"                 │
│  Key: sessionID → EditorState JSON                          │
│  TTL: 24h (auto-cleanup)                                    │
└─────────────────────────────────────────────────────────────┘
                                     │
                                     │ Query when state changes
                                     ▼
┌─────────────────────────────────────────────────────────────┐
│ PostgreSQL - Source of Truth                                │
│  - Table schemas, columns, actual data                      │
│  - Queried only when rendering (never stores UI state)      │
└─────────────────────────────────────────────────────────────┘
```

### Request Flow Examples

#### Initial Page Load (Query)
```
GET /editor
  → Render empty shell with Datastar SSE connection
  → Browser opens SSE: GET /editor/sse
  → Server gets session state from NATS KV (empty initially)
  → Render empty table view
```

#### Load Table (Command)
```
POST /api/editor/load?table=public.users
  → Get session state
  → Update state.TableName = "public.users", state.Page = 1
  → Save to NATS KV
  → Return 200 OK (no content)
  
NATS KV watch triggers in SSE handler:
  → Parse table name
  → Query Postgres for columns & page 1 data
  → Render DataTable component
  → Push via SSE: PatchElement(#table-container)
  → Browser updates display
```

#### Change Page (Command)
```
POST /api/editor/page/3
  → Get session state
  → Update state.Page = 3
  → Save to NATS KV
  → Return 200 OK
  
NATS KV watch triggers:
  → Query Postgres for page 3 data (using existing state.TableName)
  → Render DataTable component
  → Push via SSE
```

## Data Model

### EditorState (NATS KV)
```go
type EditorState struct {
    TableName string `json:"tableName"` // "schema.table" or empty
    Page      int    `json:"page"`      // 1-based page number
    Version   int    `json:"version"`   // Increment to force refresh
}
```

### NATS KV Bucket Configuration
```go
js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
    Bucket: "editor_sessions",
    TTL:    24 * time.Hour,     // Auto-cleanup after 24h
    Storage: jetstream.MemoryStorage, // Fast ephemeral storage
})
```

### Session Key Format
- Key: `{sessionID}` (from Gorilla sessions)
- Value: JSON-serialized `EditorState`
- Example: `"abc123"` → `{"tableName":"public.users","page":2,"version":1}`

## Implementation Strategy

### Phase 1: Foundation
1. Add `EditorState` struct to services package
2. Create NATS KV bucket "editor_sessions" in initialization
3. Integrate `StateStore[EditorState]` in editor feature
4. Add `Watch()` method to `StateStore[T]`

### Phase 2: Query Routes
1. Simplify `EditorPage` handler to render shell only
2. Add `EditorSSE` handler with state watching
3. Implement `renderTableFromState()` helper for SSE rendering

### Phase 3: Command Routes
1. Add `LoadTable` handler (POST /api/editor/load)
2. Add `ChangePage` handler (POST /api/editor/page/{page})
3. Ensure all commands update NATS state only (no direct rendering)

### Phase 4: Testing & Validation
1. Test SSE reconnection behavior
2. Verify session isolation (multiple users)
3. Confirm NATS KV TTL cleanup
4. Performance test with large tables

## Risks / Trade-offs

### Risk: SSE Connection Stability
**Mitigation**: Browsers auto-reconnect SSE. Add reconnection logging for debugging.

### Risk: NATS KV Overhead
**Mitigation**: Memory storage mode is fast. State is tiny (< 1KB per session).

### Trade-off: Server-Side Rendering Latency
**Impact**: Every state change requires server-side rendering + network transfer.
**Mitigation**: Templ is fast. Pages are small. Can add client-side optimistic updates later if needed.

### Risk: Session State Loss
**Impact**: If NATS restarts, active sessions lose UI state (but not data).
**Mitigation**: 24h TTL means short-lived sessions anyway. Users just re-select table.

### Trade-off: No Optimistic UI
**Impact**: UI waits for server round-trip on every action.
**Mitigation**: Acceptable for initial implementation. SSE is fast enough (<100ms typical).

## Migration Plan

### Prerequisites
- Complete `add-dynamic-table-editor` change (existing handlers/services)
- NATS JetStream running (already in project)

### Migration Steps
1. Deploy with feature flag (optional): `ENABLE_CQRS_EDITOR=true`
2. Create "editor_sessions" KV bucket on startup
3. Deploy new handlers alongside old ones
4. Switch routes to new handlers
5. Monitor SSE connection stability
6. Remove old handlers after validation

### Rollback Plan
- Keep old handlers in code (comment out) for one release
- Can switch routes back if issues found
- No data loss (Postgres untouched)
- NATS KV state is ephemeral (safe to delete bucket)

## Open Questions
None - design is straightforward and builds on existing patterns.

## Future Enhancements (Out of Scope)
- Cell editing: Add `UpdateCell` command
- Row deletion: Add `DeleteRow` command
- Filtering: Add `FilterText` to `EditorState`
- Sorting: Add `SortColumn` and `SortDirection` to state
- Multi-user presence: Broadcast cursor positions via NATS subjects
- Optimistic UI: Client-side state updates before server confirmation
