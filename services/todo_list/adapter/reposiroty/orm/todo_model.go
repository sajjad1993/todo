package orm

import (
	"github.com/sajjad1993/todo/services/todo_list/domain/todo"
	"gorm.io/gorm"
)

type TodoItem struct {
	gorm.Model
	Title      string    `json:"title" gorm:"not null;size:32"`
	Priority   uint      `json:"priority" gorm:"not null"`
	TodoList   *TodoList `json:"list"gorm:"foreignkey:TodoListID"`
	TodoListID uint      `json:"todo_list_id"`
}

func (t *TodoItem) toEntity() *todo.Item {
	if t == nil || t.ID == 0 {
		return nil
	}
	return &todo.Item{
		ID:       t.ID,
		Title:    t.Title,
		Priority: t.Priority,
		ListID:   t.TodoListID,
	}
}

func (t *TodoItem) fromEntity(e *todo.Item) *TodoItem {
	if t == nil || e == nil {
		return nil
	}
	t.ID = e.ID
	t.Title = e.Title
	t.Priority = e.Priority
	t.TodoListID = e.ListID
	return t
}
