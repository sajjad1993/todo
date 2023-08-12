package broker_utils

type DeleteTodoListMessage struct {
	ID     uint
	UserID uint
}

type DeleteTodoMessage struct {
	ID     uint
	UserID uint
}

const (
	SignUp                    = "SIGNUP"
	DoneSignUp                = "SIGNED_UP"
	CreateTodoListCommand     = "CREATE_TODO_LIST"
	DoneCreateTodoListCommand = "CREATED_TODO_LIST"
	CreateTodoCommand         = "CREATE_TODO"
	DONECreateTodoCommand     = "CREATED_TODO"
	DeleteTodoListCommand     = "DELETE_TODO_LIST"
	DoneDeleteTodoListCommand = "DELETED_TODO_LIST"
	DeleteTodoItemCommand     = "DELETE_TODO_Item"
	DoneDeleteTodoItemCommand = "DELETED_TODO_Item"
	UpdateTodoListCommand     = "UPDATE_TODO_LIST"
	DoneUpdateTodoListCommand = "UPDATED_TODO_LIST"
	UpdateTodo                = "UPDATE_TODO"
	DoneUpdateTodo            = "UPDATED_TODO"
)
