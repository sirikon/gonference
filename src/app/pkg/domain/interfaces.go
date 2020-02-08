package domain

// TalkRepository .
type TalkRepository interface {
	GetAll() []*Talk
	GetAllWithRated(visitorKey string) []*RatedTalk
	Get(id int) *Talk
	GetBySlug(slug string) *Talk
	Add(talk *Talk)
	Update(talk *Talk)
	Delete(id int)
}

type RatingRepository interface {
	Add(rating *Rating)
	GetByVisitorKey(visitorKey string) []*Rating
	GetByTalkIdAndVisitorKey(talkID int, visitorKey string) *Rating
	GetByTalkId(talkID int) []*Rating
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
