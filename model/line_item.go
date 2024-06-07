package model

import "github.com/google/uuid"

type LineItem struct {
	ItemID   uuid.UUID `json:"itemId"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
