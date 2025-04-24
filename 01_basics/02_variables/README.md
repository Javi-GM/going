# Variables in Go

Variables are used to store and manipulate data in a program. Go is a statically typed language, which means the type of a variable is known at compile time.

## Key Concepts

### Variable Declaration

There are several ways to declare variables in Go:

1. **Using `var` keyword with explicit type:**
   ```go
   var name string
   name = "Gopher"
   ```

2. **Using `var` with initialization:**
   ```go
   var name string = "Gopher"
   ```

3. **Using type inference (preferred for local variables):**
   ```go
   var name = "Gopher"  // Type string is inferred
   ```

4. **Short declaration (most common within functions):**
   ```go
   name := "Gopher"  // Declares and initializes in one step
   ```

### Multiple Variable Declarations

```go
var a, b, c int                      // Declares multiple variables of the same type
var x, y = 10, "hello"               // Declares variables with different types
firstName, lastName := "John", "Doe" // Short declaration for multiple variables
```

### Zero Values

Go initializes variables with "zero values" when declared without an explicit initial value:

- `0` for numeric types (`int`, `float`, etc.)
- `false` for booleans
- `""` (empty string) for strings
- `nil` for pointers, slices, maps, channels, functions

### Variable Scope

- Variables declared inside a function are local to that function
- Variables declared outside any function are **package-level** variables
- Variables declared within blocks (like if statements) are scoped to that block

### Constants

For values that won't change, use the `const` keyword:

```go
const Pi = 3.14159
const (
    StatusOK      = 200
    StatusNotFound = 404
)
```

## Exercise

Fix the failing test in `variables_test.go` by:

1. Creating a file `variables.go`
2. Implementing the functions required by the test:
   - `DeclareInteger()` - Declare and return an integer with value 42
   - `DeclareFloat()` - Declare and return a float with value 3.14
   - `DeclareString()` - Declare and return a string "Go is fun"
   - `DeclareBool()` - Declare and return a boolean with value true
   - `MultipleDeclarations()` - Declare multiple variables and return them as specified in the test

## Running the Test

```bash
go test
```

## Bonus Challenge

- Create a function that swaps the values of two variables without using a third variable
- Explore the scope of variables by creating functions that demonstrate variable shadowing 