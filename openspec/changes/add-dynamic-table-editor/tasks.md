# Implementation Tasks

## 1. Database Layer
- [ ] 1.1 Add SQL query to fetch paginated table rows (`GetTableData`)
- [ ] 1.2 Add SQL query to count total rows in a table (`GetTableRowCount`)
- [ ] 1.3 Run `sqlc generate` to create Go query functions

## 2. Service Layer
- [ ] 2.1 Create `internal/features/editor/services/editor_service.go`
- [ ] 2.2 Implement `GetTableData(ctx, schema, table, limit, offset)` method
- [ ] 2.3 Implement `GetTableRowCount(ctx, schema, table)` method
- [ ] 2.4 Implement table name parsing logic (split "schema.table" format)

## 3. Domain Models
- [ ] 3.1 Add `TableData` struct to represent query results
- [ ] 3.2 Add `PaginationInfo` struct for page metadata

## 4. Backend Handlers
- [ ] 4.1 Update `EditorPage` handler to accept table and page query parameters
- [ ] 4.2 Create `LoadTableData` handler for SSE data updates
- [ ] 4.3 Add pagination handler for next/previous page actions
- [ ] 4.4 Add error handling for invalid table names

## 5. Routes
- [ ] 5.1 Update editor routes to support query parameters
- [ ] 5.2 Add route for table data loading endpoint

## 6. Frontend - Table Selector
- [ ] 6.1 Add text input field to `editor.templ` for table selection
- [ ] 6.2 Add submit button or auto-load mechanism
- [ ] 6.3 Wire up Datastar actions to trigger data loading
- [ ] 6.4 Add error message display area

## 7. Frontend - Data Table
- [ ] 7.1 Update `DataTable` component to accept dynamic columns
- [ ] 7.2 Update `DataTable` component to accept dynamic row data
- [ ] 7.3 Remove hardcoded sample data from `editor.templ`

## 8. Frontend - Pagination
- [ ] 8.1 Create pagination component in `datatable.templ`
- [ ] 8.2 Add Previous button with disabled state logic
- [ ] 8.3 Add Next button with disabled state logic
- [ ] 8.4 Add page indicator (e.g., "Page 2 of 10")
- [ ] 8.5 Wire up pagination buttons to SSE handlers
- [ ] 8.6 Position pagination controls below the data table

## 9. Testing
- [ ] 9.1 Test table selection with valid table names
- [ ] 9.2 Test error handling with invalid table names
- [ ] 9.3 Test pagination navigation (next, previous, boundaries)
- [ ] 9.4 Test with empty tables
- [ ] 9.5 Test with tables having many columns

## 10. Validation
- [ ] 10.1 Run `openspec validate add-dynamic-table-editor --strict`
- [ ] 10.2 Fix any validation issues
