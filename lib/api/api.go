package api // import "gochat/lib/api"

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine) {
	r.NoRoute(NotFoundHandler())

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(ErrorsMiddleware())

	v1 := r.Group("/api/v1")

	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/register", RegisterHandler())
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
