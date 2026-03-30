package sizeModel

import (
	"time"

	"github.com/google/uuid"
)

type Sizes struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateSize struct {
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
}
