package api

import (
	"net/http"

	"lib/messages"
	"lib/tokens"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

		err = service.Post(user.Nickname, req.Text)
		if err != nil {
			ctx.Error(errors.Wrap(err, "failed to post message"))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func MessageListWebsocketHandler(service messages.Service) gin.HandlerFunc {
	upgrader := websocket.Upgrader{
		WriteBufferSize: 1024,
	}

	return func(ctx *gin.Context) {
		user, ok := tokens.GetUserFromContext(ctx)
		if !ok {
			ctx.Error(errors.New("failed to get user from context"))
			return
		}

		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.Error(errors.Wrap(err, "failed to get ws connection"))
			return
		}

		service.AddWSClient(user.Subject, conn)
	}
}
