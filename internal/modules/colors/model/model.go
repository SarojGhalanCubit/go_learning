package colorModel

import (
	"time"

	"github.com/google/uuid"
)

type Colors struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	HexCode   string    `json:"hex_code"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateColor struct {
	Name    string `json:"name"`
	HexCode string `json:"hex_code"`
}
