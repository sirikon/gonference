package middleware

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/web/session"
	"net/http"
)

func RequireAuthRole(router *gin.RouterGroup, requiredRole string) {
	router.Use(func(ctx *gin.Context) {
		s := session.GetSession(ctx)
		if s.Get(session.RoleKey) != requiredRole {
			ctx.Redirect(http.StatusFound, "/login")
			ctx.Abort()
		}
	})
}
