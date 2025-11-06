# Design: Dynamic Table Editor with Pagination

## Context
The editor feature currently displays hardcoded sample data. This change transforms it into a functional no-code table editor where users can select and view actual database tables. The implementation must handle:
- Dynamic table selection
- Efficient data loading for large tables
- Simple pagination for datasets exceeding screen capacity
- PostgreSQL-specific queries for metadata and data retrieval

## Goals / Non-Goals

### Goals
- Enable users to view data from any table in their connected database
- Provide simple pagination (previous/next) for navigating large datasets
- Maintain the existing spreadsheet-like UI with resizable columns
- Use server-rendered templates (Templ) with Datastar for reactivity
- Keep implementation minimal and straightforward

### Non-Goals
- Advanced filtering, sorting, or search capabilities (future enhancement)
- Inline cell editing (future enhancement)
- Row selection or bulk operations (future enhancement)
- Jump-to-page or configurable page size controls (fixed at 50 rows/page)
- Table dropdown/autocomplete (using simple text input for now)

## Decisions

### Decision 1: Table Selection via Text Input
**What**: Use a simple text input accepting "schema.table" format (e.g., "public.users")

**Why**:
- Minimal implementation complexity
- Direct and unambiguous table specification
- Avoids need for table enumeration/dropdown
- Can be enhanced later with autocomplete

**Alternatives considered**:
- Dropdown list of all tables: requires loading all tables upfront, complex for large databases
- Separate schema + table inputs: more UI complexity
- Reuse explorer sidebar: couples two features tightly

### Decision 2: Dynamic SQL Query Construction
**What**: Use PostgreSQL `format()` function or prepared statement with dynamic table/schema names

**Why**:
- sqlc doesn't support dynamic table names in queries
- Need to safely construct queries at runtime
- Must prevent SQL injection via proper escaping

**Implementation**:
```sql
-- Option A: Use format() with proper escaping (chosen)
SELECT * FROM %I.%I ORDER BY ... LIMIT $1 OFFSET $2

-- Option B: Direct string interpolation (rejected - unsafe)
```

**Alternatives considered**:
- Generate separate sqlc queries for each table: impractical
- Use ORM: adds complexity, against project conventions

### Decision 3: Fixed Page Size (50 rows)
**What**: Hardcode page size to 50 rows per page

**Why**:
- Simplifies UI and state management
- Adequate for most debugging/viewing use cases
- Can be made configurable later if needed

**Alternatives considered**:
- User-configurable page size: adds UI complexity
- Adaptive page size based on screen height: over-engineering

### Decision 4: Pagination Strategy
**What**: Use LIMIT/OFFSET for pagination with separate COUNT query

**Why**:
- Standard PostgreSQL pagination pattern
- Simple to implement and understand
- Adequate performance for typical table sizes

**Trade-offs**:
- OFFSET can be slow on very large tables
- COUNT(*) can be expensive, but PostgreSQL optimizes this well
- For future: consider cursor-based pagination for very large tables

**Alternatives considered**:
- Cursor-based pagination: more complex, unnecessary for initial implementation
- Load all data client-side: impractical for large tables

### Decision 5: Service Layer Architecture
**What**: Create `editor_service.go` similar to `explorer_service.go`

**Why**:
- Consistent with existing project patterns
- Encapsulates database logic
- Allows for testing and future enhancement

**Structure**:
```go
type EditorService struct {
    q *store.Queries
}

func (s *EditorService) GetTableData(ctx, schema, table, limit, offset)
func (s *EditorService) GetTableRowCount(ctx, schema, table)
```

### Decision 6: Data Type Handling
**What**: Return all cell values as strings initially

**Why**:
- Simplifies initial implementation
- Works with existing DataTable component expecting string data
- Type-aware rendering can be added later (e.g., date formatting, number alignment)

**Future enhancement**:
- Add type information to TableData struct
- Render cells differently based on PostgreSQL data type
- Special handling for JSON, arrays, timestamps

## Risks / Trade-offs

### Risk 1: SQL Injection
**Risk**: Dynamic table/schema names could allow SQL injection
**Mitigation**: 
- Use PostgreSQL `%I` format specifier (identifier escaping)
- Validate schema/table names against information_schema
- Consider allowlist if security is critical

### Risk 2: OFFSET Performance
**Risk**: Large OFFSETs can be slow on big tables
**Mitigation**:
- Document 50-row page size as reasonable for most use cases
- Future: add note about performance for million+ row tables
- Future: consider cursor-based pagination

### Risk 3: Large Row Values
**Risk**: Tables with very wide rows or many columns may cause display issues
**Mitigation**:
- Existing column resize handles this partially
- Consider column limit or horizontal scroll
- Future: add column visibility controls

## Migration Plan

This is a new capability, not a breaking change:
1. Deploy updated SQL queries and regenerate sqlc code
2. Add service layer
3. Update editor UI with new components
4. Feature is opt-in via navigation to /editor

No data migration or user action required.

## Open Questions

1. **Column ordering**: Should we preserve table's natural column order or allow custom ordering?
   - **Resolved**: Use natural order (ordinal_position) for simplicity
   
2. **Primary key detection**: Should we order by primary key automatically?
   - **Resolved**: Order by table's natural order initially; can enhance later
   
3. **NULL handling**: How should NULL values display?
   - **Resolved**: Display as "(null)" or empty cell; can refine in future
   
4. **Error recovery**: If table is dropped mid-session, what's the UX?
   - **Resolved**: Show error message, let user select different table
