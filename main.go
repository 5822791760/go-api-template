package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/5822791760/go-api-template/config"
	"github.com/5822791760/go-api-template/initials"
	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Load Config
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
		return
	}

	db, err := initials.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	initials.InitRoutes(r, db)

	fmt.Printf("\n======================================\n\n")
	fmt.Printf("Listening to port 8080")
	fmt.Printf("\n\n======================================\n\n")

	log.Fatal(http.ListenAndServe(":8080", r))
}
