package model

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ArticleComment struct {
	Model
	Id          int64  `bson:"id" json:"id"`
	ArticleId   int64  `bson:"article_id" json:"article_id"`
	Commenter   string `bson:"commenter" json:"commenter"`
	CommenterId int64  `bson:"commenter_id" json:"commenter_id"`
	Content     string `bson:"content" json:"content"`
	ParentId    int64  `bson:"parent_id" json:"parent_id"`
}

func (c *ArticleComment) collection() string {
	return "bs_article_comment"
}

func (c *ArticleComment) Create(ctx *gin.Context, col *mongo.Database) (interface{}, error) {
	_, err := col.Collection(c.collection()).InsertOne(ctx, c)
	if err != nil {
		return "", err
	}

	return c.Id, nil
}

func (c *ArticleComment) Update(ctx *gin.Context, col *mongo.Database, comment *ArticleComment) (int64, error) {
	result, err := col.Collection(c.collection()).UpdateByID(ctx, comment.Id, comment)
	if err != nil {
		return 0, err
	}
	return result.UpsertedCount, err
}

func (c *ArticleComment) FindOne(ctx *gin.Context, col *mongo.Database, id int64) (comment *ArticleComment, err error) {
	filter := bson.M{"id": id}
	err = col.Collection(c.collection()).FindOne(ctx, filter, nil).Decode(&comment)
	return
}

func (c *ArticleComment) FindByArticleId(ctx *gin.Context, col *mongo.Database, articleId int64) (comments []*ArticleComment, err error) {
	comments = make([]*ArticleComment, 0)
	filter := bson.M{"article_id": articleId}
	cur, err := col.Collection(c.collection()).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var articleComment ArticleComment
		err := cur.Decode(&articleComment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &articleComment)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return
}

func (a *ArticleComment) UpdateOne(ctx *gin.Context, col *mongo.Database, id primitive.ObjectID, data map[string]interface{}) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", id}}
	updateItems := bson.D{}
	for key, val := range data {
		updateItems = append(updateItems, bson.E{Key: key, Value: val})
	}
	update := bson.D{{"$set", updateItems}}
	result, err := col.Collection(a.collection()).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 && result.UpsertedCount == 0 {
		return errors.New("更新文档失败")
	}
	return nil
}

func (c *ArticleComment) Delete(ctx *gin.Context, col *mongo.Database, id int64) (int64, error) {
	filter := bson.M{"id": id}
	result, err := col.Collection(c.collection()).DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
