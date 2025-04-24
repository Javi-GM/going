package ifelse

// CheckNumber determines if a number is positive, negative, or zero
func CheckNumber(num int) string {
	if num > 0 {
		return "positive"
	} else if num < 0 {
		return "negative"
	} else {
		return "zero"
	}
}

// GradeScore converts a numeric score to a letter grade
func GradeScore(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

// CheckEvenOdd determines if a number is even or odd
func CheckEvenOdd(num int) string {
	if num%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

// GetDiscount calculates discount based on amount and loyalty status
func GetDiscount(amount float64, hasLoyaltyCard bool) float64 {
	var discountRate float64
	
	if amount >= 1000 {
		discountRate = 0.10 // 10% discount for amounts >= 1000
		if hasLoyaltyCard {
			discountRate += 0.05 // Additional 5% for loyalty card holders
		}
	} else if amount >= 500 {
		discountRate = 0.05 // 5% discount for amounts >= 500
		if hasLoyaltyCard {
			discountRate += 0.05 // Additional 5% for loyalty card holders
		}
	} else if hasLoyaltyCard {
		discountRate = 0.05 // 5% discount only for loyalty card holders
	}
	
	return amount * discountRate
}

// IsLeapYear determines if a year is a leap year
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
} 