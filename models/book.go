package models

import "time"

type Book struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Author    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
