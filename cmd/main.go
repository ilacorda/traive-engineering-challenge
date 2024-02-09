package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"traive-engineering-challenge/internal/api"
	"traive-engineering-challenge/internal/api/handlers"
	"traive-engineering-challenge/internal/config"
	"traive-engineering-challenge/internal/repository/postgres"
)

func main() {

	fmt.Println("Starting application ...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	dbConn, err := postgres.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to establish database connection: %v", err)
	}

	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}(dbConn)

	repo, err := postgres.NewRepository(dbConn)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	app := api.NewApplication(repo)
	router := handlers.NewRouter(app)

	const addr = ":8080"
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
