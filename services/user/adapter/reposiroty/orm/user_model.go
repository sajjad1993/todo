package orm

import (
	"github.com/sajjad1993/todo/services/user/domain/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"Index;not null;size:32"`
	Email    string `json:"email" gorm:"uniqueIndex;not null;size:64"`
	Password string `json:"password" gorm:"not null"`
}

func (u *User) toEntity() *user.User {
	if u == nil || u.ID == 0 {
		return nil
	}
	return &user.User{
		ID:             u.ID,
		Name:           u.Name,
		Email:          u.Email,
		HashedPassword: u.Password,
	}
}

func (u *User) fromEntity(e *user.User) *User {
	if u == nil || e == nil {
		return nil
	}
	u.ID = e.ID
	u.Name = e.Name
	u.Email = e.Email
	u.Password = e.HashedPassword
	return u
}
