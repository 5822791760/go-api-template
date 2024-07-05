package authors

import (
	"database/sql"
	"net/http"

	. "github.com/5822791760/go-api-template/.gen/postgres/public/table"
	"github.com/5822791760/go-api-template/api/authors/requests"
	"github.com/5822791760/go-api-template/api/authors/responses"
	"github.com/5822791760/go-api-template/libs/errors"

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

func (u *AuthorUseCase) GetAuthors(res *[]responses.GetAuthorsResponse) errors.ErrRenderer {
	stmt := SELECT(
		Authors.ID.AS("GetAuthorsResponse.ID"),
		Authors.Name.AS("GetAuthorsResponse.Name"),
		Authors.Bio.AS("GetAuthorsResponse.Bio"),
		Books.ID.AS("books.ID"),
		Books.Name.AS("books.Name"),
		Books.Bookno.AS("books.Bookno"),
	).FROM(Authors.LEFT_JOIN(Books, Books.AuthorID.EQ(Authors.ID)))

	err := stmt.Query(u.db, res)
	if err != nil {
		return errors.NewErr(err, errors.ErrQuery, http.StatusInternalServerError)
	}

	return nil
}

func (u *AuthorUseCase) CreateAuthor(body requests.CreateAuthorRequest, res *responses.CreateAuthorResponse) errors.ErrRenderer {
	if err := u.authorService.CreateAuthor(body); err != nil {
		return err
	}

	res.Success = true

	return nil
}
