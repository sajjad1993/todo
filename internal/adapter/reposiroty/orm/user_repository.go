package orm

import (
	"context"
	"errors"
	"github.com/sajjad1993/todo/internal/domain/user"
	"github.com/sajjad1993/todo/pkg/errs"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db: db}
}

func (r *userRepository) GetById(ctx context.Context, id uint) (*user.User, error) {
	model := new(User)
	err := r.db.WithContext(ctx).Where("id = ?", id).First(model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError(err.Error())
		}
		return nil, errs.NewInternalError(err.Error())
	}
	return model.toEntity(), nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	model := new(User)
	err := r.db.WithContext(ctx).Where("email = ?", email).First(model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError(err.Error())
		}
		return nil, errs.NewInternalError(err.Error())
	}
	return model.toEntity(), nil
}

func (r *userRepository) Create(ctx context.Context, userEnt *user.User) error {
	model := new(User).fromEntity(userEnt)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errs.NewDuplicateEntity(err.Error())
		}
		return errs.NewInternalError(err.Error())
	}
	return nil
}
