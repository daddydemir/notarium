package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"

	v1 "github.com/daddydemir/notarium/internal/api/v1"

	"github.com/gorilla/mux"

	"github.com/daddydemir/notarium/internal/api/v1/entries"
	"github.com/daddydemir/notarium/internal/api/v1/files"
	"github.com/daddydemir/notarium/internal/api/v1/notes"
	"github.com/daddydemir/notarium/internal/api/v1/tags"
	"github.com/daddydemir/notarium/internal/api/v1/topics"
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
	entryHandler := v1.NewHandler(entryService)

	topicHandler := v1.NewHandler(topics.NewService(topics.NewRepository(database)))
	tagHandler := v1.NewHandler(tags.NewService(tags.NewRepository(database)))
	noteHandler := v1.NewHandler(notes.NewService(notes.NewRepository(database)))
	fileHandler := v1.NewHandler(files.NewService(files.NewRepository(database)))

	r := mux.NewRouter()

	entries.NewHandler(entryService, r).RegisterRoutes()

	r.HandleFunc("/api/v1/entries", entryHandler.List).Methods("GET")
	r.HandleFunc("/api/v1/entry", entryHandler.Create).Methods("POST")
	r.HandleFunc("/api/v1/entry/{id}", entryHandler.Get).Methods("GET")
	r.HandleFunc("/api/v1/entry/{id}", entryHandler.Update).Methods("PUT")
	r.HandleFunc("/api/v1/entry/{id}", entryHandler.Delete).Methods("DELETE")

	r.HandleFunc("/api/v1/topics", topicHandler.List).Methods("GET")
	r.HandleFunc("/api/v1/topic", topicHandler.Create).Methods("POST")
	r.HandleFunc("/api/v1/topic/{id}", topicHandler.Get).Methods("GET")
	r.HandleFunc("/api/v1/topic/{id}", topicHandler.Update).Methods("PUT")
	r.HandleFunc("/api/v1/topic/{id}", topicHandler.Delete).Methods("DELETE")

	r.HandleFunc("/api/v1/tags", tagHandler.List).Methods("GET")
	r.HandleFunc("/api/v1/tag", tagHandler.Create).Methods("POST")
	r.HandleFunc("/api/v1/tag/{id}", tagHandler.Get).Methods("GET")
	r.HandleFunc("/api/v1/tag/{id}", tagHandler.Update).Methods("PUT")
	r.HandleFunc("/api/v1/tag/{id}", tagHandler.Delete).Methods("DELETE")

	r.HandleFunc("/api/v1/notes", noteHandler.List).Methods("GET")
	r.HandleFunc("/api/v1/note", noteHandler.Create).Methods("POST")
	r.HandleFunc("/api/v1/note/{id}", noteHandler.Get).Methods("GET")
	r.HandleFunc("/api/v1/note/{id}", noteHandler.Update).Methods("PUT")
	r.HandleFunc("/api/v1/note/{id}", noteHandler.Delete).Methods("DELETE")

	r.HandleFunc("/api/v1/files", fileHandler.List).Methods("GET")
	r.HandleFunc("/api/v1/file", fileHandler.Create).Methods("POST")
	r.HandleFunc("/api/v1/file/{id}", fileHandler.Get).Methods("GET")
	r.HandleFunc("/api/v1/file/{id}", fileHandler.Update).Methods("PUT")
	r.HandleFunc("/api/v1/file/{id}", fileHandler.Delete).Methods("DELETE")

	//templates.LoadTemplates("internal/web/templates")

	//homeHandler := handlers.NewHomeHandler(entryService)
	//r.Handle("/", homeHandler)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	handler := cors.AllowAll().Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))

}
