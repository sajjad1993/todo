package jwt

import "fmt"

type TokenInValid struct {
	message string
}

func (err *TokenInValid) Error() string {
	return fmt.Sprintf("token is invalid because %s", err.message)
}
