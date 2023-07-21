package dao

import (
	"fmt"
	"github.com/blog-service/global"
	"github.com/blog-service/internal/model"
	"github.com/blog-service/pkg/util"
	"github.com/gin-gonic/gin"
)

var (
	cacheArticleCount       string = "cache::article::count"
	cacheAuthorArticleCount string = "cache::article::%d::count"
)

func (d *Dao) CreateArticle(ctx *gin.Context, show, sort, authorId int, url, title, content string) (int64, error) {
	articleId := util.Uuid(ctx)
	article := model.Article{
		Id:       articleId,
		Url:      url,
		Show:     show,
		Sort:     sort,
		Title:    title,
		Content:  content,
		Pic:      "",
		AuthorId: authorId,
	}
	id, err := article.CreateArticle(ctx, d.engine)
	if err == nil {
		global.RedisDB.Incr(d.getCacheArticleKey())
		global.RedisDB.Incr(d.getCacheAuthorArticleKey(authorId))
	}
	return id, err
}

func (d *Dao) UpdateArticle(ctx *gin.Context, id int64, updates map[string]interface{}) error {
	var article *model.Article
	err := article.UpdateArticle(ctx, d.engine, id, updates)
	return err
}

func (d *Dao) ListArticle(ctx *gin.Context) (list []*model.Article, err error) {
	var article *model.Article
	list, err = article.ArticleList(ctx, d.engine)
	return list, err
}

func (d *Dao) CountArticle(ctx *gin.Context) (count int64, err error) {
	var article *model.Article
	count, err = article.ArticleCount(ctx, d.engine)
	return count, err
}

func (d *Dao) DeleteArticle(ctx *gin.Context, id int64) error {
	var article *model.Article
	articleInfo, err := article.ArticleById(ctx, d.engine, id)
	if err != nil {
		return err
	}

	err = article.ArticleDelete(ctx, d.engine, id)
	if err == nil {
		global.RedisDB.Decr(d.getCacheArticleKey())
		global.RedisDB.Decr(d.getCacheAuthorArticleKey(articleInfo.AuthorId))
	}
	return err
}

func (d *Dao) DetailArticle(ctx *gin.Context, id int64) (info *model.Article, err error) {
	var article *model.Article
	info, err = article.ArticleById(ctx, d.engine, id)
	return
}

func (d *Dao) getCacheAuthorArticleKey(authorId int) string {
	//return util.EncodeMD5(fmt.Sprintf(cacheAuthorArticleCount, authorId))
	return fmt.Sprintf(cacheAuthorArticleCount, authorId)
}

func (d *Dao) getCacheArticleKey() string {
	//return util.EncodeMD5(cacheArticleCount)
	return cacheArticleCount
}
