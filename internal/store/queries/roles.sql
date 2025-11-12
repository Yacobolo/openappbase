-- RBAC Roles queries

-- name: CreateRole :one
INSERT INTO roles (
    name,
    description,
    is_system_role
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: GetRole :one
SELECT *
FROM roles
WHERE id = ?
LIMIT 1;

-- name: GetRoleByName :one
SELECT *
FROM roles
WHERE name = ?
LIMIT 1;

-- name: ListRoles :many
SELECT *
FROM roles
ORDER BY name;

-- name: UpdateRole :one
UPDATE roles
SET 
    name = ?,
    description = ?
WHERE id = ?
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = ? AND is_system_role = 0;

-- name: AssignRoleToUser :one
INSERT INTO user_roles (
    user_id,
    role_id
) VALUES (
    ?, ?
)
RETURNING *;

-- name: RemoveRoleFromUser :exec
DELETE FROM user_roles
WHERE user_id = ? AND role_id = ?;

-- name: GetUserRoles :many
SELECT r.*
FROM roles r
INNER JOIN user_roles ur ON r.id = ur.role_id
WHERE ur.user_id = ?
ORDER BY r.name;

-- name: GetUsersWithRole :many
SELECT u.*
FROM users u
INNER JOIN user_roles ur ON u.id = ur.user_id
WHERE ur.role_id = ?
ORDER BY u.username;
