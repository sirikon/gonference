package domain

import (
	"time"
)

// Talk .
type Talk struct {
	ID           int       `json:"id"`
	Slug         string    `json:"slug"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	SpeakerName  string    `json:"speakerName"`
	SpeakerTitle string    `json:"speakerTitle"`
	Track        string    `json:"track"`
	When         time.Time `json:"when"`
}

type Rating struct {
	ID int
	TalkID int
	VisitorKey string
	Stars int
	Comment string
}

type Question struct {
	ID int `json:"id"`
	TalkID int `json:"talkId"`
	VisitorKey string `json:"visitorKey"`
	Question string `json:"question"`
}

type User struct {
	Username string
	Password string
}
