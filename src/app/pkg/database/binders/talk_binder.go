package binders

import "gonference/pkg/domain"

var TalkFields = []string{"id", "slug", "name", "description", "speaker_name", "speaker_title", "track", "when_date"}

func TalkWriter(talk *domain.Talk) []interface{} {
	return []interface{}{
		talk.ID,
		talk.Slug,
		talk.Name,
		talk.Description,
		talk.SpeakerName,
		talk.SpeakerTitle,
		talk.Track,
		talk.When,
	}
}

func TalkReader(scanner Scanner) *domain.Talk {
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
