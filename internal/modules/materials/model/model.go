package materialsModel

import (
	"github.com/google/uuid"
	"time"
)

type Material struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateMaterial struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
