package types

import (
	"strconv"
)

// IntegerTypes returns examples of different integer types
func IntegerTypes() []int {
	var int8Val int8 = 127        // max int8 value
	var uint8Val uint8 = 255      // max uint8 value
	var int16Val int16 = 32767    // max int16 value
	var int32Val int32 = 2147483647 // max int32 value
	
	// Convert all to int for returning in a slice
	return []int{int(int8Val), int(uint8Val), int(int16Val), int(int32Val)}
}

// FloatingTypes returns examples of float32 and float64
func FloatingTypes() (float32, float64) {
	var float32Val float32 = 3.14159
	var float64Val float64 = 3.141592653589793
	
	return float32Val, float64Val
}

// BooleanType returns examples of boolean values
func BooleanType() (bool, bool) {
	return true, false
}

// StringType returns an example string
func StringType() string {
	return "Go is a statically typed language"
}

// TypeConversion demonstrates converting between types
func TypeConversion(num int) (int, float64, string) {
	intVal := num
	floatVal := float64(num)
	stringVal := strconv.Itoa(num)
	
	return intVal, floatVal, stringVal
}

// UnicodeHandling demonstrates working with Unicode characters
func UnicodeHandling(s string) ([]rune, int) {
	// Convert string to rune slice
	runes := []rune(s)
	
	// Return the runes and the count of runes
	return runes, len(runes)
} 