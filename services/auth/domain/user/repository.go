package user

import "context"

type Repository interface {
	Reader
}

type Reader interface {
	GetByEmail(ctx context.Context, email string) (*User, error)
}
