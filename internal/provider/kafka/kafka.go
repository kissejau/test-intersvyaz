package kafka

import (
	"github.com/segmentio/kafka-go"
)

func NewProducer(cfg Config) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{cfg.Address},
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	})
}

func NewConsumer(cfg Config) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.Address},
		Topic:   cfg.Topic,
		GroupID: cfg.GroupId,
	})
}
