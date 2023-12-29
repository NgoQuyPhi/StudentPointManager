package main

import (
	db "Webapp/StudentPointManager/DB"
)

func main() {
	user, pass := "root", "NgocBich1609@@"

	db.CreateDB(user, pass)
	DB := db.CreateTable(user, pass)

}
