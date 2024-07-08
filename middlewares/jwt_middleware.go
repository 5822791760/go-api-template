package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/5822791760/go-api-template/libs/errs"
	"github.com/5822791760/go-api-template/libs/helpers"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			errs.RenderErrByString(w, "Authorization header is required", errs.ErrUnAuthorize, http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			errs.RenderErrByString(w, "Invalid Authorization header format", errs.ErrUnAuthorize, http.StatusUnauthorized)
			return
		}

		tokenString := bearerToken[1]

		claims, err := helpers.ParseJwt(tokenString)
		if err != nil {
			errs.RenderErr(w, err, errs.ErrUnAuthorize, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
