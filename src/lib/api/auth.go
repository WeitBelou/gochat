package api

import (
	"net/http"

	"lib/tokens"
	"lib/users"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(usersService users.Service, tokenService tokens.Service) gin.HandlerFunc {
	type request struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
		Nickname string `json:"nickname"`
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
		if req.Nickname == "" {
			req.Nickname = req.Login
		}

		user, err := usersService.Create(req.Login, req.Password, req.Nickname)
		if err == users.ErrUserExists {
			ctx.Error(validationErrorsList{
				"login": validationError{
					Error: "exists",
					Value: req.Login,
				},
			})
			return
		}
		if err != nil {
			ctx.Error(err)
			return
		}

		token, err := tokenService.GenerateToken(user.Login)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, response{
			Token: token,
		})
	}
}

func LoginHandler(usersService users.Service, tokenService tokens.Service) gin.HandlerFunc {
	type request struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
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

		user, err := usersService.CheckPassword(req.Login, req.Password)
		if err == users.ErrBadCredentials {
			ctx.Error(validationErrorsList{
				"login": validationError{
					Error: "bad_credentials",
					Value: req.Login,
				},
			})
			return
		}
		if err != nil {
			ctx.Error(err)
			return
		}

		token, err := tokenService.GenerateToken(user.Login)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, response{
			Token: token,
		})
	}
}
