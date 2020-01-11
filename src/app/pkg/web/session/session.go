package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Key string

const (
	UsernameKey Key = "username"
	RoleKey Key = "role"
)

const usernameKey = "username"
const roleKey = "role"

type Session struct {
	internalSession sessions.Session
}

func (s *Session) Get(key Key) string {
	value := s.internalSession.Get(key)
	if value == nil {
		return ""
	}
	return value.(string)
}

func (s *Session) Set(key Key, value string) {
	s.internalSession.Set(key, value)
	_ = s.internalSession.Save()
}

func GetSession(ctx *gin.Context) *Session {
	return &Session {
		internalSession: sessions.Default(ctx),
	}
}
