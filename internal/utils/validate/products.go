package utils

import (
	"github.com/google/uuid"
	"strings"
)

func ValidateProduct(name string, description string, quantity int, categoryID uuid.UUID, materialID uuid.UUID) map[string]string {
	errors := make(map[string]string)

	// Name Validation
	if strings.TrimSpace(name) == "" {
		errors["name"] = "Product name is required"
	} else if len(strings.TrimSpace(name)) < 3 {
		errors["name"] = "Product name must be at least 3 characters"
	}

	// Description Validation
	if strings.TrimSpace(description) == "" {
		errors["description"] = "Description is required"
	}

	// Price Validation
	// if price <= 0 {
	// 	errors["price"] = "Price must be greater than zero"
	// }

	// Quantity/Stock Validation
	if quantity < 0 {
		errors["quantity"] = "Quantity cannot be negative"
	}

	// Category ID Validation (Checking for nil UUID)
	if categoryID == uuid.Nil {
		errors["category_id"] = "A valid category must be selected"
	}

	// Material ID Validation
	if materialID == uuid.Nil {
		errors["material_id"] = "A valid material must be selected"
	}

	return errors
}
