package common

import (
	"errors"
	"github.com/sajjad1993/todo/pkg/errs"
	"google.golang.org/grpc/codes"
)

func GetGrpcStatusCodeByError(err error) codes.Code {
	if errors.Is(err, errs.ErrValidation) {
		return codes.InvalidArgument
	}
	if errors.Is(err, errs.ErrUnauthorized) {
		return codes.PermissionDenied
	}
	if errors.Is(err, errs.ErrNotFoundError) {
		return codes.NotFound
	}
	if errors.Is(err, errs.ErrNoSuchKey) {
		return codes.NotFound
	}
	return codes.Internal

}
