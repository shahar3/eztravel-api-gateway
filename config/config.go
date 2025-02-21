package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config holds all configuration variables.
type Config struct {
	Port                string        `mapstructure:"PORT"`
	Env                 string        `mapstructure:"ENV"`
	ReadTimeout         time.Duration `mapstructure:"READ_TIMEOUT"`
	WriteTimeout        time.Duration `mapstructure:"WRITE_TIMEOUT"`
	TripServiceEndpoint string        `mapstructure:"TRIP_SERVICE_ENDPOINT"`
}

// LoadConfig reads configuration from .env file or environment variables.
func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Provide default values
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("ENV", "development")
	viper.SetDefault("READ_TIMEOUT", "5s")
	viper.SetDefault("WRITE_TIMEOUT", "10s")

	if err = viper.ReadInConfig(); err != nil {
		log.Printf("No config file found: %v", err)
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Printf("Unable to decode into struct: %v", err)
		return nil, err
	}

	return config, nil
}
