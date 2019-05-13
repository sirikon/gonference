package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const cookieName = "gonference.session"

// Role .
type Role int

const (
	// UserRole .
	UserRole Role = 0
	// AdminRole .
	AdminRole Role = 1
)

// Session .
type Session struct {
	username string
	role     Role
}

// GetUsername .
func (s *Session) GetUsername() string {
	return s.username
}

// SetUsername .
func (s *Session) SetUsername(username string) {
	s.username = username
}

// GetRole .
func (s *Session) GetRole() Role {
	return s.role
}

// SetRole .
func (s *Session) SetRole(role Role) {
	s.role = role
}

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

func readSessionFromCookie(cookie *http.Cookie) (*Session, error) {
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		session := createSession()

		role, ok := claims["role"].(Role)
		if !ok {
			return nil, errors.New("There was a problem parsing users session")
		}
		session.role = role
		return session, nil
	}
	return nil, err
}

// func getOrCreateCookie(r *http.Request) {
// 	cookie, err := getCookieFromRequest(r)
// 	if err != nil {
// 		cookie = createSessionCookie()
// 	}
// 	return cookie
// }

func getCookieFromRequest(r *http.Request) (*http.Cookie, error) {
	return r.Cookie(cookieName)
}

// func createSessionCookie() *http.Cookie {
// 	return &http.Cookie{
// 		Name:  cookieName,
// 		Value: "asdf",
// 	}
// }
