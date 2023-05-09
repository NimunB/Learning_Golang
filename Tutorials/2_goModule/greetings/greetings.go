package greetings // Declare a greetings package to collect related functions.

import (
	"fmt"
	"errors"
)
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) { // Change the function so that it returns two values: a string and an error. Your caller will check the second value to see if an error occurred. 
		// If no name was given, return an error with a message.
    if name == "" {
			return "", errors.New("empty name")
		}
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name) // shortcut for declaring and initializing a variable in one line
    return message, nil // Add nil (meaning no error) as a second value in the successful return. That way, the caller can see that the function succeeded.
}