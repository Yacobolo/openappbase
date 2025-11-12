-- Sessions queries

-- name: CreateSession :one
INSERT INTO sessions (
    user_id,
    session_token,
    ip_address,
    user_agent,
    expires_at
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetSession :one
SELECT *
FROM sessions
WHERE id = ?
LIMIT 1;

-- name: GetSessionByToken :one
SELECT *
FROM sessions
WHERE session_token = ?
LIMIT 1;

-- name: ListSessionsByUser :many
SELECT *
FROM sessions
WHERE user_id = ?
ORDER BY created_at DESC;

-- name: ListActiveSessions :many
SELECT *
FROM sessions
WHERE expires_at > CURRENT_TIMESTAMP
ORDER BY created_at DESC;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = ?;

-- name: DeleteSessionByToken :exec
DELETE FROM sessions
WHERE session_token = ?;

-- name: DeleteSessionsByUser :exec
DELETE FROM sessions
WHERE user_id = ?;

-- name: DeleteExpiredSessions :exec
DELETE FROM sessions
WHERE expires_at <= CURRENT_TIMESTAMP;
