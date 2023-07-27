package user

import "context"

type Repository interface {
	Reader
	Writer
}

type Reader interface {
	GetById(ctx context.Context, id uint) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type Writer interface {
	Create(ctx context.Context, user *User) error
}
