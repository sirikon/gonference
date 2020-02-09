package database

import (
	"gonference/pkg/application"
	"gonference/pkg/infrastructure/database/binders"
	"gonference/pkg/infrastructure/database/client"
	"gonference/pkg/infrastructure/logger"
)

type RatingRepository struct {
	Logger logger.Logger
	DB     *client.DBClient
}

func (rr *RatingRepository) Add(rating *application.Rating) {
	rr.insertQuery(rating)
}

func (rr *RatingRepository) GetByVisitorKey(visitorKey string) []*application.Rating {
	return rr.selectQuery("WHERE visitor_key = $1", visitorKey)
}

func (rr *RatingRepository) GetByTalkId(talkID string) []*application.Rating {
	return rr.selectQuery("WHERE talk_id = $1", talkID)
}

func (rr *RatingRepository) GetByTalkIdAndVisitorKey(talkID string, visitorKey string) *application.Rating {
	return rr.selectOneQuery("WHERE talk_id = $1 AND visitor_key = $2 LIMIT 1", talkID, visitorKey)
}


func (rr *RatingRepository) selectOneQuery(extra string, args ...interface{}) *application.Rating {
	results := rr.selectQuery(extra, args...)
	if len(results) == 0 {
		return nil
	}
	return results[0]
}

func (rr *RatingRepository) selectQuery(extra string, args ...interface{}) []*application.Rating {
	rows := rr.DB.Select(binders.RatingFieldsString, "rating", extra, args...)
	ratings := make([]*application.Rating, 0)
	for rows.Next() {
		ratings = append(ratings, binders.RatingReader(rows.Scan))
	}
	return ratings
}

func (rr *RatingRepository) insertQuery(rating *application.Rating) {
	rr.DB.Insert(binders.RatingFieldsString, "rating", binders.RatingWriter(rating)...)
}
