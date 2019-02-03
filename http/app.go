package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirikon/gonference"
	log "github.com/sirupsen/logrus"
)

// Server .
type Server struct {
	TalkRepository gonference.TalkRepository
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	handleErr := func(err error) {
		log.Error(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	talks, err := s.TalkRepository.GetAll()
	if err != nil {
		handleErr(err)
		return
	}

	result, err := RenderTemplate("index", talks)
	if err != nil {
		handleErr(err)
		return
	}

	w.Write(result)
}

// Run .
func (s *Server) Run() error {
	router := httprouter.New()
	router.GET("/", s.indexHandler)
	router.ServeFiles("/assets/*filepath", http.Dir("./http/assets"))

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return err
	}
	return nil
}
