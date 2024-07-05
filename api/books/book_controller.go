package books

import (
	"net/http"

	"github.com/unrolled/render"
)

type BookController struct {
	render *render.Render
	useCase *BookUseCase
}

func NewBookController(render *render.Render, useCase *BookUseCase) *BookController {
	return &BookController{
		useCase: useCase,
		render: render,
	}
}

func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	res, err := c.useCase.GetBooks()
	if err != nil {
		err.Render(w, c.render)
		return
	}

	c.render.JSON(w, http.StatusOK, res)
}
