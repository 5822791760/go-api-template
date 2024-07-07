package initials

import (
	"database/sql"

	"github.com/5822791760/go-api-template/api/authors"
	"github.com/5822791760/go-api-template/api/auths"
	"github.com/5822791760/go-api-template/api/books"
	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
)

func InitRoutes(r *chi.Mux, db *sql.DB) {
	render := render.New()
	middlewareService := NewMiddlewareService(render)

	// SERVICES =======
	authorService := authors.NewAuthorService(db)
	bookService := books.NewBookService(db)
	authService := auths.NewAuthService(db)

	// USECASES =======
	authorUseCase := authors.NewAuthorUseCase(db, authorService)
	bookUseCase := books.NewBookUseCase(db, bookService)
	authUseCase := auths.NewAuthUseCase(db, authService)

	// CONTROLLERS =======
	authorController := authors.NewAuthorController(render, authorUseCase)
	bookController := books.NewBookController(render, bookUseCase)
	authController := auths.NewAuthController(render, authUseCase)

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
				r.Use(middlewareService.JwtMiddleware)

				r.Get("/authors", authorController.GetAuthors)
				r.Post("/authors", authorController.CreateAuthor)
				r.Put("/authors/{id}", authorController.UpdateAuthor)
				r.Get("/authors/{id}", authorController.GetAuthor)

				r.Get("/books", bookController.GetBooks)

			})

		})
	})
}
