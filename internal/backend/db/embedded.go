package database

import "embed"

//go:embed hr/migrations/*.sql
var hrembed embed.FS

func NewHrMigration() (embed.FS, string, string) {
	return hrembed, "hr/migrations", "postgres"
}
