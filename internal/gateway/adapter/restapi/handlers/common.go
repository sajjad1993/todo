package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sajjad1993/todo/internal/gateway/adapter/restapi/presenter"
	"github.com/sajjad1993/todo/internal/gateway/domain/user"
	"github.com/sajjad1993/todo/pkg/errs"
	"net/http"
)

func getStatusCodeByError(err error) int {
	if errors.Is(err, errs.ErrValidation) {
		return http.StatusBadRequest
	}
	if errors.Is(err, errs.ErrUnauthorized) {
		return http.StatusUnauthorized
	}
	if errors.Is(err, errs.ErrNotFoundError) {
		return http.StatusNotFound
	}
	if errors.Is(err, errs.ErrNoSuchKey) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError

}

func getUserToken(ctx *gin.Context) (*user.User, error) {
	value, ok := ctx.Get(presenter.UserTokenKey)
	if !ok {
		return nil, errs.NewUnauthorizedError("user")
	}
	ent, ok := value.(*user.User)
	if !ok {
		return nil, errs.NewUnauthorizedError("user")

	}
	return ent, nil
}
