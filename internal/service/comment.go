package service

import (
	"fmt"
	"github.com/blog-service/global"
	"github.com/blog-service/internal/dao"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"time"
)

type CommentCreateRequest struct {
	ArticleId   int64  `form:"article_id" binding:"required" json:"article_id"`
	CommenterId int64  `form:"commenter_id" binding:"required" json:"commenter_id"`
	Commenter   string `form:"commenter" binding:"required" json:"commenter"`
	Content     string `form:"content" binding:"required" json:"content"`
	ParentId    int    `form:"parent_id" binding:"" json:"parent_id"`
}

type CommentResponse struct {
	Id          int64     `json:"id"`
	CommenterId int64     `json:"commenter_id"`
	Commenter   string    `json:"commenter"`
	Content     string    `json:"content"`
	ParentId    int       `json:"parent_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CommentService struct {
	Service
}

func (svc *CommentService) Create(ctx *gin.Context, params *CommentCreateRequest) (interface{}, error) {
	var err error
	var comment dao.Comment
	err = copier.Copy(&comment, params)
	if err != nil {
		return "", err
	}

	id, err := svc.dao.CreateComment(ctx, &comment)
	if err != nil {
		global.Logger.Errorof(ctx, fmt.Sprintf("评论创建失败: %v", err))
		return "", err
	}
	return id, nil
}

func (svc *CommentService) Delete(ctx *gin.Context, id int64) error {
	count, err := svc.dao.DeleteComment(ctx, id)
	if err != nil {
		global.Logger.Errorof(ctx, fmt.Sprintf("评论删除失败: %v", err))
		return err
	}

	if count <= 0 {
		global.Logger.Errorof(ctx, fmt.Sprintf("评论删除失败: 评论删除行数为0"))
	}

	return nil
}

func (svc *CommentService) GetArticleComments(ctx *gin.Context, articleId int64) (comments []*CommentResponse, err error) {
	comments = make([]*CommentResponse, 0)
	list, err := svc.dao.GetArticleComments(ctx, articleId)
	if err != nil {
		global.Logger.Errorof(ctx, fmt.Sprintf("获取文章评论列表失败: %v", err))
		return nil, err
	}

	err = copier.Copy(&comments, list)
	if err != nil {
		global.Logger.Errorof(ctx, fmt.Sprintf("获取文章评论列表失败: %v", err))
		return nil, err
	}
	return comments, nil
}
