package app

import (
	"github.com/sajjad1993/todo/internal/gateway/app/command"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SignUp *command.SignUp
}

type Queries struct {
}

func New(signUp *command.SignUp) *Application {
	return &Application{
		Commands: Commands{
			SignUp: signUp,
		},
	}
}
