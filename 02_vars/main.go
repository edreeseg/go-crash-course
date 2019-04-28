package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var name string = "Edward"
	// Will infer type
	var lastName = "Reeseg"
	const isCool = true

	age := 27 // Shorthand method
	size := 1.3

	email, password := "edward@email.com", "test123"

	fmt.Println(name, lastName, age, isCool, email, password)
	fmt.Printf("%T\n", size)
}
