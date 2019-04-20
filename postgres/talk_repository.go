package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirikon/gonference"
	"github.com/sirupsen/logrus"
)

// TalkRepository .
type TalkRepository struct {
	Logger *logrus.Entry
	DB     *sqlx.DB
}

// GetAll .
func (tr *TalkRepository) GetAll() ([]gonference.Talk, error) {
	talks := []TalkModel{}
	query := "SELECT * FROM talk"

	tr.Logger.Info("Executing query '" + query + "'")
	err := tr.DB.Select(&talks, query)
	if err != nil {
		return nil, err
	}

	return TalksToDomainTalks(talks), nil
}

// Add .
func (tr *TalkRepository) Add(domainTalk gonference.Talk) error {
	talk := DomainTalkToTalk(domainTalk)
	query := "INSERT INTO talk (name, description, speaker_name, speaker_title, track, when_date) VALUES ($1, $2, $3, $4, $5, $6)"

	tr.Logger.Info("Executing query '" + query + "' with name '" + talk.Name + "'")
	_, err := tr.DB.Exec(query, talk.Name, talk.Description, talk.SpeakerName, talk.SpeakerTitle, talk.Track, talk.When)
	return err
}

// Get .
func (tr *TalkRepository) Get(id int) (gonference.Talk, error) {
	var talk TalkModel
	query := "SELECT * FROM talk WHERE id = $1 LIMIT 1"
	err := tr.DB.QueryRowx(query, id).StructScan(&talk)
	return talk.ToDomainTalk(), err
}

// Update .
func (tr *TalkRepository) Update(domainTalk gonference.Talk) error {
	talk := DomainTalkToTalk(domainTalk)
	query := "UPDATE talk SET name = $2, description = $3, speaker_name = $4, speaker_title = $5, track = $6, when_date = $7 WHERE id = $1"
	_, err := tr.DB.Exec(query, talk.ID, talk.Name, talk.Description, talk.SpeakerName, talk.SpeakerTitle, talk.Track, talk.When)
	return err
}
