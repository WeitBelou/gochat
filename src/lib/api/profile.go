package api

import (
	"net/http"

	"lib/tokens"
	"lib/users"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func ProfileEditHandler(usersService users.Service) gin.HandlerFunc {
	type request struct {
		Nickname string `json:"nickname" binding:"required"`
	}

	return func(ctx *gin.Context) {
		req := &request{}
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			ctx.Error(err)
			return
		}

		user, ok := tokens.GetUserFromContext(ctx)
		if !ok {
			ctx.Error(errors.New("failed to fetch user from context"))
			return
		}

		err = usersService.ChangeNickname(user.Login, req.Nickname)
		if err == users.ErrUserNotExists {
			ctx.Error(errors.New("user not exists for auth required route"))
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
