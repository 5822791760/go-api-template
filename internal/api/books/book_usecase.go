package books

import (
	"database/sql"
	"net/http"

	"github.com/5822791760/go-api-template/internal/api/books/reqs"
	"github.com/5822791760/go-api-template/internal/api/books/res"
	. "github.com/5822791760/go-api-template/internal/db/postgres/public/table"
	"github.com/5822791760/go-api-template/pkg/errs"
	"github.com/5822791760/go-api-template/pkg/helpers"

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

func (u *BookUseCase) GetBooks() ([]res.GetBooks, errs.ErrRenderer) {
	resp := []res.GetBooks{}
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

func (u *BookUseCase) CreateBooks(body reqs.CreateBook) (res.CreateBook, errs.ErrRenderer) {
	if err := u.bookService.CreateBook(body); err != nil {
		return res.CreateBook{}, err
	}

	return res.CreateBook{Success: true}, nil
}

func (u *BookUseCase) UpdateBooks(id int32, body reqs.UpdateBook) (res.UpdateBook, errs.ErrRenderer) {
	book, err := u.bookService.UpdateBook(id, body)
	if err != nil {
		return res.UpdateBook{}, err
	}

	return res.UpdateBook{
		ID:        book.ID,
		Name:      book.Name,
		Bookno:    book.Bookno,
		Price:     book.Price,
		Summary:   book.Summary,
		Amount:    book.Amount,
		AuthorID:  book.AuthorID,
		UpdatedAt: helpers.FormatDateTime(book.UpdatedAt),
	}, nil
}

func (u *BookUseCase) BuyBook(bookID int32, userID int32) (res.BuyBook, errs.ErrRenderer) {
	resp, err := u.bookService.BuyBook(bookID, userID)
	if err != nil {
		return res.BuyBook{}, err
	}

	return res.BuyBook{
		ID:     resp.ID,
		Name:   resp.Name,
		Amount: resp.Amount,
	}, nil
}
