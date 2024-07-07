package auths

import (
	"database/sql"

	"github.com/5822791760/go-api-template/api/auths/reqs"
	"github.com/5822791760/go-api-template/api/auths/res"
	"github.com/5822791760/go-api-template/libs/errs"
	"github.com/5822791760/go-api-template/types"
)

type AuthUseCase struct {
	db          *sql.DB
	authService *AuthService
}

func NewAuthUseCase(db *sql.DB, authService *AuthService) *AuthUseCase {
	return &AuthUseCase{
		db:          db,
		authService: authService,
	}
}

func (u *AuthUseCase) SignUp(body reqs.SignUpRequest, resp *res.SignUpResponse) errs.ErrRenderer {
	var signInToken types.SignInToken

	if err := u.authService.SignUp(types.SignUpBody(body), &signInToken); err != nil {
		return err
	}

	*resp = res.SignUpResponse(signInToken)

	return nil
}

func (u *AuthUseCase) SignIn(body reqs.SignInRequest, resp *res.SignInResponse) errs.ErrRenderer {
	var signInToken types.SignInToken

	if err := u.authService.SignInByUserEmail(body.Email, body.Password, &signInToken); err != nil {
		return err
	}

	*resp = res.SignInResponse(signInToken)

	return nil
}
