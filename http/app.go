package http

import (
	"net/http"

	"github.com/sirikon/gonference/ioc"

	"github.com/julienschmidt/httprouter"
)

// Server .
type Server struct {
	ServiceProvider *ioc.ServiceProvider
}

// Run .
func (s *Server) Run() error {

	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		s.ServiceProvider.CreateRequestScope().GetIndexController().Handler(w, r, ps)
	})
	router.ServeFiles("/assets/*filepath", http.Dir("./http/assets/public"))

	router.GET("/api/talks", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		s.ServiceProvider.CreateRequestScope().GetTalksAPIController().GetAllHandler(w, r, ps)
	})
	router.POST("/api/talks", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		s.ServiceProvider.CreateRequestScope().GetTalksAPIController().AddHandler(w, r, ps)
	})

	router.ServeFiles("/admin/*filepath", http.Dir("./http/assets/backoffice"))

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return err
	}
	return nil
}
