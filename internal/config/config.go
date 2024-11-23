package config

import (
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel   string `yaml:"log_level" env-default:"info"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"0.0.0.0"`
	Port    int    `yaml:"port" env-default:"8080"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/default.yaml"
		slog.Info("using default config path: " + configPath)
	}

	if _, err := os.Stat(configPath); err != nil {
		slog.Error("config gile does not exist: " + configPath)
		os.Exit(1)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		slog.Error("error reading config file: " + err.Error())
		os.Exit(1)
	}
	return &cfg
}
