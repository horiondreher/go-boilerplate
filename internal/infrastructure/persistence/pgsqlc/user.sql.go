// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package pgsqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user" (
	"email"
	, "password"
	, "full_name"
	, "is_staff"
	, "is_active"
	, "last_login"
	)
VALUES (
	$1
	, $2
	, $3
	, $4
	, $5
	, $6
	) RETURNING "uid"
	, "email"
	, "full_name"
	, "created_at"
	, "modified_at"
`

type CreateUserParams struct {
	Email     string
	Password  string
	FullName  string
	IsStaff   bool
	IsActive  bool
	LastLogin time.Time
}

type CreateUserRow struct {
	UID        uuid.UUID
	Email      string
	FullName   string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Email,
		arg.Password,
		arg.FullName,
		arg.IsStaff,
		arg.IsActive,
		arg.LastLogin,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.UID,
		&i.Email,
		&i.FullName,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, uid, email, password, full_name, is_staff, is_active, last_login, created_at, modified_at
FROM "user"
WHERE "email" = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UID,
		&i.Email,
		&i.Password,
		&i.FullName,
		&i.IsStaff,
		&i.IsActive,
		&i.LastLogin,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getUserByUID = `-- name: GetUserByUID :one
SELECT id, uid, email, password, full_name, is_staff, is_active, last_login, created_at, modified_at
FROM "user"
WHERE "uid" = $1 LIMIT 1
`

func (q *Queries) GetUserByUID(ctx context.Context, uid uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUID, uid)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UID,
		&i.Email,
		&i.Password,
		&i.FullName,
		&i.IsStaff,
		&i.IsActive,
		&i.LastLogin,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}
