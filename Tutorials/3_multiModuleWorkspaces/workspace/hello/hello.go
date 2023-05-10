package main

import (
    "fmt"
    "golang.org/x/example/stringutil" // resolves the golang.org/x/example import using the go.work file.
)

func main() {
    fmt.Println(stringutil.ToUpper("Hello")) // use toUpper in another module in workspace. eliminates need for multiple replaces in go.mod
}