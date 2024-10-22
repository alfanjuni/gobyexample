package main

import "fmt"

func main() {
	// A simple loop with a 'break' statement.
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // Exit the loop when i equals 5.
		}
		fmt.Println(i)
	}
}
// Output:
// 0
// 1
// 2
// 3
// 4
