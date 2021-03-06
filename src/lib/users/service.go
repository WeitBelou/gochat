package users

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var (
	ErrUserExists    = errors.New("user exists")
	ErrUserNotExists = errors.New("user not exists")
)

type User struct {
	gorm.Model

	Login        string
	Nickname     string
	PasswordHash string
}

type Service interface {
	Create(login string, password string, nickname string) (*User, error)
	ChangeNickname(login string, nickname string) (*User, error)
	CheckPassword(login string, password string) (*User, error)
}
