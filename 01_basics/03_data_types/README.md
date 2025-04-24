# Data Types in Go

Go has various built-in data types to handle different kinds of values. Understanding these types is essential for writing effective Go programs.

## Key Concepts

### Basic Types

1. **Boolean Type**
   - `bool`: Represents true or false values
   - Example: `var isActive bool = true`

2. **Numeric Types**
   - **Integer Types**
     - `int`, `int8`, `int16`, `int32`, `int64`: Signed integers of varying sizes
     - `uint`, `uint8`, `uint16`, `uint32`, `uint64`: Unsigned integers of varying sizes
     - `byte`: Alias for uint8
     - `rune`: Alias for int32, represents a Unicode code point
   - **Floating-Point Types**
     - `float32`, `float64`: Floating-point numbers
   - **Complex Types**
     - `complex64`, `complex128`: Complex numbers

3. **String Type**
   - `string`: Sequence of bytes (often representing text)
   - Strings are immutable in Go
   - Example: `var name string = "Gopher"`

### Type Conversion

Go requires explicit type conversion (no automatic conversion between types):

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

### Type Inference

Go can infer the type when you use the short declaration syntax:

```go
x := 42       // int
y := 3.14     // float64
z := "hello"  // string
b := true     // bool
```

### Zero Values

Every type has a default "zero value":
- Numeric types: `0`
- Boolean type: `false`
- String type: `""`
- Pointer, function, interface, slice, channel, map: `nil`

### Type Sizes and Ranges

Type sizes vary:
- `int`, `uint`: 32 or 64 bits depending on the system
- `int8`: -128 to 127
- `uint8`: 0 to 255
- `int16`: -32768 to 32767
- `uint16`: 0 to 65535
- ... and so on

## Exercise

Fix the failing test in `types_test.go` by:

1. Creating a file `types.go`
2. Implementing the functions required by the test:
   - `IntegerTypes()` - Return a slice containing different integer types
   - `FloatingTypes()` - Return examples of float32 and float64
   - `BooleanType()` - Return true and false values
   - `StringType()` - Return a string value
   - `TypeConversion()` - Convert values between different types

## Running the Test

```bash
go test
```

## Bonus Challenge

- Create a function that demonstrates handling Unicode characters in strings using runes
- Explore the limits of different number types by creating a function that demonstrates overflows 