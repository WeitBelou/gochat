package users

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type DB struct {
	conn *gorm.DB
}

func New(c Config) (*DB, error) {
	db, err := gorm.Open("postgres", c.DB.ToPostgresDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to open connect to db")
	}

	err = db.AutoMigrate(new(User)).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to sync db")
	}

	return &DB{
		conn: db,
	}, nil
}

func (db *DB) Register(login string, password string, nickname string) (*User, error) {
	u := &User{
		Login:        login,
		Nickname:     nickname,
		PasswordHash: getPasswordHash(password),
	}

	scope := db.conn.First(&User{Login: login})
	if scope.RecordNotFound() {
		db.conn.Create(u)
		if err := db.conn.Error; err != nil {
			return nil, errors.Wrap(err, "failed to create user")
		}
	} else {
		if err := scope.Error; err != nil {
			return nil, errors.Wrap(err, "failed to check if user exists")
		}
		return nil, ErrUserExists
	}

	return u, nil
}

func (db *DB) Login(login string, password string) (*User, error) {
	u := &User{}

	scope := db.conn.First(u, "login = ?", login)
	if scope.RecordNotFound() {
		return nil, ErrBadCredentials
	} else if err := scope.Error; err != nil {
		return nil, errors.Wrap(err, "failed to check if user exists")
	}

	passwordOk := checkPasswordHash(u.PasswordHash, password)
	if !passwordOk {
		return nil, ErrBadCredentials
	}
	return u, nil
}

func getPasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panicf("failed to generate password hash: %+v", err)
	}
	return string(hash)
}

func checkPasswordHash(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
