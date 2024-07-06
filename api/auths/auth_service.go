package auths

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/5822791760/go-api-template/api/auths/reqs"
	"github.com/5822791760/go-api-template/api/auths/res"
	"github.com/5822791760/go-api-template/libs/errs"
	"github.com/5822791760/go-api-template/libs/helpers"

	"github.com/5822791760/go-api-template/.gen/postgres/public/model"
	. "github.com/5822791760/go-api-template/.gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (s *AuthService) SignUp(body reqs.SignUpRequest, resp *res.SignUpResponse) errs.ErrRenderer {
	var user model.Users

	if err := helpers.ShouldNotExists(s.db, SELECT(Int(1)).FROM(Users).WHERE(Users.Email.EQ(String(body.Email)))); err != nil {
		return err
	}

	stmt := Users.
		INSERT(
			Users.ID,
			Users.Email,
			Users.Password,
			Users.Name,
			Users.LastSignInAt,
		).
		VALUES(
			DEFAULT,
			body.Email,
			body.Password,
			body.Name,
			nil,
		).RETURNING(Users.ID)

	if err := stmt.Query(s.db, &user); err != nil {
		return errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	if err := s.SignInUser(user.ID, &resp.SignInResponse); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) SignIn(body reqs.SignInRequest, resp *res.SignInResponse) errs.ErrRenderer {
	var user model.Users

	stmt := SELECT(Users.ID, Users.Password).
		FROM(Users).
		WHERE(
			Users.Email.EQ(String(body.Email)),
		).LIMIT(1)

	if err := stmt.Query(s.db, &user); err != nil {
		errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	if err := helpers.CheckPasswordHash(body.Password, user.Password); err != nil {
		return err
	}

	if err := s.SignInUser(user.ID, resp); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) AssignAccessToken(userId int32, token *string) errs.ErrRenderer {
	if err := helpers.EncodeJwt(helpers.JwtPayload{ID: strconv.Itoa(int(userId))}, token); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) SignInUser(userId int32, resp *res.SignInResponse) errs.ErrRenderer {
	var user model.Users
	var token string

	if err := s.AssignAccessToken(userId, &token); err != nil {
		return err
	}

	updateStmt := Users.UPDATE(Users.LastSignInAt).SET(DEFAULT).WHERE(Users.ID.EQ(Int(int64(userId)))).RETURNING(Users.LastSignInAt)
	if err := updateStmt.Query(s.db, &user); err != nil {
		errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	*resp = res.SignInResponse{
		AccessToken:  token,
		LastSignInAt: helpers.FormatDateTime(user.LastSignInAt),
	}

	return nil
}
