package service

import (
	"encoding/json"
	"github.com/blog-service/global"
	"github.com/blog-service/internal/model"
	"github.com/blog-service/internal/queue"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type ArticleCreateRequest struct {
	Show     int    `form:"show" json:"show" binding:"required"`
	Sort     int    `form:"sort" json:"sort" binding:"required"`
	AuthorId int    `form:"author_id" json:"author_id" binding:"required"`
	Url      string `form:"url" json:"url" binding:"required"`
	Title    string `form:"title" json:"title" binding:"required"`
	Content  string `form:"content" json:"content" binding:"required"`
}

type ArticleUpdateRequest struct {
	Id      int64  `form:"id" json:"id" binding:"required,gte=1"`
	Sort    int    `form:"sort" json:"sort" binding:"required"`
	Url     string `form:"url" json:"url" binding:"required"`
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

type ArticleListRequest struct {
	Id       int64  `form:"id" json:"id"`
	AuthorId int    `form:"author_id" json:"author_id"`
	Sort     int    `form:"sort" json:"sort"`
	Url      string `form:"url" json:"url"`
	Title    string `form:"title" json:"title"`
	Content  string `form:"content" json:"content"`
}

type ArticleDetail struct {
	Id         int64                    `json:"id"`
	Show       int                      `json:"show"`
	Sort       int                      `json:"sort"`
	AuthorId   int                      `json:"author_id"`
	Url        string                   `json:"url"`
	Title      string                   `json:"title"`
	Content    string                   `json:"content"`
	Categories []*model.ArticleCategory `json:"categories"`
	Comments   []*model.ArticleComment  `json:"comments"`
}

type ArticleService struct {
	Service
}

func (svc *ArticleService) Create(ctx *gin.Context, params *ArticleCreateRequest) (res int64, err error) {
	id, err := svc.dao.CreateArticle(ctx, params.Show, params.Sort, params.AuthorId, params.Url, params.Title, params.Content)
	if err != nil {
		global.Logger.Errorof(ctx, "文章创建失败: %v", err)
		return 0, err
	}

	var articleProducer queue.ArticleProduce
	payload, _ := json.Marshal(params)
	err = articleProducer.SendMessage(ctx, payload)
	if err != nil {
		global.Logger.Errorof(ctx, "文章写入消息队列失败: %v", err)
	}

	return id, nil
}

func (svc *ArticleService) Update(ctx *gin.Context, params *ArticleUpdateRequest) error {
	updates := make(map[string]interface{})
	bytes, _ := json.Marshal(params)
	_ = json.Unmarshal(bytes, &updates)
	err := svc.dao.UpdateArticle(ctx, params.Id, updates)
	if err != nil {
		global.Logger.Errorof(ctx, "文章更新失败: %v", err)
		return err
	}
	return nil
}

func (svc *ArticleService) List(ctx *gin.Context, params ArticleListRequest) (list []*ArticleDetail, err error) {
	data, err := svc.dao.ListArticle(ctx)
	if err != nil {
		global.Logger.Errorof(ctx, "获取文章列表失败: %v", err)
		return
	}

	bytes, _ := json.Marshal(data)
	err = json.Unmarshal(bytes, &list)
	if err != nil {
		global.Logger.Errorof(ctx, "获取文章列表失败: %v", err)
		return
	}
	return
}

func (svc *ArticleService) Count(ctx *gin.Context, params ArticleListRequest) (count int64, err error) {
	data, err := svc.dao.CountArticle(ctx)
	if err != nil {
		global.Logger.Errorof(ctx, "获取文章总数失败: %v", err)
		return
	}

	bytes, _ := json.Marshal(data)
	err = json.Unmarshal(bytes, &count)
	if err != nil {
		global.Logger.Errorof(ctx, "获取文章总数失败: %v", err)
		return
	}
	return
}

func (svc *ArticleService) Delete(ctx *gin.Context, id int64) error {
	err := svc.dao.DeleteArticle(ctx, id)
	if err != nil {
		global.Logger.Errorof(ctx, "删除文章失败: %v", err)
		return err
	}
	return nil
}

func (svc *ArticleService) Detail(ctx *gin.Context, id int64) (detail *ArticleDetail, err error) {
	articleInfo, err := svc.dao.DetailArticle(ctx, id)
	if err != nil {
		global.Logger.Errorof(ctx, "获取文章失败: %v", err)
		return
	}

	bytes, _ := json.Marshal(articleInfo)
	_ = json.Unmarshal(bytes, &detail)

	articleComments, err := svc.dao.GetArticleComments(ctx, id)
	if err != nil {
		global.Logger.Errorof(ctx, "获取文章的评论列表失败: %v", err)
		return
	}
	_ = copier.Copy(&detail.Comments, articleComments)

	return
}
