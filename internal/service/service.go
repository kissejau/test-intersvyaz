package service

import (
	"test-intersvyaz/internal/model"
	"test-intersvyaz/internal/repository"
	"test-intersvyaz/pkg/logger"
)

type MetricService interface {
	Track(model.Metric) error
}

type Service struct {
	MetricService
}

func New(repository *repository.Repository, log *logger.Logger) *Service {
	return &Service{
		MetricService: NewMetricService(repository.MetricRepository, log),
	}
}
