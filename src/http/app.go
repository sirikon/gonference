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

// Run .
func (s *Server) Run(port string) error {

	publicAssetsBox := packr.New("public-assets", "./assets/public")
	backofficeAssetsBox := packr.New("backoffice-assets", "./assets/backoffice")

	router := httprouter.New()

	/* Administration */
	router.GET("/admin/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data, err := backofficeAssetsBox.Find("index.html")
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
		w.Write(data)
	})
	router.ServeFiles("/admin/assets/*filepath", backofficeAssetsBox)

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

	/* Public */
	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		s.ServiceProvider.CreateRequestScope().GetIndexController().Handler(w, r, ps)
	})
	router.ServeFiles("/assets/*filepath", publicAssetsBox)

	log.Println("HTTP server listening on port " + port + ".")
	err := http.ListenAndServe("0.0.0.0:"+port, router)
	if err != nil {
		return err
	}
	return nil
}
