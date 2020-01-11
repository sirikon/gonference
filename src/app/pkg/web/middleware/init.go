package middleware

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/ioc"
)

func Init(r *gin.Engine, ctx *ioc.JobContext)  {
	r.Use(func(c *gin.Context) {
		scope := ctx.CreateScope()
		c.Set("JobContext", scope)
	})
}
