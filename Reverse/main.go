package main

import "fmt"

// Reverses a user-input string and prints the result.
func main() {
	// str := "Hello, World!"

	var str string
	fmt.Printf("Enter a String:")
	fmt.Scan(&str)
	reverse := ""

	for _, char := range str {
		reverse = string(char)+ reverse
	}

	fmt.Println(reverse)
}