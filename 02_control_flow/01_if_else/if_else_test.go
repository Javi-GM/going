package ifelse

import (
	"testing"
)

func TestCheckNumber(t *testing.T) {
	testCases := []struct {
		num      int
		expected string
	}{
		{5, "positive"},
		{0, "zero"},
		{-10, "negative"},
	}
	
	for _, tc := range testCases {
		result := CheckNumber(tc.num)
		if result != tc.expected {
			t.Errorf("CheckNumber(%d) = %s; want %s", tc.num, result, tc.expected)
		}
	}
}

func TestGradeScore(t *testing.T) {
	testCases := []struct {
		score    int
		expected string
	}{
		{95, "A"},
		{85, "B"},
		{75, "C"},
		{65, "D"},
		{55, "F"},
		{100, "A"},
		{50, "F"},
		{0, "F"},
	}
	
	for _, tc := range testCases {
		result := GradeScore(tc.score)
		if result != tc.expected {
			t.Errorf("GradeScore(%d) = %s; want %s", tc.score, result, tc.expected)
		}
	}
}

func TestCheckEvenOdd(t *testing.T) {
	testCases := []struct {
		num      int
		expected string
	}{
		{2, "even"},
		{7, "odd"},
		{0, "even"},
		{-1, "odd"},
		{-4, "even"},
	}
	
	for _, tc := range testCases {
		result := CheckEvenOdd(tc.num)
		if result != tc.expected {
			t.Errorf("CheckEvenOdd(%d) = %s; want %s", tc.num, result, tc.expected)
		}
	}
}

func TestGetDiscount(t *testing.T) {
	testCases := []struct {
		amount         float64
		hasLoyaltyCard bool
		expected       float64
	}{
		{100.0, false, 0.0},      // No discount for small amount, no loyalty card
		{100.0, true, 5.0},       // 5% loyalty discount for small amount
		{500.0, false, 25.0},     // 5% discount for large amount
		{500.0, true, 50.0},      // 10% discount for large amount with loyalty card
		{1000.0, false, 100.0},   // 10% discount for very large amount
		{1000.0, true, 150.0},    // 15% discount for very large amount with loyalty card
	}
	
	for _, tc := range testCases {
		result := GetDiscount(tc.amount, tc.hasLoyaltyCard)
		if result != tc.expected {
			t.Errorf("GetDiscount(%f, %t) = %f; want %f", 
				tc.amount, tc.hasLoyaltyCard, result, tc.expected)
		}
	}
}

// Bonus challenge test
// func TestIsLeapYear(t *testing.T) {
// 	testCases := []struct {
// 		year     int
// 		expected bool
// 	}{
// 		{2000, true},  // Divisible by 400
// 		{2020, true},  // Divisible by 4 but not by 100
// 		{1900, false}, // Divisible by 100 but not by 400
// 		{2021, false}, // Not divisible by 4
// 	}
// 	
// 	for _, tc := range testCases {
// 		result := IsLeapYear(tc.year)
// 		if result != tc.expected {
// 			t.Errorf("IsLeapYear(%d) = %t; want %t", tc.year, result, tc.expected)
// 		}
// 	}
// } 