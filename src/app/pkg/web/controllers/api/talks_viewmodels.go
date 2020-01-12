package api

import (
	"gonference/pkg/domain"
	"time"
)

// AddTalkViewModel .
type AddTalkViewModel struct {
	Name         string    `form:"name"`
	Description  string    `form:"description"`
	SpeakerName  string    `form:"speakerName"`
	SpeakerTitle string    `form:"speakerTitle"`
	Track        string    `form:"track"`
	When         time.Time `form:"when"`
}

// UpdateTalkViewModel .
type UpdateTalkViewModel struct {
	ID           int       `form:"id"`
	Name         string    `form:"name"`
	Description  string    `form:"description"`
	SpeakerName  string    `form:"speakerName"`
	SpeakerTitle string    `form:"speakerTitle"`
	Track        string    `form:"track"`
	When         time.Time `form:"when"`
}

// ToDomainTalk .
func (vm AddTalkViewModel) ToDomainTalk() domain.Talk {
	return domain.Talk{
		Name:         vm.Name,
		Description:  vm.Description,
		SpeakerName:  vm.SpeakerName,
		SpeakerTitle: vm.SpeakerTitle,
		Track:        vm.Track,
		When:         vm.When,
	}
}

// ToDomainTalk .
func (vm UpdateTalkViewModel) ToDomainTalk() domain.Talk {
	return domain.Talk{
		ID:           vm.ID,
		Name:         vm.Name,
		Description:  vm.Description,
		SpeakerName:  vm.SpeakerName,
		SpeakerTitle: vm.SpeakerTitle,
		Track:        vm.Track,
		When:         vm.When,
	}
}
