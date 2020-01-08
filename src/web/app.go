package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/sirikon/gonference/src/ioc"
	"github.com/sirikon/gonference/src/web/session"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"net/http"
)

// Server .
type Server struct {
	JobContext *ioc.JobContext
}

type WrappedHandler func(ctx *ioc.JobContext) ioc.Handler

func (s *Server) handle(wh func(ctx *ioc.JobContext) gin.HandlerFunc) gin.HandlerFunc {
	return func (c *gin.Context) {
		serviceProvider := c.MustGet("JobContext").(*ioc.JobContext)
		wh(serviceProvider)(c)
	}
}

func backofficeRoutes() gin.HandlerFunc {
	backofficeAssetsBox := packr.New("backoffice-assets", "./assets/backoffice")
	return func (c *gin.Context) {
		path := c.Params.ByName("filepath")
		if path == "/" {
			path = "/index.html"
		}
		data, err := backofficeAssetsBox.Find(path)
		if err != nil {
			_ = c.Error(err)
		}
		c.Data(http.StatusOK, "", data)
	}
}

func (s *Server) publicRoutes(r *gin.Engine) {
	publicAssetsBox := packr.New("public-assets", "./assets/public")

	r.GET("/", s.handle(ioc.IndexHandler))
	r.GET("/talk/:id", s.handle(ioc.TalkHandler))

	r.GET("/login", s.handle(ioc.LoginGetHandler))
	r.POST("/login", s.handle(ioc.LoginPostHandler))
	r.GET("/logout", s.handle(ioc.LoginLogoutHandler))

	r.StaticFS("/assets", publicAssetsBox)
}

func (s *Server) apiRoutes(r *gin.RouterGroup) {
	r.GET("/api/me", s.handle(ioc.MeAPIHandler))
	r.POST("/api/me/change-password", s.handle(ioc.MeAPIChangePasswordHandler))

	r.GET("/api/talks", s.handle(ioc.TalksAPIGetAllHandler))
	r.POST("/api/talks", s.handle(ioc.TalksAPIAddHandler))
	r.GET("/api/talks/:id", s.handle(ioc.TalksAPIGetHandler))
	r.PUT("/api/talks/:id", s.handle(ioc.TalksAPIUpdateHandler))
	r.DELETE("/api/talks/:id", s.handle(ioc.TalksAPIDeleteHandler))
}

// Run .
func (s *Server) Run(port string) error {
	log := logrus.New()
	store := cookie.NewStore([]byte("secret"))

	r := gin.New()
	r.Use(ginlogrus.Logger(log), gin.Recovery())
	r.Use(sessions.Sessions("gonference", store))

	r.Use(func(c *gin.Context) {
		scope := s.JobContext.CreateScope()
		c.Set("JobContext", scope)
	})

	secured := r.Group("")
	{
		secured.Use(func(ctx *gin.Context) {
			s := session.GetSession(ctx)
			if s.GetRole() != "admin" {
				ctx.Status(http.StatusForbidden)
				ctx.Abort()
			}
		})
		secured.GET("/admin/*filepath", backofficeRoutes())
		s.apiRoutes(secured)
	}

	s.publicRoutes(r)

	return r.Run("0.0.0.0:" + port)
}
