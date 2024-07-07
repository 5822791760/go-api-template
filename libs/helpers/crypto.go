package helpers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/5822791760/go-api-template/config"
	"github.com/5822791760/go-api-template/libs/errs"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JwtPayload struct {
	ID int32
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

func ParseJwt(tokenString string, claims *jwt.MapClaims) errs.ErrRenderer {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetJwtSecret()), nil
	})

	if err != nil {
		return errs.NewErr(err, errs.ErrGeneric, http.StatusInternalServerError)
	}

	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errs.NewErrByString("Unable to claim token", errs.ErrGeneric, http.StatusInternalServerError)
	}

	*claims = tokenClaims

	return nil
}

func ExtractUserID(r *http.Request, currentID *int32) errs.ErrRenderer {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	id, ok := claims["id"].(float64)
	if !ok {
		return errs.NewErrByString("id in token not found", errs.ErrGeneric, http.StatusBadRequest)
	}

	*currentID = int32(id)

	return nil
}
