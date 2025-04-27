package switch_statement

import (
	"errors"
	"fmt"
	"reflect"
)

// DayOfWeek returns the name of the day based on its number (1-7)
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

// GetSeasonInNorthernHemisphere returns the season based on the month name
func GetSeasonInNorthernHemisphere(month string) string {
	switch month {
	case "December", "January", "February":
		return "Winter"
	case "March", "April", "May":
		return "Spring"
	case "June", "July", "August":
		return "Summer"
	case "September", "October", "November":
		return "Fall"
	default:
		return "Unknown season"
	}
}

// CalculateGrade returns a letter grade based on a numeric score
func CalculateGrade(score int) (string, error) {
	if score < 0 || score > 100 {
		return "", errors.New("score must be between 0 and 100")
	}

	switch {
	case score >= 90:
		return "A", nil
	case score >= 80:
		return "B", nil
	case score >= 70:
		return "C", nil
	case score >= 60:
		return "D", nil
	default:
		return "F", nil
	}
}

// DetectType returns the type of the value as a string
func DetectType(value interface{}) string {
	switch v := value.(type) {
	case nil:
		return "nil"
	case int, int8, int16, int32, int64:
		return "int"
	case uint, uint8, uint16, uint32, uint64:
		return "uint"
	case float32, float64:
		return "float64"
	case bool:
		return "bool"
	case string:
		return "string"
	case []interface{}:
		return "slice"
	case map[string]interface{}:
		return "map"
	default:
		// Use reflection for complex types
		kind := reflect.ValueOf(v).Kind().String()
		switch kind {
		case "slice":
			return "slice"
		case "map":
			return "map"
		case "struct":
			return "struct"
		case "ptr":
			return "pointer"
		default:
			return kind
		}
	}
}

// Calculator performs a basic arithmetic operation
func Calculator(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", operator)
	}
}

// HTTPStatusMessage returns a human-readable message for HTTP status codes
func HTTPStatusMessage(code int) string {
	switch code {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 204:
		return "No Content"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 403:
		return "Forbidden"
	case 404:
		return "Not Found"
	case 418:
		return "I'm a teapot"
	case 500:
		return "Internal Server Error"
	case 502:
		return "Bad Gateway"
	case 503:
		return "Service Unavailable"
	default:
		return "Unknown Status Code"
	}
} 