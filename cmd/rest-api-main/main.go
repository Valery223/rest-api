package main

import (
	"learn/rest-api/internal/book/repository"
	"learn/rest-api/internal/book/service"
	"learn/rest-api/internal/book/transport"
	parseconfig "learn/rest-api/internal/parse_config"
	"learn/rest-api/internal/router"
	"learn/rest-api/internal/storage"
	"log/slog"
	"os"
)

func main() {
	cfg := parseconfig.MustLoadConfig()

	log := setupLogger(cfg.Env)
	log.Info("Starting logger...", "env", cfg.Env)

	db, err := storage.InitSQLiteDB(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to initialize storage", "error", err)
		os.Exit(1)
	}

	// создаем репорзиторий, сервис, хендлер(3 слоя) и вносим в роутер этот хендлер
	repo := repository.NewBookStorage(db)

	cvs := service.NewBookService(repo)

	router := router.NewRouter(transport.NewBookHandler(cvs))

	router.Run(cfg.Address)

}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	return log
}
