package helpers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/5822791760/go-api-template/libs/errs"
	"github.com/go-chi/chi/v5"
)

func URLIntParam(r *http.Request, id *int32) errs.ErrRenderer {
	paramReq := chi.URLParam(r, "id")

	if paramReq == "" {
		return errs.ErrRender{
			StatusText: errs.ErrParamNotFound,
			Code:       http.StatusBadRequest,
			ErrorText:  "Params id not found",
		}
	}

	paramId, err := strconv.Atoi(paramReq)
	if err != nil {
		return errs.ErrRender{
			StatusText: errs.ErrGeneric,
			Code:       http.StatusBadRequest,
			ErrorText:  err.Error(),
		}
	}

	*id = int32(paramId)

	return nil
}

func URLParam(r *http.Request, key string, param *string) errs.ErrRenderer {
	paramReq := chi.URLParam(r, key)
	if paramReq == "" {
		return errs.ErrRender{
			StatusText: errs.ErrParamNotFound,
			Code:       http.StatusBadRequest,
			ErrorText:  fmt.Sprintf("Params %s not found", key),
		}
	}

	*param = paramReq

	return nil
}
