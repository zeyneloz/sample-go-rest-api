package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connection string for postgres.
const connectionString = "host=%v port=%v user=%v dbname=%v password=%v sslmode=disable"

// Instance is singleton database instance.
var Instance *gorm.DB

// Config holds the settings for database initializaiton.
type Config struct {
	Hostname string
	Port     string
	User     string
	DBName   string
	Password string
}

// Init the database connection.
func (config *Config) Init() error {
	db, err := gorm.Open("postgres", fmt.Sprintf(connectionString, config.Hostname, config.Port, config.User, config.DBName, config.Password))
	// Do we need synchronization here?
	Instance = db
	return err
}
