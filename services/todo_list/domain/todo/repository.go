package todo

import "context"

type Repository interface {
	Reader
	Writer
}

type Reader interface {
	GetListByUserId(ctx context.Context, userId uint) ([]*List, error)
	GetListById(ctx context.Context, id uint) (*List, error)
	GetItemById(ctx context.Context, id uint) (*Item, error)
}

type Writer interface {
	CreateList(ctx context.Context, todo *List) error
	CreateItem(ctx context.Context, todo *Item) error
	DeleteItem(ctx context.Context, itemId uint) error
	UpdateItem(ctx context.Context, id uint, todo *Item) error
	UpdateList(ctx context.Context, id uint, todo *List) error
	DeleteList(ctx context.Context, listId uint) error
}
