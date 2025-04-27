# Break and Continue in Go

The `break` and `continue` statements are used to control the flow of loops in Go. They allow you to exit a loop early or skip to the next iteration based on certain conditions.

## Key Concepts

### Break Statement

The `break` statement terminates the innermost enclosing loop:

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Exit the loop when i equals 5
    }
    fmt.Println(i) // Prints 0, 1, 2, 3, 4 only
}
```

### Continue Statement

The `continue` statement skips the rest of the current iteration and proceeds to the next iteration:

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // Skip even numbers
    }
    fmt.Println(i) // Prints 1, 3, 5, 7, 9 only
}
```

### Using Break and Continue in Nested Loops

By default, `break` and `continue` apply only to the innermost loop:

```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 1 {
            break // Only breaks out of the inner loop
        }
        fmt.Printf("(%d, %d) ", i, j) // Prints (0, 0) (1, 0) (2, 0)
    }
}
```

### Labeled Statements

Labels allow `break` and `continue` to target specific outer loops:

```go
outerLoop:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i*j > 2 {
            break outerLoop // Exits both loops
        }
        fmt.Printf("(%d, %d) ", i, j)
    }
}
```

Similarly, you can use labeled `continue`:

```go
outerLoop:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j > i {
            continue outerLoop // Skip to the next iteration of the outer loop
        }
        fmt.Printf("(%d, %d) ", i, j)
    }
}
```

### Using Break in Switch and Select Statements

The `break` statement is also used in `switch` and `select` statements to exit early:

```go
switch n := 2; n {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
    break // This break is actually redundant in Go switches
case 3:
    fmt.Println("Three")
}
```

Note: Go's `switch` statements don't fall through by default, so `break` is typically unnecessary.

### Using Break with Switch Inside Loops

When a `switch` is inside a loop, `break` will only exit the `switch`, not the loop:

```go
for i := 0; i < 5; i++ {
    switch i {
    case 3:
        break // This only breaks out of the switch, not the loop
    }
    fmt.Println(i) // Prints 0, 1, 2, 3, 4
}
```

To break from both the switch and the loop, you need a labeled statement:

```go
loop:
for i := 0; i < 5; i++ {
    switch i {
    case 3:
        break loop // Exits the loop entirely
    }
    fmt.Println(i) // Prints 0, 1, 2 only
}
```

## Exercise

Fix the failing test in `break_continue_test.go` by:

1. Creating a file `break_continue.go`
2. Implementing the functions required by the test:
   - `FindFirstEven(nums []int) (int, bool)` - Find the first even number in a slice
   - `SumUntilNegative(nums []int) int` - Sum numbers until a negative value is encountered
   - `SkipMultiplesOf(n, max int) []int` - Generate a slice of numbers up to max, skipping multiples of n
   - `FindSubstring(s, sub string) int` - Find the starting index of a substring
   - `PrintTriangle(size int) []string` - Create a triangle pattern of specified size

## Running the Test

```bash
go test
```

## Bonus Challenge

- Implement a function that finds the first pair of consecutive duplicate values in a slice
- Write a function that generates all prime numbers up to n using the Sieve of Eratosthenes algorithm 