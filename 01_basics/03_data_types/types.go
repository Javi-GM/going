package types

import (
	"strconv"
)

func IntegerTypes() [4]int {
	return [4]int{1, 2, 3, 4}
}

func FloatingTypes() (float32, float64) {
	// Return both a float32 and float64 value to show precision differences
	return 1.1, 1.2
}

func BooleanType() (bool, bool) {
	return true, false
}

func StringType() string {
	return "Go is a great language!" 
}

func TypeConversion(num int) (int, float64, string) {
	intValue := num
	floatValue := float64(num)
	stringValue := strconv.Itoa(num)

	return intValue, floatValue, stringValue
}