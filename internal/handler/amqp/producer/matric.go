package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"test-intersvyaz/internal/model"
	"test-intersvyaz/pkg/logger"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
	log    *logger.Logger
}

func New(writer *kafka.Writer, log *logger.Logger) *Producer {
	return &Producer{
		writer: writer,
		log:    log,
	}
}

func (p *Producer) Track(ctx context.Context, metric model.Metric) error {
	msg, err := json.Marshal(metric)
	if err != nil {
		p.log.Warnf("json.Marshal: %s", err.Error())
		return fmt.Errorf("error while marshaling")
	}

	// p.log.Infof("sented: %v\n", metric)

	if err := p.writer.WriteMessages(ctx, kafka.Message{
		Value: msg,
	}); err != nil {
		p.log.Warnf("p.writer.WriteMessages: %s", err.Error())
		return fmt.Errorf("internal error")
	}
	return nil
}
