package helpers

import (
	"errors"
	"net/http"

	"github.com/5822791760/go-api-template/config"
	"github.com/5822791760/go-api-template/libs/errs"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JwtPayload struct {
	ID string
}

func HashPassword(password *string) errs.ErrRenderer {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 4)
	if err != nil {
		return errs.NewErr(err, errs.ErrGeneric, http.StatusInternalServerError)
	}

	*password = string(bytes)

	return nil
}

func CheckPasswordHash(password, hash string) errs.ErrRenderer {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errs.NewErr(errors.New("Password does not match"), errs.ErrValidate, http.StatusBadRequest)
	}

	return nil
}

func EncodeJwt(data JwtPayload, token *string) errs.ErrRenderer {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": data.ID,
	})

	secretKey := []byte(config.GetJwtSecret())
	tokenString, err := t.SignedString(secretKey)
	if err != nil {
		return errs.NewErr(err, errs.ErrGeneric, http.StatusInternalServerError)
	}

	*token = tokenString

	return nil
}
