package forloops

import (
	"reflect"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{5, 15},     // 1+2+3+4+5 = 15
		{10, 55},    // 1+2+...+10 = 55
		{100, 5050}, // Sum of 1 to 100
		{0, 0},      // Edge case
		{-5, 0},     // Should handle negative numbers gracefully
	}

	for _, tc := range testCases {
		result := SumNumbers(tc.n)
		if result != tc.expected {
			t.Errorf("SumNumbers(%d) = %d; want %d", tc.n, result, tc.expected)
		}
	}
}

func TestFactorial(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 1},      // 0! = 1 by definition
		{1, 1},      // 1! = 1
		{5, 120},    // 5! = 5*4*3*2*1 = 120
		{10, 3628800}, // 10!
		{-1, 0},     // Should handle negative numbers gracefully
	}

	for _, tc := range testCases {
		result := Factorial(tc.n)
		if result != tc.expected {
			t.Errorf("Factorial(%d) = %d; want %d", tc.n, result, tc.expected)
		}
	}
}

func TestReverseString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"Go programming", "gnimmargorp oG"},
		{"racecar", "racecar"}, // Palindrome
		{"", ""},               // Empty string
		{"a", "a"},             // Single character
	}

	for _, tc := range testCases {
		result := ReverseString(tc.input)
		if result != tc.expected {
			t.Errorf("ReverseString(%q) = %q; want %q", tc.input, result, tc.expected)
		}
	}
}

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{2, true},    // Smallest prime
		{3, true},    // Prime
		{4, false},   // Not prime (2*2)
		{17, true},   // Prime
		{25, false},  // Not prime (5*5)
		{97, true},   // Prime
		{1, false},   // 1 is not considered prime
		{0, false},   // 0 is not prime
		{-7, false},  // Negative numbers are not prime
	}

	for _, tc := range testCases {
		result := IsPrime(tc.n)
		if result != tc.expected {
			t.Errorf("IsPrime(%d) = %t; want %t", tc.n, result, tc.expected)
		}
	}
}

func TestFindMax(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3, 5, 7, 9}, 9},
		{[]int{-5, -3, -1}, -1},
		{[]int{42}, 42},          // Single element
		{[]int{3, 3, 3}, 3},      // All same
		{[]int{5, 4, 3, 2, 1}, 5}, // Decreasing order
		{[]int{-10, 0, 10}, 10},  // Mixed signs
	}

	for _, tc := range testCases {
		result := FindMax(tc.nums)
		if result != tc.expected {
			t.Errorf("FindMax(%v) = %d; want %d", tc.nums, result, tc.expected)
		}
	}
}

// Bonus challenge tests
func TestFibonacci(t *testing.T) {
	testCases := []struct {
		n        int
		expected []int
	}{
		{1, []int{0}},
		{2, []int{0, 1}},
		{5, []int{0, 1, 1, 2, 3}},
		{8, []int{0, 1, 1, 2, 3, 5, 8, 13}},
		{0, []int{}},
		{-1, []int{}}, // Should handle negative inputs gracefully
	}

	for _, tc := range testCases {
		result := Fibonacci(tc.n)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Fibonacci(%d) = %v; want %v", tc.n, result, tc.expected)
		}
	}
}

func TestCharFrequency(t *testing.T) {
	testCases := []struct {
		input    string
		expected map[rune]int
	}{
		{"hello", map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1}},
		{"go programming", map[rune]int{'g': 2, 'o': 1, ' ': 1, 'p': 1, 'r': 2, 'a': 1, 'm': 2, 'i': 1, 'n': 1}},
		{"", map[rune]int{}}, // Empty string
		{"   ", map[rune]int{' ': 3}}, // Just spaces
	}

	for _, tc := range testCases {
		result := CharFrequency(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("CharFrequency(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
} 