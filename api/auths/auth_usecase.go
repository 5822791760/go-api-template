package auths

import (
	"database/sql"

	"github.com/5822791760/go-api-template/api/auths/reqs"
	"github.com/5822791760/go-api-template/api/auths/res"
	"github.com/5822791760/go-api-template/libs/errs"
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

func (u *AuthUseCase) SignUp(body reqs.SignUpRequest, res *res.SignUpResponse) errs.ErrRenderer {
	if err := u.authService.SignUp(body, res); err != nil {
		return err
	}

	return nil
}

func (u *AuthUseCase) SignIn(body reqs.SignInRequest, res *res.SignInResponse) errs.ErrRenderer {
	if err := u.authService.SignIn(body, res); err != nil {
		return err
	}

	return nil
}
