package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDB(user, pass string) {

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/", user, pass)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Exec("CREATE DATABASE IF NOT EXISTS studentpointmanager")
}
