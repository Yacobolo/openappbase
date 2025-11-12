-- Connection Management queries

-- name: CreateConnection :one
INSERT INTO connections (
    name,
    host,
    port,
    database,
    username,
    encrypted_password,
    ssl_mode,
    ssl_config
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetConnection :one
SELECT *
FROM connections
WHERE id = ?
LIMIT 1;

-- name: GetConnectionByName :one
SELECT *
FROM connections
WHERE name = ?
LIMIT 1;

-- name: ListConnections :many
SELECT *
FROM connections
ORDER BY created_at DESC;

-- name: ListActiveConnections :many
SELECT *
FROM connections
WHERE is_active = 1
ORDER BY created_at DESC;

-- name: UpdateConnection :one
UPDATE connections
SET 
    name = ?,
    host = ?,
    port = ?,
    database = ?,
    username = ?,
    encrypted_password = ?,
    ssl_mode = ?,
    ssl_config = ?
WHERE id = ?
RETURNING *;

-- name: UpdateConnectionStatus :exec
UPDATE connections
SET is_active = ?
WHERE id = ?;

-- name: DeleteConnection :exec
DELETE FROM connections
WHERE id = ?;
