package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config is the server configuration structure.
type Config struct {
	ServerHost         string        // address that server will listening on
	MongoConnectionURI string        // mongo db connection uri
	DatabaseName       string        // mongo db database name
	SecretKey          string        // JWT SecretKey
	ExpirationDuration time.Duration // JWT ExpirationDuration
}

// initialize will read environment variables and save them in config structure fields
func (config *Config) initialize() {
	// read environment variables
	config.ServerHost = os.Getenv("serverHost")
	if config.ServerHost == "" {
		config.ServerHost = fmt.Sprintf("0:%s", os.Getenv("PORT"))
		fmt.Printf("Take HOST from $POSRT env variable => %s", config.ServerHost)
	}
	config.MongoConnectionURI = os.Getenv("mongoConnectionURI")
	config.DatabaseName = os.Getenv("databaseName")
	config.SecretKey = os.Getenv("secretKey")

	if dur, err := strconv.Atoi(os.Getenv("expirationDuration")); err == nil {
		config.ExpirationDuration = time.Duration(dur) * time.Second
	} else {
		config.ExpirationDuration = 300 * time.Second
	}
}

// NewConfig will create and initialize config struct
func NewConfig() *Config {
	config := &Config{}
	config.initialize()
	return config
}
