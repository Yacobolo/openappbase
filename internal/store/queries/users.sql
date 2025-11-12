-- Users queries

-- name: GetAllUsers :many
SELECT *
FROM users;

-- name: CreateUser :one
INSERT INTO users (
    email,
    username,
    display_name,
    hashed_password,
    sso_provider,
    sso_id,
    is_active,
    is_admin
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = ?
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?
LIMIT 1;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = ?
LIMIT 1;

-- name: GetUserBySSO :one
SELECT *
FROM users
WHERE sso_provider = ? AND sso_id = ?
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET 
    email = ?,
    username = ?,
    display_name = ?,
    is_active = ?,
    is_admin = ?
WHERE id = ?
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users
SET hashed_password = ?
WHERE id = ?;

-- name: UpdateUserLastLogin :exec
UPDATE users
SET last_login_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
