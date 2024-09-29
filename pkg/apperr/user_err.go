package apperr

import "net/http"

const (
	UserNotFound = "notFound"
)

func NewUserNotFoundErr() Err {
	return &errBase{
		Code:         http.StatusNotFound,
		ErrorMessage: "Not found",
		Context:      errContext{AuthorNotFound: "User not found"},
	}
}
