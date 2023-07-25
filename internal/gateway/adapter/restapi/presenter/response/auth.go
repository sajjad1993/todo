package response

import "github.com/sajjad1993/todo/internal/gateway/domain/auth"

type Login struct {
	Token *auth.Token `json:"token"`
}
