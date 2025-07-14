package main

import "fmt"

// Calculates the product of two user-provided integers and prints the result.
func main() {
	var a, b int

	fmt.Printf("Enter A:")
	fmt.Scan(&a)
	fmt.Printf("Enter B:")
	fmt.Scan(&b)

	multiply := a * b
	fmt.Println("The Multiplication  of", a, "and", b, "is:", multiply)
}