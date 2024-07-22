package helpers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/5822791760/go-api-template/pkg/errs"
	"github.com/go-chi/chi/v5"
)

func GetIDParam(r *http.Request) (int32, errs.ErrRenderer) {
	paramReq := chi.URLParam(r, "id")

	if paramReq == "" {
		return 0, errs.NewErrByString("Params id not found", http.StatusBadRequest)
	}

	id, err := strconv.Atoi(paramReq)
	if err != nil {
		return 0, errs.NewErr(err, http.StatusBadRequest)
	}

	return int32(id), nil
}

func GetURLParam(r *http.Request, key string) (string, errs.ErrRenderer) {
	param := chi.URLParam(r, key)
	if param == "" {
		return param, errs.NewErrByString(fmt.Sprintf("Params %s not found", key), http.StatusBadRequest)
	}

	return param, nil
}
