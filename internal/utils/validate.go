package utils

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func ValidateUser(name string, age int, phone string, email string, password string) map[string]string {
	errors := make(map[string]string)

	// Name validation
	if strings.TrimSpace(name) == "" {
		errors["name"] = "Name is required"
	} else if len(strings.TrimSpace(name)) < 2 {
		errors["name"] = "Name must be at least 2 characters"
	}

	// Age validation
	// if age < 18 {
	// 	errors["age"] = "Age must be greater than 18"
	// }
	//
	// Phone Validation
	// if strings.TrimSpace(phone) == "" {
	// 	errors["phone_number"] = "Phone Number is required."
	// }else if len(strings.TrimSpace(phone)) < 10 {
	// 	errors["phone_number"] = "Phone number must contain 10 numbers"
	// }

	// Email validation
	if strings.TrimSpace(email) == "" {
		errors["email"] = "Email is required"
	} else if !emailRegex.MatchString(email) {
		errors["email"] = "Invalid email format"
	}

	// Password validation
	if strings.TrimSpace(password) == "" {
		errors["password"] = "Password is required"
	} else if len(password) < 8 {
		errors["password"] = "Password must be at least 8 characters"
	} else if !strings.ContainsAny(password, "0123456789") {
		errors["password"] = "Password must contain at least one number"
	} else if !strings.ContainsAny(password, "!@#$%^&*") {
		errors["password"] = "Password must contain at least one special character (!@#$%^&*)"
	}

	return errors
}

func ValidateUpdateUser(name string, age int, phone string, email string) map[string]string {
	errors := make(map[string]string)

	// Name validation
	if strings.TrimSpace(name) == "" {
		errors["name"] = "Name is required"
	} else if len(strings.TrimSpace(name)) < 2 {
		errors["name"] = "Name must be at least 2 characters"
	}

	// Age validation
	if age < 18 {
		errors["age"] = "Age must be greater than 18"
	}

	// Phone Validation
	if strings.TrimSpace(phone) == "" {
		errors["phone_number"] = "Phone Number is required"
	} else if len(strings.TrimSpace(phone)) > 10 {

		errors["phone_number"] = "Phone number must contain 10 numbers"
	}

	// Email validation
	if strings.TrimSpace(email) == "" {
		errors["email"] = "Email is required"
	} else if !emailRegex.MatchString(email) {
		errors["email"] = "Invalid email format"
	}

	return errors
}

func ValidateLoginUser(email, password string) map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(email) == "" {
		errors["email"] = "Email is required"
	}

	if strings.TrimSpace(password) == "" {
		errors["password"] = "Password is required"
	}

	return errors
}
