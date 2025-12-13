package model

import "time"

// User ...
type User struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
