package middleware

import (
	"github.com/gin-gonic/gin"
)

func RequestLogger(r *gin.Engine)  {
	/*r.Use(func(ctx *gin.Context) {
		jobContext := ctx.MustGet("JobContext").(*ioc.JobContext)
		log := ioc.LoggerForAccess(jobContext)
		ginlogrus.Logger(log)(ctx)
	})*/
}
