package app

import (
	"context"
	"fmt"
	"github.com/sajjad1993/todo/pkg/errs"
	"github.com/sajjad1993/todo/services/auth/app/jwt"
	"github.com/sajjad1993/todo/services/auth/domain/token"
	"github.com/sajjad1993/todo/services/auth/domain/user"
)

// App  is the auth use case implementation
type App struct {
	userRepo  user.Repository
	JWTAccess jwt.JWT
}

// NewService returns a pointer to auth service, and implements the auth use case
func NewService(userRepo user.Repository, access jwt.JWT) UseCase {
	return &App{userRepo: userRepo, JWTAccess: access}
}

// SignIn validate password and create token for user
func (s *App) SignIn(ctx context.Context, user *user.User) (string, error) {
	entity, err := s.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}
	entity.Password = user.Password
	err = entity.VerifyPassword()
	if err != nil {
		return "", err
	}
	userToken := &token.User{
		Name:  entity.Name,
		Email: entity.Email,
		ID:    fmt.Sprintf("%d", entity.ID),
	}
	access, err := s.JWTAccess.GenerateToken(userToken)
	if err != nil {
		return "", errs.NewInternalError(fmt.Sprintf("token generation : %s ", err))
	}
	return access, nil

}
