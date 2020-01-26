package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gonference/pkg/domain"
	"gonference/pkg/infrastructure/logger"
	"gonference/pkg/utils"
)

type QuestionRepository struct {
	Logger logger.Logger
	DB *pgxpool.Pool
}

func (qr *QuestionRepository) GetByTalkId(talkId int) []domain.Question {
	query := "SELECT " + questionFields + " FROM question WHERE talk_id = $1 ORDER BY id"
	rows, err := qr.DB.Query(context.Background(), query, talkId); utils.Check(err)

	questions := make([]domain.Question, 0)
	for rows.Next() {
		questions = append(questions, *questionReader(rows.Scan))
	}

	return questions
}

func (qr *QuestionRepository) Add(question domain.Question) {
	sql := "INSERT INTO question (talk_id, visitor_key, question) VALUES ($1, $2, $3)"
	_, err := qr.DB.Exec(context.Background(), sql, question.TalkID, question.VisitorKey, question.Question); utils.Check(err)
}
