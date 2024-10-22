package main

import "fmt"

func main() {
	// A slice of numbers.
	numbers := []int{1, 2, 3, 4, 5}

	// Using 'range' to iterate over the slice.
	for i, num := range numbers {
		fmt.Printf("index: %d, value: %d\n", i, num)
	}
}