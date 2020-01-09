package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirikon/gonference/src/domain"
	"github.com/sirikon/gonference/src/web/session"
	"net/http"
)

// MeAPIController .
type MeAPIController struct {
	UserService domain.UserService
}

// Handler .
func (s *MeAPIController) Handler(ctx *gin.Context) {
	user := User{
		Username: session.GetSession(ctx).Get(session.UsernameKey),
	}

	ctx.JSON(http.StatusOK, user)
}

func (s *MeAPIController) ChangePasswordHandler(ctx *gin.Context) {
	var vm ChangePasswordViewModel
	err := ctx.BindJSON(&vm)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if vm.NewPassword != vm.RepeatNewPassword {
		_ = ctx.Error(errors.New("Passwords doesn't match"))
	}

	err = s.UserService.ChangePassword(session.GetSession(ctx).Get(session.UsernameKey), vm.CurrentPassword, vm.NewPassword)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
