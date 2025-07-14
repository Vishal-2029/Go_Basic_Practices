package main

import "fmt"

// greet returns a simple greeting message as a string.
func greet() string {
	return "Hello,Go!"
}

// Prints a welcome message to the console, typically used as an entry point for a programming tutorial or guide.
func Greet() {
	fmt.Println("Welcome to Programming!!")
}

// Entry point of the program, prints a greeting message and invokes the Greet function.
func main() {
	fmt.Println(greet())
	Greet()
}


