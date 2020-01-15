package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gonference/pkg/ioc"
	"gonference/pkg/web/session"
)

func Init(r *gin.Engine, jobContext *ioc.JobContext)  {
	r.Use(func(ctx *gin.Context) {
		// Ensure Visitor Key on session
		s := session.GetSession(ctx)
		visitorKey := s.Get(session.VisitorKey)
		if visitorKey == "" {
			visitorKey = uuid.New().String()
			s.Set(session.VisitorKey, visitorKey)
		}

		// Create JobContext
		scope := jobContext.CreateScope(uuid.New().String(), visitorKey)
		ctx.Set("JobContext", scope)
	})
}
