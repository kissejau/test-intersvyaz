package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"test-intersvyaz/internal/model"
	"test-intersvyaz/internal/service"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader        *kafka.Reader
	metricService service.MetricService
}

func New(reader *kafka.Reader, metricService service.MetricService) *Consumer {
	return &Consumer{
		reader:        reader,
		metricService: metricService,
	}
}

func (c *Consumer) Track() error {
	ctx := context.Background()
	msg, err := c.reader.ReadMessage(ctx)
	if err != nil {
		return fmt.Errorf("s.kafkaConsumer.ReadMessage: %w", err)
	}

	var metric model.Metric
	if err = json.Unmarshal(msg.Value, &metric); err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	// log.Printf("received: %v\n", metric)

	if err = c.metricService.Track(metric); err != nil {
		return fmt.Errorf("c.metricService.Track: %w", err)
	}
	return nil
}
