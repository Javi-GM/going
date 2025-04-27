package rangeloop

// SumSlice calculates the sum of all integers in a slice
func SumSlice(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// ContainsValue checks if a slice contains a specific string value
func ContainsValue(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// CountVowels counts the number of vowels in a string
func CountVowels(s string) int {
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	
	count := 0
	for _, char := range s {
		if vowels[char] {
			count++
		}
	}
	return count
}

// MergeMaps combines two maps into a new map
// If there are key conflicts, values from map2 take precedence
func MergeMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)
	
	// Copy all key-value pairs from map1
	for key, value := range map1 {
		result[key] = value
	}
	
	// Copy all key-value pairs from map2 (overwriting any duplicates)
	for key, value := range map2 {
		result[key] = value
	}
	
	return result
}

// FilterSlice returns a new slice containing only elements that pass the filter function
func FilterSlice(nums []int, filterFn func(int) bool) []int {
	var result []int
	
	for _, num := range nums {
		if filterFn(num) {
			result = append(result, num)
		}
	}
	
	return result
}

// MostFrequentChar finds the character that appears most frequently in a string
// If multiple characters tie for most frequent, returns the one that appears first
func MostFrequentChar(s string) rune {
	if len(s) == 0 {
		return 0
	}
	
	// Count frequency of each character
	frequency := make(map[rune]int)
	for _, char := range s {
		frequency[char]++
	}
	
	// Find the character with highest frequency
	var mostFrequent rune
	highestCount := 0
	
	// Iterate through the string again to maintain original order
	// for characters with the same frequency
	for _, char := range s {
		count := frequency[char]
		if count > highestCount {
			highestCount = count
			mostFrequent = char
		}
	}
	
	return mostFrequent
}

// ReverseMap swaps the keys and values of a map
// If multiple keys have the same value, the last key will be used
func ReverseMap(input map[string]int) map[int]string {
	result := make(map[int]string)
	
	for key, value := range input {
		result[value] = key
	}
	
	return result
} 