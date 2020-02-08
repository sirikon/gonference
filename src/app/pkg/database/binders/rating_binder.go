package binders

import (
	"gonference/pkg/domain"
	"strings"
)

var RatingFields = []string{"id", "talk_id", "visitor_key", "comment"}
var RatingFieldsString = strings.Join(RatingFields, ", ")
func RatingWriter(rating *domain.Rating) []interface{} {
	return []interface{}{
		rating.TalkID,
		rating.VisitorKey,
		rating.Stars,
		rating.Comment,
	}
}
func RatingReader(scanner Scanner) *domain.Rating {
	rating := &domain.Rating{}
	scan(scanner,
		&rating.ID,
		&rating.TalkID,
		&rating.VisitorKey,
		&rating.Stars,
		&rating.Comment)
	return rating
}
