package main

import "fmt"

// Swaps the values of two integers input by the user and prints the values before and after the swap.
func main() {
	var a, b int
	fmt.Println("Enter two numbers to swap:")
	fmt.Scan(&a, &b)

	fmt.Println("Before Swap: a =", a, ", b =", b)
	a, b = b, a
	
	fmt.Println("After Swap: a =", a, ", b =", b)
}