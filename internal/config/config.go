package config

import (
	"fmt"
	"os"
	"path/filepath"
	"test-intersvyaz/internal/provider/kafka"
	"test-intersvyaz/internal/provider/postgres"
	"test-intersvyaz/internal/server"

	"github.com/joho/godotenv"
)

const (
	ErrFieldNotFound = "field not found in environment variables"
)

type Config struct {
	Server   server.Config
	Postgres postgres.Config
	Kafka    kafka.Config
}

func New() (Config, error) {
	var (
		config         Config
		serverConfig   server.Config
		postgresConfig postgres.Config
		kafkaConfig    kafka.Config
	)

	wd, err := os.Getwd()
	if err != nil {
		return Config{}, fmt.Errorf("os.Getwd: %w", err)
	}

	envpath := filepath.Join(wd, ".env")

	err = godotenv.Load(envpath)
	if err != nil {
		return Config{}, fmt.Errorf("godotenv.Load: %w", err)
	}

	serverPort := os.Getenv("SERVER_PORT")
	if err := Validate(serverPort, "SERVER_PORT"); err != nil {
		return Config{}, fmt.Errorf("Validate: %w", err)
	}

	serverConfig = server.Config{
		Port: serverPort,
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	if err := Validate(postgresHost, "POSTGRES_HOST"); err != nil {
		return Config{}, fmt.Errorf("Validate: %w", err)
	}

	postgresPort := os.Getenv("POSTGRES_PORT")
	if err := Validate(postgresPort, "POSTGRES_PORT"); err != nil {
		return Config{}, fmt.Errorf("Validate: %w", err)
	}

	postgresUser := os.Getenv("POSTGRES_USER")
	if err := Validate(postgresUser, "POSTGRES_USER"); err != nil {
		return Config{}, fmt.Errorf("Validate: %w", err)
	}

	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if err := Validate(postgresPassword, "POSTGRES_PASSWORD"); err != nil {
		return Config{}, fmt.Errorf("Validate: %w", err)
	}

	postgresDatabase := os.Getenv("POSTGRES_DATABASE")
	if err := Validate(postgresDatabase, "POSTGRES_DATABASE"); err != nil {
		return Config{}, fmt.Errorf("Validate: %w", err)
	}

	postgresConfig = postgres.Config{
		Host:     postgresHost,
		Port:     postgresPort,
		User:     postgresUser,
		Password: postgresPassword,
		Database: postgresDatabase,
	}

	kafkaAddress := os.Getenv("KAFKA_ADDRESS")
	if err := Validate(kafkaAddress, "KAFKA_ADDRESS"); err != nil {
		return Config{}, fmt.Errorf("Validate: %w", err)
	}

	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	if err := Validate(kafkaTopic, "KAFKA_TOPIC"); err != nil {
		return Config{}, fmt.Errorf("Validate: %w", err)
	}

	kafkaGroupId := os.Getenv("KAFKA_GROUP_ID")
	if err := Validate(kafkaGroupId, "KAFKA_GROUP_ID"); err != nil {
		return Config{}, fmt.Errorf("Validate: %w", err)
	}

	kafkaConfig = kafka.Config{
		Address: kafkaAddress,
		Topic:   kafkaTopic,
		GroupId: kafkaGroupId,
	}

	config = Config{
		Server:   serverConfig,
		Postgres: postgresConfig,
		Kafka:    kafkaConfig,
	}

	return config, nil
}

func Validate(val string, field string) error {
	if len(val) == 0 {
		return fmt.Errorf("%s %s", field, ErrFieldNotFound)
	}
	return nil
}
