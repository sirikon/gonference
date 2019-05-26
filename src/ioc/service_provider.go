package ioc

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirikon/gonference/src/database"
	"github.com/sirikon/gonference/src/domain"
	"github.com/sirikon/gonference/src/utils"
	"github.com/sirikon/gonference/src/web/controllers/api"
	"github.com/sirikon/gonference/src/web/controllers/public"
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

// CreateScope .
func (sp *ServiceProvider) CreateScope() *ServiceProvider {
	return &ServiceProvider{
		dbConnection: sp.dbConnection,
		requestInfo: &RequestInfo{
			UID: utils.RandomString(32),
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
func (sp *ServiceProvider) GetTalkRepository() domain.TalkRepository {
	return &database.TalkRepository{
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

func (sp *ServiceProvider) GetLoginController() *public.LoginController {
	return &public.LoginController{
	}
}

// GetTalksAPIController .
func (sp *ServiceProvider) GetTalksAPIController() *api.TalksAPIController {
	return &api.TalksAPIController{
		TalkRepository: sp.GetTalkRepository(),
	}
}

// GetMeAPIController .
func (sp *ServiceProvider) GetMeAPIController() *api.MeAPIController {
	return &api.MeAPIController{}
}
