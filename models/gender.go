package models

import "time"

type Gender struct {
	Id        uint `gorm:"primaryKey;default:auto_random()"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
