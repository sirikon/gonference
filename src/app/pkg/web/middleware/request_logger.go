package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/toorop/gin-logrus"
	"gonference/pkg/ioc"
)

func RequestLogger(r *gin.Engine)  {
	r.Use(func(ctx *gin.Context) {
		jobContext := ctx.MustGet("JobContext").(*ioc.JobContext)
		log := ioc.LoggerForAccess(jobContext)
		ginlogrus.Logger(log)(ctx)
	})
}
