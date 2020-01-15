package database

import (
	"github.com/jmoiron/sqlx"
	"gonference/pkg/domain"
	"gonference/pkg/utils"
)

type RatingRepository struct {
	DB *sqlx.DB
}

func (rr *RatingRepository) Add(domainRating domain.Rating) {
	rating := DomainRatingToRating(domainRating)
	sql := "INSERT INTO rating (talk_id, visitor_key, stars, comment) VALUES ($1, $2, $3, $4)"
	_, err := rr.DB.Exec(sql, rating.TalkID, rating.VisitorKey, rating.Stars, rating.Comment); utils.Check(err)
}

func (rr *RatingRepository) GetByTalkIdAndVisitorKey(talkID int, visitorKey string) *domain.Rating {
	var ratings []RatingModel
	sql := "SELECT id, talk_id, visitor_key, stars, comment FROM rating WHERE talk_id = $1 AND visitor_key = $2"
	utils.Check(rr.DB.Select(&ratings, sql, talkID, visitorKey))
	if len(ratings) == 0 {
		return nil
	}
	firstRating := ratings[0].ToDomainRating()
	return &firstRating
}
