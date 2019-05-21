package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func readSessionFromCookie(cookie *http.Cookie) (*Session, error) {
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
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
	role, ok := claims["role"].(Role)
	if !ok {
		return nil, errors.New("there was a problem parsing user role")
	}
	session.role = role

	return session, nil
}
