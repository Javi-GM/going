# Defer in Go

The `defer` statement in Go schedules a function call to be executed just before the surrounding function returns. It's a powerful feature for cleanup operations and ensuring certain code is executed regardless of how a function exits.

## Key Concepts

### Basic Defer Usage

```go
func example() {
    defer fmt.Println("This will print last")
    fmt.Println("This will print first")
}
```

When executed, this outputs:
```
This will print first
This will print last
```

### Multiple Defers: LIFO Order

Multiple deferred calls are executed in Last-In-First-Out (LIFO) order:

```go
func example() {
    defer fmt.Println("This will print third")
    defer fmt.Println("This will print second")
    defer fmt.Println("This will print first")
    fmt.Println("This will print before all deferred calls")
}
```

When executed, this outputs:
```
This will print before all deferred calls
This will print first
This will print second
This will print third
```

### Evaluation of Arguments

Arguments to deferred functions are evaluated when the `defer` statement is executed, not when the function is called:

```go
func example() {
    x := 1
    defer fmt.Println("x =", x) // Will print "x = 1", not "x = 2"
    x = 2
    fmt.Println("Current x =", x) // Will print "Current x = 2"
}
```

### Common Use Cases

#### Resource Cleanup

Defer is often used to ensure resources are properly closed or cleaned up:

```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // Will be called when the function returns
    
    // Rest of the function...
    return nil
}
```

#### Mutex Unlocking

Defer helps ensure mutexes are unlocked even if a function panics:

```go
func updateSharedResource() {
    mu.Lock()
    defer mu.Unlock() // Will be called even if the function panics
    
    // Update shared resource...
}
```

#### Timing Functions

Defer is useful for timing how long a function takes to execute:

```go
func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s\n", name, elapsed)
}

func expensiveOperation() {
    defer timeTrack(time.Now(), "expensiveOperation")
    
    // Perform the operation...
}
```

### Defer with Anonymous Functions

You can use anonymous functions with defer for more complex operations:

```go
func example() {
    x := 10
    
    defer func() {
        fmt.Println("x =", x) // Captures x by reference
    }()
    
    x = 20 // This change will be reflected in the deferred function
}
```

### Defer and Return Values

Deferred functions can modify named return values before they're returned:

```go
func example() (result int) {
    defer func() {
        result *= 2 // Doubles the return value
    }()
    
    result = 5
    return // Returns 10, not 5
}
```

## Exercise

Fix the failing test in `defer_test.go` by:

1. Creating a file `defer.go`
2. Implementing the functions required by the test:
   - `ProcessFile(filename string) (string, error)` - Reads from a file and returns its content
   - `RunTransaction(operations []func() error) error` - Executes a series of operations with proper cleanup
   - `TrackTime(operation func())` - Measures execution time of a function
   - `FilterPanic(handler func(interface{}), riskyFunc func())` - Executes a function and handles any panic
   - `ModifyReturnValue() int` - Demonstrates modifying a return value with defer

## Running the Test

```bash
go test
```

## Bonus Challenge

- Implement a function that uses defer to properly close and unlock multiple resources
- Create a function that uses defer to log entry and exit from a function, including any errors 