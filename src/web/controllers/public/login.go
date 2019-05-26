package public

import (
	"github.com/gin-gonic/gin"
	"github.com/sirikon/gonference/src/web/session"
	"github.com/sirikon/gonference/src/web/templates"
	"net/http"
)

// LoginController .
type LoginController struct {}

// GetHandler .
func (l *LoginController) GetHandler(c *gin.Context) {
	templates.ReplyTemplate(c, "login", nil)
}

func (l *LoginController) PostHandler(c *gin.Context) {
	s := session.GetSession(c)

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "admin" && password == "admin" {
		s.SetRole("admin")
		c.Redirect(http.StatusFound, "/admin/")
	} else {
		templates.ReplyTemplate(c, "login", nil)
	}
}

func (l *LoginController) LogoutHandler(c *gin.Context) {
	s := session.GetSession(c)
	s.SetRole("user")
	c.Redirect(http.StatusFound, "/")
}
