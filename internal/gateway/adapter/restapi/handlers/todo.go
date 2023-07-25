package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajjad1993/todo/internal/gateway/adapter/restapi/presenter/request"
	"github.com/sajjad1993/todo/internal/gateway/domain/todo"
	"github.com/sajjad1993/todo/pkg/rest"
	"net/http"
)

func (h *Handler) CreateTodoList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.TodoList
		if err := ctx.ShouldBindJSON(&req); err != nil {
			rest.FailedResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		token, err := getUserToken(ctx)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}
		todoListEnt := &todo.List{
			Name:        req.Name,
			Description: req.Description,
			UserID:      token.ID,
		}
		err = h.application.Commands.CreateTodoList.Execute(ctx, todoListEnt)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}
		rest.OKResponse(ctx)
	}
}
