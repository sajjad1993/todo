package response

import (
	"github.com/sajjad1993/todo/internal/gateway/domain/todo"
	"sort"
)

type ListToDOListResponse struct {
	Lists []*List `json:"lists"`
}

type Item struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Priority uint   `json:"priority"`
}

type List struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Items       []*Item `json:"items"`
}

type TodoLists []*List

func (l TodoLists) FromEntity(lists []*todo.List) TodoLists {
	for _, list := range lists {
		todoList := &List{
			ID:          list.ID,
			Name:        list.Name,
			Description: list.Description,
		}
		for _, item := range list.Todos {
			todoItem := &Item{
				ID:       item.ID,
				Title:    item.Title,
				Priority: item.Priority,
			}
			todoList.Items = append(todoList.Items, todoItem)
			sort.Slice(todoList.Items,
				func(i, j int) bool {
					return todoList.Items[i].Priority < todoList.Items[j].Priority
				},
			)
		}
		l = append(l, todoList)
	}
	return l
}
