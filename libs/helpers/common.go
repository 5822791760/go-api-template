package helpers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/5822791760/go-api-template/libs/errs"
)

func DecimalToInt(s string) (int, errs.ErrRenderer) {
	s = strings.TrimRight(strings.TrimRight(s, "0"), ".")

	result, err := strconv.Atoi(s)
	if err != nil {
		return 0, errs.NewErr(err, http.StatusInternalServerError)
	}

	return result, nil
}
