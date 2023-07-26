package app

import (
	"github.com/sajjad1993/todo/internal/gateway/app/command"
	"github.com/sajjad1993/todo/internal/gateway/app/query"
)

type Application struct {
	Commands *Commands
	Queries  *Queries
}

type Commands struct {
	CreateTodoList *command.CreateTodoList
	DeleteTodoList *command.DeleteTodoList
	UpdateTodoList *command.UpdateTodoList
	CreateTodo     *command.CreateTodo
	UpdateTodo     *command.UpdateTodo
	DeleteTodo     *command.DeleteTodo
	SignUp         *command.SignUp
}

type Queries struct {
	SignIn       *query.SignIn
	CheckToken   *query.CheckToken
	ListToDoList *query.ListToDoList
}

func New(commands *Commands, queries *Queries) *Application {
	return &Application{
		Commands: commands,
		Queries:  queries,
	}
}

func NewCommands(SignUp *command.SignUp, CreateTodo *command.CreateTodo,
	createTodoList *command.CreateTodoList, UpdateTodoList *command.UpdateTodoList,
	DeleteTodoList *command.DeleteTodoList, UpdateTodo *command.UpdateTodo,
	DeleteTodo *command.DeleteTodo) *Commands {
	return &Commands{
		//t-odo
		CreateTodo: CreateTodo,
		UpdateTodo: UpdateTodo,
		DeleteTodo: DeleteTodo,

		CreateTodoList: createTodoList,
		DeleteTodoList: DeleteTodoList,
		UpdateTodoList: UpdateTodoList,

		//user
		SignUp: SignUp,
	}
}

func NewQueries(signIn *query.SignIn, checkToken *query.CheckToken, ListToDoList *query.ListToDoList) *Queries {
	return &Queries{
		SignIn:       signIn,
		CheckToken:   checkToken,
		ListToDoList: ListToDoList,
	}
}
