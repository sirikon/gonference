package application

// TalkRepository .
type TalkRepository interface {
	GetAll() []*Talk
	GetAllWithRated(visitorKey string) []*RatedTalk
	Get(id string) *Talk
	GetBySlug(slug string) *Talk
	Add(talk *Talk)
	Update(talk *Talk)
	Delete(id string)
}

type RatingRepository interface {
	Add(rating *Rating)
	GetByVisitorKey(visitorKey string) []*Rating
	GetByTalkIdAndVisitorKey(talkID string, visitorKey string) *Rating
	GetByTalkId(talkID string) []*Rating
}

type QuestionRepository interface {
	GetByTalkId(talkId string) []*Question
	Add(question *Question)
}

// UserService .
type UserService interface {
	UserExists(username string) bool
	CheckPassword(username string, password string) bool
	ChangePassword(username string, currentPassword string, newPassword string)
}
