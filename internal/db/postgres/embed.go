package postgres

import (
	"embed"
)

//go:embed migrations/*.sql
var MigrationFiles embed.FS
