package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajjad1993/todo/internal/gateway/adapter/restapi/presenter/request"
	"github.com/sajjad1993/todo/internal/gateway/domain/user"
	"github.com/sajjad1993/todo/pkg/rest"
	"net/http"
)

func (h *Handler) SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.SignUp
		if err := ctx.ShouldBindJSON(&req); err != nil {
			rest.FailedResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		userEnt := &user.User{
			Name:     req.Name,
			Password: req.Password,
			Email:    req.Email,
		}
		err := h.application.Commands.SignUp.Execute(ctx, userEnt)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}
		rest.OKResponse(ctx)
	}
}
