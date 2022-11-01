package models

import "time"

// User contains information
type User struct {
	ID           int64
	Email        string
	Relationship []Relationship
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
