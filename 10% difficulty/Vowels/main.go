package main

import (
	"fmt"
	"strings"
)


func CountVowels(s string) int {
	Vowels := "AEIOUaeiou"
	count := 0

	for _,char := range s {
		if strings.ContainsRune(Vowels,char){
			count ++
		}
	}
	return count
}



// Counts the number of vowels in a given string.
func main() {
	var str string

	fmt.Printf("Enter a String: ")
	fmt.Scan(&str)

	vowelscounts := CountVowels(str)

	fmt.Println("Number of vowels in the string:", vowelscounts)
	
}

