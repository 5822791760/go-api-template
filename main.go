package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/5822791760/go-api-template/api/authors"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://postgres:mypassword@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	authorService := authors.NewAuthorService(db)
	authorUseCase := authors.NewAuthorUseCase(db, authorService)
	authorController := authors.NewAuthorController(authorUseCase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/authors", authorController.GetAuthors)

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println("")
	fmt.Println("Listening to port 8080")
	fmt.Println("")
	fmt.Println("======================================")
	fmt.Println()

	log.Fatal(http.ListenAndServe(":8080", r))
}
