package database

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"gonference/pkg/database/binders"
	"gonference/pkg/domain"
	"gonference/pkg/infrastructure/logger"
)

// TalkRepository .
type TalkRepository struct {
	Logger logger.Logger
	DB     *pgxpool.Pool
}

func (tr *TalkRepository) GetAll() []*domain.Talk {
	return tr.selectQuery("ORDER BY when_date ASC, track ASC")
}

func (tr *TalkRepository) GetAllWithRated(visitorKey string) []*domain.RatedTalk {
	rows := query(tr.DB,`
		SELECT ` + binders.RatedTalkFieldsString + `
		FROM talk t
		LEFT JOIN rating r
			ON r.talk_id = t.id
			AND r.visitor_key = $1
		ORDER BY t.when_date ASC, t.track ASC
	`, visitorKey)

	talks := make([]*domain.RatedTalk, 0)
	for rows.Next() {
		talks = append(talks, binders.RatedTalkReader(rows.Scan))
	}
	return talks
}

func (tr *TalkRepository) GetBySlug(slug string) *domain.Talk {
	return tr.selectOneQuery("slug = $1", slug)
}

func (tr *TalkRepository) Get(id int) *domain.Talk {
	return tr.selectOneQuery("id = $1", id)
}

func (tr *TalkRepository) Add(talk *domain.Talk) {
	tr.insertQuery(talk)
}

func (tr *TalkRepository) Update(talk *domain.Talk) {
	exec(tr.DB, `
		UPDATE talk
		SET
			slug = $2,
			name = $3,
			description = $4,
			speaker_name = $5,
			speaker_title = $6,
			track = $7,
			when_date = $8
		WHERE id = $1
	`, binders.TalkWriter(talk)...)
}

func (tr *TalkRepository) Delete(id int) {
	exec(tr.DB, "DELETE FROM talk WHERE id = $1", id)
}

func (tr *TalkRepository) selectOneQuery(extra string, args ...interface{}) *domain.Talk {
	results := tr.selectQuery(extra, args...)
	if len(results) == 0 {
		return nil
	}
	return results[0]
}

func (tr *TalkRepository) selectQuery(extra string, args ...interface{}) []*domain.Talk {
	rows := selectQuery(tr.DB, binders.TalkFieldsString, "talk", extra, args...)
	talks := make([]*domain.Talk, 0)
	for rows.Next() {
		talks = append(talks, binders.TalkReader(rows.Scan))
	}
	return talks
}

func (tr *TalkRepository) insertQuery(talk *domain.Talk) {
	insertQuery(tr.DB, binders.TalkFieldsString, "talk", binders.TalkWriter(talk))
}
