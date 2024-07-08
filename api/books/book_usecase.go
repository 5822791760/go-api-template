package books

import (
	"database/sql"
	"net/http"

	. "github.com/5822791760/go-api-template/.gen/postgres/public/table"
	"github.com/5822791760/go-api-template/api/books/res"
	"github.com/5822791760/go-api-template/libs/errs"

	. "github.com/go-jet/jet/v2/postgres"
)

type BookUseCase struct {
	db          *sql.DB
	bookService *BookService
}

func NewBookUseCase(db *sql.DB, bookService *BookService) *BookUseCase {
	return &BookUseCase{
		db:          db,
		bookService: bookService,
	}
}

func (u *BookUseCase) GetBooks() ([]res.GetBooksResponse, errs.ErrRenderer) {
	resp := []res.GetBooksResponse{}
	stmt := SELECT(
		Books.ID.AS("GetBooksResponse.ID"),
		Books.Name.AS("GetBooksResponse.Name"),
		Books.Bookno.AS("GetBooksResponse.Bookno"),
		Books.Summary.AS("GetBooksResponse.Summary"),
		Authors.ID.AS("author.ID"),
		Authors.Name.AS("author.Name"),
	).FROM(Books.LEFT_JOIN(Authors, Authors.ID.EQ(Books.AuthorID)))

	if err := stmt.Query(u.db, &resp); err != nil {
		return resp, errs.NewErr(err, http.StatusInternalServerError)
	}

	return resp, nil
}
