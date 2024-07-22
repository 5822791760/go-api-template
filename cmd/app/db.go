package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/5822791760/go-api-template/internal/config"
	embed "github.com/5822791760/go-api-template/internal/db/postgres"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/pressly/goose/v3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.GetDBConnectionString())
	if err != nil {
		return nil, err
	}

	// Configure the connection pool
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(30 * time.Minute)

	// Ping the database to verify connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	goose.SetBaseFS(embed.MigrationFiles)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	postgres.SetQueryLogger(func(ctx context.Context, queryInfo postgres.QueryInfo) {
		// sql, args := queryInfo.Statement.Sql()
		// fmt.Printf("- SQL: %s Args: %v \n\n", sql, args)
		fmt.Printf("\n++++++++++++++++++++++++++++++++\n")
		fmt.Printf("%s \n", queryInfo.Statement.DebugSql())

		// Depending on how the statement is executed, RowsProcessed is:
		//   - Number of rows returned for Query() and QueryContext() methods
		//   - RowsAffected() for Exec() and ExecContext() methods
		//   - Always 0 for Rows() method.
		fmt.Printf("- Rows processed: %d\n", queryInfo.RowsProcessed)
		fmt.Printf("- Duration %s\n", queryInfo.Duration.String())
		fmt.Printf("- Execution error: %v\n", queryInfo.Err)

		callerFile, callerLine, callerFunction := queryInfo.Caller()
		fmt.Printf("- Caller file: %s, line: %d, function: %s\n\n", callerFile, callerLine, callerFunction)
		fmt.Printf("++++++++++++++++++++++++++++++++\n\n")
	})

	return db, nil
}
