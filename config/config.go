package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config is a main configuration struct that using in application
type Config struct {
	Environment string
	Debug       bool
	URL         string
	// rate in seconds
	BucketRate             uint
	BucketLoginCapacity    uint
	BucketPasswordCapacity uint
	BucketIPCapacity       uint
	// timeout in ms
	ContextTimeout uint
	Db             map[string]string
	PrometheusHost string
}

// IsProduction checks that application in production mode
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

// IsDevelopment checks that application in production mode
func (c *Config) IsDevelopment() bool {
	return c.Environment == "dev"
}

// GetConfig returns application configuration struct
func GetConfig() (*Config, error) {
	var config Config
	viper.SetConfigName(".config") // name of config file (without extension)
	viper.AddConfigPath(".")       // path to look for the config file in
	err := viper.ReadInConfig()    // Find and read the Config file
	if err != nil {                // Handle errors reading the Config file
		panic(fmt.Sprint("fatal error Config file: %w \n", err))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	if config.Environment == "" {
		log.Fatalf("environment not set")
	}

	return &config, nil
}
