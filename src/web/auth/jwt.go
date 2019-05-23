package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const signingSecret = "secret"

func serializeSession(s string) (*Session, error) {
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("something went wrong when mapping the claims")
	}

	session := createSession()
	role, ok := claims["role"].(float64)
	if !ok {
		return nil, errors.New("there was a problem parsing user role")
	}
	session.role = Role(role)

	return session, nil
}

func deserializeSession(session *Session) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": session.role,
	})
	return token.SignedString([]byte(signingSecret))
}
