-- Permissions queries

-- name: CreatePermission :one
INSERT INTO permissions (
    role_id,
    exposed_table_id,
    can_view,
    can_create,
    can_update,
    can_delete
) VALUES (
    ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetPermission :one
SELECT *
FROM permissions
WHERE id = ?
LIMIT 1;

-- name: GetPermissionByRoleAndTable :one
SELECT *
FROM permissions
WHERE role_id = ? AND exposed_table_id = ?
LIMIT 1;

-- name: ListPermissionsByRole :many
SELECT *
FROM permissions
WHERE role_id = ?
ORDER BY exposed_table_id;

-- name: ListPermissionsByTable :many
SELECT *
FROM permissions
WHERE exposed_table_id = ?
ORDER BY role_id;

-- name: UpdatePermission :one
UPDATE permissions
SET 
    can_view = ?,
    can_create = ?,
    can_update = ?,
    can_delete = ?
WHERE id = ?
RETURNING *;

-- name: DeletePermission :exec
DELETE FROM permissions
WHERE id = ?;

-- name: DeletePermissionsByRole :exec
DELETE FROM permissions
WHERE role_id = ?;

-- name: DeletePermissionsByTable :exec
DELETE FROM permissions
WHERE exposed_table_id = ?;

-- Row-Level Security Rules queries

-- name: CreateRLSRule :one
INSERT INTO row_level_security_rules (
    exposed_table_id,
    role_id,
    rule_name,
    filter_expression,
    is_active
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetRLSRule :one
SELECT *
FROM row_level_security_rules
WHERE id = ?
LIMIT 1;

-- name: ListRLSRulesByTable :many
SELECT *
FROM row_level_security_rules
WHERE exposed_table_id = ?
ORDER BY rule_name;

-- name: ListRLSRulesByRole :many
SELECT *
FROM row_level_security_rules
WHERE role_id = ?
ORDER BY rule_name;

-- name: ListActiveRLSRulesByTableAndRole :many
SELECT *
FROM row_level_security_rules
WHERE exposed_table_id = ? AND role_id = ? AND is_active = 1
ORDER BY rule_name;

-- name: UpdateRLSRule :one
UPDATE row_level_security_rules
SET 
    rule_name = ?,
    filter_expression = ?,
    is_active = ?
WHERE id = ?
RETURNING *;

-- name: DeleteRLSRule :exec
DELETE FROM row_level_security_rules
WHERE id = ?;

-- name: DeleteRLSRulesByTable :exec
DELETE FROM row_level_security_rules
WHERE exposed_table_id = ?;

-- name: DeleteRLSRulesByRole :exec
DELETE FROM row_level_security_rules
WHERE role_id = ?;
