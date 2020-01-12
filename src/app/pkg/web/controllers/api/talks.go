package api

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/utils"
	"net/http"
	"path/filepath"
	"strconv"

	"gonference/pkg/domain"
)

// TalksAPIController .
type TalksAPIController struct {
	TalkRepository domain.TalkRepository
}

// GetAllHandler .
func (s *TalksAPIController) GetAllHandler(ctx *gin.Context) {
	talks, err := s.TalkRepository.GetAll(); utils.Check(err)
	ctx.JSON(http.StatusOK, talks)
}

// AddHandler .
func (s *TalksAPIController) AddHandler(ctx *gin.Context) {
	var vm AddTalkViewModel
	utils.Check(ctx.Bind(&vm))
	talk := vm.ToDomainTalk()
	utils.Check(s.TalkRepository.Add(talk))
	ctx.Status(http.StatusOK)
}

// GetHandler .
func (s *TalksAPIController) GetHandler(ctx *gin.Context) {
	talkIDStr := ctx.Params.ByName("id")
	if talkIDStr == "" {
		ctx.AbortWithStatus(404)
		return
	}

	talkID, err := strconv.Atoi(talkIDStr); utils.Check(err)
	talk, err := s.TalkRepository.Get(talkID); utils.Check(err)

	ctx.JSON(http.StatusOK, talk)
}

// UpdateHandler .
func (s *TalksAPIController) UpdateHandler(ctx *gin.Context) {
	talkIDStr := ctx.Params.ByName("id")
	if talkIDStr == "" {
		ctx.AbortWithStatus(404)
		return
	}

	talkID, err := strconv.Atoi(talkIDStr); utils.Check(err)

	var vm UpdateTalkViewModel
	utils.Check(ctx.Bind(&vm))

	updateSpeakerImageIfPresent(talkID, ctx)

	vm.ID = talkID
	talk := vm.ToDomainTalk()

	utils.Check(s.TalkRepository.Update(talk))

	ctx.Status(http.StatusOK)
}

// DeleteHandler .
func (s *TalksAPIController) DeleteHandler(ctx *gin.Context) {
	talkIDStr := ctx.Params.ByName("id")
	if talkIDStr == "" {
		ctx.AbortWithStatus(404)
		return
	}

	talkID, err := strconv.Atoi(talkIDStr); utils.Check(err)
	utils.Check(s.TalkRepository.Delete(talkID))

	ctx.Status(http.StatusOK)
}

func updateSpeakerImageIfPresent(talkID int, ctx *gin.Context) {
	if len(ctx.Request.MultipartForm.File) == 0 { return }
	file, err := ctx.FormFile("speakerImage"); utils.Check(err)
	ext := filepath.Ext(file.Filename)
	utils.Check(ctx.SaveUploadedFile(file, "uploads/talk-" + strconv.Itoa(talkID) + "-speaker-image" + ext))
}
