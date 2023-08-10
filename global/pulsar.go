package global

import "github.com/apache/pulsar-client-go/pulsar"

var (
	PulsarClient          pulsar.Client
	ArticlePulsarProducer pulsar.Producer
)
