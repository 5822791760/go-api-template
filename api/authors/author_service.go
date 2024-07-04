package authors

import "database/sql"

type AuthorService struct {
	db *sql.DB
}

func NewAuthorService(db *sql.DB) *AuthorService {
	return &AuthorService{
		db: db,
	}
}
