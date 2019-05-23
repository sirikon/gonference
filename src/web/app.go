package web

import (
	"net/http"
	"time"

	"github.com/gobuffalo/packr/v2"
	"github.com/sirikon/gonference/src/ioc"
	"github.com/sirikon/gonference/src/web/auth"
	"github.com/sirikon/gonference/src/web/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

// Server .
type Server struct {
	ServiceProvider *ioc.ServiceProvider
}

// WrappedHandler .
type WrappedHandler func(scope *ioc.ServiceProvider) middleware.RequestHandler

// WrapHandler .
func (s *Server) WrapHandler(wh WrappedHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		scope := s.ServiceProvider.CreateRequestScope()
		logger := scope.GetLogger()

		logger.WithFields(log.Fields{
			"url":    r.URL,
			"method": r.Method,
		}).Info("Request started")

		session := auth.EnsureSession(logger, r, w)

		start := time.Now()

		context := &middleware.RequestContext{
			ResponseWritter: w,
			Request:         r,
			Params:          ps,
			Session:         session,
		}

		wh(scope)(context)

		elapsed := time.Since(start)

		logger.WithFields(log.Fields{
			"elapsed": elapsed,
		}).Info("Request finished")
	}
}

func (s *Server) backofficeRoutes(router *httprouter.Router) {
	backofficeAssetsBox := packr.New("backoffice-assets", "./assets/backoffice")

	router.GET("/admin/*filepath", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		path := ps.ByName("filepath")
		if path == "/" {
			path = "/index.html"
		}
		data, err := backofficeAssetsBox.Find(path)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
		_, _ = w.Write(data)
	})
}

func (s *Server) publicRoutes(router *httprouter.Router) {
	publicAssetsBox := packr.New("public-assets", "./assets/public")

	router.GET("/", s.WrapHandler(func(s *ioc.ServiceProvider) middleware.RequestHandler {
		return s.GetIndexController().Handler
	}))

	router.ServeFiles("/assets/*filepath", publicAssetsBox)
}

func (s *Server) apiRoutes(router *httprouter.Router) {
	router.GET("/api/me", s.WrapHandler(func(s *ioc.ServiceProvider) middleware.RequestHandler {
		return s.GetMeAPIController().Handler
	}))

	router.GET("/api/talks", s.WrapHandler(func(s *ioc.ServiceProvider) middleware.RequestHandler {
		return s.GetTalksAPIController().GetAllHandler
	}))
	router.POST("/api/talks", s.WrapHandler(func(s *ioc.ServiceProvider) middleware.RequestHandler {
		return s.GetTalksAPIController().AddHandler
	}))
	router.GET("/api/talks/:id", s.WrapHandler(func(s *ioc.ServiceProvider) middleware.RequestHandler {
		return s.GetTalksAPIController().GetHandler
	}))
	router.PUT("/api/talks/:id", s.WrapHandler(func(s *ioc.ServiceProvider) middleware.RequestHandler {
		return s.GetTalksAPIController().UpdateHandler
	}))
	router.DELETE("/api/talks/:id", s.WrapHandler(func(s *ioc.ServiceProvider) middleware.RequestHandler {
		return s.GetTalksAPIController().DeleteHandler
	}))
}

// Run .
func (s *Server) Run(port string) error {
	router := httprouter.New()

	s.backofficeRoutes(router)
	s.apiRoutes(router)
	s.publicRoutes(router)

	log.Println("HTTP server listening on port " + port + ".")
	err := http.ListenAndServe("0.0.0.0:"+port, router)
	if err != nil {
		return err
	}
	return nil
}
