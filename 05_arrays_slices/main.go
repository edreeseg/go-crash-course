package main

import "fmt"

func main() {
	// Arrays
	// Arrays must be a fixed length, you must name the types
	// var fruitArr [2] string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign

	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)

	// Slices
	// A slice is basically an array that doesn't have a fixed length

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	fruitLength := len(fruitSlice)

	fmt.Println(fruitSlice)
	fmt.Printf("Length is %d\n", fruitLength)
	fmt.Println(fruitSlice[1:3])
}
