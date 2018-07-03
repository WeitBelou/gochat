package api

import (
	"net/http"

	"lib/messages"
	"lib/tokens"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func MessagesListHandler(service messages.Service) gin.HandlerFunc {
	type response struct {
		Messages []messages.Message `json:"messages"`
	}
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, response{
			Messages: service.List(),
		})
	}
}

func MessagePostHandler(service messages.Service) gin.HandlerFunc {
	type request struct {
		Text string `json:"text" binding:"required"`
	}

	return func(ctx *gin.Context) {
		req := &request{}
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			ctx.Error(err)
			return
		}

		if len(req.Text) > 255 {
			ctx.Error(validationErrorsList{
				"text.length": validationError{
					Error: "text too long (> 255)",
					Value: len(req.Text),
				},
			})
			return
		}

		user, ok := tokens.GetUserFromContext(ctx)
		if !ok {
			ctx.Error(errors.New("failed to fetch user from context"))
			return
		}

		// TODO(i.kosolapov): Replace login with nickname
		service.Post(user.Login, req.Text)

		ctx.JSON(http.StatusOK, gin.H{})
	}
}
