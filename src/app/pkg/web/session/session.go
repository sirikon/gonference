package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gonference/pkg/utils"
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
	value := s.internalSession.Get(string(key))
	if value == nil {
		return ""
	}
	return value.(string)
}

func (s *Session) Set(key Key, value string) {
	s.internalSession.Set(string(key), value)
	utils.HandleErr(s.internalSession.Save())
}

func GetSession(ctx *gin.Context) *Session {
	return &Session {
		internalSession: sessions.Default(ctx),
	}
}
