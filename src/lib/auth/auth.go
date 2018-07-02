package auth

import (
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
	passwordHash, err := db.getPasswordHash(password)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate password hash")
	}
	u := &User{
		Login:        login,
		Nickname:     nickname,
		PasswordHash: passwordHash,
	}

	scope := db.conn.First(&User{Login: login})
	if scope.RecordNotFound() {
		db.conn.Create(u)
		if db.conn.Error != nil {
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

func (*UsersDB) Login(login string, password string) (string, error) {
	panic("implement me")
}

func (*UsersDB) Logout(token string) error {
	panic("implement me")
}

func (db *UsersDB) getToken(u *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject: u.Login,
	})
	tokenString, err := token.SignedString([]byte(db.cfg.Secret))
	if err != nil {
		return "", errors.Wrap(err, "failed to get token")
	}
	return tokenString, nil
}

func (db *UsersDB) getPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate password hash")
	}
	return string(hash), nil
}
