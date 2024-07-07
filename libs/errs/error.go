package errs

import (
	"net/http"

	"github.com/5822791760/go-api-template/libs"
)

type ErrRenderer interface {
	Render(w http.ResponseWriter)
	Error() string
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

func NewErrByString(message string, status string, code int) ErrRender {
	return ErrRender{
		StatusText: status,
		Code:       code,
		ErrorText:  message,
	}
}

func RenderErr(w http.ResponseWriter, err error, status string, code int) {
	libs.Render.JSON(w, code, ErrRender{
		StatusText: status,
		Code:       code,
		ErrorText:  err.Error(),
	})
}

func RenderErrByString(w http.ResponseWriter, message string, status string, code int) {
	libs.Render.JSON(w, code, ErrRender{
		StatusText: status,
		Code:       code,
		ErrorText:  message,
	})
}

func (e ErrRender) Render(w http.ResponseWriter) {
	libs.Render.JSON(w, e.Code, e)
	return
}

func (e ErrRender) Error() string {
	return e.ErrorText
}
