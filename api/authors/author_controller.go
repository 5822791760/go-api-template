package authors

import (
	"net/http"

	"github.com/unrolled/render"
)

type AuthorController struct {
	render *render.Render
	useCase *AuthorUseCase
}

func NewAuthorController(render *render.Render, useCase *AuthorUseCase) *AuthorController {
	return &AuthorController{
		render: render,
		useCase: useCase,
	}
}

func (c *AuthorController) GetAuthors(w http.ResponseWriter, r *http.Request) {
	res, err := c.useCase.GetAuthors()
	if err != nil {
		err.Render(w, c.render)
		return
	}

	c.render.JSON(w, http.StatusOK, res)
}
