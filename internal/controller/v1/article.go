package v1

import (
	"github.com/blog-service/internal/service"
	"github.com/blog-service/pkg/app"
	"github.com/blog-service/pkg/convert"
	"github.com/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

// Get @Summary 获取单个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (a *Article) Get(c *gin.Context) {
	response := app.NewResponse(c)
	id := c.Param("id")
	if id == "" {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.ArticleService{
		Service: service.New(c.Request.Context()),
	}
	detail, err := svc.Detail(c, convert.StrTo(id).MustInt64())
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDetailArticle)
		return
	}
	response.ToSuccessResponse(detail)
}

// List @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (a *Article) List(c *gin.Context) {
	var params service.ArticleListRequest
	response := app.NewResponse(c)
	_, errs := app.BindAndValid(c, &params)
	if errs != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.ArticleService{
		Service: service.New(c.Request.Context()),
	}
	list, err := svc.List(c, params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorListArticle)
		return
	}
	count, err := svc.Count(c, params)
	response.ToResponseList(list, count)
}

// Create @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (a *Article) Create(c *gin.Context) {
	var params *service.ArticleCreateRequest
	response := app.NewResponse(c)
	_, errs := app.BindAndValid(c, &params)
	if errs != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.ArticleService{
		Service: service.New(c.Request.Context()),
	}
	_, err := svc.Create(c, params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateArticle)
		return
	}
	response.ToSuccessResponse(nil)
}

// Update @Summary 更新标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (a *Article) Update(c *gin.Context) {
	var params *service.ArticleUpdateRequest
	response := app.NewResponse(c)
	_, errs := app.BindAndValid(c, &params)
	if errs != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.ArticleService{
		Service: service.New(c.Request.Context()),
	}
	err := svc.Update(c, params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUpdateArticle)
		return
	}
	response.ToSuccessResponse(nil)
}

// Delete @Summary 删除标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (a *Article) Delete(c *gin.Context) {
	response := app.NewResponse(c)
	id := c.Param("id")
	if id == "" {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.ArticleService{
		Service: service.New(c.Request.Context()),
	}
	err := svc.Delete(c, convert.StrTo(id).MustInt64())
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteArticle)
		return
	}
	response.ToSuccessResponse(nil)
}
