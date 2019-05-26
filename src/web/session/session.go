package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Session struct {
	internalSession sessions.Session
}

func (s *Session) GetUsername() string {
	value := s.internalSession.Get("username")
	if value == nil {
		return ""
	}
	return value.(string)
}

func (s *Session) SetUsername(username string) {
	s.internalSession.Set("username", username)
	_ = s.internalSession.Save()
}

func (s *Session) GetRole() string {
	value := s.internalSession.Get("role")
	if value == nil {
		return "user"
	}
	return value.(string)
}

func (s *Session) SetRole(role string) {
	s.internalSession.Set("role", role)
	_ = s.internalSession.Save()
}

func GetSession(ctx *gin.Context) *Session {
	return &Session {
		internalSession: sessions.Default(ctx),
	}
}
