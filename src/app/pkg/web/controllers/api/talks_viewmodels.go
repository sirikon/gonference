package api

import (
	"gonference/pkg/domain"
	"time"
)

// AddTalkViewModel .
type AddTalkViewModel struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	SpeakerName  string    `json:"speakerName"`
	SpeakerTitle string    `json:"speakerTitle"`
	Track        string    `json:"track"`
	When         time.Time `json:"when"`
}

// UpdateTalkViewModel .
type UpdateTalkViewModel struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	SpeakerName  string    `json:"speakerName"`
	SpeakerTitle string    `json:"speakerTitle"`
	Track        string    `json:"track"`
	When         time.Time `json:"when"`
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
