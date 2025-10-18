-- name: GetFullSchemaDetails :many
-- Fetches a complete overview of all schemas and their tables, including ownership
-- and row estimates. This single query powers both the sidebar navigation and the
-- "Details" tab for any selected table.
SELECT
    n.nspname AS table_schema,
    c.relname AS table_name,
    a.rolname AS table_owner,
    -- reltuples is a float4, so we cast to bigint for a clean integer count.
    c.reltuples::bigint AS estimated_row_count,
    c.relrowsecurity AS row_security_enabled,
    -- A window function to count tables per schema without a separate query.
    COUNT(*) OVER (PARTITION BY n.nspname) AS schema_table_count
FROM
    pg_class c
    JOIN
    pg_namespace n ON n.oid = c.relnamespace
    JOIN
    pg_authid a ON a.oid = c.relowner
ORDER BY 
table_schema, table_name

---

-- name: GetTableColumns :many
-- Retrieves detailed information about each column in a specific table for the "Columns" tab.
SELECT
    column_name,
    data_type,
    -- Present is_nullable as a boolean for easier use in Go.
    CASE WHEN is_nullable = 'YES' THEN true ELSE false END AS is_nullable,
    column_default
FROM
    information_schema.columns
WHERE
    table_schema = $1
    AND table_name = $2
ORDER BY
    ordinal_position;

---

-- name: GetTableIndexes :many
-- Fetches all indexes for a given table. This is a common and useful feature for a database explorer.
SELECT
    indexname,
    indexdef
FROM
    pg_indexes
WHERE
    schemaname = $1
    AND tablename = $2
ORDER BY
    indexname;