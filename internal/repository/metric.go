package repository

import (
	"fmt"
	"test-intersvyaz/internal/model"
	"test-intersvyaz/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type MetricRepositoryImpl struct {
	db  *sqlx.DB
	log *logger.Logger
}

func NewMetricRepository(db *sqlx.DB, log *logger.Logger) *MetricRepositoryImpl {
	return &MetricRepositoryImpl{
		db:  db,
		log: log,
	}
}

func (r *MetricRepositoryImpl) Track(partition int, data model.Metric) error {
	query := fmt.Sprintf(`
	INSERT INTO %s%d (id, user_id, event_id, event_name, layout_id, layout_name, created_date)
	VALUES (:id, :user_id, :event_id, :event_name, :layout_id, :layout_name, :created_date)
	`, partitionMetricsPrefix, partition)

	if _, err := r.db.NamedExec(query, data); err != nil {
		return fmt.Errorf("r.db.Exec: %w", err)
	}
	return nil
}
