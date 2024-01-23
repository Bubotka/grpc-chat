package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenAIKey string
	GRPC      GrpcConfig
}

type GrpcConfig struct {
	Host    string
	Port    string
	Timeout time.Duration
}

func MustConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		OpenAIKey: getEnv("OPENAI_KEY", ""),
		GRPC: GrpcConfig{
			Host:    getEnv("GRPC_HOST", "localhost"),
			Port:    getEnv("GRPC_PORT", "50051"),
			Timeout: 10 * time.Second,
		},
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
