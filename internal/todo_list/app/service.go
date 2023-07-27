package app

import (
	"context"
	"github.com/sajjad1993/todo/internal/todo_list/domain/todo"
	"github.com/sajjad1993/todo/pkg/errs"
)

// service is the auth use case implementation
type service struct {
	todoRepo todo.Repository
}

// NewService returns a pointer to auth service, and implements the auth use case
func NewService(todoRepo todo.Repository) UseCase {
	return &service{todoRepo: todoRepo}
}

// CreateToDoList creates new t-odo list
func (s *service) CreateToDoList(ctx context.Context, list *todo.List) error {

	err := s.todoRepo.CreateList(ctx, list)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) DeleteToDoList(ctx context.Context, listId uint, userId uint) error {
	err := s.checkUserListOwnership(ctx, listId, userId)
	if err != nil {
		return err
	}
	err = s.todoRepo.DeleteList(ctx, listId)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) UpdateToDoList(ctx context.Context, id uint, list *todo.List) error {
	err := s.checkUserListOwnership(ctx, id, list.UserID)
	if err != nil {
		return err
	}
	err = s.todoRepo.UpdateList(ctx, id, list)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateToDoItem(ctx context.Context, item *todo.Item) error {
	err := s.checkUserListOwnership(ctx, item.ListID, item.UserID)
	if err != nil {
		return err
	}
	err = s.todoRepo.CreateItem(ctx, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateToDoItem(ctx context.Context, itemId uint, item *todo.Item) error {
	todoItem, err := s.todoRepo.GetItemById(ctx, itemId)
	if err != nil {
		return err
	}
	err = s.checkUserListOwnership(ctx, todoItem.ListID, item.UserID)
	if err != nil {
		return err
	}
	err = s.todoRepo.UpdateItem(ctx, itemId, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteToDoItem(ctx context.Context, itemId uint, userId uint) error {
	todoItem, err := s.todoRepo.GetItemById(ctx, itemId)
	if err != nil {
		return err
	}
	err = s.checkUserListOwnership(ctx, todoItem.ListID, userId)
	if err != nil {
		return err
	}
	err = s.todoRepo.DeleteItem(ctx, itemId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetTodoListByUserID(ctx context.Context, userId uint) ([]*todo.List, error) {
	todoLists, err := s.todoRepo.GetListByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return todoLists, err
}
func (s *service) checkUserListOwnership(ctx context.Context, listId uint, userId uint) error {
	todoList, err := s.todoRepo.GetListById(ctx, listId)
	if err != nil {
		return err
	}
	if todoList.UserID != userId {
		return errs.NewUnauthorizedError("the todo-list dose`nt belong to this user")
	}
	return nil
}
