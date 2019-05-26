package public

import (
	"github.com/gin-gonic/gin"
	"github.com/sirikon/gonference/src/domain"
	"github.com/sirikon/gonference/src/web/session"
	"github.com/sirikon/gonference/src/web/templates"
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
		s.SetRole("admin")
		s.SetUsername(username)
		c.Redirect(http.StatusFound, "/admin/")
	} else {
		templates.ReplyTemplate(c, "login", nil)
	}
}

func (l *LoginController) LogoutHandler(c *gin.Context) {
	s := session.GetSession(c)
	s.SetRole("user")
	s.SetUsername("")
	c.Redirect(http.StatusFound, "/")
}
