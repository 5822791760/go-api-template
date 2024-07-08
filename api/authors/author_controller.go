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

func (c *AuthorController) GetAuthors(w http.ResponseWriter, r *http.Request) {
	resp, err := c.useCase.GetAuthors()
	if err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusOK, resp)
}

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

func (c *AuthorController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var body reqs.CreateAuthorRequest
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

func (c *AuthorController) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetIDParam(r)
	if err != nil {
		err.Render(w)
		return
	}

	body := reqs.UpdateAuthorRequest{}
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
