package database

import (
	"gonference/pkg/utils"
	"time"

	"gonference/pkg/domain"
)

type Scanner = func(dest ...interface{}) error

func scan(scanner Scanner, dest ...interface{}) {
	utils.Check(scanner(dest...))
}

const talkFields = "t.id, t.slug, t.name, t.description, t.speaker_name, t.speaker_title, t.track, t.when_date"
func talkReader(scanner Scanner) *domain.Talk {
	talk := &domain.Talk{}
	scan(scanner,
		&talk.ID,
		&talk.Slug,
		&talk.Name,
		&talk.Description,
		&talk.SpeakerName,
		&talk.SpeakerTitle,
		&talk.Track,
		&talk.When)
	return talk
}

const ratedTalkFields = "t.id, t.slug, t.name, t.description, t.speaker_name, t.speaker_title, t.track, t.when_date, r.id is not null as rated"
func ratedTalkReader(scanner Scanner) *domain.RatedTalk {
	talk := &domain.RatedTalk{}
	scan(scanner,
		&talk.ID,
		&talk.Slug,
		&talk.Name,
		&talk.Description,
		&talk.SpeakerName,
		&talk.SpeakerTitle,
		&talk.Track,
		&talk.When,
		&talk.Rated)
	return talk
}

const ratingFields = "id, talk_id, visitor_key, stars, comment"
func ratingReader(scanner Scanner) *domain.Rating {
	rating := &domain.Rating{}
	scan(scanner,
		&rating.ID,
		&rating.TalkID,
		&rating.VisitorKey,
		&rating.Stars,
		&rating.Comment)
	return rating
}

const questionFields = "id, talk_id, visitor_key, question"
func questionReader(scanner Scanner) *domain.Question {
	question := &domain.Question{}
	scan(scanner,
		&question.ID,
		&question.TalkID,
		&question.VisitorKey,
		&question.Question)
	return question
}





















// TalkModel .
type TalkModel struct {
	ID           int       `db:"id"`
	Slug         string    `db:"slug"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	SpeakerName  string    `db:"speaker_name"`
	SpeakerTitle string    `db:"speaker_title"`
	Track        string    `db:"track"`
	When         time.Time `db:"when_date"`
}

// ToDomainTalk .
func (tm TalkModel) ToDomainTalk() domain.Talk {
	return domain.Talk{
		ID:           tm.ID,
		Slug:         tm.Slug,
		Name:         tm.Name,
		Description:  tm.Description,
		SpeakerName:  tm.SpeakerName,
		SpeakerTitle: tm.SpeakerTitle,
		Track:        tm.Track,
		When:         tm.When,
	}
}

// TalksToDomainTalks .
func TalksToDomainTalks(talks []*TalkModel) []domain.Talk {
	result := make([]domain.Talk, len(talks))
	for i, talk := range talks {
		result[i] = talk.ToDomainTalk()
	}
	return result
}

// DomainTalkToTalk .
func DomainTalkToTalk(talk domain.Talk) TalkModel {
	return TalkModel{
		ID:           talk.ID,
		Slug:         talk.Slug,
		Name:         talk.Name,
		Description:  talk.Description,
		SpeakerName:  talk.SpeakerName,
		SpeakerTitle: talk.SpeakerTitle,
		Track:        talk.Track,
		When:         talk.When,
	}
}

type RatingModel struct {
	ID int `db:"id"`
	TalkID int `db:"talk_id"`
	VisitorKey string `db:"visitor_key"`
	Stars int `db:"stars"`
	Comment string `db:"comment"`
}

func RatingsToDomainRatings(ratings []*RatingModel) []domain.Rating {
	result := make([]domain.Rating, len(ratings))
	for i, rating := range ratings {
		result[i] = rating.ToDomainRating()
	}
	return result
}

func QuestionsToDomainQuestions(questions []QuestionModel) []domain.Question {
	result := make([]domain.Question, len(questions))
	for i, question := range questions {
		result[i] = question.ToDomainQuestion()
	}
	return result
}

func (rm RatingModel) ToDomainRating() domain.Rating {
	return domain.Rating{
		ID:         rm.ID,
		TalkID:     rm.TalkID,
		VisitorKey: rm.VisitorKey,
		Stars:      rm.Stars,
		Comment:    rm.Comment,
	}
}

func DomainRatingToRating(rating domain.Rating) RatingModel {
	return RatingModel{
		ID:         rating.ID,
		TalkID:     rating.TalkID,
		VisitorKey: rating.VisitorKey,
		Stars:      rating.Stars,
		Comment:    rating.Comment,
	}
}

type QuestionModel struct {
	ID int `db:"id"`
	TalkID int `db:"talk_id"`
	VisitorKey string `db:"visitor_key"`
	Question string `db:"question"`
}

func (qm QuestionModel) ToDomainQuestion() domain.Question {
	return domain.Question{
		ID:         qm.ID,
		TalkID:     qm.TalkID,
		VisitorKey: qm.VisitorKey,
		Question:   qm.Question,
	}
}

func DomainQuestionToQuestion(question domain.Question) QuestionModel {
	return QuestionModel{
		ID:         question.ID,
		TalkID:     question.TalkID,
		VisitorKey: question.VisitorKey,
		Question:   question.Question,
	}
}

type UserModel struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
