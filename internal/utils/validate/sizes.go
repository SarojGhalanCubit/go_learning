package utils

import (
	"strings"
)

func ValidateSize(name string, sort_order int) map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(name) == "" {
		errors["name"] = "Name is required"
	}
	if sort_order <= 0 {
		errors["sort_order"] = "Sort order must be a positive number"
	}
	return errors
}
