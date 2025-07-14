package main

import "fmt"

// Determines whether a user-inputted number is even or odd.
func main() {

	var number int
	fmt.Printf("Add Number:")
	fmt.Scan(&number)

	if number%2 ==0 {
		fmt.Println("Number is Even")
	}else {
		fmt.Println("Number is Odd")
	}
}