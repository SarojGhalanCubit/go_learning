package productModel

import "github.com/google/uuid"

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"string"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	IsActive    bool      `json:"is_active"`
	MaterialID  uuid.UUID `json:"material_id"`
	CategoryID  uuid.UUID `json:"category_id"`
}
