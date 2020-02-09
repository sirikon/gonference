package binders

import (
	"gonference/pkg/application"
	"strings"
)

var UserFields = []string{"username", "password"}
var UserFieldsString = strings.Join(UserFields, ", ")
func UserReader(scanner Scanner) *application.User {
	user := &application.User{}
	scan(scanner,
		&user.Username,
		&user.Password)
	return user
}
