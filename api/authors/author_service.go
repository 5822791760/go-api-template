package authors

import (
	"database/sql"
	"net/http"

	"github.com/5822791760/go-api-template/api/authors/requests"
	"github.com/5822791760/go-api-template/constants"
	"github.com/5822791760/go-api-template/helpers"

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

func (s *AuthorService) CreateAuthor(body requests.CreateAuthorRequest) helpers.IErrResponse {
	stmt := Authors.INSERT(Authors.ID, Authors.Name, Authors.Bio).VALUES(DEFAULT, body.Name, body.Bio)
	if _, err := stmt.Exec(s.db); err != nil {
		return helpers.NewErr(err, constants.ErrQuery, http.StatusInternalServerError)
	}

	return nil
}