package main

import "fmt"

// Demonstrates two ways to iterate over a slice of strings in Go, 
// showcasing the use of the range keyword for a more idiomatic approach 
// and a traditional indexed loop for explicit control.
func main() {
	arr := []string{"Go", "Python", "Rust", "Vue"}

	for _, val := range arr {
		fmt.Println(val)
	}
	// Also use this syntax  
	for i := 0; i < len(arr); i++{
		fmt.Println(arr[i])
	}
}

