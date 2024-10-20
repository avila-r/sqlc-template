-- name: CreateTask :one
insert into tasks (title, description)
values ($1, $2)
returning id, title, description, created_at, updated_at;

-- name: ListTasks :many
SELECT id, title, description, status, created_at, updated_at
FROM tasks
ORDER BY created_at DESC;

-- name: GetTaskByID :one
SELECT id, title, description, status, created_at, updated_at
FROM tasks
WHERE id = $1;

-- name: UpdateTaskStatus :exec
UPDATE tasks
SET status = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;

-- name: UpdateTaskDetails :exec
UPDATE tasks
SET title = $2, description = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
