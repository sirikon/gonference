package public

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/domain"
	u "gonference/pkg/utils"
	"gonference/pkg/web/templates"
)

// IndexController .
type IndexController struct {
	TalkRepository domain.TalkRepository
}

// Handler .
func (s *IndexController) Handler(c *gin.Context) {
	talks, err := s.TalkRepository.GetAll(); u.Check(err)
	templates.ReplyTemplate(c, "index", talks)
}
