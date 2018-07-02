package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundHandler() gin.HandlerFunc {
	type response struct {
		Error string `json:"error"`
	}

	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, response{
			Error: "not_found",
		})
	}
}

func ErrorsMiddleware() gin.HandlerFunc {
	type response struct {
		Error string `json:"error"`
	}

	return func(ctx *gin.Context) {
		ctx.Next()

		if ctx.Errors != nil {
			ctx.JSON(http.StatusInternalServerError, response{
				Error: "internal",
			})
		}
	}
}
