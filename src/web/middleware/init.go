package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirikon/gonference/src/ioc"
)

func Init(r *gin.Engine, ctx *ioc.JobContext)  {
	r.Use(func(c *gin.Context) {
		scope := ctx.CreateScope()
		c.Set("JobContext", scope)
	})
}
