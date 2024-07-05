package helpers

import (
	"net/http"
	"strconv"

	"github.com/5822791760/go-api-template/libs/errs"
	"github.com/go-chi/chi/v5"
)

func URLIntParam(r *http.Request, id *int64) errs.ErrRenderer {
	paramsId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return errs.ErrRender{
			StatusText: errs.ErrGeneric,
			Code:       http.StatusBadRequest,
			ErrorText:  err.Error(),
		}
	}

	*id = int64(paramsId)

	return nil
}

func URLParam(r *http.Request, param *string) errs.ErrRenderer {
	paramReq := chi.URLParam(r, "id")
	if paramReq == "" {
		return errs.ErrRender{
			StatusText: errs.ErrGeneric,
			Code:       http.StatusBadRequest,
			ErrorText:  "Wrong Params",
		}
	}

	*param = paramReq

	return nil
}
