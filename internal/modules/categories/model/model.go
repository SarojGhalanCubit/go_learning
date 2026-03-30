package categoryModel

import (
	"time"

	"github.com/google/uuid"
)

type Categories struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type CreateCategory struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	Slug     string `json:"slug"`
}
