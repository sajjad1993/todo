package query

import (
	"context"
	"github.com/sajjad1993/todo/services/gateway/domain/todo"
)

type ListToDoList struct {
	todoRepo todo.Reader
}

func (q *ListToDoList) Run(ctx context.Context, userID uint) ([]*todo.List, error) {
	todos, err := q.todoRepo.GetByUserID(ctx, userID)
	if err != nil {
		//todo handel grpc errors in grpc layers
		return nil, err
	}
	return todos, nil
}

func NewListToDoList(todoRepo todo.Reader) *ListToDoList {
	return &ListToDoList{
		todoRepo: todoRepo,
	}
}
