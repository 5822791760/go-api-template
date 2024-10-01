//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Post = newPostTable("public", "post", "")

type postTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnInteger
	Title     postgres.ColumnString
	Content   postgres.ColumnString
	CreatedAt postgres.ColumnTimestamp
	CreatedBy postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PostTable struct {
	postTable

	EXCLUDED postTable
}

// AS creates new PostTable with assigned alias
func (a PostTable) AS(alias string) *PostTable {
	return newPostTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PostTable with assigned schema name
func (a PostTable) FromSchema(schemaName string) *PostTable {
	return newPostTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PostTable with assigned table prefix
func (a PostTable) WithPrefix(prefix string) *PostTable {
	return newPostTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PostTable with assigned table suffix
func (a PostTable) WithSuffix(suffix string) *PostTable {
	return newPostTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPostTable(schemaName, tableName, alias string) *PostTable {
	return &PostTable{
		postTable: newPostTableImpl(schemaName, tableName, alias),
		EXCLUDED:  newPostTableImpl("", "excluded", ""),
	}
}

func newPostTableImpl(schemaName, tableName, alias string) postTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		TitleColumn     = postgres.StringColumn("title")
		ContentColumn   = postgres.StringColumn("content")
		CreatedAtColumn = postgres.TimestampColumn("created_at")
		CreatedByColumn = postgres.IntegerColumn("created_by")
		allColumns      = postgres.ColumnList{IDColumn, TitleColumn, ContentColumn, CreatedAtColumn, CreatedByColumn}
		mutableColumns  = postgres.ColumnList{TitleColumn, ContentColumn, CreatedAtColumn, CreatedByColumn}
	)

	return postTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Title:     TitleColumn,
		Content:   ContentColumn,
		CreatedAt: CreatedAtColumn,
		CreatedBy: CreatedByColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
