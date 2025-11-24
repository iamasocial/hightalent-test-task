package main

import (
	"context"
	"flag"

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

	log := logger.NewSlogger(env)
	log.Info(context.Background(), "logger initialized")

	ctxBkg := context.Background()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Error(ctxBkg, "failed to load config", "error", err)
		return
	}

	log.Info(ctxBkg, "config loaded successfully")

	dbConn, err := db.NewPostgresDB(cfg.DB)
	if err != nil {
		log.Error(ctxBkg, "failed to connect to db", "error", err)
		return
	}

	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Error(ctxBkg, "failed to get raw sql db", "error", err)
		return
	}

	if err := migrations.Run(sqlDB); err != nil {
		log.Error(ctxBkg, "failed to run migrations", "error", err)
		return
	}

	log.Info(ctxBkg, "migrations applied successfully")

	repo := repository.NewRepository(dbConn)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)
	router := handler.NewRouter(h)
	srv := server.NewServer(cfg.HTTP, router)

	log.Info(ctxBkg, "starting server", "port", cfg.HTTP.Port)

	if err := srv.Run(); err != nil {
		log.Error(ctxBkg, "server stopped", "error", err)
	}
}
