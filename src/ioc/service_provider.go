package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirikon/gonference/src/database"
	"github.com/sirikon/gonference/src/domain"
	"github.com/sirikon/gonference/src/utils"
	"github.com/sirikon/gonference/src/web/controllers/api"
	"github.com/sirikon/gonference/src/web/controllers/public"
	"github.com/sirupsen/logrus"
)

// JobContext .
type JobContext struct {
	UID string
	dbConnection *sqlx.DB
}

// CreateJobContext .
func CreateJobContext(dbConnection *sqlx.DB) *JobContext {
	return &JobContext{
		dbConnection: dbConnection,
	}
}

// CreateScope .
func (ctx *JobContext) CreateScope() *JobContext {
	return &JobContext{
		dbConnection: ctx.dbConnection,
		UID:          utils.RandomString(32),
	}
}

// Logger .
func Logger(ctx *JobContext) *logrus.Entry {
	return logrus.WithField("job_uid", ctx.UID)
}

// DbConnection .
func DbConnection(ctx *JobContext) *sqlx.DB {
	return ctx.dbConnection
}

// TalkRepository .
func TalkRepository(ctx *JobContext) domain.TalkRepository {
	return &database.TalkRepository{
		DB:     DbConnection(ctx),
		Logger: Logger(ctx),
	}
}

func UserService(ctx *JobContext) domain.UserService {
	return &database.UserService{
		DB:     DbConnection(ctx),
		Logger: Logger(ctx),
	}
}

// IndexController .
func IndexController(ctx *JobContext) *public.IndexController {
	return &public.IndexController{
		TalkRepository: TalkRepository(ctx),
	}
}
func IndexHandler(ctx *JobContext) gin.HandlerFunc { return IndexController(ctx).Handler }

// TalkController .
func TalkController(ctx *JobContext) *public.TalkController {
	return &public.TalkController{
		TalkRepository: TalkRepository(ctx),
	}
}
func TalkHandler(ctx *JobContext) gin.HandlerFunc { return TalkController(ctx).Handler }

func LoginController(ctx *JobContext) *public.LoginController {
	return &public.LoginController{
		UserService: UserService(ctx),
	}
}
func LoginGetHandler(ctx *JobContext) gin.HandlerFunc { return LoginController(ctx).GetHandler }
func LoginPostHandler(ctx *JobContext) gin.HandlerFunc { return LoginController(ctx).PostHandler }
func LoginLogoutHandler(ctx *JobContext) gin.HandlerFunc { return LoginController(ctx).LogoutHandler }

// TalksAPIController .
func TalksAPIController(ctx *JobContext) *api.TalksAPIController {
	return &api.TalksAPIController{
		TalkRepository: TalkRepository(ctx),
	}
}
func TalksAPIGetHandler(ctx *JobContext) gin.HandlerFunc { return TalksAPIController(ctx).GetHandler }
func TalksAPIGetAllHandler(ctx *JobContext) gin.HandlerFunc { return TalksAPIController(ctx).GetAllHandler }
func TalksAPIAddHandler(ctx *JobContext) gin.HandlerFunc { return TalksAPIController(ctx).AddHandler }
func TalksAPIUpdateHandler(ctx *JobContext) gin.HandlerFunc { return TalksAPIController(ctx).UpdateHandler }
func TalksAPIDeleteHandler(ctx *JobContext) gin.HandlerFunc { return TalksAPIController(ctx).DeleteHandler }

// MeAPIController .
func MeAPIController(ctx *JobContext) *api.MeAPIController {
	return &api.MeAPIController{
		UserService: UserService(ctx),
	}
}
func MeAPIHandler(ctx *JobContext) gin.HandlerFunc { return MeAPIController(ctx).Handler }
func MeAPIChangePasswordHandler(ctx *JobContext) gin.HandlerFunc { return MeAPIController(ctx).ChangePasswordHandler }
