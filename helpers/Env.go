package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

// Env Gets the value of an environment variable./**
func Env(key string, def any) any {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return def
}
