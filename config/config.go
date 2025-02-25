package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

// KafkaConfig holds Kafka-related settings.
type KafkaConfig struct {
	Broker string `mapstructure:"broker"`
}

// Config holds all configuration variables.
type Config struct {
	Port                string        `mapstructure:"port"`
	Env                 string        `mapstructure:"env"`
	ReadTimeout         time.Duration `mapstructure:"read_timeout"`
	WriteTimeout        time.Duration `mapstructure:"write_timeout"`
	TripServiceEndpoint string        `mapstructure:"trip_service_endpoint"`
	Kafka               KafkaConfig   `mapstructure:"kafka"`
	DBUser              string        `mapstructure:"DB_USER"`
	DBPassword          string        `mapstructure:"DB_PASSWORD"`
	JwtSecret           string        `mapstructure:"JWT_SECRET"`
}

// LoadConfig reads configuration from YAML file and environment variables.
func LoadConfig() (*Config, error) {
	viper.AutomaticEnv() // Automatically read env variables

	// Get environment mode from ENV variable (default: "development")
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	// Set the config file based on environment
	configFile := env + ".yaml"
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config") // Add current directory

	// Read YAML file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: No config file found: %v", err)
	}

	// Read .env file for credentials
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Unable to decode into struct: %v", err)
		return nil, err
	}

	return &config, nil
}
