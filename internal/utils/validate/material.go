package utils

import "strings"

func ValidateMaterial(name string) map[string]string {
	errors := make(map[string]string)

	// Name validation
	if strings.TrimSpace(name) == "" {
		errors["name"] = "Name is required"
	} else if len(strings.TrimSpace(name)) < 2 {
		errors["name"] = "Name must be at least 2 characters"
	}

	return errors

}
