package public

import (
	"github.com/gin-gonic/gin"
	"github.com/sirikon/gonference/src/domain"
	"github.com/sirikon/gonference/src/utils"
	"github.com/sirikon/gonference/src/web/templates"
)

// IndexController .
type IndexController struct {
	TalkRepository domain.TalkRepository
}

// Handler .
func (s *IndexController) Handler(c *gin.Context) {
	talks, err := s.TalkRepository.GetAll(); utils.HandleErr(err)
	templates.ReplyTemplate(c, "index", talks)
}
