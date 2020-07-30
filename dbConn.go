package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func OpenDB() *sql.DB {
	db, err := sql.open("mysql", "<USERNAME>:<PASSWORD>@tcp(127.0.0.1:3306)/<DATABASE_NAME>")

	if err != nil {
		panic(err.Error())
	}

	return db
}
