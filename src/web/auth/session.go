package auth

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
