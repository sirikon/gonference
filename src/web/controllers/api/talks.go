package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/sirikon/gonference/src/domain"
	"github.com/sirikon/gonference/src/web/middleware"
	log "github.com/sirupsen/logrus"
)

// TalksAPIController .
type TalksAPIController struct {
	TalkRepository domain.TalkRepository
}

// GetAllHandler .
func (s *TalksAPIController) GetAllHandler(ctx *middleware.RequestContext) {
	handleErr := func(err error) {
		log.Error(err)
		http.Error(ctx.ResponseWritter, "Something went wrong", http.StatusInternalServerError)
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

	ctx.ResponseWritter.Header().Add("content-type", "application/json")
	ctx.ResponseWritter.Write(result)
}

// AddHandler .
func (s *TalksAPIController) AddHandler(ctx *middleware.RequestContext) {
	handleErr := func(err error) {
		log.Error(err)
		http.Error(ctx.ResponseWritter, "Something went wrong", http.StatusInternalServerError)
	}

	decoder := json.NewDecoder(ctx.Request.Body)
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

	ctx.ResponseWritter.WriteHeader(http.StatusOK)
}

// GetHandler .
func (s *TalksAPIController) GetHandler(ctx *middleware.RequestContext) {
	handleErr := func(err error) {
		log.Error(err)
		http.Error(ctx.ResponseWritter, "Something went wrong", http.StatusInternalServerError)
	}

	talkIDStr := ctx.Params.ByName("id")
	if talkIDStr == "" {
		handleErr(errors.New("Talk id missing"))
		return
	}

	talkID, err := strconv.Atoi(talkIDStr)
	if err != nil {
		handleErr(err)
		return
	}

	talk, err := s.TalkRepository.Get(talkID)
	if err != nil {
		handleErr(err)
		return
	}

	result, err := json.Marshal(talk)
	if err != nil {
		handleErr(err)
		return
	}

	ctx.ResponseWritter.Header().Add("content-type", "application/json")
	ctx.ResponseWritter.Write(result)
}

// UpdateHandler .
func (s *TalksAPIController) UpdateHandler(ctx *middleware.RequestContext) {
	handleErr := func(err error) {
		log.Error(err)
		http.Error(ctx.ResponseWritter, "Something went wrong", http.StatusInternalServerError)
	}

	talkIDStr := ctx.Params.ByName("id")
	if talkIDStr == "" {
		handleErr(errors.New("Talk id missing"))
		return
	}

	talkID, err := strconv.Atoi(talkIDStr)
	if err != nil {
		handleErr(err)
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	var vm UpdateTalkViewModel
	err = decoder.Decode(&vm)
	if err != nil {
		handleErr(err)
		return
	}

	vm.ID = talkID

	talk := vm.ToDomainTalk()

	err = s.TalkRepository.Update(talk)
	if err != nil {
		handleErr(err)
		return
	}

	ctx.ResponseWritter.WriteHeader(http.StatusOK)
}

// DeleteHandler .
func (s *TalksAPIController) DeleteHandler(ctx *middleware.RequestContext) {
	handleErr := func(err error) {
		log.Error(err)
		http.Error(ctx.ResponseWritter, "Something went wrong", http.StatusInternalServerError)
	}

	talkIDStr := ctx.Params.ByName("id")
	if talkIDStr == "" {
		handleErr(errors.New("Talk id missing"))
		return
	}

	talkID, err := strconv.Atoi(talkIDStr)
	if err != nil {
		handleErr(err)
		return
	}

	err = s.TalkRepository.Delete(talkID)
	if err != nil {
		handleErr(err)
		return
	}

	ctx.ResponseWritter.WriteHeader(http.StatusOK)
}
