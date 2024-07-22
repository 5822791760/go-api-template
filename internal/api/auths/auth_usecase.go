package auths

import (
	"database/sql"

	"github.com/5822791760/go-api-template/internal/api/auths/reqs"
	"github.com/5822791760/go-api-template/internal/api/auths/res"
	"github.com/5822791760/go-api-template/pkg/errs"
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

func (u *AuthUseCase) SignUp(body reqs.SignUp) (res.SignUp, errs.ErrRenderer) {
	signInToken, err := u.authService.SignUp(SignUpBody(body))
	if err != nil {
		return res.SignUp{}, err
	}

	resp := res.SignUp{
		AccessToken:  signInToken.AccessToken,
		LastSignInAt: signInToken.LastSignInAt,
	}

	return resp, nil
}

func (u *AuthUseCase) SignIn(body reqs.SignIn) (res.SignIn, errs.ErrRenderer) {
	signInToken, err := u.authService.SignInByUserEmail(body.Email, body.Password)
	if err != nil {
		return res.SignIn{}, err
	}

	resp := res.SignIn{
		AccessToken:  signInToken.AccessToken,
		LastSignInAt: signInToken.LastSignInAt,
	}

	return resp, nil
}
