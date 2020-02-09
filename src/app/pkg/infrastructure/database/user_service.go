package database

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/pkg/errors"
	"gonference/pkg/application"
	"gonference/pkg/infrastructure/database/binders"
	"gonference/pkg/infrastructure/database/client"
	"gonference/pkg/infrastructure/logger"
)

type UserService struct {
	Logger logger.Logger
	DB     *client.DBClient
}

func (u *UserService) UserExists(username string) bool {
	if u.get(username) == nil {
		return false
	}
	return true
}

func (u *UserService) CheckPassword(username string, password string) bool {
	user := u.get(username)
	if user == nil {
		return false
	}
	hashedPassword := hashPassword(password)
	return hashedPassword == user.Password
}

func (u *UserService) ChangePassword(username string, currentPassword string, newPassword string) {
	if !u.CheckPassword(username, currentPassword) {
		panic(errors.New("Incorrect current password"))
	}
	u.changePassword(username, newPassword)
}

func (u *UserService) changePassword(username string, newPassword string) {
	hashedPassword := hashPassword(newPassword)
	u.DB.Exec(`UPDATE \"user\" SET password = $2 WHERE username = $1`, username, hashedPassword)
}

func (u *UserService) get(username string) *application.User {
	rows := u.DB.Select(binders.UserFieldsString, "\"user\"", "WHERE username = $1 LIMIT 1", username)
	users := make([]*application.User, 0)
	for rows.Next() {
		users = append(users, binders.UserReader(rows.Scan))
	}
	if len(users) == 0 {
		return nil
	}
	return users[0]
}

func hashPassword(password string) string {
	s := sha256.New()
	s.Write([]byte(password))
	return base64.URLEncoding.EncodeToString(s.Sum(nil))
}
