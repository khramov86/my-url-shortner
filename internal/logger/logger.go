package logger

import (
	"log"
	"log/slog"
	"os"

	"github.com/khramov86/my-url-shortner/internal/config"
)

var loglevel slog.Level

func Init(cfg *config.Config) *slog.Logger {
	switch cfg.LogLevel {
	case "debug":
		loglevel = slog.LevelDebug
	case "info":
		loglevel = slog.LevelInfo
	case "warn":
		loglevel = slog.LevelWarn
	case "error":
		loglevel = slog.LevelError
	default:
		loglevel = slog.LevelDebug
	}
	log.Printf("log level: %s", cfg.LogLevel)
	opts := &slog.HandlerOptions{
		Level: loglevel,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}
