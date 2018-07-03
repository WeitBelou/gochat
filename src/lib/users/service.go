package users

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var (
	ErrUserExists     = errors.New("user exists")
	ErrBadCredentials = errors.New("bad credentials")
)

type User struct {
	gorm.Model

	Login        string
	Nickname     string
	PasswordHash string
}

type Service interface {
	Register(login string, password string, nickname string) (*User, error)
	Login(login string, password string) (*User, error)
}
