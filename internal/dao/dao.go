package dao

import (
	"github.com/blog-service/global"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Dao struct {
	engine  *gorm.DB
	mongoDB *mongo.Database
	//ctx    context.Context
}

func New(engine *gorm.DB, mongoDB *mongo.Client) *Dao {
	return &Dao{
		engine:  engine,
		mongoDB: mongoDB.Database(global.MongoDBSetting.DBName),
		//ctx:    ctx,
	}
}
