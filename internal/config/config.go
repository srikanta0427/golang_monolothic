package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	// Env  string
}

func MustLoadConfig() Config {
	_ = godotenv.Load()
	
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT must be set")
	}
	return Config{
		Port: port,
	}
}