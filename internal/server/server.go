package server

import (
	"context"
	"net/http"

	"github.com/iamasocial/hightalent-test-task/internal/config"
)

// Server wraps an http.Server and provides run/shutdown methods
type Server struct {
	httpServer *http.Server
}

// NewServer creates a new Server instance with the given HTTP config and handler
func NewServer(cfg config.HTTP, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      handler,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
	}
}

// Run starts the HTTP server and blocks until it stops
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully stops the HTTP server with a given context
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
