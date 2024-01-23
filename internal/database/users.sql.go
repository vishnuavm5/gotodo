// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, name, username, password, created_at, updated_at,email)
VALUES ($1, $2, $3, $4, $5, $6,$7) 
RETURNING id
`

type CreateUserParams struct {
	ID        uuid.UUID
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Username,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Email,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getUser = `-- name: GetUser :one
SELECT username,name,password,id FROM users WHERE username=$1
`

type GetUserRow struct {
	Username string
	Name     string
	Password string
	ID       uuid.UUID
}

func (q *Queries) GetUser(ctx context.Context, username string) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i GetUserRow
	err := row.Scan(
		&i.Username,
		&i.Name,
		&i.Password,
		&i.ID,
	)
	return i, err
}
