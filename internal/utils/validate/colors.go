package utils

import (
	"strings"
)

func ValidateColor(name, hex_code string) map[string]string {
	errors := make(map[string]string)

	// Name validation
	if strings.TrimSpace(name) == "" {
		errors["name"] = "color name is required"
	} else if len(strings.TrimSpace(name)) < 2 {
		errors["name"] = "color name must be at least 2 characters"
	}

	// Name validation
	if strings.TrimSpace(hex_code) == "" {
		errors["name"] = "color hex code is required"
	} else if len(strings.TrimSpace(name)) < 2 {
		errors["name"] = "color hex code must be at least 6 characters"
	}

	return errors
}
