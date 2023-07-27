package broker_utils

type DeleteTodoListMessage struct {
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
	UpdateTodoListCommand     = "UPDATE_TODO_LIST"
	UpdateTodo                = "UPDATE_TODO"
)
