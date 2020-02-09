package public

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gonference/pkg/application"
	"gonference/pkg/infrastructure/web/session"
	"gonference/pkg/infrastructure/web/templates"
	"gonference/pkg/utils"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// TalkController .
type TalkController struct {
	TalkRepository     application.TalkRepository
	RatingRepository   application.RatingRepository
	QuestionRepository application.QuestionRepository
}

// Handler .
func (s *TalkController) Handler(c *gin.Context) {
	visitorKey := session.GetSession(c).Get(session.VisitorKey)
	slug := c.Param("slug")
	talk := s.TalkRepository.GetBySlug(slug)
	rating := s.RatingRepository.GetByTalkIdAndVisitorKey(talk.ID, visitorKey)
	speakerImageFileName := getSpeakerImageFileName(talk.ID)
	templates.ReplyTemplate(c, "talk", gin.H{
		"talk": talk,
		"speakerImageFileName": speakerImageFileName,
		"ratingDone": rating != nil,
		"questionReceived": c.Query("q") == "1",
	})
}

func (s *TalkController) PostRatingHandler(ctx *gin.Context) {
	slug := ctx.Param("slug")
	visitorKey := session.GetSession(ctx).Get(session.VisitorKey)
	var vm AddRatingViewModel
	utils.Check(ctx.Bind(&vm))

	if vm.Stars < 1 || vm.Stars > 5 {
		ctx.AbortWithStatus(400)
		return
	}

	talk := s.TalkRepository.GetBySlug(slug)
	rating := &application.Rating{
		ID:         createUUID(),
		TalkID:     talk.ID,
		VisitorKey: visitorKey,
		Stars:      vm.Stars,
		Comment:    vm.Comment,
	}
	s.RatingRepository.Add(rating)
	ctx.Redirect(http.StatusFound, "/talk/" + slug)
}

func (s *TalkController) PostQuestionHandler(ctx *gin.Context) {
	slug := ctx.Param("slug")
	visitorKey := session.GetSession(ctx).Get(session.VisitorKey)
	var vm AddQuestionViewModel
	utils.Check(ctx.Bind(&vm))

	talk := s.TalkRepository.GetBySlug(slug)
	question := &application.Question{
		ID:         createUUID(),
		TalkID:     talk.ID,
		VisitorKey: visitorKey,
		Question:   vm.Question,
	}
	s.QuestionRepository.Add(question)
	ctx.Redirect(http.StatusFound, "/talk/" + slug + "?q=1")
}

func (s *TalkController) GetRatingsHandler(ctx *gin.Context) {
	slug := ctx.Param("slug")
	talk := s.TalkRepository.GetBySlug(slug)
	ratings := s.RatingRepository.GetByTalkId(talk.ID)
	summary := calculateRatingsSummary(ratings)

	templates.ReplyTemplate(ctx, "talk_ratings", gin.H{
		"talk": talk,
		"ratings": ratings,
		"summary": summary,
	})
}

func getSpeakerImageFileName(talkID string) string {
	var result string
	utils.Check(filepath.Walk("uploads/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), "talk-" + talkID + "-speaker-image") {
			result = "/uploads/" + info.Name()
		}
		return nil
	}))
	return result
}

func calculateRatingsSummary(ratings []*application.Rating) RatingsSummary {
	averageCounter := 0
	result := RatingsSummary{
		Average: 0,
		Count: map[int]int{
			1: 0,
			2: 0,
			3: 0,
			4: 0,
			5: 0,
		},
	}

	if len(ratings) == 0 {
		return result
	}

	for _, rating := range ratings {
		averageCounter += rating.Stars
		result.Count[rating.Stars]++
	}

	result.Average = math.Floor((float64(averageCounter) / float64(len(ratings))) * 100)/100
	return result
}

func createUUID() string {
	return uuid.New().String()
}
