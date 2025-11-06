# Change: Add Dynamic Table Editor with Pagination

## Why
The editor currently displays hardcoded sample data. Users need the ability to select any table from their database and view/edit its actual data with pagination support for large datasets.

## What Changes
- Add table selection UI (simple text input for schema.table format)
- Fetch and display actual table data from the backend
- Implement pagination controls at the bottom of the data table
- Add SQL query to retrieve paginated table rows
- Create service layer to handle table data fetching
- Update editor handlers to support data loading and pagination

## Impact
- Affected specs: `table-editor` (new capability)
- Affected code:
  - `internal/features/editor/pages/editor.templ` - add table selector input
  - `internal/features/editor/components/datatable.templ` - add pagination controls
  - `internal/features/editor/handlers.go` - add data fetching handlers
  - `internal/features/editor/routes.go` - add new routes
  - `internal/features/editor/services/` - new service layer
  - `internal/store/query.sql` - new SQL queries for data and count
  - `internal/domain/schema.go` - potentially add TableRow type
