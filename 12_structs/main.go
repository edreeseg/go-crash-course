package main

import (
	"fmt"
	"strconv"
)

// Define person struct

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Two types of methods, value receivers and pointer receivers
// Value receivers do calculations and work with values but don't change anything
// Pointer receivers can make changes

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {
		p.lastName = spouseLastName
	}
}

func main() {
	// Initialize person using struct

	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	// Alternative
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}
	fmt.Println(person1)
	fmt.Println(person1.firstName)
	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

	fmt.Println(person2.greet())
	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
}
