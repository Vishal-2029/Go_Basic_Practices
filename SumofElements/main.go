package main

import (
	"fmt"
	
)

func main() {
	slices := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, Val := range slices {
		sum += Val
	}

	fmt.Println("Sum of elements:", sum)
}