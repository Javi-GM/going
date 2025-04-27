package switch_statement

import (
	"testing"
)

func TestDayOfWeek(t *testing.T) {
	testCases := []struct {
		day      int
		expected string
	}{
		{1, "Monday"},
		{2, "Tuesday"},
		{3, "Wednesday"},
		{4, "Thursday"},
		{5, "Friday"},
		{6, "Saturday"},
		{7, "Sunday"},
		{8, "Invalid day"},
		{0, "Invalid day"},
		{-1, "Invalid day"},
	}

	for _, tc := range testCases {
		result := DayOfWeek(tc.day)
		if result != tc.expected {
			t.Errorf("DayOfWeek(%d) = %s; want %s", tc.day, result, tc.expected)
		}
	}
}

func TestGetSeasonInNorthernHemisphere(t *testing.T) {
	testCases := []struct {
		month    string
		expected string
	}{
		{"January", "Winter"},
		{"February", "Winter"},
		{"March", "Spring"},
		{"April", "Spring"},
		{"May", "Spring"},
		{"June", "Summer"},
		{"July", "Summer"},
		{"August", "Summer"},
		{"September", "Fall"},
		{"October", "Fall"},
		{"November", "Fall"},
		{"December", "Winter"},
		{"Jan", "Unknown season"},
		{"", "Unknown season"},
	}

	for _, tc := range testCases {
		result := GetSeasonInNorthernHemisphere(tc.month)
		if result != tc.expected {
			t.Errorf("GetSeasonInNorthernHemisphere(%s) = %s; want %s", tc.month, result, tc.expected)
		}
	}
}

func TestCalculateGrade(t *testing.T) {
	testCases := []struct {
		score         int
		expected      string
		expectError   bool
	}{
		{95, "A", false},
		{85, "B", false},
		{75, "C", false},
		{65, "D", false},
		{55, "F", false},
		{101, "", true},  // Invalid score, should return error
		{-1, "", true},   // Invalid score, should return error
	}

	for _, tc := range testCases {
		result, err := CalculateGrade(tc.score)

		// Check error handling
		if tc.expectError && err == nil {
			t.Errorf("CalculateGrade(%d) expected an error but got nil", tc.score)
			continue
		}

		if !tc.expectError && err != nil {
			t.Errorf("CalculateGrade(%d) returned unexpected error: %v", tc.score, err)
			continue
		}

		// Only check result when no error is expected
		if !tc.expectError && result != tc.expected {
			t.Errorf("CalculateGrade(%d) = %s; want %s", tc.score, result, tc.expected)
		}
	}
}

// func TestDetectType(t *testing.T) {
// 	testCases := []struct {
// 		value    interface{}
// 		expected string
// 	}{
// 		{"hello", "string"},
// 		{42, "int"},
// 		{3.14, "float64"},
// 		{true, "bool"},
// 		{[]int{1, 2, 3}, "slice"},
// 		{map[string]int{"a": 1}, "map"},
// 		{struct{}{}, "struct"},
// 		{nil, "nil"},
// 	}

// 	for _, tc := range testCases {
// 		result := DetectType(tc.value)
// 		if result != tc.expected {
// 			t.Errorf("DetectType(%v) = %s; want %s", tc.value, result, tc.expected)
// 		}
// 	}
// }

// // Bonus challenge tests
// func TestCalculator(t *testing.T) {
// 	testCases := []struct {
// 		a           float64
// 		b           float64
// 		operator    string
// 		expected    float64
// 		expectError bool
// 	}{
// 		{10, 5, "+", 15, false},
// 		{10, 5, "-", 5, false},
// 		{10, 5, "*", 50, false},
// 		{10, 5, "/", 2, false},
// 		{10, 0, "/", 0, true},  // Division by zero should return error
// 		{10, 5, "%", 0, true},  // Invalid operator should return error
// 	}

// 	for _, tc := range testCases {
// 		result, err := Calculator(tc.a, tc.b, tc.operator)

// 		// Check error handling
// 		if tc.expectError && err == nil {
// 			t.Errorf("Calculator(%f, %f, %s) expected an error but got nil", 
// 				tc.a, tc.b, tc.operator)
// 			continue
// 		}

// 		if !tc.expectError && err != nil {
// 			t.Errorf("Calculator(%f, %f, %s) returned unexpected error: %v", 
// 				tc.a, tc.b, tc.operator, err)
// 			continue
// 		}

// 		// Only check result when no error is expected
// 		if !tc.expectError && result != tc.expected {
// 			t.Errorf("Calculator(%f, %f, %s) = %f; want %f", 
// 				tc.a, tc.b, tc.operator, result, tc.expected)
// 		}
// 	}
// }

// func TestHTTPStatusCode(t *testing.T) {
// 	testCases := []struct {
// 		code     int
// 		expected string
// 	}{
// 		{200, "OK"},
// 		{201, "Created"},
// 		{204, "No Content"},
// 		{400, "Bad Request"},
// 		{404, "Not Found"},
// 		{500, "Internal Server Error"},
// 		{418, "I'm a teapot"},
// 		{999, "Unknown Status Code"},
// 	}

// 	for _, tc := range testCases {
// 		result := HTTPStatusMessage(tc.code)
// 		if result != tc.expected {
// 			t.Errorf("HTTPStatusMessage(%d) = %s; want %s", tc.code, result, tc.expected)
// 		}
// 	}
// } 