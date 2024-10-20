package main

import "fmt"

// Declaring variables globally
var globalVar string = "I am a global variable."

func main() {
    fmt.Println("Welcome to the var keyword example!")

    // Declaring and initializing variables
    var name string = "Go Developer"
    var age int = 25

    // Declaring variables without initialization
    var uninitializedVar int

    fmt.Printf("Name: %s, Age: %d\n", name, age)
    fmt.Printf("Global variable: %s\n", globalVar)
    fmt.Printf("Uninitialized variable: %d (default value)\n", uninitializedVar)
}
