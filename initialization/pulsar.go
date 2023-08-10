package initialization

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/blog-service/global"
	"github.com/blog-service/pkg/logger"
	"github.com/sirupsen/logrus"
	"time"
)

const afterArticleCreate string = "after_article_create"

func SetPulsarClient() error {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  3 * time.Second,
		ConnectionTimeout: 3 * time.Second,
		Logger: logger.NewLoggerWithLogrus(logrus.StandardLogger(),
			global.AppSetting.LogSavePath+"/"+global.AppSetting.LogFileName+global.AppSetting.LogFileExt),
	})
	if err != nil {
		return err
	}
	global.PulsarClient = client
	fmt.Println("initialization.SetPulsarClient 初始化完成...")
	_ = initArticleProducer()
	return nil
}

func initArticleProducer() error {
	articleProduce, err := global.PulsarClient.CreateProducer(pulsar.ProducerOptions{
		Topic: afterArticleCreate,
	})
	if err != nil {
		global.Logger.Fatalf(context.Background(), "create producer err: %v", err)
		return err
	}
	global.ArticlePulsarProducer = articleProduce
	return nil
}
