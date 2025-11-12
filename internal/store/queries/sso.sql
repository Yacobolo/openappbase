-- SSO Providers queries

-- name: CreateSSOProvider :one
INSERT INTO sso_providers (
    provider_name,
    provider_type,
    client_id,
    encrypted_client_secret,
    auth_url,
    token_url,
    user_info_url,
    config,
    is_active
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetSSOProvider :one
SELECT *
FROM sso_providers
WHERE id = ?
LIMIT 1;

-- name: GetSSOProviderByName :one
SELECT *
FROM sso_providers
WHERE provider_name = ?
LIMIT 1;

-- name: ListSSOProviders :many
SELECT *
FROM sso_providers
ORDER BY provider_name;

-- name: ListActiveSSOProviders :many
SELECT *
FROM sso_providers
WHERE is_active = 1
ORDER BY provider_name;

-- name: UpdateSSOProvider :one
UPDATE sso_providers
SET 
    provider_name = ?,
    provider_type = ?,
    client_id = ?,
    encrypted_client_secret = ?,
    auth_url = ?,
    token_url = ?,
    user_info_url = ?,
    config = ?,
    is_active = ?
WHERE id = ?
RETURNING *;

-- name: DeleteSSOProvider :exec
DELETE FROM sso_providers
WHERE id = ?;
