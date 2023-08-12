package errs

import (
	"errors"
	"fmt"
)

var ErrValidation = errors.New("validation error")
var ErrUnauthorized = errors.New("unauthorized error")
var ErrNoSuchKey = errors.New("key not found")
var ErrNotFoundError = errors.New("not found error")
var ErrInternalError = errors.New("internal error")
var ErrDuplicateEntity = errors.New("5|duplicate entity")
var ErrTimeOut = errors.New("5|time out")

func NewValidationError(errMsg string) error {
	return fmt.Errorf("%w|%s", ErrValidation, errMsg)
}

func NewUnauthorizedError(errMsg string) error {
	return fmt.Errorf("%w|%s", ErrUnauthorized, errMsg)
}

func NewNoSuchKeyError(errMsg string) error {
	return fmt.Errorf("%w|%s", ErrNoSuchKey, errMsg)
}

func NewNotFoundError(errMsg string) error {
	return fmt.Errorf("%w|%s", ErrNotFoundError, errMsg)
}
func NewInternalError(errMsg string) error {
	return fmt.Errorf("%w|%s", ErrInternalError, errMsg)
}
func NewDuplicateEntity(errMsg string) error {
	return fmt.Errorf("%w|%s", ErrDuplicateEntity, errMsg)
}
func NewTimeOut(errMsg string) error {
	return fmt.Errorf("%w|%s", ErrTimeOut, errMsg)
}
