package public

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/domain"
	u "gonference/pkg/utils"
	"gonference/pkg/web/session"
	"gonference/pkg/web/templates"
)

// IndexController .
type IndexController struct {
	TalkRepository domain.TalkRepository
	RatingRepository domain.RatingRepository
}

// Handler .
func (s *IndexController) Handler(c *gin.Context) {
	visitorKey := session.GetSession(c).Get(session.VisitorKey)
	domainTalks, err := s.TalkRepository.GetAll(); u.Check(err)
	ratings := s.RatingRepository.GetByVisitorKey(visitorKey)
	talks := DomainTalksToTalks(domainTalks, ratings)

	templates.ReplyTemplate(c, "index", talks)
}
