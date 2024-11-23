package main

import (
	"fmt"

	"github.com/khramov86/my-url-shortner/internal/app"
	"github.com/khramov86/my-url-shortner/internal/config"
	"github.com/khramov86/my-url-shortner/internal/logger"
)

func main() {
	cfg := config.MustLoad()
	logger := logger.Init(cfg)
	logger.Info(fmt.Sprintf("starting server at %s:%d", cfg.HTTPServer.Address, cfg.HTTPServer.Port))
	app.Init(cfg)
}
