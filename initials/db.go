package initials

import (
	"database/sql"

	"github.com/5822791760/go-api-template/config"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.GetDBConnectionString())
	if err != nil {
		return nil, err
	}

	return db, nil
}