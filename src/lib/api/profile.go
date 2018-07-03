package api

import (
	"net/http"

	"lib/tokens"
	"lib/users"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func ProfileEditHandler(usersService users.Service, tokensService tokens.Service) gin.HandlerFunc {
	type request struct {
		Nickname string `json:"nickname" binding:"required"`
	}

	type response struct {
		Token string `json:"token"`
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

		u, err := usersService.ChangeNickname(user.Subject, req.Nickname)
		if err == users.ErrUserNotExists {
			ctx.Error(errors.New("user not exists for auth required route"))
			return
		}

		token, err := tokensService.GenerateToken(u.Login, u.Nickname)
		if err != nil {
			ctx.Error(errors.New("failed to generate new token"))
			return
		}

		ctx.JSON(http.StatusOK, response{
			Token: token,
		})
	}
}
