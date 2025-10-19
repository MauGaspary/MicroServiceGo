package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID   uuid.UUID `json:"order_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Items     []Item    `json:"items"`
	Total     float64   `json:"total"`
	Currency  string    `json:"currency"`
}

type Item struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Name      string    `json:"name"`
}
