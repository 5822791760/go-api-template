package dbutil

import (
	j "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

func SelectExist() j.SelectStatement {
	return j.SELECT(j.Int(1))
}

func IsExist(db qrm.DB, statement j.SelectStatement) (bool, error) {
	var data struct {
		Exists bool
	}
	stmt := j.
		SELECT(
			j.EXISTS(statement).AS("Exists"),
		)

	if err := stmt.Query(db, &data); err != nil {
		return false, err
	}

	return data.Exists, nil
}
