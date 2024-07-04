package authors

import (
	"database/sql"

	"github.com/5822791760/go-api-template/api/authors/authorres"

	. "github.com/5822791760/go-api-template/.gen/postgres/public/table"

	. "github.com/go-jet/jet/v2/postgres"
)

type AuthorUseCase struct {
	db            *sql.DB
	authorService *AuthorService
}

func NewAuthorUseCase(db *sql.DB, authorService *AuthorService) *AuthorUseCase {
	return &AuthorUseCase{
		db:            db,
		authorService: authorService,
	}
}

func (u *AuthorUseCase) GetAuthors() ([]authorres.GetAuthorsResponse, error) {
	authors := []authorres.GetAuthorsResponse{}

	stmt := SELECT(
		Authors.ID.AS("GetAuthorsResponse.ID"),
		Authors.Name.AS("GetAuthorsResponse.Name"),
		Authors.Bio.AS("GetAuthorsResponse.Bio"),
	).FROM(Authors)
	err := stmt.Query(u.db, &authors)
	if err != nil {
		return nil, err
	}

	return authors, nil
}
