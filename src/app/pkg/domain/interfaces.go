package domain

// TalkRepository .
type TalkRepository interface {
	GetAll() ([]Talk, error)
	Get(id int) (Talk, error)
	GetBySlug(slug string) Talk
	Add(talk Talk) (int, error)
	Update(talk Talk) error
	Delete(id int) error
}

type RatingRepository interface {
	Add(rating Rating)
	GetByVisitorKey(visitorKey string) []Rating
	GetByTalkIdAndVisitorKey(talkID int, visitorKey string) *Rating
}

type QuestionRepository interface {
	GetByTalkId(talkId int) []Question
	Add(question Question)
}

// UserService .
type UserService interface {
	UserExists(username string) bool
	CheckPassword(username string, password string) (bool, error)
	ChangePassword(username string, currentPassword string, newPassword string) error
}
