package postgres

import (
	"github.com/sirikon/gonference"
)

// TalkModel .
type TalkModel struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

// ToDomainTalk .
func (tm TalkModel) ToDomainTalk() gonference.Talk {
	return gonference.Talk{
		ID:   tm.ID,
		Name: tm.Name,
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
		ID:   talk.ID,
		Name: talk.Name,
	}
}
