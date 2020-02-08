package binders

import (
	"gonference/pkg/domain"
	"strings"
)

var QuestionFields = []string{"id", "talk_id", "visitor_key", "question"}
var QuestionFieldsString = strings.Join(QuestionFields, ", ")
func QuestionWriter(question *domain.Question) []interface{} {
	return []interface{}{
		question.ID,
		question.TalkID,
		question.VisitorKey,
		question.Question,
	}
}
func QuestionReader(scanner Scanner) *domain.Question {
	question := &domain.Question{}
	scan(scanner,
		&question.ID,
		&question.TalkID,
		&question.VisitorKey,
		&question.Question)
	return question
}
