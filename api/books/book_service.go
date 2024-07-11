package books

import (
	"database/sql"
	"net/http"

	"github.com/5822791760/go-api-template/api/books/reqs"
	"github.com/5822791760/go-api-template/libs/errs"
	"github.com/5822791760/go-api-template/libs/helpers"
	"github.com/5822791760/go-api-template/types"

	"github.com/5822791760/go-api-template/.gen/postgres/public/model"
	. "github.com/5822791760/go-api-template/.gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"
)

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{
		db: db,
	}
}

func (s *BookService) CreateBook(body reqs.CreateBook) errs.ErrRenderer {
	stmt := Books.INSERT(
		Books.Name,
		Books.Bookno,
		Books.Price,
		Books.Summary,
		Books.AuthorID,
		Books.Amount,
		Books.CreatedAt,
		Books.UpdatedAt,
	).
		VALUES(
			body.Name,
			body.Bookno,
			body.Price,
			*body.Summary,
			body.AuthorID,
			body.Amount,
			DEFAULT,
			DEFAULT,
		)

	if _, err := stmt.Exec(s.db); err != nil {
		return errs.NewErr(err, http.StatusInternalServerError)
	}

	return nil
}

func (s *BookService) UpdateBook(id int32, body reqs.UpdateBook) (model.Books, errs.ErrRenderer) {
	stmt := Books.UPDATE(
		Books.Name,
		Books.Bookno,
		Books.Price,
		Books.Summary,
		Books.Amount,
		Books.AuthorID,
		Books.UpdatedAt,
	).SET(
		body.Name,
		body.Bookno,
		body.Price,
		*body.Summary,
		body.Amount,
		body.AuthorID,
		DEFAULT,
	).WHERE(Books.ID.EQ(Int(int64(id)))).RETURNING(Books.AllColumns)

	var returningBook model.Books
	if err := stmt.Query(s.db, &returningBook); err != nil {
		return model.Books{}, errs.NewErr(err, http.StatusInternalServerError)
	}

	return returningBook, nil
}

func (s *BookService) BuyBook(bookID int32, userID int32) (types.BuyBook, errs.ErrRenderer) {
	tx, err := helpers.Transactions(s.db)
	if err != nil {
		return types.BuyBook{}, err
	}

	defer tx.Rollback()

	var returningBook struct {
		ID     int32
		Name   string
		Amount int32
		Price  string
	}
	stmt := Books.UPDATE(
		Books.Amount,
	).
		SET(
			Books.Amount.SUB(Int(1)),
		).
		WHERE(Books.ID.EQ(Int(int64(bookID)))).
		RETURNING(
			Books.ID.AS("ID"),
			Books.Name.AS("Name"),
			Books.Amount.AS("Amount"),
			Books.Price.AS("Price"),
		)

	if err := stmt.Query(tx, &returningBook); err != nil {
		return types.BuyBook{}, errs.NewErr(err, http.StatusInternalServerError)
	}

	if returningBook.Amount < 0 {
		return types.BuyBook{}, errs.NewErrByString("Not enough book in inventory", http.StatusBadRequest)
	}

	var returningUser struct {
		ID   int32
		Cash string
		Name string
	}
	stmt = Users.UPDATE(
		Users.Cash,
	).
		SET(
			Users.Cash.SUB(Decimal(returningBook.Price)),
		).
		WHERE(Users.ID.EQ(Int(int64(userID)))).
		RETURNING(
			Users.ID.AS("ID"),
			Users.Cash.AS("Cash"),
			Users.Name.AS("Name"),
		)

	if err := stmt.Query(tx, &returningUser); err != nil {
		return types.BuyBook{}, errs.NewErr(err, http.StatusInternalServerError)
	}

	cash, err := helpers.DecimalToInt(returningUser.Cash)
	if err != nil {
		return types.BuyBook{}, err
	}

	if cash < 0 {
		return types.BuyBook{}, errs.NewErrByString("User don't have enough cash", http.StatusBadRequest)
	}

	cStmt := BookSlips.INSERT(
		BookSlips.ID,
		BookSlips.PaidAmount,
		BookSlips.BookID,
		BookSlips.CreatedByID,
		BookSlips.CreatedAt,
		BookSlips.Buyer,
		BookSlips.BookName,
	).
		VALUES(
			DEFAULT,
			returningBook.Price,
			returningBook.ID,
			returningUser.ID,
			DEFAULT,
			returningUser.Name,
			returningBook.Name,
		)

	if _, err := cStmt.Exec(tx); err != nil {
		return types.BuyBook{}, errs.NewErr(err, http.StatusInternalServerError)
	}

	tx.Commit()

	return types.BuyBook{
		ID:     returningBook.ID,
		Name:   returningBook.Name,
		Amount: returningBook.Amount,
	}, nil
}
