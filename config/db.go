package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	dsn := "root:HYAaIBdLdAWQharILYAoPfXerMGyrGUT@tcp(autorack.proxy.rlwy.net:41210)/railway"
	db, err := sql.Open("mysql", dsn)
	return db, err
}
