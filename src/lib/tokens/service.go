package tokens

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type User struct {
	jwt.StandardClaims
	Nickname string `json:"nickname"`
}

type Service interface {
	GenerateToken(login string, nickname string) (string, error)
	CheckToken(token string) (*User, bool)
	CheckOneTimeToken(token string) (*User, bool)
	GenerateOneTimeToken(login string, nickname string) string
}

const userKey = "userKey"

func GetUserFromContext(ctx *gin.Context) (*User, bool) {
	u, ok := ctx.Value(userKey).(*User)
	return u, ok
}

func PutUserToContext(ctx *gin.Context, u *User) {
	ctx.Set(userKey, u)
}
