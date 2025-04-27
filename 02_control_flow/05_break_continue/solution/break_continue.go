package breakcontinue

import (
	"strings"
)

// FindFirstEven finds the first even number in a slice
// Returns the number and a boolean indicating if an even number was found
func FindFirstEven(nums []int) (int, bool) {
	for _, num := range nums {
		if num%2 == 0 {
			return num, true
		}
	}
	return 0, false
}

// SumUntilNegative adds all numbers in a slice until a negative number is encountered
// If no negative number is found, returns the sum of all numbers
func SumUntilNegative(nums []int) int {
	sum := 0
	for _, num := range nums {
		if num < 0 {
			break
		}
		sum += num
	}
	return sum
}

// SkipMultiplesOf returns a slice containing numbers from 1 to max,
// skipping any multiples of n
func SkipMultiplesOf(n, max int) []int {
	var result []int
	
	for i := 1; i <= max; i++ {
		if i%n == 0 {
			continue
		}
		result = append(result, i)
	}
	
	return result
}

// FindSubstring returns the index of the first occurrence of a substring
// Returns -1 if the substring is not found
func FindSubstring(s, sub string) int {
	// Handle edge cases
	if len(sub) == 0 {
		return 0
	}
	if len(s) < len(sub) {
		return -1
	}
	
	// Brute force substring search
	for i := 0; i <= len(s)-len(sub); i++ {
		found := true
		
		for j := 0; j < len(sub); j++ {
			if s[i+j] != sub[j] {
				found = false
				break
			}
		}
		
		if found {
			return i
		}
	}
	
	return -1
	
	// Alternative approach using the standard library:
	// return strings.Index(s, sub)
}

// PrintTriangle returns a slice of strings forming a triangle pattern
// of the specified size
func PrintTriangle(size int) []string {
	if size <= 0 {
		return []string{}
	}
	
	result := make([]string, size)
	
	for i := 0; i < size; i++ {
		// Create a string with (i+1) stars
		result[i] = strings.Repeat("*", i+1)
	}
	
	return result
}

// FindFirstDuplicatePair finds the first pair of consecutive duplicate values
// Returns the value and a boolean indicating if a pair was found
func FindFirstDuplicatePair(nums []int) (int, bool) {
	// Handle edge cases
	if len(nums) < 2 {
		return 0, false
	}
	
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i], true
		}
	}
	
	return 0, false
}

// SieveOfEratosthenes generates all prime numbers up to n
// using the Sieve of Eratosthenes algorithm
func SieveOfEratosthenes(n int) []int {
	// Handle edge cases
	if n < 2 {
		return []int{}
	}
	
	// Create a boolean slice to mark numbers as prime or composite
	// Initialize all as true (prime) and mark non-primes later
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	
	// Apply the Sieve algorithm
	for i := 2; i*i <= n; i++ {
		// If i is prime (not marked yet)
		if isPrime[i] {
			// Mark all multiples of i as non-prime
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	
	// Collect all prime numbers
	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	
	return primes
} 