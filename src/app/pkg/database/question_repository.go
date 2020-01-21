package database

import (
	"github.com/jmoiron/sqlx"
	"gonference/pkg/domain"
	"gonference/pkg/infrastructure/logger"
	"gonference/pkg/utils"
)

type QuestionRepository struct {
	Logger logger.Logger
	DB *sqlx.DB
}

func (qr *QuestionRepository) GetByTalkId(talkId int) []domain.Question {
	var questions []QuestionModel
	query := "SELECT * FROM question WHERE talk_id = $1 ORDER BY id"
	logSelect(qr.Logger, query)
	utils.Check(qr.DB.Select(&questions, query, talkId))
	return QuestionsToDomainQuestions(questions)
}

func (qr *QuestionRepository) Add(domainQuestion domain.Question) {
	question := DomainQuestionToQuestion(domainQuestion)
	sql := "INSERT INTO question (talk_id, visitor_key, question) VALUES ($1, $2, $3)"
	logMutation(qr.Logger, sql)
	_, err := qr.DB.Exec(sql, question.TalkID, question.VisitorKey, question.Question); utils.Check(err)
}
