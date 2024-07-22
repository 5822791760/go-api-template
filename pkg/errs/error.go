package errs

import (
	"net/http"

	"github.com/5822791760/go-api-template/pkg/consts"
)

type ErrRenderer interface {
	Render(w http.ResponseWriter)
	Error() string
}

type ErrRender struct {
	Code      int             `json:"code"`  // application-specific error code
	ErrorText string          `json:"error"` // application-level error message, for debugging
	Context   map[string]bool `json:"context"`
}

func NewErr(err error, code int) *ErrRender {
	return &ErrRender{
		Code:      code,
		ErrorText: err.Error(),
	}
}

func NewErrByString(message string, code int) *ErrRender {
	return &ErrRender{
		Code:      code,
		ErrorText: message,
	}
}

func (e *ErrRender) Render(w http.ResponseWriter) {
	consts.Render.JSON(w, e.Code, e)
	return
}

func (e *ErrRender) Error() string {
	return e.ErrorText
}

func (e *ErrRender) WithContext(context map[string]bool) *ErrRender {
	e.Context = context
	return e
}
