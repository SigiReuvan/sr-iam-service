package config

import (
	"os"

	"github.com/go-kit/log"
)

// Config holds configuration values for the application.
type Config struct {
	DBUser        string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	RedisAddr     string
	RedisPassword string
}

// Load reads configuration from environment variables (with defaults)
// and uses the provided logger for logging warnings.
func Load(logger log.Logger) *Config {
	return &Config{
		DBUser:        getEnv(logger, "DB_USER", "postgres"),
		DBPassword:    getEnv(logger, "DB_PASSWORD", "postgres"),
		DBHost:        getEnv(logger, "DB_HOST", "localhost"),
		DBPort:        getEnv(logger, "DB_PORT", "5432"),
		DBName:        getEnv(logger, "DB_NAME", "postgres"),
		RedisAddr:     getEnv(logger, "REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv(logger, "REDIS_PASSWORD", ""),
	}
}

// getEnv retrieves the environment variable by key or returns the fallback value,
// logging a warning using logger.Log if the variable is not set.
func getEnv(logger log.Logger, key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	logger.Log("msg", "environment variable not set, using default", "key", key, "default", fallback)
	return fallback
}
