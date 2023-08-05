package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	HTTPServer HTTPServer `yaml:"HTTPServer"`
	Database   Database   `yaml:"database"`
}

type Database struct {
	Host     string `yaml:"host" env-default:"localhost:5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func Load() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		slog.Error("not load CONFIG_PATH!")
		return nil
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		slog.Error("config file does not exist!")
		return nil
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		slog.Error("error read config file!", slog.String("path", configPath))
		return nil
	}

	return &cfg
}
