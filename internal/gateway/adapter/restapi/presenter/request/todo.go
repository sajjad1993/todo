package request

type TodoList struct {
	Name        string `json:"name" validation:"required"`
	Description string `json:"description" validation:"required"`
}
