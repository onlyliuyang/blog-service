package v1

import (
	"github.com/blog-service/internal/service"
	"github.com/blog-service/pkg/app"
	"github.com/blog-service/pkg/convert"
	"github.com/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Category struct{}

func (c *Category) Create(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	var params *service.CategoryCreateRequest
	_, errs := app.BindAndValid(ctx, &params)
	if errs != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.CategoryService{
		Service: service.New(ctx),
	}
	err := svc.Create(ctx, params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateCategory)
		return
	}
	response.ToSuccessResponse(nil)
}

func (c *Category) Update(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	var params *service.CategoryUpdateRequest
	_, errs := app.BindAndValid(ctx, &params)
	if errs != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.CategoryService{
		Service: service.New(ctx),
	}
	err := svc.Update(ctx, params.Id, params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUpdateCategory.WithDetails(err.Error()))
		return
	}
	response.ToSuccessResponse(nil)
}

func (c *Category) Delete(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	id := ctx.Param("id")
	if id == "" {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.CategoryService{
		Service: service.New(ctx),
	}
	err := svc.Delete(ctx, convert.StrTo(id).MustInt64())
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteCategory.WithDetails(err.Error()))
		return
	}
	response.ToSuccessResponse(nil)
}

func (c *Category) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageIndex := app.GetPage(ctx)
	pageSize := app.GetPageSize(ctx)
	svc := service.CategoryService{Service: service.New(ctx)}
	data, err := svc.List(ctx, pageIndex, pageSize)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetCategoryList.WithDetails(err.Error()))
		return
	}
	count, _ := svc.Count(ctx)
	response.ToResponseList(data, count)
}
