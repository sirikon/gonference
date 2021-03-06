package ioc

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/application"
	"gonference/pkg/infrastructure/database"
	"gonference/pkg/infrastructure/database/client"
	"gonference/pkg/infrastructure/logger"
	"gonference/pkg/infrastructure/web/controllers/api"
	"gonference/pkg/infrastructure/web/controllers/public"
)

// JobContext .
type JobContext struct {
	UID        string
	VisitorKey string
	db         *client.DBClient
}

// CreateJobContext .
func CreateJobContext(db *client.DBClient) *JobContext {
	return &JobContext{
		db: db,
	}
}

// CreateScope .
func (ctx *JobContext) CreateScope(uid, visitorKey string) *JobContext {
	return &JobContext{
		db:         ctx.db,
		UID:        uid,
		VisitorKey: visitorKey,
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

// TalkRepository .
func TalkRepository(ctx *JobContext) application.TalkRepository {
	return &database.TalkRepository{
		DB:     ctx.db,
		Logger: Logger(ctx),
	}
}

func RatingRepository(ctx *JobContext) application.RatingRepository {
	return &database.RatingRepository{
		Logger: Logger(ctx),
		DB:     ctx.db,
	}
}

func QuestionRepository(ctx *JobContext) application.QuestionRepository {
	return &database.QuestionRepository{
		DB:     ctx.db,
		Logger: Logger(ctx),
	}
}

func UserService(ctx *JobContext) application.UserService {
	return &database.UserService{
		DB:     ctx.db,
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
func TalkGetRatingsHandler(ctx *JobContext) gin.HandlerFunc {
	return TalkController(ctx).GetRatingsHandler
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
		QuestionRepository: QuestionRepository(ctx),
	}
}
func TalksAPIGetHandler(ctx *JobContext) gin.HandlerFunc { return TalksAPIController(ctx).GetHandler }
func TalksAPIGetAllHandler(ctx *JobContext) gin.HandlerFunc {
	return TalksAPIController(ctx).GetAllHandler
}
func TalksAPIGetTalkQuestionsHandler(ctx *JobContext) gin.HandlerFunc {
	return TalksAPIController(ctx).GetTalkQuestionsHandler
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
