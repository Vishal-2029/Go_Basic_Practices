package main

import "fmt"

//! Determine if an array reads the same backward as forward.
//? Check if an Array is Palindromic

func IsPalindromic(arr []int) bool{
	n := len(arr)
	for i := 0; i < n/2; i++{
		if arr[i] != arr[n-1-i]{
			return false
		}
	}
	return true
}

//! Write a function to reverse the contents of an array in-place.
//? Reverse an Array

func ReversArr(str []string)[]string{
	n := len(str)
	reversed := make([]string, n)
	for i := 0; i < n; i++ {
		reversed[i] = str[n-1-i]
	}
	return reversed
}

func main(){
	arr := []int{1,2,3,2,1}
	str := []string{"A", "p", "p", "l", "e"}
	result := IsPalindromic(arr)
	revers := ReversArr(str)

	fmt.Println("Array: ", result)
	fmt.Println("Reverse Array: ",revers)

}