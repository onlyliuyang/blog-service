package model

import (
	"github.com/blog-service/pkg/app"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Author struct {
	Id          int64  `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Password    string `gorm:"column:password"`
	Mobile      string `gorm:"column:mobile"`
	HeadUrl     string `gorm:"column:head_url"`
	CountryCode int    `gorm:"column:country_code"`
	Status      int    `gorm:"column:status"`
	*Model
}

func (a *Author) TableName() string {
	return "bs_author"
}

/*
创建作者
*/
func (a *Author) CreateAuthor(ctx *gin.Context, db *gorm.DB) (int64, error) {
	var err error
	err = db.WithContext(ctx).Table(a.TableName()).Create(a).Error
	return a.Id, err
}

/*
更新作者
*/
func (a *Author) UpdateAuthor(ctx *gin.Context, db *gorm.DB, id int64, updates map[string]interface{}) error {
	var err error
	err = db.WithContext(ctx).Table(a.TableName()).Where("id = ?", id).Updates(updates).Error
	return err
}

/*
获取作者列表
*/
func (a *Author) AuthorList(ctx *gin.Context, db *gorm.DB) (list []*Author, err error) {
	if a.Id != 0 {
		db.Where("id = ?", a.Id)
	}

	if a.Name != "" {
		db.Where(" name like ?", a.Name)
	}

	if a.Mobile != "" {
		db.Where("mobile = ?", a.Mobile)
	}

	offset := app.GetPageOffset(app.GetPage(ctx), app.GetPageSize(ctx))
	limit := app.GetPageSize(ctx)
	err = db.WithContext(ctx).Table(a.TableName()).Offset(offset).Limit(limit).Find(&list).Error
	return
}

/*
获取作者总数
*/
func (a *Author) AuthorCount(ctx *gin.Context, db *gorm.DB) (count int64, err error) {
	err = db.WithContext(ctx).Table(a.TableName()).Count(&count).Error
	return
}

/*
删除作者
*/
func (a *Author) AuthorDelete(ctx *gin.Context, db *gorm.DB, id int64) error {
	err := db.WithContext(ctx).Table(a.TableName()).Where("id = ?", id).Delete(nil).Error
	return err
}

/*
获取作者详情
*/
func (a *Author) AuthorById(ctx *gin.Context, db *gorm.DB, id int64) (info *Author, err error) {
	err = db.WithContext(ctx).Where("aid = ?", id).Find(&info).Error
	return
}

func (a *Author) AuthorByName(ctx *gin.Context, db *gorm.DB, name, password string) (info *Author, err error) {
	err = db.WithContext(ctx).Where("name", name).Where("password", password).Find(&info).Error
	return
}
