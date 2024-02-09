package main

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	_ "traive-engineering-challenge/docs"
	"traive-engineering-challenge/internal/api"
	"traive-engineering-challenge/internal/api/handlers"
	"traive-engineering-challenge/internal/config"
	"traive-engineering-challenge/internal/repository/postgres"
)

// @title Transaction API
// @description API for creating and retrieving transactions
// @version 1.0
// @BasePath /v1
func main() {

	log.SetFormatter(&log.JSONFormatter{})

	fmt.Println("Starting application ...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.WithError(err).Fatal("Failed to load configuration")
	}

	dbConn, err := postgres.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.WithError(err).Fatal("Failed to establish database connection")
	}

	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.WithError(err).Fatal("Failed to close database connection")
		}
	}(dbConn)

	repo, err := postgres.NewRepository(dbConn)
	if err != nil {
		log.WithError(err).Fatal("Failed to initialize repository")
	}

	app := api.NewApplication(repo)
	router := handlers.NewRouter(app)

	const addr = ":8080"
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}
