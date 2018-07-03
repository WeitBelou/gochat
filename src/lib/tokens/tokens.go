package tokens

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type JWT struct {
	secret string
}

func New(cfg Config) *JWT {
	return &JWT{
		secret: string(cfg.Secret),
	}
}

func (t *JWT) GenerateToken(login string, nickname string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, User{
		StandardClaims: jwt.StandardClaims{
			Subject:   login,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
		Nickname: nickname,
	})
	tokenString, err := token.SignedString([]byte(t.secret))
	if err != nil {
		return "", errors.Wrap(err, "failed to get token")
	}
	return tokenString, nil
}

func (t *JWT) CheckToken(token string) (*User, bool) {
	parsedToken, err := jwt.ParseWithClaims(token, &User{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.secret), nil
	})
	if err != nil {
		return &User{}, false
	}
	if claims, ok := parsedToken.Claims.(*User); ok && parsedToken.Valid {
		return claims, true
	}

	return nil, false
}
