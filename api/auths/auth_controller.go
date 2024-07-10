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

// SignUp godoc
//
//	@Description	Register User using user info
//	@Tags			auths
//	@Accept			json
//	@Produce		json
//
//	@Param			request	body		reqs.SignUp	true	"Sign up information"
//
//	@Success		200		{object}	res.SignUp
//	@Failure		400		{object}	errs.ErrRender
//	@Router			/public/auths/sign_up [post]
func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	body := reqs.SignUp{}
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

// SignIn godoc
//
//	@Description	Signin using email, password
//	@Tags			auths
//	@Accept			json
//	@Produce		json
//
//	@Param			request		body		reqs.SignIn	true	"Sign in information"
//
//	@Success		200			{object}	res.SignIn
//	@Failure		400,500,401	{object}	errs.ErrRender
//	@Router			/public/auths/sign_in [post]
func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	body := reqs.SignIn{}
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
