package database

import (
	"database/sql"
	"my-app-server/helpers"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	return sql.Open("mysql", helpers.ConnectionString)
}
