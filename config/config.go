package config

import (
	"os"
)

// Config is the server configuration structure.
type Config struct {
	ServerHost         string // address that server will listening on
	MongoConnectionURI string // mongo db connection uri
	DatabaseName       string // mongo db database name
}

// initialize will read environment variables and save them in config structure fields
func (config *Config) initialize() {
	// read environment variables
	config.ServerHost = os.Getenv("serverHost")
	config.MongoConnectionURI = os.Getenv("mongoConnectionURI")
	config.DatabaseName = os.Getenv("databaseName")
}

// NewConfig will create and initialize config struct
func NewConfig() *Config {
	config := &Config{}
	config.initialize()
	return config
}
