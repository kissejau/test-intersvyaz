package server

import (
	"context"
	"fmt"
	"net/http"
	"test-intersvyaz/pkg/logger"
)

type Server struct {
	srv *http.Server
	log *logger.Logger
}

func NewServer(config Config, handler http.Handler, log *logger.Logger) Server {
	return Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%s", config.Port),
			Handler: handler,
		},
		log: log,
	}
}

func (s Server) Run() error {
	s.log.Infof("Server is starting on port %s", s.srv.Addr)
	if err := s.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("s.srv.ListenAndServer: %w", err)
	}
	return nil
}

func (s Server) Shutdown(ctx context.Context) error {
	s.log.Info("Server is shutting down...")
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("s.srv.Shutdown: %w", err)
	}
	return nil
}
