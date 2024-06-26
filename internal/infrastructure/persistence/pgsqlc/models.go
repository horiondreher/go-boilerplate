// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package pgsqlc

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           int64
	UID          uuid.UUID
	UserEmail    string
	RefreshToken string
	UserAgent    string
	ClientIP     string
	IsBlocked    bool
	ExpiresAt    time.Time
	CreatedAt    time.Time
}

type User struct {
	ID         int64
	UID        uuid.UUID
	Email      string
	Password   string
	FullName   string
	IsStaff    bool
	IsActive   bool
	LastLogin  time.Time
	CreatedAt  time.Time
	ModifiedAt time.Time
}
