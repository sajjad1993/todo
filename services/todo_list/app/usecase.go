package app

import (
	"context"
	"github.com/sajjad1993/todo/services/todo_list/domain/todo"
)

type UseCase interface {
	// CreateToDoList creates new t-odo list in t-odo repository
	CreateToDoList(ctx context.Context, list *todo.List) error
	UpdateToDoList(ctx context.Context, id uint, list *todo.List) error
	DeleteToDoList(ctx context.Context, listId uint, userId uint) error
	CreateToDoItem(ctx context.Context, item *todo.Item) error
	UpdateToDoItem(ctx context.Context, itemId uint, item *todo.Item) error
	DeleteToDoItem(ctx context.Context, itemId uint, userId uint) error
	GetTodoListByUserID(ctx context.Context, userId uint) ([]*todo.List, error)
}
