package books

import (
	"net/http"

	"github.com/5822791760/go-api-template/libs"
)

type BookController struct {
	useCase *BookUseCase
}

func NewBookController(useCase *BookUseCase) *BookController {
	return &BookController{
		useCase: useCase,
	}
}

func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	resp, err := c.useCase.GetBooks()
	if err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusOK, resp)
}
