package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"github.com/sajjad1993/todo/pkg/rest"
	"github.com/sajjad1993/todo/services/gateway/domain/todo"
	"github.com/sajjad1993/todo/services/gateway/infrastructure/restapi/presenter/request"
	"github.com/sajjad1993/todo/services/gateway/infrastructure/restapi/presenter/response"
	"net/http"
	"strconv"
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
		err = h.controller.CreateTodoList(ctx, todoListEnt)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}
		rest.OKResponse(ctx)
	}
}

func (h *Handler) ListTodoList() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token, err := getUserToken(ctx)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}
		todos, err := h.application.Queries.ListToDoList.Run(ctx, token.ID)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}

		var result response.TodoLists
		rest.GeneralResponse(ctx, http.StatusOK, true, "",
			response.ListToDOListResponse{Lists: result.FromEntity(todos)}, nil)
	}
}

func (h *Handler) UpdateTodoList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.TodoList
		if err := ctx.ShouldBindJSON(&req); err != nil {
			rest.FailedResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		param := ctx.Param("id")
		todoListId, err := strconv.Atoi(param)

		if err != nil {
			rest.FailedResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		token, err := getUserToken(ctx)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}
		todoListEnt := &todo.List{
			ID:          uint(todoListId),
			Name:        req.Name,
			Description: req.Description,
			UserID:      token.ID,
		}
		commandMessage := command_utils.NewCommandMessage("", command_utils.SuccessStatus, todoListEnt)
		commandChanel := h.application.Commands.UpdateTodoList.Execute(ctx, commandMessage)
		select {
		case <-ctx.Done():
			rest.FailedResponse(ctx, http.StatusGatewayTimeout, "")
		case message := <-commandChanel:
			err = message.GetError()
			if err != nil {
				rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
				return
			}
			rest.OKResponse(ctx)
		}
	}
}
func (h *Handler) DeleteTodoList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("id")
		todoListId, err := strconv.Atoi(param)

		if err != nil {
			rest.FailedResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		token, err := getUserToken(ctx)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}
		err = h.controller.DeleteTodoList(ctx, uint(todoListId), token.ID)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}
		rest.OKResponse(ctx)
	}
}

func (h *Handler) CreateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.Todo
		if err := ctx.ShouldBindJSON(&req); err != nil {
			rest.FailedResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		token, err := getUserToken(ctx)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}
		todoEnt := &todo.Item{
			Title:    req.Title,
			Priority: req.Priority,
			ListId:   req.ListID,
			UserId:   token.ID,
		}
		err = h.controller.CreateTodoItem(ctx, todoEnt)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}
		rest.OKResponse(ctx)
	}
}

func (h *Handler) UpdateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.UpdateTodo
		if err := ctx.ShouldBindJSON(&req); err != nil {
			rest.FailedResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		param := ctx.Param("id")
		id, err := strconv.Atoi(param)

		if err != nil {
			rest.FailedResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		token, err := getUserToken(ctx)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}
		todoEnt := &todo.Item{
			ID:       uint(id),
			Title:    req.Title,
			Priority: req.Priority,
			UserId:   token.ID,
		}
		commandMessage := command_utils.NewCommandMessage("", command_utils.SuccessStatus, todoEnt)
		commandChanel := h.application.Commands.UpdateTodo.Execute(ctx, commandMessage)
		select {
		case <-ctx.Done():
			rest.FailedResponse(ctx, http.StatusGatewayTimeout, "")
		case message := <-commandChanel:
			err := message.GetError()
			if err != nil {
				rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
				return
			}
			rest.OKResponse(ctx)
		}
	}
}
func (h *Handler) DeleteTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("id")
		todoListId, err := strconv.Atoi(param)

		if err != nil {
			rest.FailedResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		token, err := getUserToken(ctx)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}

		err = h.controller.DeleteTodoList(ctx, uint(todoListId), token.ID)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}
		rest.OKResponse(ctx)
	}
}
