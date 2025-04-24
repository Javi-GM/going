package variables

// DeclareInteger declares and returns an integer with value 42
func DeclareInteger() int {
	var number int = 42
	return number
}

// DeclareFloat declares and returns a float with value 3.14
func DeclareFloat() float64 {
	pi := 3.14
	return pi
}

// DeclareString declares and returns a string "Go is fun"
func DeclareString() string {
	var message = "Go is fun"
	return message
}

// DeclareBool declares and returns a boolean with value true
func DeclareBool() bool {
	isAwesome := true
	return isAwesome
}

// MultipleDeclarations declares multiple variables and returns them
func MultipleDeclarations() (int, float64, string) {
	x, y, z := 10, 15.5, "multiple"
	return x, y, z
}

// SwapVariables swaps the values of two integers without using a third variable
func SwapVariables(a, b int) (int, int) {
	// Using arithmetic operations to swap without a third variable
	a = a + b
	b = a - b
	a = a - b
	return a, b
} 