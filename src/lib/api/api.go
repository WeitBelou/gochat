package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Services struct {
	Auth authService
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
		authGroup.POST("/register", RegisterHandler(services.Auth))
		authGroup.POST("/login", LoginHandler())
		authGroup.POST("/logout", LogoutHandler())
	}

	profileGroup := v1.Group("/profile")
	{
		profileGroup.POST("/edit", ProfileEditHandler())
	}

	messagesGroup := v1.Group("/messages")
	{
		messagesGroup.GET("", MessagesListHandler())
		messagesGroup.POST("", MessagePostHandler())
	}
}