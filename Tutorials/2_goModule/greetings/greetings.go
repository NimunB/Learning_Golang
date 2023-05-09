package greetings // Declare a greetings package to collect related functions.

import (
	"fmt"
	"errors"
	"math/rand"
  "time"
)
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) { // Change the function so that it returns two values: a string and an error. Your caller will check the second value to see if an error occurred. 
		// If no name was given, return an error with a message.
    if name == "" {
			return "", errors.New("empty name")
		}
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name) // shortcut for declaring and initializing a variable in one line
    return message, nil // Add nil (meaning no error) as a second value in the successful return. That way, the caller can see that the function succeeded.
}
// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano()) // using time to set the seed of rand
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats (datatype is string).
	formats := []string{
			"Hi, %v. Welcome!",
			"Great to see you, %v!",
			"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}