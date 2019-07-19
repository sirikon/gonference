package public

import (
	"github.com/gin-gonic/gin"
	"github.com/sirikon/gonference/src/domain"
	"github.com/sirikon/gonference/src/web/templates"
	"strconv"
)

// TalkController .
type TalkController struct {
	TalkRepository domain.TalkRepository
}

// Handler .
func (s *TalkController) Handler(c *gin.Context) {
	handleErr := func(err error) {
		_ = c.Error(err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleErr(err)
		return
	}

	talk, err := s.TalkRepository.Get(id)
	if err != nil {
		handleErr(err)
		return
	}

	templates.ReplyTemplate(c, "talk", talk)
}
