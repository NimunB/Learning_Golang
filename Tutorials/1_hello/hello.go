package main // declare package (way to group functions) - made up of all the files in the same directory

import "fmt" // Import the popular fmt package, which contains functions for formatting text

import "rsc.io/quote" // import the rsc.io/quote package and add a call to its Go function

func main() {
    //fmt.Println("Hello, World!")
		fmt.Println(quote.Go())
}