// This file contains the configuration created for the application in various configuration files.
package consts

import (
	"log"

	"github.com/spf13/viper"
)

// This would be the configuration object which would be used throughout the application
var loadedConfig *Configuration = nil

type Configuration struct {
	// Port is the port number the server will listen on
	Port int
	// Database is the connection string for the database
	DatabaseURL string
	// Secret is the secret key used for JWT token signing
	Secret string
	// TokenDuration is the duration for which the JWT token is valid
	TokenDuration int
	// RefreshTokenDuration is the duration for which the refresh token is valid
	RefreshTokenDuration int
	// ENV is the environment the application is running in
	ENV string
	// CORSOrigins is the list of origins that are allowed to make requests to the server
	CORSOrigins []string
	//AllowedOrigins is the list of origins that are allowed to make requests to the server
	AllowedOrigins []string
}

// GetConfig returns the configuration object
func (c *Configuration) LoadConfig(path string) *Configuration {
	// Load from viper
	// Set the loadedConfig to the configuration object
	// now load viper
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(c)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	loadedConfig = c
	return loadedConfig
}
