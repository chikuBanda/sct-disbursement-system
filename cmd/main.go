package main

import (
	"sct-system/database"
	"sct-system/router"
)

func main() {
	database.StartDbConnection()
	router.StartRouting()
}
