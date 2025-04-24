# If/Else Statements in Go

Conditional statements allow your program to make decisions based on conditions. Go offers `if`, `else if`, and `else` statements for this purpose.

## Key Concepts

### Basic Syntax

```go
if condition {
    // Code executed if condition is true
} else if anotherCondition {
    // Code executed if anotherCondition is true
} else {
    // Code executed if all conditions are false
}
```

### Condition Evaluation

- Go conditions don't need parentheses, but the braces are required
- The condition must evaluate to a boolean value
- Common comparison operators: `==`, `!=`, `<`, `>`, `<=`, `>=`
- Logical operators: `&&` (AND), `||` (OR), `!` (NOT)

### Short Statement

Go allows a short statement before the condition:

```go
if x := calculateValue(); x < 10 {
    // Use x
} else {
    // x is also available here
}
// x is NOT available here
```

This is useful for:
- Declaring variables used only in the if-else block
- Performing an operation and checking its result in one line

### Scope of Variables

Variables declared in the short statement are only in scope within the if-else blocks.

### No Ternary Operator

Go does not have a ternary conditional operator (`condition ? value1 : value2`).

## Exercise

Fix the failing test in `if_else_test.go` by:

1. Creating a file `if_else.go`
2. Implementing the functions required by the test:
   - `CheckNumber(num int) string` - Return "positive", "negative", or "zero" based on the input
   - `GradeScore(score int) string` - Return a letter grade based on a numeric score
   - `CheckEvenOdd(num int) string` - Return "even" or "odd" based on the input
   - `GetDiscount(amount float64, hasLoyaltyCard bool) float64` - Calculate discount based on amount and loyalty status

## Running the Test

```bash
go test
```

## Bonus Challenge

- Implement a function that determines if a year is a leap year using nested if statements
- Create a function that validates if a string is a valid email address using logical operators 