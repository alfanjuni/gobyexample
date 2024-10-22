package main

import "fmt"

func main() {
	number := 3

	// Use 'switch' with multiple 'case' conditions.
	switch number {
	case 1:
		fmt.Println("This is number 1")
	case 2:
		fmt.Println("This is number 2")
	case 3:
		fmt.Println("This is number 3")
	default:
		fmt.Println("Number not recognized")
	}
}