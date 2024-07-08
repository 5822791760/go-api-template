package auths

import (
	"net/http"

	"github.com/5822791760/go-api-template/api/auths/reqs"
	"github.com/5822791760/go-api-template/libs"
	"github.com/5822791760/go-api-template/libs/helpers"
)

type AuthController struct {
	useCase *AuthUseCase
}

func NewAuthController(useCase *AuthUseCase) *AuthController {
	return &AuthController{
		useCase: useCase,
	}
}

func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	body := reqs.SignUpRequest{}
	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w)
		return
	}

	hashedPassword, err := helpers.GetHashedPassword(body.Password)
	if err != nil {
		err.Render(w)
	}

	*&body.Password = hashedPassword

	resp, err := c.useCase.SignUp(body)
	if err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusCreated, resp)
}

func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	body := reqs.SignInRequest{}
	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w)
		return
	}

	resp, err := c.useCase.SignIn(body)
	if err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusOK, resp)
}
