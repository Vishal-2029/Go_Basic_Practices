package main

import "fmt"

func main() {

	var num int

	fmt.Printf("Check Number:")
	fmt.Scan(&num)

	if num >0{
		fmt.Println(num,"is positive")
	}else if num < 0 {
		fmt.Println(num,"is Negative")
	} else {
		fmt.Println(num,"is Zero")
	}
}