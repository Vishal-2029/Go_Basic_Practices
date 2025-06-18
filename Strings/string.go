package main

import (
	"fmt"
	"strings"
)

//! Write a function to reverse a string.

func reversestr(s string)string{
	reversed := ""

	for _, char := range s{
		reversed = string(char) + reversed
	}
	return reversed
}

//! Count the number of vowels in a string.
func vowelsstr(str string )int{
	vowels :="aeiouAEIOU"
	count := 0

	for _, char := range str{
		if strings.ContainsRune(vowels,char){
			count++
		}
	}
	return count
}


func main(){
	str := "Apple"
	
	 
	fmt.Println("String: ",str)

	reverseString := reversestr(str)
	vowelsString := vowelsstr(str)

	fmt.Println("Reverse String:",reverseString)
	fmt.Println("Vowels String:", vowelsString)
}