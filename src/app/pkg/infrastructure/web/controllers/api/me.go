package api

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/application"
	"gonference/pkg/infrastructure/web/models"
	"gonference/pkg/infrastructure/web/session"
	"gonference/pkg/utils"
	"net/http"
)

// MeAPIController .
type MeAPIController struct {
	UserService application.UserService
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
	utils.Check(ctx.BindJSON(&vm))

	checkNewPasswordMatchesWithRepetition(vm)

	username := session.GetSession(ctx).Get(session.UsernameKey)
	s.UserService.ChangePassword(username, vm.CurrentPassword, vm.NewPassword)

	ctx.Status(http.StatusOK)
}

func checkNewPasswordMatchesWithRepetition(vm ChangePasswordViewModel)  {
	if vm.NewPassword != vm.RepeatNewPassword {
		panic(models.UserError{Message: "New passwords doesn't match"})
	}
}
