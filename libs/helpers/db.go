package helpers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/5822791760/go-api-template/libs/errs"
	. "github.com/go-jet/jet/v2/postgres"
)

type RowExist struct {
	Exists bool `json:"exists"`
}

func CheckAffectedRow(res sql.Result) errs.ErrRenderer {
	if a, _ := res.RowsAffected(); a == 0 {
		return errs.NewErr(errors.New("Updated row not found"), errs.ErrQuery, http.StatusNotFound)
	}

	return nil
}

func ShouldNotExists(db *sql.DB, statement SelectStatement) errs.ErrRenderer {
	var rowExists RowExist

	stmt := SELECT(EXISTS(statement).AS("RowExist.Exists"))
	if err := stmt.Query(db, &rowExists); err != nil {
		return errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	if rowExists.Exists == true {
		return errs.NewErr(errors.New("This Data already exist"), errs.ErrQuery, http.StatusBadRequest)
	}

	return nil
}

func ShouldNotExistsTx(db *sql.Tx, statement SelectStatement) errs.ErrRenderer {
	var rowExists RowExist

	stmt := SELECT(EXISTS(statement).AS("RowExist.Exists"))
	if err := stmt.Query(db, &rowExists); err != nil {
		return errs.NewErr(err, errs.ErrQuery, http.StatusInternalServerError)
	}

	if rowExists.Exists == true {
		return errs.NewErr(errors.New("This Data already exist"), errs.ErrQuery, http.StatusBadRequest)
	}

	return nil
}

func FormatDateTime(date interface{ Format(string) string }) string {
	return date.Format(time.RFC3339)
}
