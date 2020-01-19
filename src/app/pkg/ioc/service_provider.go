package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gonference/pkg/database"
	"gonference/pkg/domain"
	"gonference/pkg/infrastructure/logger"
	"gonference/pkg/web/controllers/api"
	"gonference/pkg/web/controllers/public"
)

// JobContext .
type JobContext struct {
	UID          string
	VisitorKey   string
	dbConnection *sqlx.DB
}

// CreateJobContext .
func CreateJobContext(dbConnection *sqlx.DB) *JobContext {
	return &JobContext{
		dbConnection: dbConnection,
	}
}

// CreateScope .
func (ctx *JobContext) CreateScope(uid, visitorKey string) *JobContext {
	return &JobContext{
		dbConnection: ctx.dbConnection,
		UID:          uid,
		VisitorKey:   visitorKey,
	}
}

// Logger .
func Logger(ctx *JobContext) logger.Logger {
	return contextualizeLogger(logger.Instance, ctx)
}

func LoggerForAccess(ctx *JobContext) logger.Logger {
	return contextualizeLogger(logger.InstanceForAccess, ctx)
}

func contextualizeLogger(logger logger.Logger, ctx *JobContext) logger.Logger {
	return logger.
		WithField("job_uid", ctx.UID).
		WithField("visitor_key", ctx.VisitorKey)
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

func RatingRepository(ctx *JobContext) domain.RatingRepository {
	return &database.RatingRepository{
		Logger: Logger(ctx),
		DB:     DbConnection(ctx),
	}
}

func QuestionRepository(ctx *JobContext) domain.QuestionRepository {
	return &database.QuestionRepository{
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
		TalkRepository:   TalkRepository(ctx),
		RatingRepository: RatingRepository(ctx),
	}
}
func IndexHandler(ctx *JobContext) gin.HandlerFunc { return IndexController(ctx).Handler }

// TalkController .
func TalkController(ctx *JobContext) *public.TalkController {
	return &public.TalkController{
		TalkRepository:     TalkRepository(ctx),
		RatingRepository:   RatingRepository(ctx),
		QuestionRepository: QuestionRepository(ctx),
	}
}
func TalkHandler(ctx *JobContext) gin.HandlerFunc { return TalkController(ctx).Handler }
func TalkPostRatingHandler(ctx *JobContext) gin.HandlerFunc {
	return TalkController(ctx).PostRatingHandler
}
func TalkPostQuestionHandler(ctx *JobContext) gin.HandlerFunc {
	return TalkController(ctx).PostQuestionHandler
}

func LoginController(ctx *JobContext) *public.LoginController {
	return &public.LoginController{
		UserService: UserService(ctx),
	}
}
func LoginGetHandler(ctx *JobContext) gin.HandlerFunc    { return LoginController(ctx).GetHandler }
func LoginPostHandler(ctx *JobContext) gin.HandlerFunc   { return LoginController(ctx).PostHandler }
func LoginLogoutHandler(ctx *JobContext) gin.HandlerFunc { return LoginController(ctx).LogoutHandler }

// TalksAPIController .
func TalksAPIController(ctx *JobContext) *api.TalksAPIController {
	return &api.TalksAPIController{
		TalkRepository: TalkRepository(ctx),
	}
}
func TalksAPIGetHandler(ctx *JobContext) gin.HandlerFunc { return TalksAPIController(ctx).GetHandler }
func TalksAPIGetAllHandler(ctx *JobContext) gin.HandlerFunc {
	return TalksAPIController(ctx).GetAllHandler
}
func TalksAPIAddHandler(ctx *JobContext) gin.HandlerFunc { return TalksAPIController(ctx).AddHandler }
func TalksAPIUpdateHandler(ctx *JobContext) gin.HandlerFunc {
	return TalksAPIController(ctx).UpdateHandler
}
func TalksAPIDeleteHandler(ctx *JobContext) gin.HandlerFunc {
	return TalksAPIController(ctx).DeleteHandler
}

// MeAPIController .
func MeAPIController(ctx *JobContext) *api.MeAPIController {
	return &api.MeAPIController{
		UserService: UserService(ctx),
	}
}
func MeAPIHandler(ctx *JobContext) gin.HandlerFunc { return MeAPIController(ctx).Handler }
func MeAPIChangePasswordHandler(ctx *JobContext) gin.HandlerFunc {
	return MeAPIController(ctx).ChangePasswordHandler
}
