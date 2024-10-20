package main

import "fmt"

// Every function starts with the 'func' keyword.
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(5, 7)
    fmt.Printf("The sum of 5 and 7 is: %d\n", result)
}
