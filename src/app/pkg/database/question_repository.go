package database

import (
	"github.com/jmoiron/sqlx"
	"gonference/pkg/domain"
	"gonference/pkg/utils"
)

type QuestionRepository struct {
	DB *sqlx.DB
}

func (qr *QuestionRepository) Add(domainQuestion domain.Question) {
	question := DomainQuestionToQuestion(domainQuestion)
	sql := "INSERT INTO question (talk_id, visitor_key, question) VALUES ($1, $2, $3)"
	_, err := qr.DB.Exec(sql, question.TalkID, question.VisitorKey, question.Question); utils.Check(err)
}
