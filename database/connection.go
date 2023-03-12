package database

import (
	"database/sql"
	"fmt"
)

var db *sql.DB
var err error

func ConnectToDB() *sql.DB {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/go") 
	if err != nil {
		fmt.Println("Server could connect with database")
	}
	return db
}
