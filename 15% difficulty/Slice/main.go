package main

import "fmt"

// Finds the minimum and maximum values in a given slice of integers.
func main() {
	slice := []int{24, 87, 53, 69, 52, 32, 56, 98, 65, 43}

	min := slice[0]
	max := slice[0]

	for _, val := range slice {
		if val < min {
			min = val
		}

		if val > max {
			max = val
		}
	}

	fmt.Println("Slice:", slice)
	fmt.Printf("Minimum value: %d\n", min)
	fmt.Printf("Maximum value: %d\n", max)
}
