-- name: CreateUser :one
INSERT INTO users (id, name, username, password, created_at, updated_at,email)
VALUES ($1, $2, $3, $4, $5, $6,$7) 
RETURNING id;


-- name: GetUser :one
SELECT username,name,password,id FROM users WHERE username=$1;
