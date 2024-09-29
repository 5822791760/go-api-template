package coreutil

import (
	"strings"

	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/georgysavva/scany/v2/sqlscan"
)

func NewScanner() (*sqlscan.API, error) {
	dbscanAPI, err := sqlscan.NewDBScanAPI(
		dbscan.WithFieldNameMapper(strings.ToLower),
		dbscan.WithStructTagKey("db"),
	)
	if err != nil {
		return nil, err
	}
	scan, err := sqlscan.NewAPI(dbscanAPI)
	if err != nil {
		return nil, err
	}

	return scan, nil
}
