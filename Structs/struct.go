package main

import "fmt"

//! Create a Book struct and store a slice of books.
type Books struct {
	Name  string
	Price float64
}

//! Write a method to increase the age of a person.

type person struct {
	Name string
	Age  int
}

func (p *person) IncreaseAge(){
	p.Age++
}

func main() {
	b := []Books{{"name", 2024}, {"name2", 2004}, {"name3", 2006}}
	fmt.Println("Books:", b)

	Person := person{"Vishal",20}
	fmt.Println("Befor Increase:",Person)

	Person.IncreaseAge()

	fmt.Println("After Increase:",Person)
}
	