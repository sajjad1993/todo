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
	CreateList(ctx context.Context, todo *List) error
	CreateItem(ctx context.Context, todo *Item) error
	DeleteItem(ctx context.Context, itemId uint) error
	UpdateItem(ctx context.Context, id uint, todo *Item) error
	UpdateList(ctx context.Context, id uint, todo *List) error
	DeleteList(ctx context.Context, listId uint) error
}
