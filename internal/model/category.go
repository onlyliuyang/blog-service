package model

import (
	"github.com/blog-service/pkg/app"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Category struct {
	Id       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Level    int    `gorm:"column:level"`
	ParentId int    `gorm:"column:parent_id"`
	Status   int    `gorm:"column:status"`
	*Model
}

func (c *Category) TableName() string {
	return "bs_category"
}

/*
新增标签
*/
func (c *Category) Create(ctx *gin.Context, db *gorm.DB) (id int64, err error) {
	err = db.WithContext(ctx).Table(c.TableName()).Create(c).Error
	return c.Id, err
}

/*
更新标签
*/
func (c *Category) Update(ctx *gin.Context, db *gorm.DB, id int64, updates map[string]interface{}) error {
	var err error
	err = db.WithContext(ctx).Table(c.TableName()).Where("id = ?", id).Updates(updates).Error
	return err
}

/*
删除标签
*/
func (c *Category) Delete(ctx *gin.Context, db *gorm.DB, id int64) error {
	var err error
	err = db.WithContext(ctx).Table(c.TableName()).Where("id = ?", id).Error
	return err
}

/*
批量获取标签
*/
func (c *Category) MultiGet(ctx *gin.Context, db *gorm.DB, ids []int64) (list []*Category, err error) {
	err = db.WithContext(ctx).Table(c.TableName()).Where(" id in ?", ids).Find(&list).Error
	return
}

/*
获取单个标签
*/
func (c *Category) Get(ctx *gin.Context, db *gorm.DB, id int64) (info *Category, err error) {
	err = db.WithContext(ctx).Table(c.TableName()).Where("id = ?", id).Find(&info).Error
	return
}

/*
获取标签列表
*/
func (c *Category) List(ctx *gin.Context, db *gorm.DB, pageIndex, pageSize int) (list []*Category, err error) {
	offset := app.GetPageOffset(pageIndex, pageSize)
	limit := pageSize
	err = db.WithContext(ctx).Table(c.TableName()).Offset(offset).Limit(limit).Find(&list).Error
	return
}

/*
获取标签总数
*/
func (c *Category) Count(ctx *gin.Context, db *gorm.DB) (count int64, err error) {
	err = db.WithContext(ctx).Table(c.TableName()).Count(&count).Error
	return
}
