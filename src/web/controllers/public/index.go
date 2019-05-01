package public

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirikon/gonference/src/domain"
	"github.com/sirikon/gonference/src/web/templates"
	log "github.com/sirupsen/logrus"
)

// IndexController .
type IndexController struct {
	TalkRepository domain.TalkRepository
}

// Handler .
func (s *IndexController) Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	handleErr := func(err error) {
		log.Error(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
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

	w.Write(result)
}
