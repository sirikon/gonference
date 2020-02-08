package binders

import (
	"gonference/pkg/domain"
	"strings"
)

var UserFields = []string{"username", "password"}
var UserFieldsString = strings.Join(UserFields, ", ")
func UserReader(scanner Scanner) *domain.User {
	user := &domain.User{}
	scan(scanner,
		&user.Username,
		&user.Password)
	return user
}
