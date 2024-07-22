package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/5822791760/go-api-template/docs"
	"github.com/5822791760/go-api-template/internal/config"
	_ "github.com/pressly/goose"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@host		localhost:3000
//	@BasePath	/api/v1

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
func main() {
	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	const Port = 3000

	r.Get("/api/documentation/*", httpSwagger.Handler(
		httpSwagger.URL("/api/documentation/doc.json"),
		httpSwagger.UIConfig(map[string]string{
			"persistAuthorization": "true",
		}),
	))

	if err := InitRoutes(r, db); err != nil {
		log.Fatalf("Error printing routes: %s", err)
		return
	}

	fmt.Printf("\n======================================\n\n")
	fmt.Printf("Listening to port %d", Port)
	fmt.Printf("\n\n======================================\n\n")

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", Port), r)
}
