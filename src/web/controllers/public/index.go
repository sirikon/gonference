package public

import (
	"net/http"

	"github.com/sirikon/gonference/src/domain"
	"github.com/sirikon/gonference/src/web/middleware"
	"github.com/sirikon/gonference/src/web/templates"
	log "github.com/sirupsen/logrus"
)

// IndexController .
type IndexController struct {
	TalkRepository domain.TalkRepository
}

// Handler .
func (s *IndexController) Handler(ctx *middleware.RequestContext) {
	handleErr := func(err error) {
		log.Error(err)
		http.Error(ctx.ResponseWritter, "Something went wrong", http.StatusInternalServerError)
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

	ctx.ResponseWritter.Write(result)
}
