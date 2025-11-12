-- User Preferences queries

-- name: CreateUserPreference :one
INSERT INTO user_preferences (
    user_id,
    exposed_table_id,
    preference_key,
    preference_value
) VALUES (
    ?, ?, ?, ?
)
RETURNING *;

-- name: GetUserPreference :one
SELECT *
FROM user_preferences
WHERE user_id = ? AND exposed_table_id = ? AND preference_key = ?
LIMIT 1;

-- name: ListUserPreferencesByUser :many
SELECT *
FROM user_preferences
WHERE user_id = ?
ORDER BY preference_key;

-- name: ListUserPreferencesByTable :many
SELECT *
FROM user_preferences
WHERE user_id = ? AND exposed_table_id = ?
ORDER BY preference_key;

-- name: UpdateUserPreference :one
UPDATE user_preferences
SET preference_value = ?
WHERE user_id = ? AND exposed_table_id = ? AND preference_key = ?
RETURNING *;

-- name: DeleteUserPreference :exec
DELETE FROM user_preferences
WHERE user_id = ? AND exposed_table_id = ? AND preference_key = ?;

-- name: DeleteUserPreferencesByUser :exec
DELETE FROM user_preferences
WHERE user_id = ?;

-- name: DeleteUserPreferencesByTable :exec
DELETE FROM user_preferences
WHERE exposed_table_id = ?;
