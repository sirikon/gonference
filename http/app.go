package http

import (
	"encoding/json"
	"net/http"

	"github.com/sirikon/gonference"
)

// Server .
type Server struct {
	TalkRepository gonference.TalkRepository
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	talks, err := s.TalkRepository.GetAll()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	message, err := json.Marshal(talks)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	w.Write(message)
}

// Run .
func (s *Server) Run() error {
	http.HandleFunc("/", s.indexHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return err
	}
	return nil
}
