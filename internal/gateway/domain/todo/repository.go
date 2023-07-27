package todo

import (
	"context"
)

type Repository interface {
	Reader
	Writer
}

type Reader interface {
	GetByUserID(ctx context.Context, userID uint) ([]*List, error)
}
type Writer interface {
}
