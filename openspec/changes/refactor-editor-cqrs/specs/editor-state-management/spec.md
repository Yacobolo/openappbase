# Editor State Management Specification

## ADDED Requirements

### Requirement: Editor Session State Structure
The editor SHALL maintain session-specific UI state separate from database data.

#### Scenario: Session state includes table selection and pagination
- **GIVEN** a user is interacting with the editor
- **WHEN** the system needs to track their current view
- **THEN** the system SHALL store an `EditorState` containing:
  - `TableName` (string): the selected table in "schema.table" format or empty
  - `Page` (integer): the current 1-based page number
  - `Version` (integer): a version counter for triggering manual refreshes

#### Scenario: Session state is isolated per user
- **GIVEN** multiple users are using the editor simultaneously
- **WHEN** each user interacts with different tables or pages
- **THEN** each user's session SHALL have independent state
- **AND** changes to one user's state SHALL NOT affect other users

### Requirement: NATS KV State Storage
The editor SHALL use NATS JetStream Key-Value bucket to store ephemeral session state.

#### Scenario: State stored in dedicated KV bucket
- **GIVEN** the editor feature is initialized
- **WHEN** the application starts
- **THEN** a NATS JetStream KV bucket named "editor_sessions" SHALL be created
- **AND** the bucket SHALL use memory storage for fast access
- **AND** the bucket SHALL have a TTL of 24 hours for automatic cleanup

#### Scenario: State keyed by session ID
- **GIVEN** a user has an active session
- **WHEN** the editor state is saved
- **THEN** the state SHALL be stored with the session ID as the key
- **AND** the state SHALL be serialized as JSON

#### Scenario: Stale sessions automatically cleaned up
- **GIVEN** a user's session has been inactive for 24 hours
- **WHEN** the NATS KV TTL expires
- **THEN** the session state SHALL be automatically deleted
- **AND** no manual cleanup SHALL be required

### Requirement: CQRS Route Separation
The editor SHALL separate query routes from command routes following CQRS principles.

#### Scenario: Query routes are read-only
- **GIVEN** the editor is accessed via query routes
- **WHEN** a request is made to `/editor` or `/editor/sse`
- **THEN** the request SHALL use GET method
- **AND** the request SHALL NOT modify any state
- **AND** the request SHALL return current state representation

#### Scenario: Command routes modify state
- **GIVEN** the editor receives a state-changing action
- **WHEN** a request is made to `/api/editor/*` routes
- **THEN** the request SHALL use POST, PUT, or DELETE methods
- **AND** the request SHALL modify session state in NATS KV
- **AND** the request SHALL NOT directly render or return UI components

#### Scenario: Command routes trigger SSE updates
- **GIVEN** a command successfully updates state
- **WHEN** the state is saved to NATS KV
- **THEN** any active SSE connections for that session SHALL be notified
- **AND** the SSE handler SHALL re-render affected components

### Requirement: SSE-Based State Watching
The editor SHALL provide an SSE endpoint that watches for state changes and pushes updates to the browser.

#### Scenario: SSE connection establishes watch
- **GIVEN** a user loads the editor page
- **WHEN** the browser connects to `/editor/sse`
- **THEN** the server SHALL retrieve the current session state
- **AND** the server SHALL establish a watch on the NATS KV for that session
- **AND** the server SHALL send the initial table rendering

#### Scenario: State change triggers re-render
- **GIVEN** an SSE connection is actively watching a session
- **WHEN** the session state is updated in NATS KV
- **THEN** the SSE watcher SHALL receive a notification
- **AND** the server SHALL query Postgres for current table data based on new state
- **AND** the server SHALL render the appropriate component
- **AND** the server SHALL push the component update via SSE
- **AND** Datastar SHALL swap out the updated portion of the DOM

#### Scenario: SSE handles empty table state
- **GIVEN** an SSE connection is active
- **WHEN** the session state has no table selected
- **THEN** the server SHALL render an empty table view component
- **AND** the empty view SHALL be pushed via SSE

#### Scenario: SSE handles errors gracefully
- **GIVEN** an SSE connection is active
- **WHEN** an error occurs loading table data (invalid table, permissions, etc.)
- **THEN** the server SHALL render an error view component
- **AND** the error view SHALL display a user-friendly message
- **AND** the error SHALL be pushed via SSE without closing the connection

### Requirement: Load Table Command
The editor SHALL provide a command to select which table to view.

#### Scenario: Load table via API
- **GIVEN** a user wants to view a specific table
- **WHEN** a POST request is made to `/api/editor/load?table=schema.table`
- **THEN** the system SHALL validate the session
- **AND** the system SHALL update the session state with the table name
- **AND** the system SHALL reset the page number to 1
- **AND** the system SHALL save the state to NATS KV
- **AND** the system SHALL return HTTP 200 OK

#### Scenario: Invalid table name rejected
- **GIVEN** a user attempts to load a table
- **WHEN** the table name is empty or missing
- **THEN** the system SHALL return HTTP 400 Bad Request
- **AND** the system SHALL NOT modify the session state

### Requirement: Change Page Command
The editor SHALL provide a command to navigate between pages.

#### Scenario: Navigate to specific page via API
- **GIVEN** a user is viewing a paginated table
- **WHEN** a POST request is made to `/api/editor/page/{page}`
- **THEN** the system SHALL validate the page parameter is a positive integer
- **AND** the system SHALL update the session state with the new page number
- **AND** the system SHALL save the state to NATS KV
- **AND** the system SHALL return HTTP 200 OK

#### Scenario: Invalid page number rejected
- **GIVEN** a user attempts to change pages
- **WHEN** the page parameter is not a valid positive integer
- **THEN** the system SHALL return HTTP 400 Bad Request
- **AND** the system SHALL NOT modify the session state

### Requirement: State Store Integration
The editor SHALL use the generic `StateStore[T]` pattern for session state management.

#### Scenario: Initialize state store with EditorState type
- **GIVEN** the editor feature is being set up
- **WHEN** dependencies are initialized
- **THEN** a `StateStore[EditorState]` SHALL be created
- **AND** the state store SHALL be configured with the "editor_sessions" KV bucket
- **AND** the state store SHALL be passed to the editor handlers

#### Scenario: State store provides Get/Save operations
- **GIVEN** the editor needs to access session state
- **WHEN** handlers call state store methods
- **THEN** `Get(ctx, sessionID)` SHALL retrieve state or return empty state if not found
- **AND** `Save(ctx, sessionID, state)` SHALL persist state to NATS KV
- **AND** errors SHALL be propagated to the caller

### Requirement: State Store Watch Capability
The state store SHALL support watching for state changes via NATS KV.

#### Scenario: Watch method enables state observation
- **GIVEN** a handler needs to observe state changes
- **WHEN** the handler calls `Watch(ctx, sessionID)`
- **THEN** the system SHALL create a NATS KV watcher for that specific key
- **AND** the system SHALL return a channel that receives state updates
- **AND** the watcher SHALL continue until the context is canceled

#### Scenario: Watch receives updates on state changes
- **GIVEN** a watcher is active for a session
- **WHEN** another handler saves state for that session
- **THEN** the watcher SHALL receive the updated state
- **AND** the update SHALL be delivered as deserialized EditorState
- **AND** the watcher SHALL continue watching for subsequent changes

#### Scenario: Watch cleanup on context cancellation
- **GIVEN** a watcher is active
- **WHEN** the context is canceled (e.g., SSE connection closes)
- **THEN** the watcher SHALL stop receiving updates
- **AND** the watcher SHALL clean up resources
- **AND** no goroutine leaks SHALL occur

### Requirement: Editor Page Query Handler
The editor page handler SHALL render a minimal shell that connects to SSE.

#### Scenario: Editor page returns HTML shell
- **GIVEN** a user navigates to `/editor`
- **WHEN** a GET request is received
- **THEN** the system SHALL render the editor page template
- **AND** the page SHALL include a container for table content
- **AND** the page SHALL include Datastar SSE connection to `/editor/sse`
- **AND** the page SHALL NOT query Postgres or load table data

### Requirement: Reactive Component Rendering
The SSE handler SHALL render components based on current state and push via Datastar.

#### Scenario: Render table data from state
- **GIVEN** the SSE handler receives a state update with a valid table name
- **WHEN** the handler processes the update
- **THEN** the handler SHALL parse the schema and table from state
- **AND** the handler SHALL query Postgres for columns and data for the current page
- **AND** the handler SHALL render the DataTable component
- **AND** the handler SHALL push the component with `sse.PatchElementTempl()` targeting `#table-container`

#### Scenario: Render error state
- **GIVEN** the SSE handler encounters an error loading data
- **WHEN** the error occurs (table not found, permissions, etc.)
- **THEN** the handler SHALL render an error view component
- **AND** the handler SHALL include the error message
- **AND** the handler SHALL push the error view to `#table-container`
- **AND** the SSE connection SHALL remain open

### Requirement: Command Response Behavior
Command endpoints SHALL return simple HTTP status codes without rendering components.

#### Scenario: Successful command returns 200 OK
- **GIVEN** a command endpoint processes a valid request
- **WHEN** the state is successfully updated
- **THEN** the endpoint SHALL return HTTP 200 OK
- **AND** the response body SHALL be empty or contain minimal JSON
- **AND** the endpoint SHALL NOT render or return HTML/components

#### Scenario: Invalid command returns 4xx error
- **GIVEN** a command endpoint receives an invalid request
- **WHEN** validation fails (bad parameters, missing data, etc.)
- **THEN** the endpoint SHALL return HTTP 400 Bad Request
- **AND** the response MAY include an error message
- **AND** the session state SHALL remain unchanged

#### Scenario: Server error returns 500
- **GIVEN** a command endpoint encounters an internal error
- **WHEN** the error occurs during state save or NATS operations
- **THEN** the endpoint SHALL return HTTP 500 Internal Server Error
- **AND** the error SHALL be logged with structured logging
- **AND** the session state SHALL remain in its previous valid state

## MODIFIED Requirements

None - This is a new capability introducing CQRS pattern to the editor.

## REMOVED Requirements

None - Existing table data display and pagination requirements remain unchanged.
