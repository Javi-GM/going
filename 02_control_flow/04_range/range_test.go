package rangeloop

import (
	"reflect"
	"testing"
)

func TestSumSlice(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{-1, -2, -3}, -6},
		{[]int{}, 0},            // Empty slice
		{[]int{42}, 42},         // Single element
		{[]int{10, -10}, 0},     // Sum to zero
	}

	for _, tc := range testCases {
		result := SumSlice(tc.nums)
		if result != tc.expected {
			t.Errorf("SumSlice(%v) = %d; want %d", tc.nums, result, tc.expected)
		}
	}
}

func TestContainsValue(t *testing.T) {
	testCases := []struct {
		slice    []string
		value    string
		expected bool
	}{
		{[]string{"apple", "banana", "orange"}, "banana", true},
		{[]string{"apple", "banana", "orange"}, "grape", false},
		{[]string{}, "apple", false},                      // Empty slice
		{[]string{"apple"}, "apple", true},                // Single element
		{[]string{"apple", "banana"}, "APPLE", false},     // Case sensitive
	}

	for _, tc := range testCases {
		result := ContainsValue(tc.slice, tc.value)
		if result != tc.expected {
			t.Errorf("ContainsValue(%v, %q) = %t; want %t", tc.slice, tc.value, result, tc.expected)
		}
	}
}

func TestCountVowels(t *testing.T) {
	testCases := []struct {
		str      string
		expected int
	}{
		{"hello", 2},
		{"Go Programming", 4},
		{"rhythm", 0},           // No vowels
		{"", 0},                 // Empty string
		{"aeiouAEIOU", 10},      // All vowels
		{"123!@#", 0},           // No vowels
	}

	for _, tc := range testCases {
		result := CountVowels(tc.str)
		if result != tc.expected {
			t.Errorf("CountVowels(%q) = %d; want %d", tc.str, result, tc.expected)
		}
	}
}

func TestMergeMaps(t *testing.T) {
	testCases := []struct {
		map1     map[string]int
		map2     map[string]int
		expected map[string]int
	}{
		{
			map[string]int{"a": 1, "b": 2},
			map[string]int{"c": 3, "d": 4},
			map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
		},
		{
			map[string]int{"a": 1, "b": 2},
			map[string]int{"b": 3, "c": 4},
			map[string]int{"a": 1, "b": 3, "c": 4}, // map2 values overwrite map1 on conflict
		},
		{
			map[string]int{},
			map[string]int{"a": 1, "b": 2},
			map[string]int{"a": 1, "b": 2},
		},
		{
			map[string]int{"a": 1, "b": 2},
			map[string]int{},
			map[string]int{"a": 1, "b": 2},
		},
		{
			map[string]int{},
			map[string]int{},
			map[string]int{},
		},
	}

	for _, tc := range testCases {
		// Make copies of the maps to ensure the function doesn't modify the originals
		map1Copy := make(map[string]int)
		for k, v := range tc.map1 {
			map1Copy[k] = v
		}
		map2Copy := make(map[string]int)
		for k, v := range tc.map2 {
			map2Copy[k] = v
		}

		result := MergeMaps(map1Copy, map2Copy)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("MergeMaps(%v, %v) = %v; want %v", tc.map1, tc.map2, result, tc.expected)
		}
	}
}

func TestFilterSlice(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	isPositive := func(n int) bool { return n > 0 }
	isNegative := func(n int) bool { return n < 0 }

	testCases := []struct {
		nums      []int
		filterFn  func(int) bool
		expected  []int
		testName  string
	}{
		{[]int{1, 2, 3, 4, 5}, isEven, []int{2, 4}, "even numbers"},
		{[]int{-3, -2, -1, 0, 1, 2, 3}, isPositive, []int{1, 2, 3}, "positive numbers"},
		{[]int{-3, -2, -1, 0, 1, 2, 3}, isNegative, []int{-3, -2, -1}, "negative numbers"},
		{[]int{}, isEven, []int{}, "empty slice"},
		{[]int{1, 3, 5}, isEven, []int{}, "no matching elements"},
		{[]int{2, 4, 6}, isEven, []int{2, 4, 6}, "all matching elements"},
	}

	for _, tc := range testCases {
		result := FilterSlice(tc.nums, tc.filterFn)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("FilterSlice(%v, %s) = %v; want %v", tc.nums, tc.testName, result, tc.expected)
		}
	}
}

// Bonus challenge tests
func TestMostFrequentChar(t *testing.T) {
	testCases := []struct {
		str      string
		expected rune
	}{
		{"hello", 'l'},             // 'l' appears twice
		{"programming", 'g'},       // 'g' and 'r' and 'm' appear twice, but 'g' is first
		{"aaaaabbbcc", 'a'},        // 'a' appears 5 times
		{"", 0},                    // Empty string
		{"abcde", 'a'},             // All chars appear once, return first
	}

	for _, tc := range testCases {
		result := MostFrequentChar(tc.str)
		if result != tc.expected {
			t.Errorf("MostFrequentChar(%q) = %q; want %q", tc.str, string(result), string(tc.expected))
		}
	}
}

func TestReverseMap(t *testing.T) {
	testCases := []struct {
		input    map[string]int
		expected map[int]string
	}{
		{
			map[string]int{"one": 1, "two": 2, "three": 3},
			map[int]string{1: "one", 2: "two", 3: "three"},
		},
		{
			map[string]int{},
			map[int]string{},
		},
		{
			map[string]int{"a": 1, "b": 1, "c": 2},
			map[int]string{1: "b", 2: "c"}, // For duplicate values, last key wins
		},
	}

	for _, tc := range testCases {
		result := ReverseMap(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("ReverseMap(%v) = %v; want %v", tc.input, result, tc.expected)
		}
	}
} 