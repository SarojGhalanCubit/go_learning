package productModel

import "github.com/google/uuid"

type ProductResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"string"`
	Slug         string    `json:"slug"`
	Description  string    `json:"description"`
	Quantity     int       `json:"quantity"`
	IsActive     bool      `json:"is_active"`
	MaterialName string    `json:"material_name"`
	CategoryName string    `json:"category_name"`
}

type CreateProduct struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Slug        string    `json:"slug"`
	IsActive    bool      `json:"is_active"`
	MaterialID  uuid.UUID `json:"material_id"`
	CategoryID  uuid.UUID `json:"category_id"`
}
