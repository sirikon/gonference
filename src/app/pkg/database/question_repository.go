package database

import (
	"gonference/pkg/database/binders"
	"gonference/pkg/database/client"
	"gonference/pkg/domain"
	"gonference/pkg/infrastructure/logger"
)

type QuestionRepository struct {
	Logger logger.Logger
	DB *client.DBClient
}

func (qr *QuestionRepository) GetByTalkId(talkId string) []*domain.Question {
	rows := qr.DB.Select(
		binders.QuestionFieldsString, "question", "WHERE talk_id = $1 ORDER BY id", talkId)

	questions := make([]*domain.Question, 0)
	for rows.Next() {
		questions = append(questions, binders.QuestionReader(rows.Scan))
	}
	return questions
}

func (qr *QuestionRepository) Add(question *domain.Question) {
	qr.DB.Insert(binders.QuestionFieldsString, "question", binders.QuestionWriter(question)...)
}
