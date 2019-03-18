package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/zeyneloz/sample-go-rest-api/internal/app"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/config"
	"github.com/zeyneloz/sample-go-rest-api/internal/svc/db"
	"github.com/zeyneloz/sample-go-rest-api/internal/svc/redis"
)

// Retry interval for sevices, in seconds.
const retryInterval = 5

// Initer object is an object that must be initialized to successfully run the server.
// If Init returns an error, the server stops working
type Initer interface {
	Init() error
}

func main() {
	// Iniitalize and load config.
	cnf := config.NewConfig()
	cnf.LoadConfig()

	// Database service initializer config.
	dbIniter := &db.Config{
		Hostname: cnf.DBHost,
		Port:     cnf.DBPort,
		DBName:   cnf.DBName,
		User:     cnf.DBUser,
		Password: cnf.DBPassword,
	}

	// Redis service initializer config.
	redisIniter := &redis.Config{
		Hostname: cnf.RedisHost,
		Port:     cnf.RedisPort,
		DB:       cnf.RedisDB,
		Password: cnf.RedisPasssword,
	}

	// Initialize all of the required services.
	initService("Database", dbIniter)
	initService("Redis", redisIniter)

	// Migrate models.
	app.MigrateModels()

	// Get main router and start the http server.
	log.Printf("Server is running on port %v \n", cnf.ServerPort)
	http.ListenAndServe(fmt.Sprintf(":%v", cnf.ServerPort), app.GetRouter())
}

// Initialize service with infinite retry.
// Since `service` s are a must to run the server, we must wait for them to init.
func initService(name string, initer Initer) {
	err := initer.Init()
	for err != nil {
		log.Printf("%v service is failed to init, retry in %v seconds\n", name, retryInterval)
		time.Sleep(retryInterval * time.Second)
		err = initer.Init()
	}
	log.Printf("%v service is initialized successfully.\n", name)
}
