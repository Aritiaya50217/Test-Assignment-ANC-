package models

import "time"

type Color struct {
	Id        uint      `gorm:"primaryKey;default:auto_random()"`
	Name      string    `gorm:"name"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
