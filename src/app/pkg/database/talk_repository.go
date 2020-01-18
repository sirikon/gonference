package database

import (
	"github.com/jmoiron/sqlx"
	"gonference/pkg/domain"
	"github.com/sirupsen/logrus"
	"gonference/pkg/utils"
)

// TalkRepository .
type TalkRepository struct {
	Logger *logrus.Entry
	DB     *sqlx.DB
}

// GetAll .
func (tr *TalkRepository) GetAll() ([]domain.Talk, error) {
	var talks []TalkModel
	query := "SELECT * FROM talk ORDER BY when_date ASC"

	tr.Logger.Info("Executing query '" + query + "'")
	err := tr.DB.Select(&talks, query)
	if err != nil {
		return nil, err
	}

	return TalksToDomainTalks(talks), nil
}

// Add .
func (tr *TalkRepository) Add(domainTalk domain.Talk) (int, error) {
	talk := DomainTalkToTalk(domainTalk)
	var id int
	query := "INSERT INTO talk (slug, name, description, speaker_name, speaker_title, track, when_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	tr.Logger.Info("Executing query '" + query + "' with name '" + talk.Name + "'")
	err := tr.DB.
		QueryRow(query, talk.Slug, talk.Name, talk.Description, talk.SpeakerName, talk.SpeakerTitle, talk.Track, talk.When).
		Scan(&id)
	return id, err
}

// Get .
func (tr *TalkRepository) Get(id int) (domain.Talk, error) {
	var talk TalkModel
	query := "SELECT * FROM talk WHERE id = $1 LIMIT 1"
	err := tr.DB.QueryRowx(query, id).StructScan(&talk)
	return talk.ToDomainTalk(), err
}

func (tr *TalkRepository) GetBySlug(slug string) domain.Talk {
	var talk TalkModel
	query := "SELECT * FROM talk WHERE slug = $1 LIMIT 1"
	utils.Check(tr.DB.QueryRowx(query, slug).StructScan(&talk))
	return talk.ToDomainTalk()
}

// Update .
func (tr *TalkRepository) Update(domainTalk domain.Talk) error {
	talk := DomainTalkToTalk(domainTalk)
	query := "UPDATE talk SET slug = $2, name = $3, description = $4, speaker_name = $5, speaker_title = $6, track = $7, when_date = $8 WHERE id = $1"
	_, err := tr.DB.Exec(query, talk.ID, talk.Slug, talk.Name, talk.Description, talk.SpeakerName, talk.SpeakerTitle, talk.Track, talk.When)
	return err
}

// Delete .
func (tr *TalkRepository) Delete(id int) error {
	query := "DELETE FROM talk WHERE id = $1"
	_, err := tr.DB.Exec(query, id)
	return err
}
