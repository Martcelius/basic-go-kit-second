package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// GetDBconn exported function for open connection
func GetDBconn() *sql.DB {
	db, err := sql.Open("mysql", "root:<password>@tcp(127.0.0.1:3306)/gofirst")
	if err != nil {
		panic(err)
	}

	return db
}
