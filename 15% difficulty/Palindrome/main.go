package main

import "fmt"

func isPalindrome(s string) bool {
	ranes := []rune(s)
	n := len(ranes)

	for i := 0; i < n/2; i++ {
		if ranes[i] != ranes[n-1-i] {
			return false
		}
	}
	return true
}

// Checks if the input string is a palindrome and prints the result.
func main() {
	var input string
	fmt.Printf("Enter a string to check if it's a palindrome: ")
	fmt.Scan(&input)

	if isPalindrome(input){
		fmt.Printf("'%s' is a palindrome.\n", input)
	} else {
		fmt.Printf("'%s' is not a palindrome.\n", input)
	}
}