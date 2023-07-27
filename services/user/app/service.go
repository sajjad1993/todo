package app

import (
	"context"
	"github.com/sajjad1993/todo/services/user/domain/user"
)

// service is the auth use case implementation
type service struct {
	userRepo user.Repository
}

func (s *service) GetUser(ctx context.Context, email string) (*user.User, error) {
	return s.userRepo.GetByEmail(ctx, email)
}

// NewService returns a pointer to auth service, and implements the auth use case
func NewService(userRepo user.Repository) UseCase {
	return &service{userRepo: userRepo}
}

// SignUp creates a new user in user repository
func (s *service) SignUp(ctx context.Context, user *user.User) error {
	err := user.HashPassword()
	if err != nil {
		return err
	}
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil

}
