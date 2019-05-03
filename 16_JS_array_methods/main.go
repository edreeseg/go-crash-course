package main

import "fmt"

func doubler(n int) int {
	return n * 2
}

func sum(a, b int) int {
	return a + b
}

func difference(a, b int) int {
	return a - b
}

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 != 0
}

// Only work for ints

func forEach(arr []int, cb func(n int) int) {
	for i, value := range arr {
		arr[i] = cb(value)
	}
}

func arrMap(arr []int, cb func(n int) int) []int {
	var newArr []int
	for _, value := range arr {
		newVal := cb(value)
		newArr = append(newArr, newVal)
	}
	return newArr
}

func filter(arr []int, cb func(n int) bool) []int {
	var filteredArr []int
	for _, value := range arr {
		if cb(value) {
			filteredArr = append(filteredArr, value)
		}
	}
	return filteredArr
}

func reduce(arr []int, cb func(a, b int) int) int {
	accumulator := arr[0]
	rest := arr[1:]
	for _, value := range rest {
		accumulator = cb(accumulator, value)
	}
	return accumulator
}

func indexOf(arr []int, n int) int {
	for index, value := range arr {
		if value == n {
			return index
		}
	}
	return -1
}

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func reverseString(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sum := reduce(arr, sum)
	difference := reduce(arr, difference)
	filtered := filter(arr, isOdd)
	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(filtered)
	fmt.Println(indexOf(arr, 99))
	reverseArr := []int{5, 4, 3, 2, 1}
	fmt.Println(reverse(reverseArr))
	fmt.Println(reverseString("Hello World"))
}
