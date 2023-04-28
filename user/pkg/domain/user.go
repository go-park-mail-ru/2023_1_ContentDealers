package domain

import (
	"errors"
)

var (
	ErrWrongCredentials  = errors.New("wrong credentials")
	ErrNotValidEmail     = errors.New("email is not valid")
	ErrNotValidPassword  = errors.New("password is not valid")
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type User struct {
	ID           uint64
	Email        string
	PasswordHash string
	AvatarURL    string
}
