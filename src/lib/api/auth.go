package api

import (
	"net/http"

	"lib/auth"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(service auth.Service) gin.HandlerFunc {
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

		token, err := service.Register(req.Login, req.Password, req.Nickname)
		if err == auth.ErrUserExists {
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

		ctx.JSON(http.StatusOK, response{
			Token: token,
		})
	}
}

func LoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		panic("not implemented")
	}
}

func LogoutHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		panic("not implemented")
	}
}
