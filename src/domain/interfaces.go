package domain

// TalkRepository .
type TalkRepository interface {
	GetAll() ([]Talk, error)
	Get(int) (Talk, error)
	Add(Talk) error
	Update(Talk) error
	Delete(int) error
}
