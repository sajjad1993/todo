package handlers

import "github.com/sajjad1993/todo/services/gateway/app"

type Handler struct {
	application *app.Application
}

func NewHandler(application *app.Application) *Handler {
	return &Handler{application: application}
}
