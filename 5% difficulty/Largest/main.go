package main

import "fmt"

// Compares two user-input integers and prints the larger of the two.
func main() {
	var a, b int

	fmt.Printf("Enter A:")
	fmt.Scan(&a)
	fmt.Printf("Enter B:")
	fmt.Scan(&b)

	if a < b {
		fmt.Println("The largest Number is:", b)
	} else {
		fmt.Println("The largest Number is:",a)
	}
}