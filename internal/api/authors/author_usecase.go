package authors

import (
	"database/sql"
	"net/http"

	"github.com/5822791760/go-api-template/internal/api/authors/reqs"
	"github.com/5822791760/go-api-template/internal/api/authors/res"
	. "github.com/5822791760/go-api-template/internal/db/postgres/public/table"
	"github.com/5822791760/go-api-template/pkg/errs"

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

func (u *AuthorUseCase) GetAuthors() ([]res.GetAuthors, errs.ErrRenderer) {
	resp := []res.GetAuthors{}
	stmt := SELECT(
		Authors.ID.AS("GetAuthorsResponse.ID"),
		Authors.Name.AS("GetAuthorsResponse.Name"),
		Authors.Bio.AS("GetAuthorsResponse.Bio"),
		Books.ID.AS("books.ID"),
		Books.Name.AS("books.Name"),
		Books.Bookno.AS("books.Bookno"),
	).FROM(Authors.LEFT_JOIN(Books, Books.AuthorID.EQ(Authors.ID)))

	err := stmt.Query(u.db, &resp)
	if err != nil {
		return []res.GetAuthors{}, errs.NewErr(err, http.StatusInternalServerError)
	}

	return resp, nil
}

func (u *AuthorUseCase) GetAuthor(id int32) (res.GetAuthor, errs.ErrRenderer) {
	var resp res.GetAuthor
	stmt := SELECT(
		Authors.ID.AS("GetAuthorResponse.ID"),
		Authors.Name.AS("GetAuthorResponse.Name"),
		Authors.Bio.AS("GetAuthorResponse.Bio"),
		Books.ID.AS("getAuthorBooks.ID"),
		Books.Name.AS("getAuthorBooks.Name"),
		Books.Bookno.AS("getAuthorBooks.Bookno"),
	).
		FROM(
			Authors.
				LEFT_JOIN(
					Books,
					Books.AuthorID.EQ(Authors.ID)),
		).
		WHERE(
			Authors.ID.EQ(Int(int64(id))),
		)

	err := stmt.Query(u.db, &resp)
	if err != nil {
		return res.GetAuthor{}, errs.NewErr(err, http.StatusInternalServerError)
	}

	return resp, nil
}

func (u *AuthorUseCase) CreateAuthor(body reqs.CreateAuthor) (res.CreateAuthor, errs.ErrRenderer) {
	if err := u.authorService.CreateAuthor(body); err != nil {
		return res.CreateAuthor{}, err
	}

	resp := res.CreateAuthor{
		Success: true,
	}

	return resp, nil
}

func (u *AuthorUseCase) UpdateAuthor(id int32, body reqs.UpdateAuthor) (res.UpdateAuthor, errs.ErrRenderer) {
	if err := u.authorService.UpdateAuthor(id, body); err != nil {
		return res.UpdateAuthor{}, err
	}

	return res.UpdateAuthor{
		Success: true,
	}, nil
}
