# Range in Go

The `range` keyword in Go is used to iterate over elements in various data structures. It's a powerful feature that simplifies looping over collections.

## Key Concepts

### Basic Range Syntax

```go
for index, value := range collection {
    // Use index and value
}
```

The variables returned by range depend on the type of collection:

### Arrays and Slices

When ranging over arrays or slices, range returns the index and value:

```go
fruits := []string{"apple", "banana", "orange"}
for i, fruit := range fruits {
    fmt.Printf("Index: %d, Value: %s\n", i, fruit)
}
```

Output:
```
Index: 0, Value: apple
Index: 1, Value: banana
Index: 2, Value: orange
```

### Strings

When ranging over strings, range returns the index and Unicode code point (rune):

```go
for i, char := range "Hello, 世界" {
    fmt.Printf("Index: %d, Character: %c, Unicode: %U\n", i, char, char)
}
```

Note: Indexes may not increment by 1 for Unicode characters that use multiple bytes.

### Maps

When ranging over maps, range returns the key and value:

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
}
for name, age := range ages {
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

Note: Map iteration order is not guaranteed.

### Channels

When ranging over channels, range returns only the value:

```go
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
}()

for num := range ch {
    fmt.Println(num)
}
```

Range will exit after the channel is closed.

### Using Only One Variable

You can omit variables you don't need:

```go
// Using only the index
for i := range collection {
    // Use only i
}

// Using only the value (with blank identifier)
for _, value := range collection {
    // Use only value
}
```

### Range and Reference Types

For slices and maps, range iterates over a copy of the index and values. Modifying the value variable doesn't affect the original collection:

```go
numbers := []int{1, 2, 3}
for _, num := range numbers {
    num *= 2 // This doesn't modify the original slice
}
// numbers is still [1, 2, 3]

// To modify the original, use the index:
for i := range numbers {
    numbers[i] *= 2
}
// numbers is now [2, 4, 6]
```

## Exercise

Fix the failing test in `range_test.go` by:

1. Creating a file `range.go`
2. Implementing the functions required by the test:
   - `SumSlice(nums []int) int` - Sum all elements in a slice
   - `ContainsValue(slice []string, value string) bool` - Check if a slice contains a value
   - `CountVowels(s string) int` - Count the vowels in a string
   - `MergeMaps(map1, map2 map[string]int) map[string]int` - Merge two maps
   - `FilterSlice(nums []int, filterFn func(int) bool) []int` - Filter a slice based on a function

## Running the Test

```bash
go test
```

## Bonus Challenge

- Write a function to find the most frequent character in a string
- Implement a function to reverse a map (swap keys and values) 