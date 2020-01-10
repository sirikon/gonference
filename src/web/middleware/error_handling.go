package middleware

import "github.com/gin-gonic/gin"

func ErrorHandling(r *gin.Engine)  {
	r.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.String(500, "boom: " + err.(error).Error())
			}
		}()
		ctx.Next()
	})
}
