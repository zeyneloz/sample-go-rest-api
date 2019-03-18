package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

const envPrefix = ""

// Config of the application.
type Config struct {
	ServerPort     string `envconfig:"SERVER_PORT"`
	DBHost         string `envconfig:"POSTGRES_HOST"`
	DBPort         string `envconfig:"POSTGRES_PORT"`
	DBName         string `envconfig:"POSTGRES_DB"`
	DBUser         string `envconfig:"POSTGRES_USER"`
	DBPassword     string `envconfig:"POSTGRES_PASSWORD"`
	RedisHost      string `envconfig:"REDIS_HOST"`
	RedisPort      string `envconfig:"REDIS_PORT"`
	RedisDB        int    `envconfig:"REDIS_DB"`
	RedisPasssword string `envconfig:"REDIS_PASSWORD"`
	SecretKey      string `envconfig:"SECRET_KEY"`
	JwtExpire      int    // JWT exprie timeout in minutes
}

// Singleton config.
var config *Config

// NewConfig returns an empty config
func NewConfig() *Config {
	// Defaults for non environment settings goes here.
	config = &Config{
		JwtExpire: 120,
	}
	return config
}

// LoadConfig from environmental variables.
func (cnf *Config) LoadConfig() {
	err := envconfig.Process(envPrefix, cnf)
	if err != nil {
		log.Fatal("Failed to load config from env: ", err)
	}
}

// GetConfig returns singleton config.
func GetConfig() *Config {
	return config
}
