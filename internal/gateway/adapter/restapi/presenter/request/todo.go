package request

type TodoList struct {
	Name        string `json:"name" validation:"required"`
	Description string `json:"description" validation:"required"`
}

type Todo struct {
	Title    string `json:"title" validation:"required" `
	Priority uint   `json:"priority" validation:"required"  `
	ListID   uint   `json:"listID" validation:"required"`
}

type UpdateTodo struct {
	Title    string `json:"title" validation:"required" `
	Priority uint   `json:"priority" validation:"required"  `
}
