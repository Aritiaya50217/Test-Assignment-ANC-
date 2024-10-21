package models

import "time"

type Size struct {
	Id        uint `gorm:"primaryKey;default:auto_random()"`
	Size      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
