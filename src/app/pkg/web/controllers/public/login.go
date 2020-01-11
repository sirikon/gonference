package public

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/domain"
	"gonference/pkg/web/session"
	"gonference/pkg/web/templates"
	"net/http"
)

// LoginController .
type LoginController struct {
	UserService domain.UserService
}

// GetHandler .
func (l *LoginController) GetHandler(c *gin.Context) {
	templates.ReplyTemplate(c, "login", nil)
}

func (l *LoginController) PostHandler(c *gin.Context) {
	s := session.GetSession(c)

	username := c.PostForm("username")
	password := c.PostForm("password")

	if !l.UserService.UserExists(username) {
		templates.ReplyTemplate(c, "login", nil)
		return
	}

	result, err := l.UserService.CheckPassword(username, password)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if result {
		s.Set(session.RoleKey, "admin")
		s.Set(session.UsernameKey, username)
		c.Redirect(http.StatusFound, "/admin/")
	} else {
		templates.ReplyTemplate(c, "login", nil)
	}
}

func (l *LoginController) LogoutHandler(c *gin.Context) {
	s := session.GetSession(c)
	s.Set(session.RoleKey, "user")
	s.Set(session.UsernameKey, "")
	c.Redirect(http.StatusFound, "/")
}