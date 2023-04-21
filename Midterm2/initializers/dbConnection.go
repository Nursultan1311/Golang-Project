package initializers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:root@tcp(127.0.0.1:8889)/go"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
