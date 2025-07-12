package helpers

import (

	"regexp"
)

// Validate email
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}


func ValidatePhone(phone string) bool {
	re := regexp.MustCompile(`^\d{10,15}$`)
	return re.MatchString(phone)
}
