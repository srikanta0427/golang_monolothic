package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	// Env  string
	DatabaseURL string
}

func MustLoadConfig() Config {
	_ = godotenv.Load()
	
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT must be set")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		panic("DATABASE_URL must be set")
	}

	return Config{
		Port: port,
		DatabaseURL: databaseURL,
	}
}