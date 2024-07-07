package authors

import (
	"net/http"

	"github.com/5822791760/go-api-template/api/authors/reqs"
	"github.com/5822791760/go-api-template/api/authors/res"
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
	res := []res.GetAuthorsResponse{}
	var currentUserID int32

	if err := helpers.ExtractUserID(r, &currentUserID); err != nil {
		err.Render(w)
		return
	}

	if err := c.useCase.GetAuthors(&res); err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusOK, res)
}

func (c *AuthorController) GetAuthor(w http.ResponseWriter, r *http.Request) {
	res := res.GetAuthorResponse{}
	var id int32

	if err := helpers.URLIntParam(r, &id); err != nil {
		err.Render(w)
		return
	}

	if err := c.useCase.GetAuthor(id, &res); err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusOK, res)
}

func (c *AuthorController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var body reqs.CreateAuthorRequest
	var res res.CreateAuthorResponse

	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w)
		return
	}

	if err := c.useCase.CreateAuthor(body, &res); err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusOK, res)
}

func (c *AuthorController) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var id int32
	body := reqs.UpdateAuthorRequest{}
	res := res.UpdateAuthorResponse{}

	if err := helpers.URLIntParam(r, &id); err != nil {
		err.Render(w)
		return
	}

	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w)
		return
	}

	if err := c.useCase.UpdateAuthor(int32(id), body, &res); err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusOK, res)
}
