package initials

import (
	"context"
	"net/http"
	"strings"

	"github.com/5822791760/go-api-template/libs/errs"
	"github.com/5822791760/go-api-template/libs/helpers"
	"github.com/golang-jwt/jwt"
	"github.com/unrolled/render"
)

type MiddlewareService struct {
	render *render.Render
}

func NewMiddlewareService(render *render.Render) *MiddlewareService {
	return &MiddlewareService{
		render: render,
	}
}

func (m *MiddlewareService) JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			errs.RenderErrByString(w, m.render, "Authorization header is required", errs.ErrUnAuthorize, http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			errs.RenderErrByString(w, m.render, "Invalid Authorization header format", errs.ErrUnAuthorize, http.StatusUnauthorized)
			return
		}

		tokenString := bearerToken[1]

		var claims jwt.MapClaims
		if err := helpers.ParseJwt(tokenString, &claims); err != nil {
			errs.RenderErr(w, m.render, err, errs.ErrUnAuthorize, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
