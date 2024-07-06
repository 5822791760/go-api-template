package auths

import (
	"net/http"

	"github.com/5822791760/go-api-template/api/auths/reqs"
	"github.com/5822791760/go-api-template/api/auths/res"
	"github.com/5822791760/go-api-template/libs/helpers"
	"github.com/unrolled/render"
)

type AuthController struct {
	render  *render.Render
	useCase *AuthUseCase
}

func NewAuthController(render *render.Render, useCase *AuthUseCase) *AuthController {
	return &AuthController{
		useCase: useCase,
		render:  render,
	}
}

func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	body := reqs.SignUpRequest{}
	res := res.SignUpResponse{}

	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w, c.render)
		return
	}

	if err := helpers.HashPassword(&body.Password); err != nil {
		err.Render(w, c.render)
		return
	}

	if err := c.useCase.SignUp(body, &res); err != nil {
		err.Render(w, c.render)
		return
	}

	c.render.JSON(w, http.StatusCreated, res)
}

func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	body := reqs.SignInRequest{}
	resp := res.SignInResponse{}

	if err := helpers.Decode(r, &body); err != nil {
		err.Render(w, c.render)
		return
	}

	if err := c.useCase.SignIn(body, &resp); err != nil {
		err.Render(w, c.render)
		return
	}

	c.render.JSON(w, http.StatusOK, resp)
}
