package database

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
	"gonference/pkg/domain"
	"gonference/pkg/infrastructure/logger"
)

type RatingRepository struct {
	Logger logger.Logger
	NewPool *pgxpool.Pool
	DB *sqlx.DB
}

func (rr *RatingRepository) Add(rating *domain.Rating) {
	rr.insertQuery(rating)
}

func (rr *RatingRepository) GetByVisitorKey(visitorKey string) []*domain.Rating {
	return rr.selectQuery("WHERE visitor_key = $1", visitorKey)
}

func (rr *RatingRepository) GetByTalkId(talkID int) []*domain.Rating {
	return rr.selectQuery("WHERE talk_id = $1", talkID)
}

func (rr *RatingRepository) GetByTalkIdAndVisitorKey(talkID int, visitorKey string) *domain.Rating {
	return rr.selectOneQuery("WHERE talk_id = $1 AND visitor_key = $2 LIMIT 1", talkID, visitorKey)
}


func (rr *RatingRepository) selectOneQuery(extra string, args ...interface{}) *domain.Rating {
	results := rr.selectQuery(extra, args...)
	if len(results) == 0 {
		return nil
	}
	return results[0]
}

func (rr *RatingRepository) selectQuery(extra string, args ...interface{}) []*domain.Rating {
	rows := selectQuery(rr.NewPool, ratingFields, "rating", extra, args...)
	ratings := make([]*domain.Rating, 0)
	for rows.Next() {
		ratings = append(ratings, ratingReader(rows.Scan))
	}
	return ratings
}

func (rr *RatingRepository) insertQuery(rating *domain.Rating) {
	insertQuery(rr.NewPool, ratingFields, "rating", ratingWriter(rating))
}
