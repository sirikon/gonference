package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// MeAPIController .
type MeAPIController struct {
}

// Handler .
func (s *MeAPIController) Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)

	user := User{
		Username: session.Get("username").(string),
	}

	ctx.JSON(http.StatusOK, user)
}
