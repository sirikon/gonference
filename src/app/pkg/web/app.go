package web

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/assets"
	"gonference/pkg/ioc"
	"gonference/pkg/utils"
	"gonference/pkg/web/middleware"
	"net/http"
)

// Server .
type Server struct {
	JobContext *ioc.JobContext
}

// Run .
func (s *Server) Run(port string) error {
	router := gin.New()
	middleware.Sessions(router)
	middleware.Init(router, s.JobContext)
	middleware.ErrorHandling(router)
	middleware.RequestLogger(router)

	s.adminRoutes(router.Group(""))
	s.publicRoutes(router)

	return router.Run("0.0.0.0:" + port)
}

func (s *Server) publicRoutes(r *gin.Engine) {
	r.Use(customStatic)
	r.GET("/", handle(ioc.IndexHandler))
	r.GET("/talk/:slug", handle(ioc.TalkHandler))
	r.POST("/talk/:slug/rating", handle(ioc.TalkPostRatingHandler))
	r.POST("/talk/:slug/question", handle(ioc.TalkPostQuestionHandler))

	r.GET("/login", handle(ioc.LoginGetHandler))
	r.POST("/login", handle(ioc.LoginPostHandler))
	r.GET("/logout", handle(ioc.LoginLogoutHandler))

	r.StaticFS("/assets", assets.FrontStyle)
	r.StaticFS("/uploads", http.Dir("uploads/"))
}

func (s *Server) adminRoutes(r *gin.RouterGroup) {
	middleware.RequireAuthRole(r, "admin")
	r.GET("/admin/*filepath", backofficeAssets())

	api := r.Group("/api")
	middleware.APIErrorHandling(api)
	api.GET("/me", handle(ioc.MeAPIHandler))
	api.POST("/me/change-password", handle(ioc.MeAPIChangePasswordHandler))

	api.GET("/talks", handle(ioc.TalksAPIGetAllHandler))
	api.POST("/talks", handle(ioc.TalksAPIAddHandler))
	api.GET("/talks/:id", handle(ioc.TalksAPIGetHandler))
	api.PUT("/talks/:id", handle(ioc.TalksAPIUpdateHandler))
	api.DELETE("/talks/:id", handle(ioc.TalksAPIDeleteHandler))
}

func backofficeAssets() gin.HandlerFunc {
	return func (c *gin.Context) {
		path := c.Params.ByName("filepath")
		if path == "/" {
			path = "/index.html"
		}
		data, err := assets.BackofficeUI.Find(path)
		if err != nil {
			_ = c.Error(err)
		}
		c.Data(http.StatusOK, "", data)
	}
}

func customStatic(ctx *gin.Context)  {
	filePath := "custom/static" + ctx.Request.URL.String()
	if utils.FileExists(filePath) {
		ctx.File(filePath)
		ctx.Abort()
	}
}

func handle(wh func(ctx *ioc.JobContext) gin.HandlerFunc) gin.HandlerFunc {
	return func (c *gin.Context) {
		jobContext := c.MustGet("JobContext").(*ioc.JobContext)
		wh(jobContext)(c)
	}
}
