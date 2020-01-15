package middleware

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/web/models"
	"runtime/debug"
)

func ErrorHandling(r *gin.Engine)  {
	errorHandling(&r.RouterGroup, func(ctx *gin.Context, code int, message string) {
		ctx.String(code, "Unexpected error: " + message)
	})
}

func APIErrorHandling(r *gin.RouterGroup)  {
	errorHandling(r, func(ctx *gin.Context, code int, message string) {
		ctx.AbortWithStatusJSON(code, gin.H{"message": message})
	})
}

func errorHandling(r *gin.RouterGroup, replier func(ctx *gin.Context, code int, message string))  {
	r.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				code, message := getResponseForError(err)
				replier(ctx, code, message)
			}
		}()
		ctx.Next()
	})
}

func getResponseForError(err interface{}) (int, string) {
	var code int
	var message string

	switch v := err.(type) {
	case models.UserError:
		code = 400
		message = v.Message
	case error:
		code = 500
		message = v.Error() + "\n" + string(debug.Stack())
	}

	return code, message
}
