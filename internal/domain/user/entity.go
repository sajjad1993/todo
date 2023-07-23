package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Entity struct {
	ID       int
	Name     string
	Email    string
	Password string
}

// SetPassword Hashes the given password and store in it
func (e *Entity) SetPassword(password string) error {
	if len(password) == 0 {
		//todo build an error package .
		return errors.New("password cant be nil")
	}
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	e.Password = string(b)
	return nil
}

// VerifyPassword compares the given password with the user's PasswordHash and checks if they match.
func (e *Entity) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password))
}
