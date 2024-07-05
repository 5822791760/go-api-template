package initials

import (
	"database/sql"

	"github.com/5822791760/go-api-template/api/authors"
	"github.com/5822791760/go-api-template/api/books"
	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
)

func InitRoutes(r *chi.Mux, db *sql.DB) {
	renderer := render.New()

	// SERVICES =======
	authorService := authors.NewAuthorService(db)
	bookService := books.NewBookService(db)

	// USECASES =======
	authorUseCase := authors.NewAuthorUseCase(db, authorService)
	bookUseCase := books.NewBookUseCase(db, bookService)

	// CONTROLLERS =======
	authorController := authors.NewAuthorController(renderer, authorUseCase)
	bookController := books.NewBookController(renderer, bookUseCase)

	// ROUTES =======
	r.Get("/authors", authorController.GetAuthors)
	r.Post("/authors", authorController.CreateAuthor)
	r.Get("/books", bookController.GetBooks)
}