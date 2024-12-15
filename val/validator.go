package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {

	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("length must be between %d - %d", minLength, maxLength)
	}
	return nil

}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 20); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("username must be alphanumeric and underscore only")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 8, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("invalid mail address")
	}
	return nil
}
func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 20); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("full name should only contain letters and spaces")
	}
	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}
