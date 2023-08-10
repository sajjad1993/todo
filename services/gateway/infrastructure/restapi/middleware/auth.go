package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sajjad1993/todo/pkg/rest"
	"github.com/sajjad1993/todo/services/gateway/app/query"
	"github.com/sajjad1993/todo/services/gateway/domain/auth"
	"github.com/sajjad1993/todo/services/gateway/infrastructure/restapi/presenter"
	"net/http"
	"strings"
)

const bearerPrefix = "Bearer "

// CheckToken just check token => save user in header if token exist
func CheckToken(authService *query.CheckToken) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			rest.FailedResponse(ctx, http.StatusUnauthorized, "auth required")
			return
		}
		token := auth.Token(strings.TrimPrefix(authHeader, bearerPrefix))
		tokenUser, err := authService.Run(ctx, &token)
		if err != nil {
			rest.FailedResponse(ctx, http.StatusUnauthorized, "auth required")
			return
		}
		if tokenUser != nil {
			ctx.Set(presenter.UserTokenKey, tokenUser)
			ctx.Next()
		} else {
			rest.FailedResponse(ctx, http.StatusUnauthorized, "auth required")
			return
		}

	}

}
