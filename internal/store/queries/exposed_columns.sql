-- Exposed Columns queries

-- name: CreateExposedColumn :one
INSERT INTO exposed_columns (
    exposed_table_id,
    column_name,
    display_name,
    data_type,
    is_visible,
    is_editable,
    display_order,
    validation_rules
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetExposedColumn :one
SELECT *
FROM exposed_columns
WHERE id = ?
LIMIT 1;

-- name: ListExposedColumnsByTable :many
SELECT *
FROM exposed_columns
WHERE exposed_table_id = ?
ORDER BY display_order, column_name;

-- name: ListVisibleColumnsByTable :many
SELECT *
FROM exposed_columns
WHERE exposed_table_id = ? AND is_visible = 1
ORDER BY display_order, column_name;

-- name: UpdateExposedColumn :one
UPDATE exposed_columns
SET 
    display_name = ?,
    is_visible = ?,
    is_editable = ?,
    display_order = ?,
    validation_rules = ?
WHERE id = ?
RETURNING *;

-- name: DeleteExposedColumn :exec
DELETE FROM exposed_columns
WHERE id = ?;

-- name: DeleteExposedColumnsByTable :exec
DELETE FROM exposed_columns
WHERE exposed_table_id = ?;
