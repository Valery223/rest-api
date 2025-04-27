package main

import (
	parseconfig "learn/rest-api/internal/parse_config"
	"learn/rest-api/internal/storage/sqlite"
	"log/slog"
	"os"
)

func main() {
	cfg := parseconfig.MustLoadConfig()

	log := setupLogger(cfg.Env)
	log.Info("Starting logger...", "env", cfg.Env)

	// TODO init storage: sqlite
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to initialize storage", "error", err)
		os.Exit(1)
	}

	url, err := storage.GetURL("ex")
	if err != nil {
		log.Error("error for get url", "err", err)
		return
	}

	log.Info("get url", "url", url)

	// TODO: init router:gin

	// TODO: run server
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
