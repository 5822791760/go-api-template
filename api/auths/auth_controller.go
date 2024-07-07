package auths

import (
	"net/http"

	"github.com/5822791760/go-api-template/api/auths/reqs"
	"github.com/5822791760/go-api-template/api/auths/res"
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
	res := res.SignUpResponse{}

	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w)
		return
	}

	if err := helpers.HashPassword(&body.Password); err != nil {
		err.Render(w)
		return
	}

	if err := c.useCase.SignUp(body, &res); err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusCreated, res)
}

func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	body := reqs.SignInRequest{}
	resp := res.SignInResponse{}

	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w)
		return
	}

	if err := c.useCase.SignIn(body, &resp); err != nil {
		err.Render(w)
		return
	}

	libs.Render.JSON(w, http.StatusOK, resp)
}
