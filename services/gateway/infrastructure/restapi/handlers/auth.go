package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sajjad1993/todo/pkg/rest"
	"github.com/sajjad1993/todo/services/gateway/domain/user"
	"github.com/sajjad1993/todo/services/gateway/infrastructure/restapi/presenter/request"
	"github.com/sajjad1993/todo/services/gateway/infrastructure/restapi/presenter/response"
	"net/http"
)

func (h *Handler) SignIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.Login
		if err := ctx.ShouldBindJSON(&req); err != nil {
			rest.FailedResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		userEnt := &user.User{
			Password: req.Password,
			Email:    req.Email,
		}
		token, err := h.application.Queries.SignIn.Run(ctx, userEnt)
		if err != nil {
			rest.FailedResponse(ctx, getStatusCodeByError(err), err.Error())
			return
		}
		rest.SuccessResponse(ctx, "", response.Login{Token: token})
	}
}
