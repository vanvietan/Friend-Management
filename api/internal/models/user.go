package models

// User contains information
type User struct {
	ID           int64
	Email        string
	Relationship []Relationship `gorm:"-"`
	//CreatedAt    time.Time
	//UpdatedAt    time.Time
}
