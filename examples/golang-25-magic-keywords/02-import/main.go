package main

import (
    "fmt"
    "time"
)

func main() {
    currentTime := time.Now()
    fmt.Println("Welcome to the import keyword example!")
    fmt.Printf("The current time is: %s\n", currentTime.Format("2006-01-02 15:04:05"))
}