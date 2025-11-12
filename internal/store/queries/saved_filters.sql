-- Saved Filters queries

-- name: CreateSavedFilter :one
INSERT INTO saved_filters (
    user_id,
    exposed_table_id,
    filter_name,
    filter_config,
    is_public
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetSavedFilter :one
SELECT *
FROM saved_filters
WHERE id = ?
LIMIT 1;

-- name: ListSavedFiltersByUser :many
SELECT *
FROM saved_filters
WHERE user_id = ?
ORDER BY filter_name;

-- name: ListSavedFiltersByTable :many
SELECT *
FROM saved_filters
WHERE user_id = ? AND exposed_table_id = ?
ORDER BY filter_name;

-- name: ListPublicSavedFiltersByTable :many
SELECT *
FROM saved_filters
WHERE exposed_table_id = ? AND is_public = 1
ORDER BY filter_name;

-- name: UpdateSavedFilter :one
UPDATE saved_filters
SET 
    filter_name = ?,
    filter_config = ?,
    is_public = ?
WHERE id = ?
RETURNING *;

-- name: DeleteSavedFilter :exec
DELETE FROM saved_filters
WHERE id = ?;

-- name: DeleteSavedFiltersByUser :exec
DELETE FROM saved_filters
WHERE user_id = ?;

-- name: DeleteSavedFiltersByTable :exec
DELETE FROM saved_filters
WHERE exposed_table_id = ?;
