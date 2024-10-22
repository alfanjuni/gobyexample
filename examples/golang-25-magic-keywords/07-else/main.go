package main

import "fmt"

func main() {
	x := 10
	y := 15

	// Check if x is greater than y.
	if x > y {
		fmt.Println("10 is greater than 15")
	} else {
		// Else block is executed if the condition is false.
		fmt.Println("10 is less than 15")
	}
}