package orm

import (
	"github.com/sajjad1993/todo/internal/todo_list/domain/todo"
	"gorm.io/gorm"
)

type TodoItem struct {
	gorm.Model
	Title    string    `json:"title" gorm:"not null;size:32"`
	Priority uint      `json:"priority" gorm:"not null"`
	TodoList *TodoList `json:"list"`
	ListId   uint      `json:"list-id"`
}

func (t *TodoItem) toEntity() *todo.Item {
	if t == nil || t.ID == 0 {
		return nil
	}
	return &todo.Item{
		ID:       t.ID,
		Title:    t.Title,
		Priority: t.Priority,
		ListID:   t.ListId,
	}
}

func (t *TodoItem) fromEntity(e *todo.Item) *TodoItem {
	if t == nil || e == nil {
		return nil
	}
	t.ID = e.ID
	t.Title = e.Title
	t.Priority = e.Priority
	t.ListId = e.ListID
	return t
}
