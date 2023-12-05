package queue

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/blog-service/global"
	"github.com/gin-gonic/gin"
)

type ArticleProduce struct {
}

func (p *ArticleProduce) SendMessage(ctx *gin.Context, payload []byte) error {
	messageId, err := global.ArticlePulsarProducer.Send(ctx, &pulsar.ProducerMessage{
		Payload: payload,
	})
	global.Logger.Infof(ctx, "发送article message:%s, result:%s", string(payload), messageId.String())
	return err
}
