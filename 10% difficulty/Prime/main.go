package main

import "fmt"

func isPrime(n int) bool {
	if n<= 1 {
		return false
	}

	for i := 2; i*i <=  n; i++{
		if n%i ==0{
			return false
		}
	}
	return true
}

// Checks if a given number is prime and prints the result to the console.
func main() {
	var num int

	fmt.Printf("Check Number:")
	fmt.Scan(&num)

	if isPrime(num){
		fmt.Println(num, "is a prime number")
	}else {
		fmt.Println(num, "is not a prime number")
	}	
}