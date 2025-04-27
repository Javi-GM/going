# For Loops in Go

Go has only one looping construct: the `for` loop. However, it's flexible enough to handle all types of iteration.

## Key Concepts

### Basic For Loop

The most common loop syntax with initialization, condition, and post statement:

```go
for i := 0; i < 10; i++ {
    // Code to repeat 10 times
    // i will be 0, 1, 2, ..., 9
}
```

### While Loop Equivalent

Go doesn't have a `while` keyword, but you can use `for` with just a condition:

```go
i := 0
for i < 10 {
    // Code to repeat while i < 10
    i++
}
```

### Infinite Loop

You can create an infinite loop by omitting the condition:

```go
for {
    // Code that repeats forever
    // (use break to exit)
}
```

### For-Range Loop

For iterating over collections (arrays, slices, strings, maps, channels):

```go
// For arrays, slices, and strings, range provides the index and value
for index, value := range collection {
    // Use index and value
}

// For maps, range provides the key and value
for key, value := range myMap {
    // Use key and value
}

// For channels, range provides just the value
for value := range myChannel {
    // Use value
}
```

### Using Only the Index

If you only need the index, you can omit the value variable:

```go
for i, _ := range collection {
    // Use only i
}

// Shorthand
for i := range collection {
    // Use only i
}
```

### Using Only the Value

If you only need the value, you can use the blank identifier for the index:

```go
for _, value := range collection {
    // Use only value
}
```

### Break and Continue

- `break`: Exits the innermost loop immediately
- `continue`: Skips to the next iteration of the loop

### Labeled Loops

Go allows labeling loops to break or continue outer loops:

```go
outerLoop:
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if i*j > 10 {
            break outerLoop // Breaks out of the outer loop
        }
    }
}
```

## Exercise

Fix the failing test in `for_loops_test.go` by:

1. Creating a file `for_loops.go`
2. Implementing the functions required by the test:
   - `SumNumbers(n int) int` - Calculate the sum of numbers from 1 to n
   - `Factorial(n int) int` - Calculate the factorial of n
   - `ReverseString(s string) string` - Reverse a string
   - `IsPrime(n int) bool` - Check if a number is prime
   - `FindMax(nums []int) int` - Find the maximum value in a slice

## Running the Test

```bash
go test
```

## Bonus Challenge

- Implement the Fibonacci sequence up to n terms
- Create a function that counts the frequency of each character in a string 