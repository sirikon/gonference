package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirikon/gonference/src/web/auth"
)

// RequestContext .
type RequestContext struct {
	ResponseWritter http.ResponseWriter
	Request         *http.Request
	Params          httprouter.Params
	Session         *auth.Session
}

// RequestHandler .
type RequestHandler func(*RequestContext)
