package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"test-intersvyaz/internal/config"
	"test-intersvyaz/internal/handler/amqp/consumer"
	"test-intersvyaz/internal/handler/amqp/producer"
	handler "test-intersvyaz/internal/handler/http"
	"test-intersvyaz/internal/provider/kafka"
	"test-intersvyaz/internal/provider/postgres"
	"test-intersvyaz/internal/repository"
	"test-intersvyaz/internal/server"
	"test-intersvyaz/internal/service"
	"test-intersvyaz/pkg/logger"

	"github.com/golang-migrate/migrate"
	postgresDB "github.com/golang-migrate/migrate/database/postgres"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	ctx := context.Background()

	logger := logger.NewLogger()

	cfg, err := config.New()
	if err != nil {
		logger.Errorf("config.New: %s", err.Error())
	}

	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		logger.Errorf("postgres.New: %s", err.Error())
	}
	logger.Info("app connected to postgres successfully")

	UpMigrations(db, logger)

	repository := repository.New(db, logger)
	service := service.New(repository, logger)

	kafkaProducer, kafkaConsumer := kafka.NewProducer(cfg.Kafka), kafka.NewConsumer(cfg.Kafka)
	amqpProducer, amqpConsumer := producer.New(kafkaProducer, logger), consumer.New(kafkaConsumer, service.MetricService)

	handler := handler.New(amqpProducer, logger)
	routes := handler.Routes()

	srv := server.NewServer(cfg.Server, routes, logger)

	func() {
		go func() {
			for {
				if err = amqpConsumer.Track(); err != nil {
					logger.Warnf("amqpConsumer.Track: %s", err.Error())
				}
			}
		}()
		go func() {
			if err := srv.Run(); err != nil {
				logger.Errorf("srv.Run: %s", err.Error())
			}
		}()
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down app..")

	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("srv.Shutdown: %s", err.Error())
	}
}

func UpMigrations(db *sqlx.DB, log *logger.Logger) {
	driver, err := postgresDB.WithInstance(db.DB, &postgresDB.Config{})
	if err != nil {
		log.Errorf("postgres.WithInstance: %s", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations/postgres", "postgres", driver)
	if err != nil {
		log.Errorf("migrate.NewWithDatabaseInstance: %s", err.Error())
	}

	if err := m.Up(); err != nil {
		log.Infof("m.Up: %s", err.Error())
	} else {
		log.Infof("Database migration was run successfully")
	}
}
