package config

import "github.com/spf13/viper"

type Config interface {
	GetString(key string) string
	GetInt(key string) int
}

type ViperConfig struct {
	v *viper.Viper
}
