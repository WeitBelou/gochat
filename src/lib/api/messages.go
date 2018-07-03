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
		messagesList, err := service.List()
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, response{
			Messages: messagesList,
		})
	}
}

func MessagePostHandler(service messages.Service) gin.HandlerFunc {
	type request struct {
		Text string `json:"text" binding:"required,len(256)"`
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

		// TODO(i.kosolapov): Replace login with nickname
		err = service.Post(user.Login, req.Text)
		if err != nil {
			ctx.Error(errors.Wrap(err, "failed to post new message"))
		}

		ctx.JSON(http.StatusOK, gin.H{})
	}
}
