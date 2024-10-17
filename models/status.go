package models

import (
	"time"
)

type Status struct {
	Id        uint      `gorm:"primaryKey;default:auto_random()"`
	Status    string    `gorm:"status"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
