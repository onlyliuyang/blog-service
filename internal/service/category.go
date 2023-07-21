package service

import (
	"github.com/blog-service/global"
	"github.com/blog-service/internal/dao"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"time"
)

type CategoryCreateRequest struct {
	Name     string `form:"name" binding:"required"`
	Level    int    `form:"level" binding:"required"`
	ParentId int    `form:"parent_id" binding:"required"`
	Status   int    `form:"status" binding:"required"`
}

type CategoryUpdateRequest struct {
	Id       int64  `form:"id" binding:"required"`
	Name     string `form:"name"`
	Level    int    `form:"level"`
	ParentId int    `form:"parent_id"`
	Status   int    `form:"status"`
}

type CategoryListResponse struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Level     int       `json:"level"`
	ParentId  int       `json:"parent_id"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryService struct {
	Service
}

func (svc *CategoryService) Create(ctx *gin.Context, params *CategoryCreateRequest) error {
	var err error
	var category dao.Category
	err = copier.Copy(&category, params)
	if err != nil {
		global.Logger.Errorof(ctx, "创建标签失败: %s", err.Error())
		return err
	}

	err = svc.dao.CreateCategory(ctx, &category)
	if err != nil {
		global.Logger.Errorof(ctx, "创建标签失败: %s", err.Error())
		return err
	}
	return nil
}

func (svc *CategoryService) Update(ctx *gin.Context, id int64, params *CategoryUpdateRequest) error {
	var err error
	var category dao.Category
	err = copier.Copy(&category, params)
	if err != nil {
		global.Logger.Errorof(ctx, "更新标签失败: %s", err.Error())
		return err
	}

	err = svc.dao.UpdateCategory(ctx, id, &category)
	if err != nil {
		global.Logger.Errorof(ctx, "更新标签失败: %s", err.Error())
		return err
	}
	return nil
}

func (svc *CategoryService) Delete(ctx *gin.Context, id int64) error {
	var err error
	_, err = svc.dao.GetCategory(ctx, id)
	if err != nil {
		global.Logger.Errorof(ctx, "删除标签失败，标签查询错误: %s", err.Error())
		return err
	}

	err = svc.dao.DeleteCategory(ctx, id)
	if err != nil {
		global.Logger.Errorof(ctx, "删除标签失败: %s", err)
		return err
	}
	return nil
}

func (svc *CategoryService) List(ctx *gin.Context, pageIndex, pageSize int) (list []*CategoryListResponse, err error) {
	list = make([]*CategoryListResponse, 0)
	data, err := svc.dao.ListCategory(ctx, pageIndex, pageSize)
	if err != nil {
		global.Logger.Errorof(ctx, "获取标签列表失败: %s", err.Error())
		return
	}

	err = copier.Copy(&list, data)
	if err != nil {
		global.Logger.Errorof(ctx, "获取标签列表失败: %s", err.Error())
		return
	}
	return
}

func (svc *CategoryService) Count(ctx *gin.Context) (count int64, err error) {
	return svc.dao.CountCategory(ctx)
}
