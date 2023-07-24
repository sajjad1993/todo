package handlers

import (
	"errors"
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
