package main

import (
	"New_Book_Management/API"
	"New_Book_Management/Database"
	"log"
)

func main() {
	Database.ConnectDatabase()
	var apiBookManager = &API.BookManagementAPI{DB: Database.Db}
	router := API.SetupRouter(apiBookManager)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
