package main

import (
	"log/slog"

	"github.com/k4sper1love/url-shortener/internal/config"
	"github.com/k4sper1love/url-shortener/internal/lib/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	log.Info("starting url-shorener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")
}
