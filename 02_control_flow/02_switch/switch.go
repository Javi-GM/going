package switch_statement

import "fmt"

func DayOfWeek(day int) string {
	switch day {
		case 1:
			return "Monday"
		case 2:
			return "Tuesday"
		case 3:
			return "Wednesday"
		case 4:
			return "Thursday"
		case 5:
			return "Friday"
		case 6:
			return "Saturday"
		case 7:
			return "Sunday"
		default:
			return "Invalid day"
	}
}

func GetSeasonInNorthernHemisphere(month string) string {
	switch month {
		case "December": 
			fallthrough
		case "January": 
			fallthrough
		case "February": 
			return "Winter"
		case "March": 
			fallthrough
		case "April": 
			fallthrough
		case "May": 
			return "Spring"
		case "June": 
			fallthrough
		case "July": 
			fallthrough
		case "August": 
			return "Summer"
		case "September": 
			fallthrough
		case "October": 
			fallthrough
		case "November": 
			return "Fall"
		default:
			return "Unknown season"
	}
}

func CalculateGrade(score int) (string, error) {
	switch {
		case score >= 95 && score <= 100:
			return "A", nil
		case score >= 85 && score <= 94:
			return "B", nil
		case score >= 75 && score <= 84:
			return "C", nil
		case score >= 65 && score <= 74:
			return "D", nil
		case score >= 0 && score <= 64:
			return "F", nil
		default:
			return "", fmt.Errorf("expected score should be between 0 and 100 but got %d", score)
	}
}