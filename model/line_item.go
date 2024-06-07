package model

import "github.com/google/uuid"

type LineItem struct {
	ItemID   uuid.UUID
	Quantity uint
	Price    uint
}
