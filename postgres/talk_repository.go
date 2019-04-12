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
