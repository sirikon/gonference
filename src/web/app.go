package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/sirikon/gonference/src/ioc"
	"net/http"
)

// Server .
type Server struct {
	ServiceProvider *ioc.ServiceProvider
}

type WrappedHandler func(scope *ioc.ServiceProvider) gin.HandlerFunc

func (s *Server) wrapHandler(wh WrappedHandler) gin.HandlerFunc {
	return func (c *gin.Context) {
		serviceProvider := c.MustGet("ServiceProvider").(*ioc.ServiceProvider)
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

	r.GET("/", s.wrapHandler(func (sp *ioc.ServiceProvider) gin.HandlerFunc {
		return sp.GetIndexController().Handler
	}))

	r.StaticFS("/assets", publicAssetsBox)
}

func (s *Server) apiRoutes(r *gin.Engine) {
	r.GET("/api/me", s.wrapHandler(func(sp *ioc.ServiceProvider) gin.HandlerFunc {
		return sp.GetMeAPIController().Handler
	}))

	r.GET("/api/talks", s.wrapHandler(func(s *ioc.ServiceProvider) gin.HandlerFunc {
		return s.GetTalksAPIController().GetAllHandler
	}))
	r.POST("/api/talks", s.wrapHandler(func(s *ioc.ServiceProvider) gin.HandlerFunc {
		return s.GetTalksAPIController().AddHandler
	}))
	r.GET("/api/talks/:id", s.wrapHandler(func(s *ioc.ServiceProvider) gin.HandlerFunc {
		return s.GetTalksAPIController().GetHandler
	}))
	r.PUT("/api/talks/:id", s.wrapHandler(func(s *ioc.ServiceProvider) gin.HandlerFunc {
		return s.GetTalksAPIController().UpdateHandler
	}))
	r.DELETE("/api/talks/:id", s.wrapHandler(func(s *ioc.ServiceProvider) gin.HandlerFunc {
		return s.GetTalksAPIController().DeleteHandler
	}))
}

// Run .
func (s *Server) Run(port string) error {
	r := gin.New()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("gonference", store))

	r.Use(func(c *gin.Context) {
		scope := s.ServiceProvider.CreateScope()
		c.Set("ServiceProvider", scope)
		session := sessions.Default(c)
		if session.Get("username") == nil {
			session.Set("username", "user")
			_ = session.Save()
		}
	})

	r.GET("/admin/*filepath", backofficeRoutes())

	s.publicRoutes(r)
	s.apiRoutes(r)

	return r.Run("0.0.0.0:" + port)
}

//
//// WrappedHandler .
//type WrappedHandler func(scope *ioc.ServiceProvider) middleware.RequestHandler
//
//// WrapHandler .
//func (s *Server) WrapHandler(wh WrappedHandler) httprouter.Handle {
//	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//		scope := s.ServiceProvider.CreateRequestScope()
//		logger := scope.GetLogger()
//
//		logger.WithFields(log.Fields{
//			"url":    r.URL,
//			"method": r.Method,
//		}).Info("Request started")
//
//		session := auth.EnsureSession(logger, r, w)
//
//		start := time.Now()
//
//		context := &middleware.RequestContext{
//			ResponseWritter: w,
//			Request:         r,
//			Params:          ps,
//			Session:         session,
//		}
//
//		wh(scope)(context)
//
//		elapsed := time.Since(start)
//
//		logger.WithFields(log.Fields{
//			"elapsed": elapsed,
//		}).Info("Request finished")
//	}
//}
//
//
//
