package authors

import (
	"encoding/json"
	"net/http"
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
	res, err := c.useCase.GetAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}
