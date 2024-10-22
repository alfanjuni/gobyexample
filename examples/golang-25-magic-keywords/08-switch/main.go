package main

import "fmt"

func main() {
	number := 2

	// Use 'switch' to handle multiple conditions.
	switch number {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown number")
	}
}
