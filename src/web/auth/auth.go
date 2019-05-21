package auth

import (
	"net/http"
)

const cookieName = "gonference.session"

// GetSession .
func GetSession(r *http.Request) *Session {
	cookie, err := getCookieFromRequest(r)
	if err != nil {
		return createSession()
	}

	session, err := readSessionFromCookie(cookie)
	if err != nil {
		return createSession()
	}

	return session
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
