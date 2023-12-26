package service

import (
	"fmt"
	"test-intersvyaz/internal/model"
	"test-intersvyaz/internal/repository"
	"test-intersvyaz/pkg/logger"
	"time"

	"github.com/google/uuid"
)

const (
	errTrack = "error while saving metrics into db"
)

type MetricServiceImpl struct {
	metricsRepo repository.MetricRepository
	log         *logger.Logger
}

func NewMetricService(metricsRepo repository.MetricRepository, log *logger.Logger) *MetricServiceImpl {
	return &MetricServiceImpl{
		metricsRepo: metricsRepo,
		log:         log,
	}
}

func (s *MetricServiceImpl) Track(data model.Metric) error {
	data.Id = uuid.NewString()
	data.CreatedDate = time.Now()

	partion := data.CreatedDate.Day()

	if err := s.metricsRepo.Track(partion, data); err != nil {
		s.log.Warnf("s.metricsRepo.Track: %s", err.Error())
		return fmt.Errorf(errTrack)
	}

	return nil
}
