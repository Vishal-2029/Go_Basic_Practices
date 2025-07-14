package main

import "fmt"

func isfactorial(n int) int {
	fac := 1

	for i := 1; i <= n; i++ {
		fac *= i
	}

	return fac
}

func main() {
	var num int

	fmt.Printf("Enter a number to calculate its factorial: ")
	fmt.Scan(&num)

	factorial:= isfactorial(num)
	fmt.Println("Factorial of", num, "is", factorial)
}