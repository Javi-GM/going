package breakcontinue

import (
	"reflect"
	"testing"
)

func TestFindFirstEven(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
		found    bool
	}{
		{[]int{1, 3, 5, 6, 7, 8}, 6, true},
		{[]int{2, 4, 6, 8}, 2, true},
		{[]int{1, 3, 5, 7}, 0, false},
		{[]int{}, 0, false},
	}

	for _, tc := range testCases {
		result, found := FindFirstEven(tc.nums)
		if found != tc.found {
			t.Errorf("FindFirstEven(%v) found = %v; want %v", tc.nums, found, tc.found)
		}
		if found && result != tc.expected {
			t.Errorf("FindFirstEven(%v) = %d; want %d", tc.nums, result, tc.expected)
		}
	}
}

func TestSumUntilNegative(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, -1, 4, 5}, 6},
		{[]int{-1, 2, 3, 4}, 0},
		{[]int{1, 2, 3, 4}, 10},
		{[]int{}, 0},
	}

	for _, tc := range testCases {
		result := SumUntilNegative(tc.nums)
		if result != tc.expected {
			t.Errorf("SumUntilNegative(%v) = %d; want %d", tc.nums, result, tc.expected)
		}
	}
}

func TestSkipMultiplesOf(t *testing.T) {
	testCases := []struct {
		n        int
		max      int
		expected []int
	}{
		{3, 10, []int{1, 2, 4, 5, 7, 8, 10}},
		{2, 10, []int{1, 3, 5, 7, 9}},
		{5, 20, []int{1, 2, 3, 4, 6, 7, 8, 9, 11, 12, 13, 14, 16, 17, 18, 19}},
		{7, 6, []int{1, 2, 3, 4, 5, 6}},
		{1, 5, []int{}}, // All numbers are multiples of 1
	}

	for _, tc := range testCases {
		result := SkipMultiplesOf(tc.n, tc.max)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("SkipMultiplesOf(%d, %d) = %v; want %v", tc.n, tc.max, result, tc.expected)
		}
	}
}

func TestFindSubstring(t *testing.T) {
	testCases := []struct {
		s        string
		sub      string
		expected int
	}{
		{"Hello, world!", "world", 7},
		{"Go is awesome", "is", 3},
		{"Programming", "gram", 3},
		{"Testing", "xyz", -1}, // Substring not found
		{"", "test", -1},       // Empty string
		{"Test", "", 0},        // Empty substring (matches at the beginning)
	}

	for _, tc := range testCases {
		result := FindSubstring(tc.s, tc.sub)
		if result != tc.expected {
			t.Errorf("FindSubstring(%q, %q) = %d; want %d", tc.s, tc.sub, result, tc.expected)
		}
	}
}

func TestPrintTriangle(t *testing.T) {
	testCases := []struct {
		size     int
		expected []string
	}{
		{1, []string{"*"}},
		{3, []string{"*", "**", "***"}},
		{5, []string{"*", "**", "***", "****", "*****"}},
		{0, []string{}},
	}

	for _, tc := range testCases {
		result := PrintTriangle(tc.size)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("PrintTriangle(%d) = %v; want %v", tc.size, result, tc.expected)
		}
	}
}

// Bonus challenge tests
func TestFindFirstDuplicatePair(t *testing.T) {
	testCases := []struct {
		nums      []int
		expected  int
		foundPair bool
	}{
		{[]int{1, 2, 2, 3, 4}, 2, true},
		{[]int{1, 3, 5, 5, 7}, 5, true},
		{[]int{5, 5, 1, 2, 3}, 5, true},
		{[]int{1, 2, 3, 4, 5}, 0, false}, // No duplicates
		{[]int{1}, 0, false},             // Too short
		{[]int{}, 0, false},              // Empty slice
	}

	for _, tc := range testCases {
		result, found := FindFirstDuplicatePair(tc.nums)
		if found != tc.foundPair {
			t.Errorf("FindFirstDuplicatePair(%v) found = %v; want %v", tc.nums, found, tc.foundPair)
		}
		if found && result != tc.expected {
			t.Errorf("FindFirstDuplicatePair(%v) = %d; want %d", tc.nums, result, tc.expected)
		}
	}
}

func TestSieveOfEratosthenes(t *testing.T) {
	testCases := []struct {
		n        int
		expected []int
	}{
		{10, []int{2, 3, 5, 7}},
		{20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{1, []int{}},  // No primes less than or equal to 1
		{2, []int{2}}, // Just 2
	}

	for _, tc := range testCases {
		result := SieveOfEratosthenes(tc.n)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("SieveOfEratosthenes(%d) = %v; want %v", tc.n, result, tc.expected)
		}
	}
} 