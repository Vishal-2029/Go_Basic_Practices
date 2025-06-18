package main

import "fmt"

//! Swap two numbers using pointers.

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

//! Modify an array element using pointer.
func main() {
	x := 10
	y := 20
	arr := []int{10,20,30,40}

	fmt.Println(arr)
	ptr := &arr[3]
	*ptr = 400
	fmt.Println("New Array:", arr)

	//* Swap two number
	fmt.Println("x=", x, "y=", y)
	swap(&x, &y)
	fmt.Println("x=", x, "y=", y)
}
