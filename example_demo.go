package main

import (
	"book-tracker/models"
	"book-tracker/services"
	"fmt"
	"log"
)

// Example demonstrating the book service functionality
func main() {
	// Initialize the service
	bookService := services.NewInMemoryBookService()

	fmt.Println("=== Book Tracker Service Demo ===\n")

	// Create some books
	fmt.Println("1. Creating books...")
	
	book1, err := bookService.CreateBook(models.CreateBookRequest{
		Title:  "Clean Code",
		Author: "Robert C. Martin",
		Status: models.StatusToBeRead,
	})
	if err != nil {
		log.Fatal("Failed to create book1:", err)
	}
	fmt.Printf("Created: %s by %s (Status: %s)\n", book1.Title, book1.Author, book1.Status)

	book2, err := bookService.CreateBook(models.CreateBookRequest{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		Status: models.StatusCurrentlyReading,
	})
	if err != nil {
		log.Fatal("Failed to create book2:", err)
	}
	fmt.Printf("Created: %s by %s (Status: %s)\n", book2.Title, book2.Author, book2.Status)

	book3, err := bookService.CreateBook(models.CreateBookRequest{
		Title:  "Design Patterns",
		Author: "Gang of Four",
		Status: models.StatusRead,
	})
	if err != nil {
		log.Fatal("Failed to create book3:", err)
	}
	fmt.Printf("Created: %s by %s (Status: %s)\n", book3.Title, book3.Author, book3.Status)

	// Get all books
	fmt.Println("\n2. Getting all books...")
	allBooks, err := bookService.GetAllBooks()
	if err != nil {
		log.Fatal("Failed to get all books:", err)
	}
	fmt.Printf("Total books: %d\n", len(allBooks))
	for _, book := range allBooks {
		fmt.Printf("- %s by %s (Status: %s)\n", book.Title, book.Author, book.Status)
	}

	// Get books by status
	fmt.Println("\n3. Getting books currently being read...")
	currentlyReading, err := bookService.GetBooksByStatus(models.StatusCurrentlyReading)
	if err != nil {
		log.Fatal("Failed to get currently reading books:", err)
	}
	fmt.Printf("Currently reading (%d books):\n", len(currentlyReading))
	for _, book := range currentlyReading {
		fmt.Printf("- %s by %s\n", book.Title, book.Author)
	}

	// Update a book status
	fmt.Println("\n4. Updating book status...")
	statusRead := models.StatusRead
	updatedBook, err := bookService.UpdateBook(book1.ID, models.UpdateBookRequest{
		Status: &statusRead,
	})
	if err != nil {
		log.Fatal("Failed to update book:", err)
	}
	fmt.Printf("Updated: %s status changed to %s\n", updatedBook.Title, updatedBook.Status)

	// Try to create a duplicate book
	fmt.Println("\n5. Testing duplicate book prevention...")
	_, err = bookService.CreateBook(models.CreateBookRequest{
		Title:  "Clean Code",
		Author: "Robert C. Martin",
		Status: models.StatusToBeRead,
	})
	if err != nil {
		fmt.Printf("Expected error: %s\n", err.Error())
	}

	// Get final state
	fmt.Println("\n6. Final library state...")
	finalBooks, err := bookService.GetAllBooks()
	if err != nil {
		log.Fatal("Failed to get final books:", err)
	}
	for _, book := range finalBooks {
		fmt.Printf("- %s by %s (Status: %s)\n", book.Title, book.Author, book.Status)
	}

	fmt.Println("\n=== Demo Complete ===")
}