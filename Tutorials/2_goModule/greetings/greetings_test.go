package greetings // Implement test functions in the same package as the code you're testing.

import (
    "testing"
    "regexp"
)

// Test functions take a pointer to the testing package's testing.T type as a parameter. 
// You use this parameter's methods for reporting and logging from your test.

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)  // sth name sth
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil { // test failed
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want) // print a message to the console and end execution
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}