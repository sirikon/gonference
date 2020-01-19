package database

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gonference/pkg/infrastructure/logger"
)

type UserService struct {
	Logger logger.Logger
	DB     *sqlx.DB
}

func (u *UserService) UserExists(username string) bool {
	_, err := u.get(username)
	if err != nil {
		return false
	}
	return true
}

func (u *UserService) CheckPassword(username string, password string) (bool, error) {
	user, err := u.get(username)
	if err != nil {
		return false, err
	}
	hashedPassword := hashPassword(password)
	return hashedPassword == user.Password, nil
}

func (u *UserService) ChangePassword(username string, currentPassword string, newPassword string) error {
	result, err := u.CheckPassword(username, currentPassword)
	if err != nil {
		return err
	}

	if !result {
		return errors.New("wrong current password")
	}

	return u.changePassword(username, newPassword)
}

func (u *UserService) get(username string) (UserModel, error) {
	var user UserModel
	query := "SELECT * FROM \"user\" WHERE username = $1 LIMIT 1"
	logSelect(u.Logger, query)
	err := u.DB.QueryRowx(query, username).StructScan(&user)
	return user, err
}

func (u *UserService) changePassword(username string, newPassword string) error {
	hashedPassword := hashPassword(newPassword)
	query := "UPDATE \"user\" SET password = $2 WHERE username = $1"
	logMutation(u.Logger, query)
	_, err := u.DB.Exec(query, username, hashedPassword)
	return err
}

func hashPassword(password string) string {
	s := sha256.New()
	s.Write([]byte(password))
	return base64.URLEncoding.EncodeToString(s.Sum(nil))
}
