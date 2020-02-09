package database

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"gonference/pkg/database/binders"
	"gonference/pkg/domain"
	"gonference/pkg/infrastructure/logger"
)

type QuestionRepository struct {
	Logger logger.Logger
	DB *pgxpool.Pool
}

func (qr *QuestionRepository) GetByTalkId(talkId string) []*domain.Question {
	rows := selectQuery(
		qr.DB, binders.QuestionFieldsString, "question", "WHERE talk_id = $1 ORDER BY id", talkId)

	questions := make([]*domain.Question, 0)
	for rows.Next() {
		questions = append(questions, binders.QuestionReader(rows.Scan))
	}
	return questions
}

func (qr *QuestionRepository) Add(question *domain.Question) {
	insertQuery(qr.DB, binders.QuestionFieldsString, "question", binders.QuestionWriter(question)...)
}
