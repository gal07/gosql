package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar_go")
	if err != nil {
		return nil, err
	}
	return db, nil
}
