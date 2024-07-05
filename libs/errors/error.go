package errors

import (
	"net/http"

	"github.com/unrolled/render"
)

type ErrRenderer interface {
	Render(w http.ResponseWriter, render *render.Render)
}

type ErrRender struct {
	StatusText string `json:"status"`          // user-level status message
	Code       int    `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func NewErr(err error, status string, code int) ErrRender {
	return ErrRender{
		StatusText: status,
		Code:       code,
		ErrorText:  err.Error(),
	}
}

func (e ErrRender) Render(w http.ResponseWriter, render *render.Render) {
	render.JSON(w, e.Code, e)
	return
}
