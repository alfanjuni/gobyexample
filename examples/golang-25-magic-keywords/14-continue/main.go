package main

import "fmt"

func main() {
	// A loop with a 'continue' statement.
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue // Skip the iteration when i equals 3.
		}
		fmt.Println(i)
	}
}