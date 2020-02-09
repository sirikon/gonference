package binders

import (
	"gonference/pkg/application"
	"strings"
)

var QuestionFields = []string{"id", "talk_id", "visitor_key", "question"}
var QuestionFieldsString = strings.Join(QuestionFields, ", ")
func QuestionWriter(question *application.Question) []interface{} {
	return []interface{}{
		question.ID,
		question.TalkID,
		question.VisitorKey,
		question.Question,
	}
}
func QuestionReader(scanner Scanner) *application.Question {
	question := &application.Question{}
	scan(scanner,
		&question.ID,
		&question.TalkID,
		&question.VisitorKey,
		&question.Question)
	return question
}
