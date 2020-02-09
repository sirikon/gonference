package binders

import (
	"gonference/pkg/application"
	"strings"
)

var TalkFields = []string{"id", "slug", "name", "description", "speaker_name", "speaker_title", "track", "when_date"}
var TalkFieldsString = strings.Join(TalkFields, ", ")
func TalkWriter(talk *application.Talk) []interface{} {
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
func TalkReader(scanner Scanner) *application.Talk {
	talk := &application.Talk{}
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

var RatedTalkFields = []string{
	"t.id", "t.slug", "t.name", "t.description", "t.speaker_name", "t.speaker_title", "t.track", "t.when_date",
	"r.id is not null as rated"}
var RatedTalkFieldsString = strings.Join(RatedTalkFields, ", ")
func RatedTalkReader(scanner Scanner) *application.RatedTalk {
	talk := &application.RatedTalk{}
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
