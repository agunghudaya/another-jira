package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

// NewConfig initializes and returns a new Viper configuration instance
func NewConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config") // config.json or config.yaml
	v.SetConfigType("json")   // Change to "yaml" if using YAML
	v.AddConfigPath("./internal/infrastructure/config")

	viper.SetDefault("fe.url", "http://localhost:3000")

	// Read configuration file
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
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
