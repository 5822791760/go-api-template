package authors

import (
	"net/http"

	"github.com/5822791760/go-api-template/api/authors/requests"
	"github.com/5822791760/go-api-template/api/authors/responses"
	"github.com/5822791760/go-api-template/libs/helpers"
	"github.com/unrolled/render"
)

type AuthorController struct {
	render  *render.Render
	useCase *AuthorUseCase
}

func NewAuthorController(render *render.Render, useCase *AuthorUseCase) *AuthorController {
	return &AuthorController{
		render:  render,
		useCase: useCase,
	}
}

func (c *AuthorController) GetAuthors(w http.ResponseWriter, r *http.Request) {
	res := []responses.GetAuthorsResponse{}

	if err := c.useCase.GetAuthors(&res); err != nil {
		err.Render(w, c.render)
		return
	}

	c.render.JSON(w, http.StatusOK, res)
}

func (c *AuthorController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var createAuthorReq requests.CreateAuthorRequest
	var res responses.CreateAuthorResponse

	if err := helpers.Decode(r, &createAuthorReq); err != nil {
		err.Render(w, c.render)
		return
	}

	if err := c.useCase.CreateAuthor(createAuthorReq, &res); err != nil {
		err.Render(w, c.render)
		return
	}

	c.render.JSON(w, http.StatusOK, res)
}
