package public

import (
	"gonference/pkg/domain"
	"time"
)

type Talk struct {
	ID           int
	Slug         string
	Name         string
	Description  string
	SpeakerName  string
	SpeakerTitle string
	Track        string
	When         time.Time
	Rated        bool
}

func DomainTalksToTalks(talks []*domain.Talk, ratings []*domain.Rating) []Talk {
	ratingMap := ratingMap(ratings)
	result := make([]Talk, len(talks))
	for i, talk := range talks {
		_, exists := ratingMap[talk.ID]
		result[i] = DomainTalkToTalk(talk, exists)
	}
	return result
}

func DomainTalkToTalk(talk *domain.Talk, rated bool) Talk {
	return Talk{
		ID:           talk.ID,
		Slug:         talk.Slug,
		Name:         talk.Name,
		Description:  talk.Description,
		SpeakerName:  talk.SpeakerName,
		SpeakerTitle: talk.SpeakerTitle,
		Track:        talk.Track,
		When:         talk.When,
		Rated:        rated,
	}
}

func ratingMap(ratings []*domain.Rating) map[int]*domain.Rating {
	result := make(map[int]*domain.Rating)
	for _, rating := range ratings {
		result[rating.TalkID] = rating
	}
	return result
}
