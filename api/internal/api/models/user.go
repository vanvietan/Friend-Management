package models

import "time"

// User contains information
type User struct {
	ID         int64
	Email      string
	Friends    []string
	Subscribed []string
	Block      []string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
