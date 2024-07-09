package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/5822791760/go-api-template/libs/initials"
	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if err := initials.InitConfig(); err != nil {
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

	const Port = 8080

	fmt.Printf("\n======================================\n\n")
	fmt.Printf("Listening to port %d", Port)
	fmt.Printf("\n\n======================================\n\n")

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", Port), r)
}
