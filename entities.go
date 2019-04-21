package gonference

import (
	"time"
)

// Talk .
type Talk struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	SpeakerName  string    `json:"speakerName"`
	SpeakerTitle string    `json:"speakerTitle"`
	Track        string    `json:"track"`
	When         time.Time `json:"when"`
}
