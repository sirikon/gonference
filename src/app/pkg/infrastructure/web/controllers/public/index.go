package public

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/application"
	"gonference/pkg/infrastructure/web/session"
	"gonference/pkg/infrastructure/web/templates"
)

// IndexController .
type IndexController struct {
	TalkRepository   application.TalkRepository
	RatingRepository application.RatingRepository
}

// Handler .
func (s *IndexController) Handler(c *gin.Context) {
	visitorKey := session.GetSession(c).Get(session.VisitorKey)
	talks := s.TalkRepository.GetAllWithRated(visitorKey)
	templates.ReplyTemplate(c, "index", talks)
}
