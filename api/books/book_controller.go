package books

import (
	"net/http"

	"github.com/5822791760/go-api-template/api/books/responses"
	"github.com/unrolled/render"
)

type BookController struct {
	render  *render.Render
	useCase *BookUseCase
}

func NewBookController(render *render.Render, useCase *BookUseCase) *BookController {
	return &BookController{
		useCase: useCase,
		render:  render,
	}
}

func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	res := []responses.GetBooksResponse{}

	if err := c.useCase.GetBooks(&res); err != nil {
		err.Render(w, c.render)
		return
	}

	c.render.JSON(w, http.StatusOK, res)
}
