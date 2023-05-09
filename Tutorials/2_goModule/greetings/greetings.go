package greetings // Declare a greetings package to collect related functions.

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name) // shortcut for declaring and initializing a variable in one line
    return message
}