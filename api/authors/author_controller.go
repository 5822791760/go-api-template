package authors

import (
	"net/http"

	"github.com/5822791760/go-api-template/api/authors/reqs"
	"github.com/5822791760/go-api-template/api/authors/res"
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
	res := []res.GetAuthorsResponse{}

	if err := c.useCase.GetAuthors(&res); err != nil {
		err.Render(w, c.render)
		return
	}

	c.render.JSON(w, http.StatusOK, res)
}

func (c *AuthorController) GetAuthor(w http.ResponseWriter, r *http.Request) {
	res := res.GetAuthorResponse{}
	var id int32

	if err := helpers.URLIntParam(r, &id); err != nil {
		err.Render(w, c.render)
		return
	}

	if err := c.useCase.GetAuthor(id, &res); err != nil {
		err.Render(w, c.render)
		return
	}

	c.render.JSON(w, http.StatusOK, res)
}

func (c *AuthorController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var body reqs.CreateAuthorRequest
	var res res.CreateAuthorResponse

	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w, c.render)
		return
	}

	if err := c.useCase.CreateAuthor(body, &res); err != nil {
		err.Render(w, c.render)
		return
	}

	c.render.JSON(w, http.StatusOK, res)
}

func (c *AuthorController) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var id int32
	body := reqs.UpdateAuthorRequest{}
	res := res.UpdateAuthorResponse{}

	if err := helpers.URLIntParam(r, &id); err != nil {
		err.Render(w, c.render)
		return
	}

	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w, c.render)
		return
	}

	if err := c.useCase.UpdateAuthor(int32(id), body, &res); err != nil {
		err.Render(w, c.render)
		return
	}

	c.render.JSON(w, http.StatusOK, res)
}
