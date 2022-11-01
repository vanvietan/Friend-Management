package models

import "time"

// Relationship describes type of relationship
type Relationship struct {
	ID          int64
	AddresseeID int64
	RequesterID int64
	Type        Type
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
