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
	query := "INSERT INTO talk (name) VALUES ($1)"

	tr.Logger.Info("Executing query '" + query + "' with name '" + talk.Name + "'")
	_, err := tr.DB.Exec(query, talk.Name)
	if err != nil {
		return err
	}

	return nil
}
