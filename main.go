package main

import (
	"New_Book_Management/API"
	"New_Book_Management/Database"
	"log"
)

//func main() { //For Console Using
//	Database.ConnectDatabase()
//	bookManager := &Console.BookManagementConsole{DB: Database.Db}
//
//	for {
//		fmt.Println("\nChoose an option:")
//		fmt.Println("1. Show all books")
//		fmt.Println("2. Add a new book")
//		fmt.Println("3. Delete a book")
//		fmt.Println("4. Update a book")
//		fmt.Println("4. Exit")
//
//		var choice int
//		fmt.Scan(&choice)
//
//		switch choice {
//		case 1:
//			bookManager.ShowBooks()
//		case 2:
//			bookManager.AddNewBook()
//		case 3:
//			bookManager.RemoveBook()
//		case 4:
//			bookManager.UpdateBookDetails()
//		case 5:
//			fmt.Println("Exiting...")
//			return
//		default:
//			fmt.Println("Invalid option. Please try again.")
//		}
//	}
//}

func main() { ///////// For API Usage
	Database.ConnectDatabase()
	var apiBookManager = &APIpackage.BookManagementAPI{DB: Database.Db}
	router := APIpackage.SetupRouter(apiBookManager)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
