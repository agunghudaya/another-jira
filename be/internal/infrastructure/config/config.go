package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func NewConfig() (*ViperConfig, error) {
	v := viper.New()

	// Set defaults
	v.SetDefault("fe.url", "http://localhost:3000")

	// Load environment variables from .env file
	v.SetConfigFile(".env")
	v.AddConfigPath("/app")
	if err := v.ReadInConfig(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// Enable automatic environment variable binding (e.g., DB_HOST, DB_PORT)
	v.AutomaticEnv()

	// Load JSON config
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath("./internal/infrastructure/config")

	// Read configuration file
	if err := v.MergeInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config.json: %w", err)
	}

	return &ViperConfig{v: v}, nil
}

// GetString retrieves a string value from the configuration
func (c *ViperConfig) GetString(key string) string {
	return c.v.GetString(key)
}

// GetInt retrieves an integer value from the configuration
func (c *ViperConfig) GetInt(key string) int {
	return c.v.GetInt(key)
}
