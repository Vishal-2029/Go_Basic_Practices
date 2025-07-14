package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	var a, b int
	fmt.Printf("Enter two integers (a b):\n")
	fmt.Scan(&a,&b)
	
	result, err := divide(a, b)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Division successful, result:", result)
	}
	
}
