package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID     uint64     `json:"orderId"`
	CustomerID  uuid.UUID  `json:"customerId"`
	LineItems   []LineItem `json:"listItems"`
	CreatedAt   *time.Time `json:"createdAt"`
	ShippedAt   *time.Time `json:"shippedAt"`
	CompletedAt *time.Time `json:"completedAt"`
}
