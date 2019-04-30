package main

import (
	"fmt"
	"math/rand"
)

func rand1() int {
	rand.Seed(85)
	return rand.Intn(10) // 7
}

func rand2() int {
	rand.Seed(29)
	return rand.Intn(10) // 3
}

func main() {
	fmt.Println(rand1())
	fmt.Println(rand2())
}
