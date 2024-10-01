package repos

import (
	"database/sql"
	"errors"
)

func notFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
