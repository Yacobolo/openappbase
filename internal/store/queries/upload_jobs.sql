-- Upload Jobs queries

-- name: CreateUploadJob :one
INSERT INTO upload_jobs (
    user_id,
    exposed_table_id,
    filename,
    row_count,
    status,
    mapping_config
) VALUES (
    ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetUploadJob :one
SELECT *
FROM upload_jobs
WHERE id = ?
LIMIT 1;

-- name: ListUploadJobsByUser :many
SELECT *
FROM upload_jobs
WHERE user_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: ListUploadJobsByTable :many
SELECT *
FROM upload_jobs
WHERE exposed_table_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: ListUploadJobsByStatus :many
SELECT *
FROM upload_jobs
WHERE status = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: UpdateUploadJobStatus :exec
UPDATE upload_jobs
SET 
    status = ?,
    error_message = ?
WHERE id = ?;

-- name: CompleteUploadJob :exec
UPDATE upload_jobs
SET 
    status = 'completed',
    completed_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteUploadJob :exec
DELETE FROM upload_jobs
WHERE id = ?;

-- name: DeleteOldUploadJobs :exec
DELETE FROM upload_jobs
WHERE created_at < ? AND status IN ('completed', 'failed');
