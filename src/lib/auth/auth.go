package auth

import (
	"log"
	"time"

	"lib/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	Secret config.Secret
	DB     config.DB
}

type UsersDB struct {
	cfg  Config
	conn *gorm.DB
}

func New(c Config) (*UsersDB, error) {
	db, err := gorm.Open("postgres", c.DB.ToPostgresDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to open connect to db")
	}

	err = db.AutoMigrate(new(User)).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to sync db")
	}

	return &UsersDB{
		cfg:  c,
		conn: db,
	}, nil
}

func (db *UsersDB) Register(login string, password string, nickname string) (string, error) {
	u := &User{
		Login:        login,
		Nickname:     nickname,
		PasswordHash: db.getPasswordHash(password),
	}

	scope := db.conn.First(&User{Login: login})
	if scope.RecordNotFound() {
		db.conn.Create(u)
		if err := db.conn.Error; err != nil {
			return "", errors.Wrap(err, "failed to create user")
		}
	} else {
		if err := scope.Error; err != nil {
			return "", errors.Wrap(err, "failed to check if user exists")
		}
		return "", ErrUserExists
	}

	token, err := db.getToken(u)
	if err != nil {
		return "", errors.Wrap(err, "failed to return token on registration")
	}
	return token, nil
}

func (db *UsersDB) Login(login string, password string) (string, error) {
	u := &User{}

	scope := db.conn.First(u, "login = ?", login)
	if scope.RecordNotFound() {
		return "", ErrBadCredentials
	} else if err := scope.Error; err != nil {
		return "", errors.Wrap(err, "failed to check if user exists")
	}

	passwordOk := db.checkPasswordHash(u.PasswordHash, password)
	if !passwordOk {
		return "", ErrBadCredentials
	}

	token, err := db.getToken(u)
	if err != nil {
		return "", errors.Wrap(err, "failed to return token on login")
	}
	return token, nil
}

func (db *UsersDB) getToken(u *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   u.Login,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenString, err := token.SignedString([]byte(db.cfg.Secret))
	if err != nil {
		return "", errors.Wrap(err, "failed to get token")
	}
	return tokenString, nil
}

func (db *UsersDB) getPasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panicf("failed to generate password hash: %+v", err)
	}
	return string(hash)
}

func (db *UsersDB) checkPasswordHash(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
