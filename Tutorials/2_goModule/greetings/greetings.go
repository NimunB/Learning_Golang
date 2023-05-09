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
		//message := fmt.Sprint(randomFormat()) // failure test
    return message, nil // Add nil (meaning no error) as a second value in the successful return. That way, the caller can see that the function succeeded.
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) { // parameter is a slice of names rather than a single name. Also, you change one of its return types from a string to a map so you can return names mapped to greeting messages.
	// A map to associate names with messages.
	messages := make(map[string]string) // make(map[key-type]value-type) - Create a messages map to associate each of the received names (as a key) with a generated message
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names { // range returns two values: the index of the current item in the loop and a copy of the item's value. You don't need the index, so you use the Go blank identifier (an underscore) to ignore it.
			message, err := Hello(name)
			if err != nil {
					return nil, err
			}
			// In the map, associate the retrieved message with
			// the name.
			messages[name] = message
	}
	return messages, nil
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