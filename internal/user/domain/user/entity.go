package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uint
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	HashedPassword string `json:"hashed_password"`
}

// HashPassword Hashes the given password and store in it
func (e *User) HashPassword() error {
	if len(e.Password) == 0 {
		//todo build an error package .
		return errors.New("password cant be nil")
	}
	b, err := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	e.HashedPassword = string(b)
	return nil
}

// VerifyPassword compares the given password with the user's PasswordHash and checks if they match.
func (e *User) VerifyPassword() error {
	return bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(e.HashedPassword))
}
