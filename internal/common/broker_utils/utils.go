package broker_utils

type DeleteTodoListMessage struct {
	ID     uint
	UserID uint
}

const SignUp = "SIGNUP"
const DoneSignUp = "SIGNED_UP"
const CreateTodoListCommand = "CREATE_TODO_LIST"
const DoneCreateTodoListCommand = "CREATED_TODO_LIST"
const CreateTodoCommand = "CREATE_TODO"
const DONECreateTodoCommand = "CREATED_TODO"
const DeleteTodoListCommand = "DELETE_TODO_LIST"
const DoneDeleteTodoListCommand = "DELETED_TODO_LIST"
