package main

import (
	"New_Book_Management/Console"
	"New_Book_Management/Database"
	"fmt"
)

func main() {
	Database.ConnectDatabase()
	bookManager := &Console.BookManagementConsole{DB: Database.Db}

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Show all books")
		fmt.Println("2. Add a new book")
		fmt.Println("3. Delete a book")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			bookManager.ShowBooks()
		case 2:
			bookManager.AddNewBook()
		case 3:
			bookManager.RemoveBook()
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
