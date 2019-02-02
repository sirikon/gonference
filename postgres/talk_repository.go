package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirikon/gonference"
)

// TalkRepository .
type TalkRepository struct {
	DB *sqlx.DB
}

// GetAll .
func (tr *TalkRepository) GetAll() ([]gonference.Talk, error) {
	talks := []TalkModel{}
	err := tr.DB.Select(&talks, "SELECT * FROM talk")
	if err != nil {
		return nil, err
	}
	return TalksToDomainTalks(talks), nil
}
