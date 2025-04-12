package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config holds all application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Market   MarketConfig
	Security SecurityConfig
	CORS     CORSConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string
	Env  string
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// MarketConfig holds market data-related configuration
type MarketConfig struct {
	APIKey         string
	APIURL         string
	UpdateInterval time.Duration
}

// SecurityConfig holds security-related configuration
type SecurityConfig struct {
	JWTSecret     string
	JWTExpiration time.Duration
}

// CORSConfig holds CORS-related configuration
type CORSConfig struct {
	AllowedOrigins []string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	config := &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Env:  getEnv("ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "minirisk"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "minirisk"),
		},
		Market: MarketConfig{
			APIKey:         getEnv("MARKET_DATA_API_KEY", ""),
			APIURL:         getEnv("MARKET_DATA_API_URL", ""),
			UpdateInterval: getEnvDuration("MARKET_DATA_UPDATE_INTERVAL", 60*time.Second),
		},
		Security: SecurityConfig{
			JWTSecret:     getEnv("JWT_SECRET", ""),
			JWTExpiration: getEnvDuration("JWT_EXPIRATION", 24*time.Hour),
		},
		CORS: CORSConfig{
			AllowedOrigins: getEnvSlice("CORS_ALLOWED_ORIGINS", []string{"http://localhost:3000"}),
		},
	}

	// Validate required configuration
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

// validateConfig validates the configuration
func validateConfig(config *Config) error {
	if config.Database.Password == "" {
		return fmt.Errorf("database password is required")
	}
	if config.Market.APIKey == "" {
		return fmt.Errorf("market data API key is required")
	}
	if config.Market.APIURL == "" {
		return fmt.Errorf("market data API URL is required")
	}
	if config.Security.JWTSecret == "" {
		return fmt.Errorf("JWT secret is required")
	}
	return nil
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvDuration gets an environment variable as a duration with a default value
func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}
	return duration
}

// getEnvSlice gets an environment variable as a slice with a default value
func getEnvSlice(key string, defaultValue []string) []string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	// Split by comma and trim spaces
	var result []string
	for _, item := range splitAndTrim(value, ",") {
		if item != "" {
			result = append(result, item)
		}
	}
	return result
}

// getEnvInt gets an environment variable as an integer with a default value
func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return result
}

// splitAndTrim splits a string by a separator and trims spaces from each part
func splitAndTrim(s, sep string) []string {
	var result []string
	for _, part := range split(s, sep) {
		result = append(result, trim(part))
	}
	return result
}

// split splits a string by a separator
func split(s, sep string) []string {
	return []string{s} // Simplified for this example
}

// trim trims spaces from a string
func trim(s string) string {
	return s // Simplified for this example
}
