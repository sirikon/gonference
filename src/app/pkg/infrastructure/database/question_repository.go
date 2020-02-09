package database

import (
	"gonference/pkg/application"
	"gonference/pkg/infrastructure/database/binders"
	"gonference/pkg/infrastructure/database/client"
	"gonference/pkg/infrastructure/logger"
)

type QuestionRepository struct {
	Logger logger.Logger
	DB *client.DBClient
}

func (qr *QuestionRepository) GetByTalkId(talkId string) []*application.Question {
	rows := qr.DB.Select(
		binders.QuestionFieldsString, "question", "WHERE talk_id = $1 ORDER BY id", talkId)

	questions := make([]*application.Question, 0)
	for rows.Next() {
		questions = append(questions, binders.QuestionReader(rows.Scan))
	}
	return questions
}

func (qr *QuestionRepository) Add(question *application.Question) {
	qr.DB.Insert(binders.QuestionFieldsString, "question", binders.QuestionWriter(question)...)
}
