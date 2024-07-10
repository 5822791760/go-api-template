package authors

import (
	"net/http"

	"github.com/5822791760/go-api-template/api/authors/reqs"
	"github.com/5822791760/go-api-template/libs"
	"github.com/5822791760/go-api-template/libs/helpers"
)

type AuthorController struct {
	useCase *AuthorUseCase
}

func NewAuthorController(useCase *AuthorUseCase) *AuthorController {
	return &AuthorController{
		useCase: useCase,
	}
}

// GetAuthors godoc
//
//	@Description	List All Authors
//	@Tags			authors
//	@Accept			json
//	@Produce		json
//
//	@Success		200		{object}	[]res.GetAuthors
//	@Failure		500,401	{object}	errs.ErrRender
//
//	@Security		Bearer
//
//	@Router			/authors [get]
func (c *AuthorController) GetAuthors(w http.ResponseWriter, r *http.Request) {
	resp, err := c.useCase.GetAuthors()
	if err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusOK, resp)
}

// GetAuthors godoc
//
//	@Description	Get One Author
//	@Tags			authors
//	@Accept			json
//	@Produce		json
//
//	@Param			id		path		int	true	"Author ID"
//
//	@Success		200		{object}	res.GetAuthor
//	@Failure		500,401	{object}	errs.ErrRender
//
//	@Security		Bearer
//
//	@Router			/authors/{id} [get]
func (c *AuthorController) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDParam(r)
	if err != nil {
		err.Render(w)
		return
	}

	resp, err := c.useCase.GetAuthor(id)
	if err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusOK, resp)
}

// CreateAuthor godoc
//
//	@Description	Create an Author
//	@Tags			authors
//	@Accept			json
//	@Produce		json
//
//	@Param			request		body		reqs.CreateAuthor	true	"Author info"
//
//	@Success		200			{object}	res.CreateAuthor
//	@Failure		500,401,400	{object}	errs.ErrRender
//
//	@Security		Bearer
//
//	@Router			/authors [post]
func (c *AuthorController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var body reqs.CreateAuthor
	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w)
		return
	}

	resp, err := c.useCase.CreateAuthor(body)
	if err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusOK, resp)
}

// UpdateAuthor godoc
//
//	@Description	Update an Author
//	@Tags			authors
//	@Accept			json
//	@Produce		json
//
//	@Param			request	body		reqs.UpdateAuthor	true	"Author update info"
//
//	@Success		200		{object}	res.UpdateAuthor
//	@Failure		500,401	{object}	errs.ErrRender
//
//	@Security		Bearer
//
//	@Router			/authors/{id} [patch]
func (c *AuthorController) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDParam(r)
	if err != nil {
		err.Render(w)
		return
	}

	body := reqs.UpdateAuthor{}
	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w)
		return
	}

	resp, err := c.useCase.UpdateAuthor(int32(id), body)
	if err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusOK, resp)
}
