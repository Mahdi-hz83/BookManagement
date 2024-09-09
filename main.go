package main

import (
	"New_Book_Management/API"
	"New_Book_Management/Console"
	"New_Book_Management/Database"
	"fmt"
	"log"
)

func main() {
	Database.ConnectDatabase()
	// For Console Mode:
	//runConsoleApp()

	// For API Mode:
	runAPIApp()
}

func runConsoleApp() {
	bookManager := &Console.BookManagementConsole{DB: Database.Db}

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Show all books")
		fmt.Println("2. Retrieve a book by ISBN")
		fmt.Println("3. Add a new book")
		fmt.Println("4. Delete a book")
		fmt.Println("5. Update a book")
		fmt.Println("6. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			bookManager.ShowBooks()
		case 2:
			bookManager.ShowBookByISBN()
		case 3:
			bookManager.AddNewBook()
		case 4:
			bookManager.RemoveBook()
		case 5:
			bookManager.UpdateBookDetails()
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func runAPIApp() {
	apiBookManager := &APIpackage.BookManagementAPI{DB: Database.Db}
	router := APIpackage.SetupRouter(apiBookManager)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
