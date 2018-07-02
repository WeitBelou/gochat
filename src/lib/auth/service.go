package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var (
	ErrUserExists = errors.New("user exists")
)

type User struct {
	gorm.Model

	Login        string
	Nickname     string
	PasswordHash string
}

type Service interface {
	Register(login string, password string, nickname string) (token string, err error)
	Login(login string, password string) (token string, err error)
	Logout(token string) error
}
