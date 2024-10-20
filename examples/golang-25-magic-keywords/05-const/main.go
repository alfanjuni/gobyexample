package main

import "fmt"

// Declaring constants
const Pi float64 = 3.1415
const WelcomeMessage string = "Hello, Go developer!"

func main() {
    fmt.Println("Welcome to the const keyword example!")

    // Using constants
    fmt.Printf("Value of Pi: %f\n", Pi)
    fmt.Printf("Message: %s\n", WelcomeMessage)

    // Uncommenting the next line will cause an error because constants cannot be modified.
    // Pi = 3.14
}