package model

import (
	"github.com/blog-service/pkg/app"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Article struct {
	Id       int64  `gorm:"column:article_id"`
	Url      string `gorm:"column:article_url"`
	Show     int    `gorm:"column:article_show"`
	Sort     int    `gorm:"column:article_sort"`
	Title    string `gorm:"column:article_title"`
	Content  string `gorm:"column:article_content"`
	Pic      string `gorm:"column:article_pic"`
	AuthorId int    `gorm:"column:author_id"`
	*Model
}

func (a *Article) TableName() string {
	return "bs_article"
}

/*
创建文章
*/
func (a *Article) CreateArticle(ctx *gin.Context, db *gorm.DB) (int64, error) {
	var err error
	err = db.WithContext(ctx).Table(a.TableName()).Create(a).Error
	return a.Id, err
}

/*
更新文章
*/
func (a *Article) UpdateArticle(ctx *gin.Context, db *gorm.DB, id int64, updates map[string]interface{}) error {
	var err error
	err = db.WithContext(ctx).Table(a.TableName()).Where("article_id = ?", id).Updates(updates).Error
	return err
}

/*
获取文章列表
*/
func (a *Article) ArticleList(ctx *gin.Context, db *gorm.DB) (list []*Article, err error) {
	offset := app.GetPageOffset(app.GetPage(ctx), app.GetPageSize(ctx))
	limit := app.GetPageSize(ctx)
	err = db.WithContext(ctx).Table(a.TableName()).Offset(offset).Limit(limit).Find(&list).Error
	return
}

/*
获取文章总数
*/
func (a *Article) ArticleCount(ctx *gin.Context, db *gorm.DB) (count int64, err error) {
	err = db.WithContext(ctx).Table(a.TableName()).Count(&count).Error
	return
}

/*
删除文章
*/
func (a *Article) ArticleDelete(ctx *gin.Context, db *gorm.DB, id int64) error {
	err := db.WithContext(ctx).Table(a.TableName()).Where("article_id = ?", id).Delete(nil).Error
	return err
}

/*
获取文章详情
*/
func (a *Article) ArticleById(ctx *gin.Context, db *gorm.DB, id int64) (info *Article, err error) {
	err = db.WithContext(ctx).Where("article_id = ?", id).Find(&info).Error
	return
}
