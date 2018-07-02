package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type authService interface {
	Register(login string, password string, nickname string) (exists bool, token string, err error)
	Login(login string, password string) (exists bool, token string, err error)

	Logout(token string) (exists bool, err error)
}

func RegisterHandler(service authService) gin.HandlerFunc {
	type request struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
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

		exists, token, err := service.Register(req.Login, req.Password, req.Nickname)
		if err != nil {
			ctx.Error(err)
			return
		}
		if exists {
			ctx.Error(validationErrorsList{
				"login": validationError{
					Error: "exists",
					Value: req.Login,
				},
			})
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
