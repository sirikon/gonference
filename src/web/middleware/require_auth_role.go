package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirikon/gonference/src/web/session"
	"net/http"
)

func RequireAuthRole(router *gin.RouterGroup, requiredRole string) {
	router.Use(func(ctx *gin.Context) {
		s := session.GetSession(ctx)
		if s.Get(session.RoleKey) != requiredRole {
			ctx.Status(http.StatusForbidden)
			ctx.Abort()
		}
	})
}
