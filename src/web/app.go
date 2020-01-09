package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/sirikon/gonference/src/ioc"
	"github.com/sirikon/gonference/src/web/middleware"
	"net/http"
)

// Server .
type Server struct {
	JobContext *ioc.JobContext
}

// Run .
func (s *Server) Run(port string) error {
	router := gin.New()
	middleware.Init(router, s.JobContext)
	middleware.RequestLogger(router)
	middleware.Sessions(router)

	s.adminRoutes(router.Group(""))
	s.publicRoutes(router)

	return router.Run("0.0.0.0:" + port)
}

func (s *Server) publicRoutes(r *gin.Engine) {
	publicAssetsBox := packr.New("public-assets", "./assets/public")

	r.GET("/", handle(ioc.IndexHandler))
	r.GET("/talk/:id", handle(ioc.TalkHandler))

	r.GET("/login", handle(ioc.LoginGetHandler))
	r.POST("/login", handle(ioc.LoginPostHandler))
	r.GET("/logout", handle(ioc.LoginLogoutHandler))

	r.StaticFS("/assets", publicAssetsBox)
}

func (s *Server) adminRoutes(r *gin.RouterGroup) {
	middleware.RequireAuthRole(r, "admin")
	r.GET("/admin/*filepath", backofficeAssets())

	r.GET("/api/me", handle(ioc.MeAPIHandler))
	r.POST("/api/me/change-password", handle(ioc.MeAPIChangePasswordHandler))

	r.GET("/api/talks", handle(ioc.TalksAPIGetAllHandler))
	r.POST("/api/talks", handle(ioc.TalksAPIAddHandler))
	r.GET("/api/talks/:id", handle(ioc.TalksAPIGetHandler))
	r.PUT("/api/talks/:id", handle(ioc.TalksAPIUpdateHandler))
	r.DELETE("/api/talks/:id", handle(ioc.TalksAPIDeleteHandler))
}

func backofficeAssets() gin.HandlerFunc {
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

func handle(wh func(ctx *ioc.JobContext) gin.HandlerFunc) gin.HandlerFunc {
	return func (c *gin.Context) {
		jobContext := c.MustGet("JobContext").(*ioc.JobContext)
		wh(jobContext)(c)
	}
}
