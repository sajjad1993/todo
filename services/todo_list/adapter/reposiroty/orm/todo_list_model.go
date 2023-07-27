package orm

import (
	"github.com/sajjad1993/todo/services/todo_list/domain/todo"
	"gorm.io/gorm"
)

type TodoList struct {
	gorm.Model
	Name        string `json:"name" gorm:"Index;not null;size:32"`
	Description string `json:"description" gorm:"not null"`
	UserID      uint   `json:"user_id" gorm:"Index;not null"`
	TodoItems   []TodoItem
}

func (t *TodoList) toEntity() *todo.List {
	var items []*todo.Item
	for _, item := range t.TodoItems {
		items = append(items, item.toEntity())
	}
	return &todo.List{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		UserID:      t.UserID,
		Todos:       items,
	}
}

func (t *TodoList) fromEntity(e *todo.List) *TodoList {

	t.ID = e.ID
	t.Name = e.Name
	t.Description = e.Description
	t.UserID = e.UserID
	return t
}
