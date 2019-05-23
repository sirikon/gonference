package auth

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

const cookieName = "gonference.session"

// ReadSession .
func EnsureSession(logger *logrus.Entry, r *http.Request, w http.ResponseWriter) *Session {
	var session *Session
	cookie, err := getCookieFromRequest(r)
	if err != nil {
		session = createSession()
		WriteSession(logger, w, session)
		return session
	}

	session, err = serializeSession(cookie.Value)
	if err != nil {
		session = createSession()
		WriteSession(logger, w, session)
		return session
	}

	return session
}

// WriteSession .
func WriteSession(logger *logrus.Entry, w http.ResponseWriter, session *Session) {
	s, err := deserializeSession(session)
	if err != nil {
		logger.Error(err)
	}
	cookie := &http.Cookie{
		Name: cookieName,
		Value: s,
	}
	http.SetCookie(w, cookie)
}

func createSession() *Session {
	return &Session{
		username: "admin",
		role:     UserRole,
	}
}

func getCookieFromRequest(r *http.Request) (*http.Cookie, error) {
	return r.Cookie(cookieName)
}
