# Hello World in Go

The traditional first program in any language is a "Hello World" program that simply outputs text to the screen.

## Key Concepts

### Go Program Structure

A Go program typically starts with a package declaration, followed by imports, and then functions and other declarations.

```go
// Package declaration
package main

// Import statements
import "fmt"

// Main function - entry point of the program
func main() {
    // Code to be executed
    fmt.Println("Hello, World!")
}
```

### The `main` Package

- Every executable Go program must have a `main` package
- The `main` package must contain a `main` function, which serves as the entry point
- Go automatically calls the `main` function when you run the program

### The `fmt` Package

- Part of Go's standard library
- Provides formatting and printing functions
- `fmt.Println()` prints a line of text followed by a newline

## Exercise

Fix the failing test in `hello_test.go` by:

1. Creating a new file named `hello.go`
2. Implementing a function called `Hello()` that returns the string "Hello, World!"
3. Making sure your code is properly formatted and structured

## Running the Test

```bash
go test
```

## Bonus Challenge

- Modify your `Hello()` function to accept a name parameter and return "Hello, [name]!"
- Update the tests to verify this functionality 