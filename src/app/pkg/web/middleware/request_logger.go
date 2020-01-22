package middleware

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/ioc"
)

func RequestLogger(r *gin.Engine)  {
	r.Use(func(ctx *gin.Context) {
		jobContext := ctx.MustGet("JobContext").(*ioc.JobContext)
		log := ioc.LoggerForAccess(jobContext)
		go log.
			WithField("clientIP", ctx.ClientIP()).
			WithField("url", ctx.Request.URL.Path).
			Info()
	})
}
