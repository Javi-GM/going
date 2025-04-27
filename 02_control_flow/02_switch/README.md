# Switch Statements in Go

Switch statements provide a cleaner way to express multiple conditional branches compared to if-else chains. They're especially useful when you need to check a variable against several possible values.

## Key Concepts

### Basic Syntax

```go
switch expression {
case value1:
    // Code executed if expression == value1
case value2, value3:
    // Code executed if expression == value2 OR expression == value3
default:
    // Code executed if no case matches
}
```

### Features and Behavior

- Go switch statements don't "fall through" by default (unlike C/C++/Java)
- Each case automatically breaks after execution
- Multiple values can be specified in a single case with commas
- The `default` case is optional and executes when no other case matches
- The switch expression is optional (defaults to `true`)

### Fallthrough Statement

If you do want a case to continue to the next case, use the `fallthrough` keyword:

```go
switch n {
case 1:
    fmt.Println("Case 1")
    fallthrough
case 2:
    fmt.Println("Case 2") // This will execute if n == 1 OR n == 2
}
```

### Expression-less Switch

Go allows switches without an expression, which is equivalent to `switch true`:

```go
switch {
case condition1:
    // Executes if condition1 is true
case condition2:
    // Executes if condition2 is true
default:
    // Executes if all conditions are false
}
```

This is often cleaner than writing multiple if-else statements.

### Type Switch

Go also allows switching on the type of an interface value:

```go
switch v := interface{}.(type) {
case string:
    fmt.Println("v is a string:", v)
case int:
    fmt.Println("v is an int:", v)
default:
    fmt.Println("v is another type")
}
```

## Exercise

Fix the failing test in `switch_test.go` by:

1. Creating a file `switch.go`
2. Implementing the functions required by the test:
   - `DayOfWeek(day int) string` - Return the name of the day based on its number (1-7)
   - `GetSeasonInNorthernHemisphere(month string) string` - Return "Winter", "Spring", "Summer", or "Fall" based on the month
   - `CalculateGrade(score int) (string, error)` - Return a letter grade based on a score, return error for invalid scores
   - `DetectType(value interface{}) string` - Return the type of the value as a string ("string", "int", "bool", etc.)

## Running the Test

```bash
go test
```

## Bonus Challenge

- Implement a simple calculator function that takes two numbers and an operator (+, -, *, /) as a string
- Write a function that converts HTTP status codes to human-readable messages using a switch statement 