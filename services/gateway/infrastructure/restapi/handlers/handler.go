package handlers

import (
	"github.com/sajjad1993/todo/services/gateway/adapter/controller"
	"github.com/sajjad1993/todo/services/gateway/app"
)

type Handler struct {
	application *app.Application
	controller  *controller.Commands
}

func NewHandler(application *app.Application, controller *controller.Commands) *Handler {
	return &Handler{application: application, controller: controller}
}
