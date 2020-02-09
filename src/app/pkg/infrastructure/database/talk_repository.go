package database

import (
	"gonference/pkg/application"
	"gonference/pkg/infrastructure/database/binders"
	"gonference/pkg/infrastructure/database/client"
	"gonference/pkg/infrastructure/logger"
)

// TalkRepository .
type TalkRepository struct {
	Logger logger.Logger
	DB     *client.DBClient
}

func (tr *TalkRepository) GetAll() []*application.Talk {
	return tr.selectQuery("ORDER BY when_date ASC, track ASC")
}

func (tr *TalkRepository) GetAllWithRated(visitorKey string) []*application.RatedTalk {
	rows := tr.DB.Query(`
		SELECT ` +binders.RatedTalkFieldsString+ `
		FROM talk t
		LEFT JOIN rating r
			ON r.talk_id = t.id
			AND r.visitor_key = $1
		ORDER BY t.when_date ASC, t.track ASC
	`, visitorKey)

	talks := make([]*application.RatedTalk, 0)
	for rows.Next() {
		talks = append(talks, binders.RatedTalkReader(rows.Scan))
	}
	return talks
}

func (tr *TalkRepository) GetBySlug(slug string) *application.Talk {
	return tr.selectOneQuery("WHERE slug = $1", slug)
}

func (tr *TalkRepository) Get(id string) *application.Talk {
	return tr.selectOneQuery("WHERE id = $1", id)
}

func (tr *TalkRepository) Add(talk *application.Talk) {
	tr.insertQuery(talk)
}

func (tr *TalkRepository) Update(talk *application.Talk) {
	tr.DB.Exec(`
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

func (tr *TalkRepository) Delete(id string) {
	tr.DB.Exec("DELETE FROM talk WHERE id = $1", id)
}

func (tr *TalkRepository) selectOneQuery(extra string, args ...interface{}) *application.Talk {
	results := tr.selectQuery(extra, args...)
	if len(results) == 0 {
		return nil
	}
	return results[0]
}

func (tr *TalkRepository) selectQuery(extra string, args ...interface{}) []*application.Talk {
	rows := tr.DB.Select(binders.TalkFieldsString, "talk", extra, args...)
	talks := make([]*application.Talk, 0)
	for rows.Next() {
		talks = append(talks, binders.TalkReader(rows.Scan))
	}
	return talks
}

func (tr *TalkRepository) insertQuery(talk *application.Talk) {
	tr.DB.Insert(binders.TalkFieldsString, "talk", binders.TalkWriter(talk)...)
}
