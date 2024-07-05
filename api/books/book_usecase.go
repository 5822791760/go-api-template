package books

import (
	"database/sql"
	"net/http"

	. "github.com/5822791760/go-api-template/.gen/postgres/public/table"
	"github.com/5822791760/go-api-template/api/books/responses"
	"github.com/5822791760/go-api-template/libs/reserrors"

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

func (u *BookUseCase) GetBooks(res *[]responses.GetBooksResponse) reserrors.ErrRenderer {
	stmt := SELECT(
		Books.ID.AS("GetBooksResponse.ID"),
		Books.Name.AS("GetBooksResponse.Name"),
		Books.Bookno.AS("GetBooksResponse.Bookno"),
		Authors.ID.AS("author.ID"),
		Authors.Name.AS("author.Name"),
	).FROM(Books.LEFT_JOIN(Authors, Authors.ID.EQ(Books.AuthorID)))

	if err := stmt.Query(u.db, res); err != nil {
		return reserrors.NewErr(err, reserrors.ErrQuery, http.StatusInternalServerError)
	}

	return nil
}
