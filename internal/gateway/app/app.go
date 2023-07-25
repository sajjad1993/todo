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
	SignUp         *command.SignUp
	CreateTodoList *command.CreateTodoList
}

type Queries struct {
	SignIn     *query.SignIn
	CheckToken *query.CheckToken
}

func New(commands *Commands, queries *Queries) *Application {
	return &Application{
		Commands: commands,
		Queries:  queries,
	}
}

func NewCommands(signup *command.SignUp, createTodoList *command.CreateTodoList) *Commands {
	return &Commands{
		SignUp:         signup,
		CreateTodoList: createTodoList,
	}
}

func NewQueries(signIn *query.SignIn, checkToken *query.CheckToken) *Queries {
	return &Queries{
		SignIn:     signIn,
		CheckToken: checkToken,
	}
}
