package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/5822791760/go-api-template/internal/api/authors"
	"github.com/5822791760/go-api-template/internal/api/auths"
	"github.com/5822791760/go-api-template/internal/api/books"
	"github.com/5822791760/go-api-template/internal/middlewares"
	"github.com/go-chi/chi/v5"
)

func InitRoutes(r *chi.Mux, db *sql.DB) error {
	// SERVICES =======
	authorService := authors.NewAuthorService(db)
	bookService := books.NewBookService(db)
	authService := auths.NewAuthService(db)

	// USECASES =======
	authorUseCase := authors.NewAuthorUseCase(db, authorService)
	bookUseCase := books.NewBookUseCase(db, bookService)
	authUseCase := auths.NewAuthUseCase(db, authService)

	// CONTROLLERS =======
	authorController := authors.NewAuthorController(authorUseCase)
	bookController := books.NewBookController(bookUseCase)
	authController := auths.NewAuthController(authUseCase)

	// ROUTES =======

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {

			// PUBLIC
			r.Route("/public", func(r chi.Router) {

				r.Post("/auths/sign_up", authController.SignUp)
				r.Post("/auths/sign_in", authController.SignIn)

			})

			// JWT Protected
			r.Route("/", func(r chi.Router) {
				r.Use(middlewares.JwtAuth)

				r.Get("/authors", authorController.GetAuthors)
				r.Post("/authors", authorController.CreateAuthor)
				r.Put("/authors/{id}", authorController.UpdateAuthor)
				r.Get("/authors/{id}", authorController.GetAuthor)

				r.Get("/books", bookController.GetBooks)
				r.Post("/books", bookController.CreateBook)
				r.Put("/books/{id}", bookController.UpdateBook)
				r.Patch("/books/{id}/buy", bookController.BuyBook)
			})

		})
	})

	if err := PrintRoutes(r); err != nil {
		return err
	}

	return nil
}

func PrintRoutes(r chi.Router) error {
	fmt.Println()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		return err
	}

	return nil
}
