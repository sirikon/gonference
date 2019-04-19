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

// AddTalkViewModel .
type AddTalkViewModel struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	SpeakerName  string `json:"speakerName"`
	SpeakerTitle string `json:"speakerTitle"`
	Track        string `json:"track"`
	When         string `json:"when"`
}

// ToDomainTalk .
func (vm AddTalkViewModel) ToDomainTalk() gonference.Talk {
	return gonference.Talk{
		Name:         vm.Name,
		Description:  vm.Description,
		SpeakerName:  vm.SpeakerName,
		SpeakerTitle: vm.SpeakerTitle,
		Track:        vm.Track,
		When:         vm.When,
	}
}

// AddHandler .
func (s *TalksAPIController) AddHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	handleErr := func(err error) {
		log.Error(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	decoder := json.NewDecoder(r.Body)
	var vm AddTalkViewModel
	err := decoder.Decode(&vm)
	if err != nil {
		handleErr(err)
		return
	}

	talk := vm.ToDomainTalk()

	err = s.TalkRepository.Add(talk)
	if err != nil {
		handleErr(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
