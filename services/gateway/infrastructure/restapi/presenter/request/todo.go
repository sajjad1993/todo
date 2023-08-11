package request

type TodoList struct {
	Name        string `json:"name" validation:"required" binding:"required"`
	Description string `json:"description" validation:"required" binding:"required"`
}

type Todo struct {
	Title    string `json:"title" validation:"required"  binding:"required"`
	Priority uint   `json:"priority" validation:"required"   binding:"required"`
	ListID   uint   `json:"list_id" validation:"required" binding:"required"`
}

type UpdateTodo struct {
	Title    string `json:"title" validation:"required"  binding:"required"`
	Priority uint   `json:"priority" validation:"required"   binding:"required"`
}
