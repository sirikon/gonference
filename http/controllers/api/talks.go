package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirikon/gonference"
	log "github.com/sirupsen/logrus"
)

// TalksAPIController .
type TalksAPIController struct {
	TalkRepository gonference.TalkRepository
}

// GetAllHandler .
func (s *TalksAPIController) GetAllHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	handleErr := func(err error) {
		log.Error(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	talks, err := s.TalkRepository.GetAll()
	if err != nil {
		handleErr(err)
		return
	}

	result, err := json.Marshal(talks)
	if err != nil {
		handleErr(err)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(result)
}
