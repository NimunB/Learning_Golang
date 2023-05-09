package main

import (
    "fmt"
		"log" // for logging
    "example.com/greetings" // importing greeting package
)

func main() { // Declare a main package. In Go, code executed as an application must be in a main package
		// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ") // Configure the log package to print the command name ("greetings: ") at the start of its log messages
    log.SetFlags(0) // without a time stamp or source file information.

    // ----Request greeting message------

		// Check for error
		message, err := greetings.Hello("") // Assign both of the Hello return values, including the error, to variables.

    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}