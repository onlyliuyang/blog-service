package v1

import (
	"github.com/blog-service/internal/service"
	"github.com/blog-service/pkg/app"
	"github.com/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Author struct {
}

func (a *Author) Login(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	var authorLoginRequest *service.AuthorLoginRequest
	_, errs := app.BindAndValid(ctx, &authorLoginRequest)
	if errs != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.AuthorService{
		Service: service.New(ctx.Request.Context()),
	}
	authorInfo, err := svc.GetAuthor(ctx, authorLoginRequest)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorNotFoundAuthor)
		return
	}

	token, err := app.GenerateToken(strconv.FormatInt(authorInfo.Id, 10), authorInfo.Password)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorAccountLogin)
		return
	}

	response.ToSuccessResponse(token)
}

func (a *Author) Logout(ctx *gin.Context) {

}
