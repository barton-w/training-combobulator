package main

import (
	"log"
	"training-combobulator/config"
	"training-combobulator/dal"
	"training-combobulator/handlers"
)

func main() {
	// initialize config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %s", err.Error())
	}
	log.Println("configuration initialized")

	// connect to the database
	dbClient, err := dal.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("failed to connected to database: %s", err.Error())
	}
	log.Println("connected to database")

	handlers.SomeHandler(dbClient)
}
