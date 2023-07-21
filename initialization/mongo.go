package initialization

import (
	"context"
	"fmt"
	"github.com/blog-service/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongoDB() error {
	var err error
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s", global.MongoDBSetting.UserName, global.MongoDBSetting.Password, global.MongoDBSetting.Host)
	clientOptions := options.Client().ApplyURI(mongoURI)

	global.MongoDB, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	err = global.MongoDB.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	fmt.Println("initialization.SetupMongoDB 初始化完成...")
	return nil
}
