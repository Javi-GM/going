package forloops

import (
	"math"
)

// SumNumbers calculates the sum of all integers from 1 to n
func SumNumbers(n int) int {
	// Return 0 for invalid inputs
	if n <= 0 {
		return 0
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum

	// Alternative solution using mathematical formula:
	// return n * (n + 1) / 2
}

// Factorial calculates n!
func Factorial(n int) int {
	// Special cases
	if n < 0 {
		return 0 // Invalid input
	}
	if n == 0 {
		return 1 // 0! = 1 by definition
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// ReverseString reverses the characters in a string
func ReverseString(s string) string {
	// Convert string to rune slice to handle Unicode correctly
	runes := []rune(s)
	length := len(runes)
	
	// Create a new slice to store the result
	reversed := make([]rune, length)
	
	// Fill the reversed slice
	for i := 0; i < length; i++ {
		reversed[i] = runes[length-1-i]
	}
	
	// Convert back to string
	return string(reversed)
	
	// Alternative approach using in-place swap:
	// for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
	//     runes[i], runes[j] = runes[j], runes[i]
	// }
	// return string(runes)
}

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
	// Special cases
	if n <= 1 {
		return false // By definition, 0 and 1 are not prime, and negative numbers aren't prime
	}
	if n <= 3 {
		return true // 2 and 3 are prime
	}
	if n%2 == 0 || n%3 == 0 {
		return false // Multiples of 2 or 3 aren't prime
	}

	// Check all potential factors up to sqrt(n)
	limit := int(math.Sqrt(float64(n)))
	
	// We only need to check numbers of form 6k±1 up to sqrt(n)
	for i := 5; i <= limit; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	
	return true
}

// FindMax finds the largest number in a slice
func FindMax(nums []int) int {
	// Handle empty slice (not in the test cases but good practice)
	if len(nums) == 0 {
		panic("Cannot find maximum of empty slice")
	}

	// Start with the first element as the maximum
	max := nums[0]
	
	// Compare each element to the current maximum
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	
	return max
}

// Fibonacci generates the first n numbers in the Fibonacci sequence
func Fibonacci(n int) []int {
	// Handle invalid inputs
	if n <= 0 {
		return []int{}
	}
	
	// Initialize the result slice
	fib := make([]int, n)
	
	// Fill the first two elements (special cases)
	if n >= 1 {
		fib[0] = 0
	}
	if n >= 2 {
		fib[1] = 1
	}
	
	// Generate the remaining Fibonacci numbers
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	
	return fib
}

// CharFrequency counts the frequency of each character in a string
func CharFrequency(s string) map[rune]int {
	// Initialize an empty map
	freq := make(map[rune]int)
	
	// Count each character
	for _, char := range s {
		freq[char]++
	}
	
	return freq
} 