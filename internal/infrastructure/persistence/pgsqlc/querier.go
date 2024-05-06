// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package pgsqlc

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	GetSession(ctx context.Context, uid uuid.UUID) (Session, error)
	GetUser(ctx context.Context, email string) (User, error)
}

var _ Querier = (*Queries)(nil)
