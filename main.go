package main

import (
	"golang-todolist/src/config"
	"golang-todolist/src/routes"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	// run all routes
	routes.Routes()
}
