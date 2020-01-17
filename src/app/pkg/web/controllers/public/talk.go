package public

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/domain"
	"gonference/pkg/utils"
	"gonference/pkg/web/session"
	"gonference/pkg/web/templates"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// TalkController .
type TalkController struct {
	TalkRepository domain.TalkRepository
	RatingRepository domain.RatingRepository
	QuestionRepository domain.QuestionRepository
}

// Handler .
func (s *TalkController) Handler(c *gin.Context) {
	visitorKey := session.GetSession(c).Get(session.VisitorKey)
	id, err := strconv.Atoi(c.Param("id")); utils.Check(err)
	talk, err := s.TalkRepository.Get(id); utils.Check(err)
	rating := s.RatingRepository.GetByTalkIdAndVisitorKey(id, visitorKey)
	speakerImageFileName := getSpeakerImageFileName(id)
	templates.ReplyTemplate(c, "talk", gin.H{
		"talk": talk,
		"speakerImageFileName": speakerImageFileName,
		"ratingDone": rating != nil,
		"questionReceived": c.Query("q") == "1",
	})
}

func (s *TalkController) PostRatingHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id")); utils.Check(err)
	visitorKey := session.GetSession(ctx).Get(session.VisitorKey)
	var vm AddRatingViewModel
	utils.Check(ctx.Bind(&vm))

	rating := domain.Rating{
		ID:         0,
		TalkID:     id,
		VisitorKey: visitorKey,
		Stars:      vm.Stars,
		Comment:    vm.Comment,
	}
	s.RatingRepository.Add(rating)
	ctx.Redirect(http.StatusFound, "/talk/" + strconv.Itoa(id))
}

func (s *TalkController) PostQuestionHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id")); utils.Check(err)
	visitorKey := session.GetSession(ctx).Get(session.VisitorKey)
	var vm AddQuestionViewModel
	utils.Check(ctx.Bind(&vm))

	question := domain.Question{
		ID:         0,
		TalkID:     id,
		VisitorKey: visitorKey,
		Question:   vm.Question,
	}
	s.QuestionRepository.Add(question)
	ctx.Redirect(http.StatusFound, "/talk/" + strconv.Itoa(id) + "?q=1")
}

func getSpeakerImageFileName(talkID int) string {
	var result string
	utils.Check(filepath.Walk("uploads/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), "talk-" + strconv.Itoa(talkID) + "-speaker-image") {
			result = "/uploads/" + info.Name()
		}
		return nil
	}))
	return result
}
