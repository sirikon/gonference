package http

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/sirikon/gonference/src/ioc"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Server .
type Server struct {
	ServiceProvider *ioc.ServiceProvider
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
		w.Write(data)
	})
}

func (s *Server) publicRoutes(router *httprouter.Router) {
	publicAssetsBox := packr.New("public-assets", "./assets/public")

	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		s.ServiceProvider.CreateRequestScope().GetIndexController().Handler(w, r, ps)
	})
	router.ServeFiles("/assets/*filepath", publicAssetsBox)
}

func (s *Server) apiRoutes(router *httprouter.Router) {
	router.GET("/api/talks", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		s.ServiceProvider.CreateRequestScope().GetTalksAPIController().GetAllHandler(w, r, ps)
	})
	router.POST("/api/talks", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		s.ServiceProvider.CreateRequestScope().GetTalksAPIController().AddHandler(w, r, ps)
	})
	router.GET("/api/talks/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		s.ServiceProvider.CreateRequestScope().GetTalksAPIController().GetHandler(w, r, ps)
	})
	router.PUT("/api/talks/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		s.ServiceProvider.CreateRequestScope().GetTalksAPIController().UpdateHandler(w, r, ps)
	})
	router.DELETE("/api/talks/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		s.ServiceProvider.CreateRequestScope().GetTalksAPIController().DeleteHandler(w, r, ps)
	})
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
