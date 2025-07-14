package main

import "fmt"

// Creates a map to store student information and iterates over it to print each student's name and age.
func main() {
	student := make(map[string]int)

	student["Vishal"] = 20
	student["Jeel"] = 22
	student["Dhruv"] = 21

	fmt.Printf("Student map: %v\n", student)

	for name,age := range student{
		fmt.Printf("%s: %d\n", name, age)
	}
}