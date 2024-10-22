package main

import "fmt"

func main() {
	// 'defer' postpones this function until main completes.
	defer fmt.Println("This will be printed last")

	// This will be printed first.
	fmt.Println("This will be printed first")
}