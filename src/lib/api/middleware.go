package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

func NotFoundHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not_found"})
	}
}

type validationError struct {
	Error string      `json:"error"`
	Value interface{} `json:"value"`
}

type validationErrorsList map[string]validationError

func (v validationErrorsList) Error() string {
	return fmt.Sprintf("validation error: %+v", map[string]validationError(v))
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) != 0 {
			validation := make(validationErrorsList)

			for _, err := range ctx.Errors {
				if vErrs, ok := err.Err.(validator.ValidationErrors); ok {
					for _, v := range vErrs {
						name := v.Namespace()
						if strings.Contains(name, ".") {
							parts := strings.SplitN(name, ".", 2)
							name = parts[1]
						}
						validation[name] = validationError{
							Error: v.Tag(),
							Value: v.Value(),
						}
					}
				} else if vErrs, ok := err.Err.(validationErrorsList); ok {
					for k, v := range vErrs {
						validation[k] = v
					}
				} else if err.Err == io.EOF {
					validation = validationErrorsList{
						"body": validationError{
							Error: "required",
							Value: nil,
						},
					}
					break
				} else if strings.Contains(err.Error(), "error found") {
					validation = validationErrorsList{
						"body": validationError{
							Error: "invalid",
							Value: "<hidden>",
						},
					}
					break
				} else {
					log.Printf("internal error: %+v", err)
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
					return
				}
			}
			ctx.JSON(http.StatusBadRequest, validation)
		}
	}
}
