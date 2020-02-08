package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gonference/pkg/domain"
	"gonference/pkg/infrastructure/logger"
	"gonference/pkg/utils"
)

// TalkRepository .
type TalkRepository struct {
	Logger logger.Logger
	DB     *pgxpool.Pool
}

func (tr *TalkRepository) GetAll() []*domain.Talk {
	query := "SELECT " + talkFields + " FROM talk t ORDER BY when_date ASC, track ASC"
	rows, err := tr.DB.Query(context.Background(), query); utils.Check(err)
	talks := make([]*domain.Talk, 0)
	for rows.Next() {
		talks = append(talks, talkReader(rows.Scan))
	}
	return talks
}

func (tr *TalkRepository) GetAllWithRated(visitorKey string) []*domain.RatedTalk {
	query := "SELECT " + ratedTalkFields + " FROM talk t LEFT JOIN rating r ON r.talk_id = t.id AND r.visitor_key = $1 ORDER BY when_date ASC, track ASC"
	rows, err := tr.DB.Query(context.Background(), query, visitorKey); utils.Check(err)
	talks := make([]*domain.RatedTalk, 0)
	for rows.Next() {
		talks = append(talks, ratedTalkReader(rows.Scan))
	}
	return talks
}

func (tr *TalkRepository) GetBySlug(slug string) domain.Talk {
	return *tr.selectOneQuery("slug = $1", slug)
}

func (tr *TalkRepository) Get(id int) (domain.Talk, error) {
	return *tr.selectOneQuery("id = $1", id), nil
}

func (tr *TalkRepository) Add(talk domain.Talk) (int, error) {
	var id int
	query := "INSERT INTO talk (slug, name, description, speaker_name, speaker_title, track, when_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	err := tr.DB.QueryRow(context.Background(), query, talk.Slug, talk.Name, talk.Description, talk.SpeakerName, talk.SpeakerTitle, talk.Track, talk.When).
		Scan(&id)
	return id, err
}

func (tr *TalkRepository) Update(talk domain.Talk) error {
	query := "UPDATE talk SET slug = $2, name = $3, description = $4, speaker_name = $5, speaker_title = $6, track = $7, when_date = $8 WHERE id = $1"
	_, err := tr.DB.Exec(context.Background(), query, talk.ID, talk.Slug, talk.Name, talk.Description, talk.SpeakerName, talk.SpeakerTitle, talk.Track, talk.When)
	return err
}

func (tr *TalkRepository) Delete(id int) error {
	query := "DELETE FROM talk WHERE id = $1"
	_, err := tr.DB.Exec(context.Background(), query, id)
	return err
}

func (tr *TalkRepository) selectOneQuery(extra string, args ...interface{}) *domain.Talk {
	results := tr.selectQuery(extra, args...)
	if len(results) == 0 {
		return nil
	}
	return results[0]
}

func (tr *TalkRepository) selectQuery(extra string, args ...interface{}) []*domain.Talk {
	rows := selectQuery(tr.DB, talkFields, "talk", extra, args...)
	ratings := make([]*domain.Talk, 0)
	for rows.Next() {
		ratings = append(ratings, talkReader(rows.Scan))
	}
	return ratings
}

func (tr *TalkRepository) insertQuery(rating *domain.Rating) {
	insertQuery(tr.DB, talkFields, "talk", talkWriter(rating))
}
