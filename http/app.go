package http

import (
	"net/http"

	"github.com/sirikon/gonference"
	log "github.com/sirupsen/logrus"
)

// Server .
type Server struct {
	TalkRepository gonference.TalkRepository
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
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
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./http/assets"))))
	http.HandleFunc("/", s.indexHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return err
	}
	return nil
}
