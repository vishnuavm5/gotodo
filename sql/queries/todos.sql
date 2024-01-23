-- name: CreateTodo :one
INSERT INTO todos (id,title,description,created_at,updated_at,user_id)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING title,description;

-- name: GetTodoList :many
SELECT title,id,description FROM todos WHERE user_id=$1;

-- name: GetTodoById :one
SELECT title,id,description FROM todos WHERE id=$1;

-- name: UpdateTodoById :one
UPDATE todos
SET title=$1,description=$2 WHERE id=$3 AND user_id=$4
RETURNING title,description,id;

-- name: DeleteTodoById :exec
DELETE FROM todos WHERE id=$1 and user_id=$2;