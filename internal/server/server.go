package server

import (
	"context"
	"net/http"

	"github.com/iamasocial/hightalent-test-task/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.HTTP, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
