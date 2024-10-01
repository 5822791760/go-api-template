package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/5822791760/hr/internal/backend/handlers/httpv1"
	"github.com/5822791760/hr/internal/backend/repos"
	"github.com/5822791760/hr/internal/backend/usecases/userusecase"
	"github.com/5822791760/hr/pkg/coreutil"
	"github.com/go-chi/chi/v5"
)

func InitRoutes(r *chi.Mux, db *sql.DB) error {
	clock := coreutil.NewClock()

	// Repo
	userRepo := repos.NewUserRepo(clock)

	// Use Case
	userUsecase := userusecase.NewUserUsecase(userRepo)

	// Handlers
	authorHandler := httpv1.NewAuthorHandler(db, userUsecase)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/users", authorHandler.FindAll)
			r.Get("/users/{id}", authorHandler.FindOne)
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
