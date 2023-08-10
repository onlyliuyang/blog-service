package dao

import (
	"github.com/blog-service/internal/model"
	"github.com/blog-service/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ArticleId   int64
	Commenter   string
	CommenterId int64
	Content     string
	ParentId    int64
}

func (d *Dao) CreateComment(ctx *gin.Context, params *Comment) (interface{}, error) {
	commentId := util.Uuid(ctx)
	c := model.ArticleComment{Id: commentId}
	//c.CreatedAt = time.Now().Format(time.DateTime)
	//c.UpdatedAt = time.Now().Format(time.DateTime)
	_ = copier.Copy(&c, params)
	id, err := c.Create(ctx, d.mongoDB)
	return id, err
}

func (d *Dao) UpdateComment(ctx *gin.Context, params *Comment) error {
	c := model.ArticleComment{}
	id, _ := primitive.ObjectIDFromHex("64ad0d63b8d6d4755db39b15")
	update := map[string]interface{}{"content": "我饿了"}
	err := c.UpdateOne(ctx, d.mongoDB, id, update)
	return err
}

func (d *Dao) DeleteComment(ctx *gin.Context, id int64) (int64, error) {
	var c model.ArticleComment
	count, err := c.Delete(ctx, d.mongoDB, id)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (d *Dao) GetArticleComments(ctx *gin.Context, articleId int64) (comments []*Comment, err error) {
	var c model.ArticleComment
	comments = make([]*Comment, 0)
	list, err := c.FindByArticleId(ctx, d.mongoDB, articleId)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&comments, list)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
