package database

import (
	"database/sql"
	"fmt"
)

var db *sql.DB
var err error

// In this page we connect go project with database. For this we enter name of sql user, his password then database name which you want to connect 

func ConnectToDB() *sql.DB {
	// Connecting to database.  "username:password@(127.0.0.1:8889)/databasename"
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/go") 
	if err != nil {
		fmt.Println("Server could connect with database")
	}
	return db
}
