package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	JWTSecret string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		slog.Info("no .env file; using process environment")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &Config{Port: port, JWTSecret: os.Getenv("JWT_SECRET")}
}
