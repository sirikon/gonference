package postgres

import (
	"github.com/sirikon/gonference"
)

// TalkModel .
type TalkModel struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	Description  string `db:"description"`
	SpeakerName  string `db:"speaker_name"`
	SpeakerTitle string `db:"speaker_title"`
	Track        string `db:"track"`
	When         string `db:"when_date"`
}

// ToDomainTalk .
func (tm TalkModel) ToDomainTalk() gonference.Talk {
	return gonference.Talk{
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
func TalksToDomainTalks(talks []TalkModel) []gonference.Talk {
	result := []gonference.Talk{}
	for _, talk := range talks {
		result = append(result, talk.ToDomainTalk())
	}
	return result
}

// DomainTalkToTalk .
func DomainTalkToTalk(talk gonference.Talk) TalkModel {
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
