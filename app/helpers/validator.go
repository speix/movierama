package helpers

import "regexp"

// InBetween checks if a given number is between two others.
func InBetween(number, min, max int) bool {

	if (number >= min) && (number <= max) {
		return true
	}

	return false
}

// ValidateEmail checks for a valid string representation of email address.
func ValidEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}
