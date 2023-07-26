package orm

import (
	"context"
	"errors"
	"github.com/sajjad1993/todo/internal/todo_list/domain/todo"
	"github.com/sajjad1993/todo/pkg/errs"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todo.Repository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) CreateItem(ctx context.Context, item *todo.Item) error {
	model := new(TodoItem).fromEntity(item)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errs.NewDuplicateEntity(err.Error())
		}
		return errs.NewInternalError(err.Error())
	}
	return nil
}

func (r *TodoRepository) GetListByUserId(ctx context.Context, userId uint) ([]*todo.List, error) {
	var result []TodoList
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Preload("TodoItems").Find(&result).Error
	if err != nil {
		return nil, err
	}
	var todoLists []*todo.List
	for _, p := range result {
		todoLists = append(todoLists, p.toEntity())
	}
	return todoLists, nil
}

func (r *TodoRepository) GetListById(ctx context.Context, id uint) (*todo.List, error) {
	model := new(TodoList)
	err := r.db.WithContext(ctx).Where("id = ?", id).First(model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError(err.Error())
		}
		return nil, errs.NewInternalError(err.Error())
	}
	return model.toEntity(), nil
}

func (r *TodoRepository) CreateList(ctx context.Context, todoList *todo.List) error {
	model := new(TodoList).fromEntity(todoList)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errs.NewDuplicateEntity(err.Error())
		}
		return errs.NewInternalError(err.Error())
	}
	return nil
}

func (r *TodoRepository) UpdateList(ctx context.Context, id uint, list *todo.List) error {
	model := new(TodoList)
	err := r.db.WithContext(ctx).Where("id = ?", id).First(model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError(err.Error())
		}
		return errs.NewInternalError(err.Error())
	}
	updateList := struct {
		ID          uint
		Name        string
		Description string
	}{
		ID: id, Name: list.Name, Description: list.Description,
	}
	err = r.db.Model(model).Updates(updateList).Error
	return err
}

func (r *TodoRepository) DeleteList(ctx context.Context, listId uint) error {
	err := r.db.WithContext(ctx).Where("todo_list_id = ?", listId).Delete(&TodoItem{}).Error

	err = r.db.WithContext(ctx).Where("id = ?", listId).Delete(&TodoList{}).Error
	return err
}

func (r *TodoRepository) GetItemById(ctx context.Context, id uint) (*todo.Item, error) {
	model := new(TodoItem)
	err := r.db.WithContext(ctx).Where("id = ?", id).First(model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError(err.Error())
		}
		return nil, errs.NewInternalError(err.Error())
	}
	return model.toEntity(), nil
}

func (r *TodoRepository) UpdateItem(ctx context.Context, id uint, todo *todo.Item) error {
	model := new(TodoItem)
	err := r.db.WithContext(ctx).Where("id = ?", id).First(model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError(err.Error())
		}
		return errs.NewInternalError(err.Error())
	}
	updateItem := struct {
		ID       uint
		Title    string
		Priority uint
	}{
		ID: id, Title: todo.Title, Priority: todo.Priority,
	}
	err = r.db.Model(model).Updates(updateItem).Error
	return err
}
func (r *TodoRepository) DeleteItem(ctx context.Context, itemId uint) error {
	err := r.db.WithContext(ctx).Where("id = ?", itemId).Delete(&TodoItem{}).Error
	return err
}
