package todo_list

import "github.com/sajjad1993/todo/internal/domain/todo"

type Entity struct {
	Name        string
	Description string
	Todos       []todo.Entity
}
