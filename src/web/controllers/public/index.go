package public

import (
	"github.com/gin-gonic/gin"
	"github.com/sirikon/gonference/src/domain"
	"github.com/sirikon/gonference/src/web/templates"
)

// IndexController .
type IndexController struct {
	TalkRepository domain.TalkRepository
}

// Handler .
func (s *IndexController) Handler(c *gin.Context) {
	handleErr := func(err error) {
		_ = c.Error(err)
	}

	talks, err := s.TalkRepository.GetAll()
	if err != nil {
		handleErr(err)
		return
	}

	templates.ReplyTemplate(c, "index", talks)
}
