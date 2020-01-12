package api

import (
	"github.com/gin-gonic/gin"
	"gonference/pkg/domain"
	"gonference/pkg/utils"
	"gonference/pkg/web/models"
	"gonference/pkg/web/session"
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
	utils.Check(ctx.BindJSON(&vm))

	checkNewPasswordMatchesWithRepetition(vm)

	username := session.GetSession(ctx).Get(session.UsernameKey)
	utils.Check(s.UserService.ChangePassword(username, vm.CurrentPassword, vm.NewPassword))

	ctx.Status(http.StatusOK)
}

func checkNewPasswordMatchesWithRepetition(vm ChangePasswordViewModel)  {
	if vm.NewPassword != vm.RepeatNewPassword {
		panic(models.UserError{"New passwords doesn't match"})
	}
}
