package tokens

import (
	"math/rand"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type JWT struct {
	mu            sync.Mutex
	oneTimeTokens map[string]*User

	oneTimeTokensTTL time.Duration
	secret           string
}

func (t *JWT) CheckOneTimeToken(token string) (*User, bool) {
	t.mu.Lock()
	defer t.mu.Unlock()

	u, ok := t.oneTimeTokens[token]
	if ok {
		delete(t.oneTimeTokens, token)
		return u, true
	}

	return nil, false
}

func (t *JWT) GenerateOneTimeToken(login string, nickname string) string {
	t.mu.Lock()
	defer t.mu.Unlock()

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	tokenBytes := make([]byte, 10)
	for i := range tokenBytes {
		tokenBytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	token := string(tokenBytes)

	t.oneTimeTokens[token] = &User{
		StandardClaims: jwt.StandardClaims{
			Subject: login,
		},
		Nickname: nickname,
	}
	go func() {
		<-time.After(t.oneTimeTokensTTL)
		t.mu.Lock()
		delete(t.oneTimeTokens, token)
		t.mu.Unlock()
	}()

	return token
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

func New(cfg Config) *JWT {
	return &JWT{
		oneTimeTokens: make(map[string]*User),

		secret:           string(cfg.Secret),
		oneTimeTokensTTL: cfg.OneTimeTokensTTL,
	}
}
