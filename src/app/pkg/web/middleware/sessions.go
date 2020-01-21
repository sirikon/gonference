package middleware

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
)

func Sessions(router *gin.Engine)  {
	store := cookie.NewStore([]byte(getCookieSecret()))
	router.Use(sessions.Sessions("gonference", store))
}

func getCookieSecret() string {
	envValue := os.Getenv("COOKIE_SECRET")
	if envValue != "" {
		return envValue
	}
	panic(errors.New("COOKIE_SECRET missing"))
}
