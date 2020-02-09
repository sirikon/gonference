package binders

import (
	"gonference/pkg/application"
	"strings"
)

var RatingFields = []string{"id", "talk_id", "visitor_key", "stars", "comment"}
var RatingFieldsString = strings.Join(RatingFields, ", ")
func RatingWriter(rating *application.Rating) []interface{} {
	return []interface{}{
		rating.ID,
		rating.TalkID,
		rating.VisitorKey,
		rating.Stars,
		rating.Comment,
	}
}
func RatingReader(scanner Scanner) *application.Rating {
	rating := &application.Rating{}
	scan(scanner,
		&rating.ID,
		&rating.TalkID,
		&rating.VisitorKey,
		&rating.Stars,
		&rating.Comment)
	return rating
}
