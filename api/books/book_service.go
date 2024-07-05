package books

import "database/sql"

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{
		db: db,
	}
}
