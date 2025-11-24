package main

import (
	"context"
	"flag"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/iamasocial/hightalent-test-task/internal/config"
	"github.com/iamasocial/hightalent-test-task/internal/db"
	"github.com/iamasocial/hightalent-test-task/internal/logger"
	"github.com/iamasocial/hightalent-test-task/internal/migrations"
	"github.com/iamasocial/hightalent-test-task/internal/repository"
	"github.com/iamasocial/hightalent-test-task/internal/server"
	"github.com/iamasocial/hightalent-test-task/internal/service"
	handler "github.com/iamasocial/hightalent-test-task/internal/transport/http"
)

func main() {
	var configPath, env string

	flag.StringVar(&configPath, "config", "config/config.yaml", "path to config file")
	flag.StringVar(&env, "env", "local", "application environment: local, dev, prod")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	log := logger.NewSlogger(env)
	log.Info(ctx, "logger initialized")

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Error(ctx, "failed to load config", "error", err)
		return
	}
	log.Info(ctx, "config loaded successfully")

	dbConn, err := db.NewPostgresDB(cfg.DB)
	if err != nil {
		log.Error(ctx, "failed to connect to db", "error", err)
		return
	}

	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Error(ctx, "failed to get raw sql db", "error", err)
		return
	}

	if err := migrations.Run(sqlDB); err != nil {
		log.Error(ctx, "failed to run migrations", "error", err)
		return
	}
	log.Info(ctx, "migrations applied successfully")

	repo := repository.NewRepository(dbConn)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)
	router := handler.NewRouter(h)
	srv := server.NewServer(cfg.HTTP, router)

	go func() {
		if err := srv.Run(); err != nil && err != http.ErrServerClosed {
			log.Error(ctx, "server stopped", "error", err)
			cancel()
		}
	}()

	log.Info(ctx, "server started", "port", cfg.HTTP.Port)

	<-ctx.Done()
	log.Info(ctx, "shutting down gracefully")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), cfg.HTTP.ShutdownTimeout)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error(shutdownCtx, "failed to shutdown server", "error", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Error(shutdownCtx, "failed to close database", "error", err)
	}
}
