package handlers

import (
	"github.com/sajjad1993/todo/services/gateway/adapter/controller/commands"
	"github.com/sajjad1993/todo/services/gateway/app"
)

type Handler struct {
	application *app.Application
	controller  *commands.Commands
}

func NewHandler(application *app.Application, controller *commands.Commands) *Handler {
	return &Handler{application: application, controller: controller}
}
