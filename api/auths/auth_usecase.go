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

func (u *AuthUseCase) SignUp(body reqs.SignUpRequest) (res.SignUpResponse, errs.ErrRenderer) {
	signInToken, err := u.authService.SignUp(types.SignUpBody(body))
	if err != nil {
		return res.SignUpResponse{}, err
	}

	resp := res.SignUpResponse{
		AccessToken:  signInToken.AccessToken,
		LastSignInAt: signInToken.LastSignInAt,
	}

	return resp, nil
}

func (u *AuthUseCase) SignIn(body reqs.SignInRequest) (res.SignInResponse, errs.ErrRenderer) {
	signInToken, err := u.authService.SignInByUserEmail(body.Email, body.Password)
	if err != nil {
		return res.SignInResponse{}, err
	}

	resp := res.SignInResponse{
		AccessToken:  signInToken.AccessToken,
		LastSignInAt: signInToken.LastSignInAt,
	}

	return resp, nil
}
