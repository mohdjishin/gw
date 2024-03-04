package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Services       []ServiceConfig `mapstructure:"services"`
	PrivateKeyPath string          `mapstructure:"private-key-path"`
}

type ServiceConfig struct {
	Name string `mapstructure:"name"`
	URL  string `mapstructure:"url"`
}

func LoadConfig() Config {
	var c Config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return c
}
