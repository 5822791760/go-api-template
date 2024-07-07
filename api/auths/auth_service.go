package auths

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/5822791760/go-api-template/libs/errs"
	"github.com/5822791760/go-api-template/libs/helpers"
	"github.com/5822791760/go-api-template/types"

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

func (s *AuthService) SignUp(body types.SignUpBody, resp *types.SignInToken) errs.ErrRenderer {
	var token string

	tx, _ := s.db.Begin()
	defer tx.Rollback()

	if err := helpers.ShouldNotExistsTx(
		tx,
		SELECT(Int(1)).
			FROM(Users).
			WHERE(Users.Email.EQ(String(body.Email))),
	); err != nil {
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
		).RETURNING(Users.ID.AS("ID"), Users.Password.AS("Password"))

	var returningA struct {
		ID       int32
		Password string
	}
	if err := stmt.Query(tx, &returningA); err != nil {
		return errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	if err := s.AssignAccessToken(returningA.ID, &token); err != nil {
		return err
	}

	updateStmt := Users.
		UPDATE(Users.LastSignInAt).
		SET(DEFAULT).
		WHERE(Users.ID.EQ(Int(int64(returningA.ID)))).
		RETURNING(Users.LastSignInAt.AS("LastSignInAt"))

	var returningB struct {
		LastSignInAt *time.Time
	}
	if err := updateStmt.Query(tx, &returningB); err != nil {
		return errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	*resp = types.SignInToken{
		AccessToken:  token,
		LastSignInAt: helpers.FormatDateTime(returningB.LastSignInAt),
	}

	tx.Commit()

	return nil
}

func (s *AuthService) AssignAccessToken(userID int32, token *string) errs.ErrRenderer {
	if err := helpers.EncodeJwt(helpers.JwtPayload{ID: userID}, token); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) SignInByUserID(userID int32, password string, resp *types.SignInToken) errs.ErrRenderer {
	var token string

	tx, err := s.db.Begin()
	if err != nil {
		return errs.NewErr(err, errs.ErrGeneric, http.StatusInternalServerError)
	}
	defer tx.Rollback()

	if err := s.CheckPasswordByUserId(userID, password); err != nil {
		return err
	}

	if err := s.AssignAccessToken(userID, &token); err != nil {
		return err
	}

	updateStmt := Users.
		UPDATE(Users.LastSignInAt).
		SET(DEFAULT).
		WHERE(Users.ID.EQ(Int(int64(userID)))).
		RETURNING(Users.LastSignInAt.AS("LastSignInAt"))

	var returningA struct {
		LastSignInAt *time.Time
	}
	if err := updateStmt.Query(tx, &returningA); err != nil {
		return errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	tx.Commit()

	*resp = types.SignInToken{
		AccessToken:  token,
		LastSignInAt: helpers.FormatDateTime(returningA.LastSignInAt),
	}

	return nil
}

func (s *AuthService) SignInByUserEmail(email string, password string, resp *types.SignInToken) errs.ErrRenderer {
	var token string

	tx, err := s.db.Begin()
	if err != nil {
		return errs.NewErr(err, errs.ErrGeneric, http.StatusInternalServerError)
	}
	defer tx.Rollback()

	stmt := SELECT(Users.ID.AS("ID"), Users.Password.AS("Password")).
		FROM(Users).
		WHERE(Users.Email.EQ(String(email))).
		LIMIT(1)

	var selectA struct {
		ID       int32
		Password string
	}
	if err := stmt.Query(s.db, &selectA); err != nil {
		return errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	if err := helpers.CheckPasswordHash(password, selectA.Password); err != nil {
		return err
	}

	if err := s.AssignAccessToken(selectA.ID, &token); err != nil {
		return err
	}

	updateStmt := Users.
		UPDATE(Users.LastSignInAt).
		SET(DEFAULT).
		WHERE(Users.ID.EQ(Int(int64(selectA.ID)))).
		RETURNING(Users.LastSignInAt.AS("LastSignInAt"))

	var returningA struct {
		LastSignInAt *time.Time
	}
	if err := updateStmt.Query(s.db, &returningA); err != nil {
		return errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	tx.Commit()

	*resp = types.SignInToken{
		AccessToken:  token,
		LastSignInAt: helpers.FormatDateTime(returningA.LastSignInAt),
	}

	return nil
}

func (s *AuthService) CheckPasswordByUserId(userID int32, password string) errs.ErrRenderer {
	stmt := SELECT(Users.ID.AS("ID"), Users.Password.AS("Password")).
		FROM(Users).
		WHERE(Users.ID.EQ(Int(int64(userID)))).
		LIMIT(1)

	var selectA struct {
		ID       int32
		Password string
	}
	if err := stmt.Query(s.db, &selectA); err != nil {
		return errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	if err := helpers.CheckPasswordHash(password, selectA.Password); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) CheckPasswordByUserEmail(email string, password string) errs.ErrRenderer {
	stmt := SELECT(Users.ID.AS("ID"), Users.Password.AS("Password")).
		FROM(Users).
		WHERE(Users.Email.EQ(String(email))).
		LIMIT(1)

	var selectA struct {
		ID       int32
		Password string
	}
	if err := stmt.Query(s.db, &selectA); err != nil {
		return errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	if err := helpers.CheckPasswordHash(password, selectA.Password); err != nil {
		return err
	}

	return nil
}
