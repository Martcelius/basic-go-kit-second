package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDBconn() *sql.DB {
	db, err := sql.Open("mysql", "root:<password>@tcp(127.0.0.1:3306)/<your database>")
	if err != nil {
		panic(err)
	}

	return db
}
