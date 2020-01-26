package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gonference/pkg/infrastructure/config"
)

func Sessions(router *gin.Engine)  {
	store := cookie.NewStore([]byte(getCookieSecret()))
	router.Use(sessions.Sessions("gonference", store))
}

func getCookieSecret() string {
	return config.Config.Web.CookieSecret
}
