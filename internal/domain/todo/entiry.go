package todo

import "github.com/sajjad1993/todo/internal/domain/todo_list"

type Entity struct {
	Title    string
	Priority uint
	TodoList *todo_list.Entity
}
