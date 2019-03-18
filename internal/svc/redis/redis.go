package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Client is a singleton redis client instance.
var Client *redis.Client

// Config holds the settings for redis initializaiton.
type Config struct {
	Hostname string
	Port     string
	DB       int
	Password string
}

// Init the redis connection.
func (config *Config) Init() error {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Hostname, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	// Test if connection is actually working.
	_, err := client.Ping().Result()
	if err == nil {
		// Do we need synchronization here?
		Client = client
	}
	return err
}
