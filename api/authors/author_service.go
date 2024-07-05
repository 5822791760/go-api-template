package authors

import (
	"database/sql"
	"net/http"

	"github.com/5822791760/go-api-template/api/authors/requests"
	"github.com/5822791760/go-api-template/libs/errs"

	. "github.com/5822791760/go-api-template/.gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"
)

type AuthorService struct {
	db *sql.DB
}

func NewAuthorService(db *sql.DB) *AuthorService {
	return &AuthorService{
		db: db,
	}
}

func (s *AuthorService) CreateAuthor(body requests.CreateAuthorRequest) errs.ErrRenderer {
	stmt := Authors.INSERT(Authors.ID, Authors.Name, Authors.Bio).VALUES(DEFAULT, body.Name, body.Bio)
	if _, err := stmt.Exec(s.db); err != nil {
		return errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	return nil
}

func (s *AuthorService) UpdateAuthor(id int64, body requests.UpdateAuthorRequest) errs.ErrRenderer {
	stmt := Authors.
		UPDATE(
			Authors.Name,
			Authors.Bio,
		).
		SET(
			body.Name,
			body.Bio,
		).
		WHERE(Authors.ID.EQ(Int(id)))

	if _, err := stmt.Exec(s.db); err != nil {
		return errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	return nil
}
