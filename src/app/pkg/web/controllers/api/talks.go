package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gonference/pkg/utils"
	"net/http"
	"path/filepath"
	"strconv"

	"gonference/pkg/domain"
)

// TalksAPIController .
type TalksAPIController struct {
	TalkRepository domain.TalkRepository
	QuestionRepository domain.QuestionRepository
}

// GetAllHandler .
func (s *TalksAPIController) GetAllHandler(ctx *gin.Context) {
	talks := s.TalkRepository.GetAll()
	ctx.JSON(http.StatusOK, talks)
}

// AddHandler .
func (s *TalksAPIController) AddHandler(ctx *gin.Context) {
	var vm AddTalkViewModel
	utils.Check(ctx.Bind(&vm))
	talk := vm.ToDomainTalk()
	talk.ID = createUUID()
	s.TalkRepository.Add(talk)
	updateSpeakerImageIfPresent(talk.ID, ctx)
	ctx.Status(http.StatusOK)
}

// GetHandler .
func (s *TalksAPIController) GetHandler(ctx *gin.Context) {
	talkID := ctx.Params.ByName("id")
	if talkID == "" {
		ctx.AbortWithStatus(404)
		return
	}

	talk := s.TalkRepository.Get(talkID)

	ctx.JSON(http.StatusOK, talk)
}

// UpdateHandler .
func (s *TalksAPIController) UpdateHandler(ctx *gin.Context) {
	talkID := ctx.Params.ByName("id")
	if talkID == "" {
		ctx.AbortWithStatus(404)
		return
	}

	var vm UpdateTalkViewModel
	utils.Check(ctx.Bind(&vm))

	updateSpeakerImageIfPresent(talkID, ctx)

	vm.ID = talkID
	talk := vm.ToDomainTalk()

	s.TalkRepository.Update(talk)

	ctx.Status(http.StatusOK)
}

// DeleteHandler .
func (s *TalksAPIController) DeleteHandler(ctx *gin.Context) {
	talkID := ctx.Params.ByName("id")
	if talkID == "" {
		ctx.AbortWithStatus(404)
		return
	}

	s.TalkRepository.Delete(talkID)

	ctx.Status(http.StatusOK)
}

func (s *TalksAPIController) GetTalkQuestionsHandler(ctx *gin.Context) {
	talkIDStr := ctx.Params.ByName("id")
	if talkIDStr == "" {
		ctx.AbortWithStatus(404)
		return
	}

	talkID, err := strconv.Atoi(talkIDStr); utils.Check(err)
	questions := s.QuestionRepository.GetByTalkId(talkID)

	ctx.JSON(http.StatusOK, questions)
}

func updateSpeakerImageIfPresent(talkID string, ctx *gin.Context) {
	if len(ctx.Request.MultipartForm.File) == 0 { return }
	file, err := ctx.FormFile("speakerImage"); utils.Check(err)
	ext := filepath.Ext(file.Filename)
	utils.Check(ctx.SaveUploadedFile(file, "uploads/talk-" + talkID + "-speaker-image" + ext))
}

func createUUID() string {
	return uuid.New().String()
}
