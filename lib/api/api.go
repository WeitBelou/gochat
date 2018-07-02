package api // import "gochat/lib/api"

import "github.com/gin-gonic/gin"

func Register(r gin.IRouter) {
	v1 := r.Group("/api/v1")

	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/register", func(ctx *gin.Context) { panic("not implemented") })
		authGroup.POST("/login", func(ctx *gin.Context) { panic("not implemented") })
		authGroup.POST("/logout", func(ctx *gin.Context) { panic("not implemented") })
	}

	userGroup := v1.Group("/profile")
	{
		userGroup.POST("/edit", func(ctx *gin.Context) { panic("not implemented") })
	}

	messagesGroup := v1.Group("/messages")
	{
		messagesGroup.GET("", func(ctx *gin.Context) { panic("not implemented") })
		messagesGroup.POST("", func(ctx *gin.Context) { panic("not implemented") })
	}
}
