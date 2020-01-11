package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Sessions(router *gin.Engine)  {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("gonference", store))
}
