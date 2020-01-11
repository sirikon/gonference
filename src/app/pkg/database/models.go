package database

import (
	"time"

	"gonference/pkg/domain"
)

// TalkModel .
type TalkModel struct {
	ID           int       `db:"id"`
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
		Name:         tm.Name,
		Description:  tm.Description,
		SpeakerName:  tm.SpeakerName,
		SpeakerTitle: tm.SpeakerTitle,
		Track:        tm.Track,
		When:         tm.When,
	}
}

// TalksToDomainTalks .
func TalksToDomainTalks(talks []TalkModel) []domain.Talk {
	result := []domain.Talk{}
	for _, talk := range talks {
		result = append(result, talk.ToDomainTalk())
	}
	return result
}

// DomainTalkToTalk .
func DomainTalkToTalk(talk domain.Talk) TalkModel {
	return TalkModel{
		ID:           talk.ID,
		Name:         talk.Name,
		Description:  talk.Description,
		SpeakerName:  talk.SpeakerName,
		SpeakerTitle: talk.SpeakerTitle,
		Track:        talk.Track,
		When:         talk.When,
	}
}

type UserModel struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
