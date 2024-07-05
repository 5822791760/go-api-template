package initials

import (
	"database/sql"

	"github.com/5822791760/go-api-template/api/authors"
	"github.com/5822791760/go-api-template/api/books"
	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
)

func InitRoutes(r *chi.Mux, db *sql.DB) {
	render := render.New()

	// SERVICES =======
	authorService := authors.NewAuthorService(db)
	bookService := books.NewBookService(db)

	// USECASES =======
	authorUseCase := authors.NewAuthorUseCase(db, authorService)
	bookUseCase := books.NewBookUseCase(db, bookService)

	// CONTROLLERS =======
	authorController := authors.NewAuthorController(render, authorUseCase)
	bookController := books.NewBookController(render, bookUseCase)

	// ROUTES =======

	r.Route("/authors", func(r chi.Router) {
		r.Get("/", authorController.GetAuthors)
		r.Post("/", authorController.CreateAuthor)

		r.Route("/{id}", func(r chi.Router) {
			r.Put("/", authorController.UpdateAuthor)
			r.Get("/", authorController.GetAuthor)
		})
	})

	r.Route("/books", func(r chi.Router) {
		r.Get("/", bookController.GetBooks)
	})
}
