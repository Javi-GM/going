package types

import (
	"testing"
)

func TestIntegerTypes(t *testing.T) {
	result := IntegerTypes()
	
	if len(result) != 4 {
		t.Errorf("Expected 4 integer values, got %d", len(result))
	}
	
	// Check that each value is of the correct type and range
	if result[0] < -128 || result[0] > 127 {
		t.Errorf("First value %d outside range of int8", result[0])
	}
	
	if result[1] < 0 || result[1] > 255 {
		t.Errorf("Second value %d outside range of uint8", result[1])
	}
	
	if result[2] < -32768 || result[2] > 32767 {
		t.Errorf("Third value %d outside range of int16", result[2])
	}
	
	if result[3] <= 0 {
		t.Errorf("Fourth value %d should be a positive int32", result[3])
	}
}

func TestFloatingTypes(t *testing.T) {
	float32Val, float64Val := FloatingTypes()
	
	// Test that float32 has less precision than float64
	const float32Test = 0.1234567890123456789
	if float32(float32Test) == float32Test {
		t.Error("float32 should not have full precision of a long decimal")
	}
	
	// Simple range test
	if float32Val <= 0 || float64Val <= 0 {
		t.Error("Floating values should be positive")
	}
}

func TestBooleanType(t *testing.T) {
	trueVal, falseVal := BooleanType()
	
	if !trueVal {
		t.Error("Expected true value")
	}
	
	if falseVal {
		t.Error("Expected false value")
	}
}

func TestStringType(t *testing.T) {
	result := StringType()
	
	if result == "" {
		t.Error("Expected non-empty string")
	}
	
	if len(result) < 5 {
		t.Errorf("Expected string with at least 5 characters, got %d", len(result))
	}
}

func TestTypeConversion(t *testing.T) {
	intVal, floatVal, stringVal := TypeConversion(42)
	
	if intVal != 42 {
		t.Errorf("Expected int value 42, got %d", intVal)
	}
	
	if floatVal != 42.0 {
		t.Errorf("Expected float value 42.0, got %f", floatVal)
	}
	
	if stringVal != "42" {
		t.Errorf("Expected string value \"42\", got %q", stringVal)
	}
}

// Bonus challenge test
// func TestUnicodeHandling(t *testing.T) {
// 	input := "Hello, 世界"
// 	chars, length := UnicodeHandling(input)
// 	
// 	// The length in runes (Unicode characters) should be 9
// 	if length != 9 {
// 		t.Errorf("Expected rune count 9, got %d", length)
// 	}
// 	
// 	// The length in bytes would be more than 9
// 	if len(input) <= length {
// 		t.Error("Expected byte length to be longer than rune length")
// 	}
// 	
// 	// Check if we got the right runes
// 	if len(chars) != length {
// 		t.Errorf("Expected %d runes, got %d", length, len(chars))
// 	}
// } 