package helpers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/5822791760/go-api-template/libs/errs"
	"github.com/go-chi/chi/v5"
)

func GetIDParam(r *http.Request) (int32, errs.ErrRenderer) {
	paramReq := chi.URLParam(r, "id")

	if paramReq == "" {
		return 0, errs.ErrRender{
			StatusText: errs.ErrParamNotFound,
			Code:       http.StatusBadRequest,
			ErrorText:  "Params id not found",
		}
	}

	id, err := strconv.Atoi(paramReq)
	if err != nil {
		return 0, errs.ErrRender{
			StatusText: errs.ErrGeneric,
			Code:       http.StatusBadRequest,
			ErrorText:  err.Error(),
		}
	}

	return int32(id), nil
}

func GetURLParam(r *http.Request, key string) (string, errs.ErrRenderer) {
	param := chi.URLParam(r, key)
	if param == "" {
		return param, errs.ErrRender{
			StatusText: errs.ErrParamNotFound,
			Code:       http.StatusBadRequest,
			ErrorText:  fmt.Sprintf("Params %s not found", key),
		}
	}

	return param, nil
}
