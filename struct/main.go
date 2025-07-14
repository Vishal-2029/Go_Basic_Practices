package main

import "fmt"

// Define the Book struct
type Book struct {
	title  string
	author string
	price  float64
}

func main() {
	// Create a Book instance
	book1 := Book{
		title:  "The Go Programming Language",
		author: "Alan A. A. Donovan",
		price:  39.99,
	}

	// Print the Book details
	fmt.Println("Book Title:", book1.title)
	fmt.Println("Author:", book1.author)
	fmt.Println("Price:", book1.price)
}
