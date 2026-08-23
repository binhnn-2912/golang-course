package config

import "os"

// Config holds all application configuration.
type Config struct {
	Port     string
	Env      string
	LogLevel string
	DBDSN    string
}

// Load reads configuration from environment variables with sensible defaults.
func Load() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		Env:      getEnv("APP_ENV", "development"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
		DBDSN:    getEnv("DB_DSN", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
