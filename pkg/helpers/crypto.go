package helpers

import (
	"fmt"
	"net/http"

	"github.com/5822791760/go-api-template/internal/config"
	"github.com/5822791760/go-api-template/pkg/errs"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JwtPayload struct {
	ID int32
}

func GetHashedPassword(password string) (string, errs.ErrRenderer) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "", errs.NewErr(err, http.StatusInternalServerError)
	}

	hashedPassword := string(bytes)

	return hashedPassword, nil
}

func CheckPasswordHash(password, hash string) errs.ErrRenderer {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errs.NewErrByString("Password does not match", http.StatusBadRequest)
	}

	return nil
}

func GetEncodedJwt(data JwtPayload) (string, errs.ErrRenderer) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": data.ID,
	})

	secretKey := []byte(config.GetJwtSecret())
	encodedToken, err := t.SignedString(secretKey)
	if err != nil {
		return encodedToken, errs.NewErr(err, http.StatusInternalServerError)
	}

	return encodedToken, nil
}

func ParseJwt(tokenString string) (jwt.MapClaims, errs.ErrRenderer) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetJwtSecret()), nil
	})

	if err != nil {
		return jwt.MapClaims{}, errs.NewErr(err, http.StatusInternalServerError)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, errs.NewErrByString("Unable to claim token", http.StatusInternalServerError)
	}

	return claims, nil
}

func CurrentUserID(r *http.Request) (int32, errs.ErrRenderer) {
	claims := r.Context().Value("claims").(jwt.MapClaims)

	id, ok := claims["id"].(float64)
	if !ok {
		return 0, errs.NewErrByString("id in token not found", http.StatusBadRequest)
	}

	return int32(id), nil
}
