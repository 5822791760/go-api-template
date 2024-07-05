package helpers

import (
	"net/http"
	"strconv"

	"github.com/5822791760/go-api-template/libs/reserrors"
	"github.com/go-chi/chi/v5"
)

func SetURLIntParam(r *http.Request, id *int64) reserrors.ErrRenderer {
	paramsId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return reserrors.ErrRender{
			StatusText: reserrors.ErrGeneric,
			Code:       http.StatusBadRequest,
			ErrorText:  err.Error(),
		}
	}

	*id = int64(paramsId)

	return nil
}

func SetURLParam(r *http.Request, param *string) reserrors.ErrRenderer {
	paramReq := chi.URLParam(r, "id")
	if paramReq == "" {
		return reserrors.ErrRender{
			StatusText: reserrors.ErrGeneric,
			Code:       http.StatusBadRequest,
			ErrorText:  "Wrong Params",
		}
	}

	*param = paramReq

	return nil
}
