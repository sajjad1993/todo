package handlers

import (
	"github.com/sajjad1993/todo/services/gateway/adapter/controller/commands"
	"github.com/sajjad1993/todo/services/gateway/adapter/controller/queries"
	"github.com/sajjad1993/todo/services/gateway/app"
)

type Handler struct {
	application       *app.Application
	commandController *commands.Commands
	queryController   *queries.Queries
}

func NewHandler(application *app.Application, commandController *commands.Commands, queryController *queries.Queries) *Handler {
	return &Handler{application: application,
		commandController: commandController,
		queryController:   queryController}
}
