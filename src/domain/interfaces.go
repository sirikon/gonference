package domain

// TalkRepository .
type TalkRepository interface {
	GetAll() ([]Talk, error)
	Get(int) (Talk, error)
	Add(Talk) error
	Update(Talk) error
	Delete(int) error
}

// UserService .
type UserService interface {
	UserExists(string) bool
	CheckPassword(string, string) (bool, error)
}
