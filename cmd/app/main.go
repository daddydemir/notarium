package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/daddydemir/notarium/internal/api/v1/entries"
	"github.com/daddydemir/notarium/internal/db"
	"github.com/daddydemir/notarium/internal/utils"
)

func main() {
	config, err := utils.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Config load error: %v", err)
	}

	database, err := db.NewDB(config)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	entryRepo := entries.NewRepository(database)
	entryService := entries.NewService(entryRepo)

	entryHandler := entries.NewHandler(entryService)

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/entries", entryHandler.ListEntries).Methods("GET")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	log.Fatal(http.ListenAndServe(":8080", r))

}
