package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

// NewConfig initializes and returns a new Viper configuration instance
func NewConfig() (*Config, error) {
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

	// Print all configuration values for debugging
	log.Println("Loaded Configuration:")
	for _, key := range v.AllKeys() {
		log.Printf("%s: %v", key, v.Get(key))
	}

	return &Config{v: v}, nil
}

// GetViper exposes the Viper instance
func (c *Config) GetViper() *viper.Viper {
	return c.v
}

// Helper functions to access specific config values
func (c *Config) GetString(key string) string { return c.v.GetString(key) }
func (c *Config) GetInt(key string) int       { return c.v.GetInt(key) }
