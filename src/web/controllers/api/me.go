package api

import (
	"encoding/json"
	"net/http"

	"github.com/sirikon/gonference/src/web/middleware"
	log "github.com/sirupsen/logrus"
)

// MeAPIController .
type MeAPIController struct {
}

// Handler .
func (s *MeAPIController) Handler(ctx *middleware.RequestContext) {
	handleErr := func(err error) {
		log.Error(err)
		http.Error(ctx.ResponseWritter, "Something went wrong", http.StatusInternalServerError)
	}

	user := User{
		Username: ctx.Session.GetUsername(),
	}

	result, err := json.Marshal(user)
	if err != nil {
		handleErr(err)
		return
	}

	ctx.ResponseWritter.Header().Add("content-type", "application/json")
	ctx.ResponseWritter.Write(result)
}
