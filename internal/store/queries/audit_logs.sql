-- Audit Logs queries

-- name: CreateAuditLog :one
INSERT INTO audit_logs (
    user_id,
    exposed_table_id,
    operation_type,
    row_identifier,
    changed_data,
    ip_address,
    user_agent
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetAuditLog :one
SELECT *
FROM audit_logs
WHERE id = ?
LIMIT 1;

-- name: ListAuditLogsByUser :many
SELECT *
FROM audit_logs
WHERE user_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: ListAuditLogsByTable :many
SELECT *
FROM audit_logs
WHERE exposed_table_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: ListAuditLogsByOperation :many
SELECT *
FROM audit_logs
WHERE operation_type = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: ListAuditLogsByDateRange :many
SELECT *
FROM audit_logs
WHERE created_at >= ? AND created_at <= ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: DeleteOldAuditLogs :exec
DELETE FROM audit_logs
WHERE created_at < ?;
