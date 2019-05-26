package public

import (
	"github.com/gin-gonic/gin"
	"github.com/sirikon/gonference/src/domain"
	"github.com/sirikon/gonference/src/web/templates"
	"net/http"
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

	result, err := templates.RenderTemplate("index", talks)
	if err != nil {
		handleErr(err)
		return
	}

	c.Data(http.StatusOK, "text/html", result)
}
