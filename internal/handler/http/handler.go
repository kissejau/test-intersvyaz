package handler

import (
	"net/http"
	"test-intersvyaz/internal/handler/amqp/producer"
	"test-intersvyaz/pkg/logger"
)

const (
	methodPost = "POST"

	errInvalidBody   = "invalid body"
	errIncorrectBody = "incorrect body"
)

type Handler struct {
	amqpProducer *producer.Producer
	log          *logger.Logger
}

func New(amqpProducer *producer.Producer, log *logger.Logger) *Handler {
	return &Handler{
		amqpProducer: amqpProducer,
		log:          log,
	}
}

func (h *Handler) Routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/track", h.Track)

	return router
}
