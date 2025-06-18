package main

import "fmt"

//! Given two slices, find and return the elements that exist in both.
//? Find Common Elements Between Two Slices

func commonEle(a,b []int) []int{
	seen := make(map[int]bool)
	result := []int{}

	for _,val := range a {
		seen[val] = true
	}

	for _,val := range b {
		if seen[val]{
			result = append(result,val)
		}
	}
	return result
}

//! Write a function that returns a new slice containing only even numbers.
//? Filter Even Numbers from a Slice

func EvenNum(a []int)[]int{
	var even []int

	for _, num := range a {
		if num%2 == 0{
			even = append(even,num)
		}
	}
	return even
}

func main(){
	a := []int{1,2,3,4,5,8,9,0}
	b := []int{4,5,6,7,8}

	result := commonEle(a,b)
	evenNum := EvenNum(a)

	fmt.Println("a=",a  ,"b=", b)

	fmt.Println("Common Element: ", result)
	fmt.Println("Even Number: ", evenNum)
}