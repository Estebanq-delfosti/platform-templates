package config

import "os"

type Config struct {
	LogLevel    string
	Environment string
}

func Load() *Config {
	return &Config{
		LogLevel:    getEnv("LOG_LEVEL", "INFO"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
