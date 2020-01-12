package middleware

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/web/controllers/api"
)

func ErrorHandling(r *gin.Engine)  {
	r.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				code, message := getResponseForError(err)
				ctx.AbortWithStatusJSON(code, gin.H{"message": message})
			}
		}()
		ctx.Next()
	})
}

func getResponseForError(err interface{}) (int, string) {
	var code int
	var message string

	switch v := err.(type) {
	case api.UserError:
		code = 400
		message = v.Message
	case error:
		code = 500
		message = v.Error()
	}

	return code, message
}
