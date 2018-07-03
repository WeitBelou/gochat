package api

import (
	"lib/tokens"
	"lib/users"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Services struct {
	Auth   users.Service
	Tokens tokens.Service
}

func Register(r *gin.Engine, services Services) {
	binding.Validator = NewValidator()

	r.NoRoute(NotFoundHandler())

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(ErrorMiddleware())

	v1 := r.Group("/api/v1")

	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/register", RegisterHandler(services.Auth, services.Tokens))
		authGroup.POST("/login", LoginHandler(services.Auth, services.Tokens))
	}

	profileGroup := v1.Group("/profile", AuthMiddleware(services.Tokens))
	{
		profileGroup.POST("/edit", ProfileEditHandler())
	}

	messagesGroup := v1.Group("/messages", AuthMiddleware(services.Tokens))
	{
		messagesGroup.GET("", MessagesListHandler())
		messagesGroup.POST("", MessagePostHandler())
	}
}
