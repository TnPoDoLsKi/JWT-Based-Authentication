package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func NewDatabase() (*sqlx.DB, error) {
	var err error
	DB, err = sqlx.Connect("mysql", "root:root@/go")

	return DB, err
}
