package middleware

import (
	"github.com/gin-gonic/gin"
)

func Measure(r *gin.Engine)  {
	r.Use(func(ctx *gin.Context) {
		//start := time.Now()
		ctx.Next()
		//elapsed := time.Since(start)
		//fmt.Printf("Request to %s took %s\n", ctx.Request.URL.String(), elapsed)
	})
}
