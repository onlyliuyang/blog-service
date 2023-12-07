package v1

import (
	"github.com/blog-service/internal/service"
	"github.com/blog-service/pkg/app"
	"github.com/blog-service/pkg/convert"
	"github.com/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"strings"
)

type UsersAccount struct {
}

/**
用户注册逻辑
*/
func (u *UsersAccount) Register(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	var registerRequest *service.UserAccountRegisterRequest
	_, errs := app.BindAndValid(ctx, &registerRequest)
	if errs != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.UserAccountService{
		Service: service.New(ctx),
	}

	err := svc.Resister(ctx, registerRequest)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorAccountRegister.WithDetails(errs.Errors()...))
		return
	}
	response.ToSuccessResponse(nil)
}

/**
获取用户信息
*/
func (u *UsersAccount) GetUserInfo(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	userId, _ := ctx.GetQuery("user_id")
	userId = strings.Trim(userId, " ")

	if userId == "" || len(userId) <= 0 {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails())
		return
	}

	svc := service.UserAccountService{
		Service: service.New(ctx),
	}

	userInfo, err := svc.GetUserInfo(ctx, convert.StrTo(userId).MustInt())
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetUserInfo.WithDetails(err.Error()))
		return
	}
	response.ToSuccessResponse(userInfo)
}

/**
导入一下用户列表
*/
func (u *UsersAccount) ImportUserList(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.UserAccountService{
		Service: service.New(ctx),
	}
	svc.ImportUserList(ctx)
	response.ToSuccessResponse(nil)
}

/**
获取用户列表
*/
func (u *UsersAccount) GetUserList(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	var listRequest *service.UserAccountListRequest
	_, errs := app.BindAndValid(ctx, &listRequest)
	if errs != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.UserAccountService{
		Service: service.New(ctx),
	}

	userList, err := svc.GetUserList(ctx, listRequest)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUserList.WithDetails(errs.Errors()...))
		return
	}
	response.ToSuccessResponse(userList)
}
