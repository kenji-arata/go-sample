package main

import (
	"fmt"
	"go-sample/db"
	"go-sample/model"
)

func main() {
	fmt.Println("Migration script started")
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
