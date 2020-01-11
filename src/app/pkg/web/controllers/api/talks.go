package api

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"gonference/pkg/domain"
)

// TalksAPIController .
type TalksAPIController struct {
	TalkRepository domain.TalkRepository
}

// GetAllHandler .
func (s *TalksAPIController) GetAllHandler(ctx *gin.Context) {
	handleErr := func(err error) {
		_ = ctx.Error(err)
	}

	talks, err := s.TalkRepository.GetAll()
	if err != nil {
		handleErr(err)
		return
	}

	ctx.JSON(http.StatusOK, talks)
}

// AddHandler .
func (s *TalksAPIController) AddHandler(ctx *gin.Context) {
	handleErr := func(err error) {
		_ = ctx.Error(err)
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

	ctx.Status(http.StatusOK)
}

// GetHandler .
func (s *TalksAPIController) GetHandler(ctx *gin.Context) {
	handleErr := func(err error) {
		_ = ctx.Error(err)
	}

	talkIDStr := ctx.Params.ByName("id")
	if talkIDStr == "" {
		handleErr(errors.New("talk id missing"))
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

	ctx.JSON(http.StatusOK, talk)
}

// UpdateHandler .
func (s *TalksAPIController) UpdateHandler(ctx *gin.Context) {
	handleErr := func(err error) {
		_ = ctx.Error(err)
	}

	talkIDStr := ctx.Params.ByName("id")
	if talkIDStr == "" {
		handleErr(errors.New("talk id missing"))
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

	ctx.Status(http.StatusOK)
}

// DeleteHandler .
func (s *TalksAPIController) DeleteHandler(ctx *gin.Context) {
	handleErr := func(err error) {
		_ = ctx.Error(err)
	}

	talkIDStr := ctx.Params.ByName("id")
	if talkIDStr == "" {
		handleErr(errors.New("talk id missing"))
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

	ctx.Status(http.StatusOK)
}
