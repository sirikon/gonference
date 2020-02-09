package api

import (
	"gonference/pkg/application"
	"time"
)

// AddTalkViewModel .
type AddTalkViewModel struct {
	Slug         string    `form:"slug"`
	Name         string    `form:"name"`
	Description  string    `form:"description"`
	SpeakerName  string    `form:"speakerName"`
	SpeakerTitle string    `form:"speakerTitle"`
	Track        string    `form:"track"`
	When         time.Time `form:"when"`
}

// UpdateTalkViewModel .
type UpdateTalkViewModel struct {
	ID           string       `form:"id"`
	Slug         string    `form:"slug"`
	Name         string    `form:"name"`
	Description  string    `form:"description"`
	SpeakerName  string    `form:"speakerName"`
	SpeakerTitle string    `form:"speakerTitle"`
	Track        string    `form:"track"`
	When         time.Time `form:"when"`
}

// ToDomainTalk .
func (vm AddTalkViewModel) ToDomainTalk() *application.Talk {
	return &application.Talk{
		Slug:         vm.Slug,
		Name:         vm.Name,
		Description:  vm.Description,
		SpeakerName:  vm.SpeakerName,
		SpeakerTitle: vm.SpeakerTitle,
		Track:        vm.Track,
		When:         vm.When,
	}
}

// ToDomainTalk .
func (vm UpdateTalkViewModel) ToDomainTalk() *application.Talk {
	return &application.Talk{
		ID:           vm.ID,
		Slug:         vm.Slug,
		Name:         vm.Name,
		Description:  vm.Description,
		SpeakerName:  vm.SpeakerName,
		SpeakerTitle: vm.SpeakerTitle,
		Track:        vm.Track,
		When:         vm.When,
	}
}
