package ioc

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirikon/gonference"
	"github.com/sirikon/gonference/http/controllers/api"
	"github.com/sirikon/gonference/http/controllers/public"
	"github.com/sirikon/gonference/postgres"
	"github.com/sirikon/gonference/util"
	"github.com/sirupsen/logrus"
)

// RequestInfo .
type RequestInfo struct {
	UID string
}

// ServiceProvider .
type ServiceProvider struct {
	dbConnection *sqlx.DB
	requestInfo  *RequestInfo
}

// CreateServiceProvider .
func CreateServiceProvider(dbConnection *sqlx.DB) *ServiceProvider {
	return &ServiceProvider{
		dbConnection: dbConnection,
	}
}

// CreateRequestScope .
func (sp *ServiceProvider) CreateRequestScope() *ServiceProvider {
	return &ServiceProvider{
		dbConnection: sp.dbConnection,
		requestInfo: &RequestInfo{
			UID: util.RandomString(32),
		},
	}
}

// GetLogger .
func (sp *ServiceProvider) GetLogger() *logrus.Entry {
	return logrus.WithField("request_uid", sp.requestInfo.UID)
}

// GetDbConnection .
func (sp *ServiceProvider) GetDbConnection() *sqlx.DB {
	return sp.dbConnection
}

// GetTalkRepository .
func (sp *ServiceProvider) GetTalkRepository() gonference.TalkRepository {
	return &postgres.TalkRepository{
		DB:     sp.GetDbConnection(),
		Logger: sp.GetLogger(),
	}
}

// GetIndexController .
func (sp *ServiceProvider) GetIndexController() *public.IndexController {
	return &public.IndexController{
		TalkRepository: sp.GetTalkRepository(),
	}
}

// GetTalksAPIController .
func (sp *ServiceProvider) GetTalksAPIController() *api.TalksAPIController {
	return &api.TalksAPIController{
		TalkRepository: sp.GetTalkRepository(),
	}
}
