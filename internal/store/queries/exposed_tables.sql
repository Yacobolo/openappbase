-- Exposed Tables queries

-- name: CreateExposedTable :one
INSERT INTO exposed_tables (
    connection_id,
    schema_name,
    table_name,
    display_name,
    description,
    is_active
) VALUES (
    ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetExposedTable :one
SELECT *
FROM exposed_tables
WHERE id = ?
LIMIT 1;

-- name: ListExposedTablesByConnection :many
SELECT *
FROM exposed_tables
WHERE connection_id = ?
ORDER BY display_name;

-- name: ListActiveExposedTablesByConnection :many
SELECT *
FROM exposed_tables
WHERE connection_id = ? AND is_active = 1
ORDER BY display_name;

-- name: UpdateExposedTable :one
UPDATE exposed_tables
SET 
    display_name = ?,
    description = ?,
    is_active = ?
WHERE id = ?
RETURNING *;

-- name: DeleteExposedTable :exec
DELETE FROM exposed_tables
WHERE id = ?;

-- name: DeleteExposedTablesByConnection :exec
DELETE FROM exposed_tables
WHERE connection_id = ?;
