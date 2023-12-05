package v1

import (
	"github.com/blog-service/internal/service"
	"github.com/blog-service/pkg/app"
	"github.com/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
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
