package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/5822791760/go-api-template/pkg/errs"
	"github.com/5822791760/go-api-template/pkg/helpers"
)

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			errs.NewErrByString("Authorization header is required", http.StatusUnauthorized).Render(w)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			errs.NewErrByString("Invalid Authorization header format", http.StatusUnauthorized).Render(w)
			return
		}

		tokenString := bearerToken[1]

		claims, err := helpers.ParseJwt(tokenString)
		if err != nil {
			errs.NewErr(err, http.StatusUnauthorized).Render(w)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
