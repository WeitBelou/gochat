package tokens

import "github.com/gin-gonic/gin"

type User struct {
	Login string
}

type Service interface {
	GenerateToken(login string) (string, error)
	CheckToken(token string) (*User, bool)
}

const userKey = "userKey"

func GetUserFromContext(ctx *gin.Context) (*User, bool) {
	u, ok := ctx.Value(userKey).(*User)
	return u, ok
}

func PutUserToContext(ctx *gin.Context, u *User) {
	ctx.Set(userKey, u)
}
