package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderResponse struct {
	ID uuid.UUID `json:"id,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type OrderCreateRequest struct {
	ID uuid.UUID
}
