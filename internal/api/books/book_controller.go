package books

import (
	"net/http"

	"github.com/5822791760/go-api-template/internal/api/books/reqs"
	"github.com/5822791760/go-api-template/pkg/consts"
	"github.com/5822791760/go-api-template/pkg/helpers"
)

type BookController struct {
	useCase *BookUseCase
}

func NewBookController(useCase *BookUseCase) *BookController {
	return &BookController{
		useCase: useCase,
	}
}

// GetBooks godoc
//
//	@Description	List All Books
//	@Tags			books
//	@Accept			json
//	@Produce		json
//
//	@Success		200		{object}	[]res.GetBooks
//	@Failure		500,401	{object}	errs.ErrRender
//
//	@Security		Bearer
//
//	@Router			/books [get]
func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	resp, err := c.useCase.GetBooks()
	if err != nil {
		err.Render(w)
		return
	}

	consts.Render.JSON(w, http.StatusOK, resp)
}

// CreateBook godoc
//
//	@Description	Create a Book
//	@Tags			books
//	@Accept			json
//	@Produce		json
//
//	@Param			request		body		reqs.CreateBook	true	"Book info"
//
//	@Success		200			{object}	res.CreateBook
//	@Failure		500,401,400	{object}	errs.ErrRender
//
//	@Security		Bearer
//
//	@Router			/books [post]
func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var body reqs.CreateBook

	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w)
		return
	}

	resp, err := c.useCase.CreateBooks(body)
	if err != nil {
		err.Render(w)
		return
	}

	consts.Render.JSON(w, http.StatusCreated, resp)
}

// UpdateBook godoc
//
//	@Description	Update a Book
//	@Tags			books
//	@Accept			json
//	@Produce		json
//
//	@Param			id			path		int				true	"Book ID"
//	@Param			request		body		reqs.UpdateBook	true	"Book info"
//
//	@Success		200			{object}	res.UpdateBook
//	@Failure		500,401,400	{object}	errs.ErrRender
//
//	@Security		Bearer
//
//	@Router			/books/{id} [put]
func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var body reqs.UpdateBook

	id, err := helpers.GetIDParam(r)
	if err != nil {
		err.Render(w)
		return
	}

	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w)
		return
	}

	resp, err := c.useCase.UpdateBooks(id, body)
	if err != nil {
		err.Render(w)
		return
	}

	consts.Render.JSON(w, http.StatusOK, resp)
}

// BuyBook godoc
//
//	@Description	Buy a Book
//	@Tags			books
//	@Accept			json
//	@Produce		json
//
//	@Param			id			path		int	true	"Book ID"
//
//	@Success		200			{object}	res.BuyBook
//	@Failure		500,401,400	{object}	errs.ErrRender
//
//	@Security		Bearer
//
//	@Router			/books/{id}/buy [patch]
func (c *BookController) BuyBook(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDParam(r)
	if err != nil {
		err.Render(w)
		return
	}

	userID, err := helpers.CurrentUserID(r)
	if err != nil {
		err.Render(w)
		return
	}

	resp, err := c.useCase.BuyBook(id, userID)
	if err != nil {
		err.Render(w)
		return
	}

	consts.Render.JSON(w, http.StatusOK, resp)
}
