package helpers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	return sql.Open("mysql", ConnectionString)
}
