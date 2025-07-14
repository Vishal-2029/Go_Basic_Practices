package main

import "fmt"

func Fibonacci(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		fmt.Print(a, "")
		next := a + b
		a = b
		b = next
	}
	return a
}

// Prints the Fibonacci series up to the nth term, where n is provided by the user.
func main() {
	var n int
	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&n)

	
	fmt.Printf("Fibonacci series (%d terms):\n", n)
	Fibonacci(n)
}
