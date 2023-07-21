package model

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleCategory struct {
	Id         int64 `bson:"id"`
	ArticleId  int64 `bson:"article_id"`
	CategoryId int64 `bson:"category_id"`
	Model
}

func (a *ArticleCategory) collection() string {
	return "bs_article_category"
}

func (a *ArticleCategory) Create(ctx *gin.Context, col *mongo.Database) (interface{}, error) {
	result, err := col.Collection(a.collection()).InsertOne(ctx, a, nil)
	if err != nil {
		return "", err
	}
	return result.InsertedID, nil
}
