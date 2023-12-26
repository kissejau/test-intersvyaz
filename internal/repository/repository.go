package repository

import (
	"test-intersvyaz/internal/model"
	"test-intersvyaz/pkg/logger"

	"github.com/jmoiron/sqlx"
)

const (
	tableMetrics           = "metrics"
	partitionMetricsPrefix = "metrics_"
)

type MetricRepository interface {
	Track(int, model.Metric) error
}

type Repository struct {
	MetricRepository
}

func New(db *sqlx.DB, log *logger.Logger) *Repository {
	return &Repository{
		MetricRepository: NewMetricRepository(db, log),
	}
}
