package v1

import (
	"github.com/blog-service/internal/service"
	"github.com/blog-service/pkg/app"
	"github.com/blog-service/pkg/convert"
	"github.com/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Comment struct {
}

func (c *Comment) Create(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.CommentService{
		Service: service.New(ctx),
	}

	var params service.CommentCreateRequest
	_, errs := app.BindAndValid(ctx, &params)
	if errs != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	id, err := svc.Create(ctx, &params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateComment.WithDetails(err.Error()))
		return
	}
	data := map[string]interface{}{"id": id}
	response.ToSuccessResponse(data)
}

func (c *Comment) Delete(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.CommentService{
		Service: service.New(ctx),
	}

	commentId := ctx.Query("comment_id")
	if commentId == "" {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails("comment_id不能为空"))
	}

	errs := svc.Delete(ctx, convert.StrTo(commentId).MustInt64())
	if errs != nil {
		response.ToErrorResponse(errcode.ErrorDeleteComment.WithDetails(errs.Error()))
		return
	}
	response.ToSuccessResponse(nil)
}
