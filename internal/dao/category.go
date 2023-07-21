package dao

import (
	"encoding/json"
	"github.com/blog-service/internal/model"
	"github.com/blog-service/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type Category struct {
	Name     string `gorm:"column:name" json:"name"`
	Level    int    `gorm:"column:level" json:"level"`
	ParentId int    `gorm:"column:parent_id" json:"parent_id"`
	Status   int    `gorm:"column:status" json:"status"`
}

func (d *Dao) CreateCategory(ctx *gin.Context, params *Category) error {
	categoryId := util.Uuid(ctx)
	c := &model.Category{Id: categoryId}
	_ = copier.Copy(&c, &params)
	_, err := c.Create(ctx, d.engine)
	return err
}

func (d *Dao) UpdateCategory(ctx *gin.Context, id int64, params *Category) error {
	var c *model.Category
	updates := make(map[string]interface{})
	bytes, err := json.Marshal(params)
	if err != nil {
		return err
	}
	_ = json.Unmarshal(bytes, &updates)
	err = c.Update(ctx, d.engine, id, updates)
	return err
}

func (d *Dao) GetCategory(ctx *gin.Context, id int64) (info *model.Category, err error) {
	var c *model.Category
	info, err = c.Get(ctx, d.engine, id)
	return
}

func (d *Dao) MultiGetCategory(ctx *gin.Context, ids []int64) (list []*model.Category, err error) {
	var c *model.Category
	list, err = c.MultiGet(ctx, d.engine, ids)
	return
}

func (d *Dao) DeleteCategory(ctx *gin.Context, id int64) error {
	var c *model.Category
	err := c.Delete(ctx, d.engine, id)
	return err
}

func (d *Dao) ListCategory(ctx *gin.Context, pageIndex, pageSize int) (list []*model.Category, err error) {
	var c *model.Category
	list, err = c.List(ctx, d.engine, pageIndex, pageSize)
	return
}

func (d *Dao) CountCategory(ctx *gin.Context) (count int64, err error) {
	var c *model.Category
	count, err = c.Count(ctx, d.engine)
	return
}
