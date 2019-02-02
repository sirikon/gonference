package gonference

// TalkRepository .
type TalkRepository interface {
	GetAll() ([]Talk, error)
}
