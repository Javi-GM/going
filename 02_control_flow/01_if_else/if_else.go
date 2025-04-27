package ifelse

import "fmt"

const POSITIVE, NEGATIVE, ZERO = "positive", "negative", "zero"
const A, B, C, D, E, F = "A", "B", "C", "D", "E", "F"
const EVEN, ODD = "EVEN", "ODD"

func CheckNumber(value int) string {
	if value > 0 { 
		return POSITIVE
	}
	if value < 0 {
		return NEGATIVE
	} 

	return ZERO
}

func GradeScore(gradeScore int) (string, error)  {
	if gradeScore > 100 || gradeScore < 0 {
		return "", fmt.Errorf("grade score %d is not valid: should be between 0 and 100", gradeScore)
	}

	if gradeScore > 85 {
		return A, nil
	} else if gradeScore > 75 {
		return B, nil
	} else if gradeScore > 65 {
		return C, nil
	} else if gradeScore > 55 {
		return D, nil
	}
	return F, nil
}

func CheckEvenOdd(value int) string {
	if (value % 2 == 0) {
		return EVEN
	}
	return ODD
}

func GetDiscount(amount float64, hasLoyaltyCard bool) float64 {
	const SMALL_AMOUNT_DISCOUNT = 0.0;
	const SMALL_AMOUNT_DISCOUNT_WITH_LOYALTY_CARD = 0.05;
	const LARGE_AMOUNT_DISCOUNT = 0.05;
	const LARGE_AMOUNT_DISCOUNT_WITH_LOYALTY_CARD = 0.1;
	const VERY_LARGE_AMOUNT_DISCOUNT = 0.1;
	const VERY_LARGE_AMOUNT_DISCOUNT_WITH_LOYALTY_CARD = 0.15;

	if amount <= 100 {
		if !hasLoyaltyCard {
			return amount * SMALL_AMOUNT_DISCOUNT
		} else {
			return amount * SMALL_AMOUNT_DISCOUNT_WITH_LOYALTY_CARD
		}
	}
	if amount <= 500 {
		if !hasLoyaltyCard {
			return amount * LARGE_AMOUNT_DISCOUNT
		} else {
			return amount * LARGE_AMOUNT_DISCOUNT_WITH_LOYALTY_CARD
		}
	}
	if amount > 500 {
		if !hasLoyaltyCard {
			return amount * VERY_LARGE_AMOUNT_DISCOUNT
		} else {
			return amount * VERY_LARGE_AMOUNT_DISCOUNT_WITH_LOYALTY_CARD
		}
	}

	return 0.0
}