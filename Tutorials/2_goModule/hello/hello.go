package main

import (
    "fmt"

    "example.com/greetings" // importing greeting package
)

func main() { // Declare a main package. In Go, code executed as an application must be in a main package
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}