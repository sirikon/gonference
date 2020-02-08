package domain

import (
	"time"
)

// Talk .
type Talk struct {
	ID           string       `json:"id"`
	Slug         string    `json:"slug"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	SpeakerName  string    `json:"speakerName"`
	SpeakerTitle string    `json:"speakerTitle"`
	Track        string    `json:"track"`
	When         time.Time `json:"when"`
}

type RatedTalk struct {
	Talk
	Rated bool
}

type Rating struct {
	ID string
	TalkID string
	VisitorKey string
	Stars int
	Comment string
}

type Question struct {
	ID string `json:"id"`
	TalkID string `json:"talkId"`
	VisitorKey string `json:"visitorKey"`
	Question string `json:"question"`
}

type User struct {
	Username string
	Password string
}
