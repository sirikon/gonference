package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/sirikon/gonference"
	apiControllers "github.com/sirikon/gonference/http/controllers/api"
	publicControllers "github.com/sirikon/gonference/http/controllers/public"
)

// Server .
type Server struct {
	TalkRepository gonference.TalkRepository
}

// Run .
func (s *Server) Run() error {

	indexController := publicControllers.IndexController{
		TalkRepository: s.TalkRepository,
	}

	talksController := apiControllers.TalksController{
		TalkRepository: s.TalkRepository,
	}

	router := httprouter.New()
	router.GET("/", indexController.Handler)
	router.ServeFiles("/assets/*filepath", http.Dir("./http/assets/public"))

	router.GET("/api/talks", talksController.GetAllHandler)
	router.ServeFiles("/admin/*filepath", http.Dir("./http/assets/backoffice"))

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return err
	}
	return nil
}
