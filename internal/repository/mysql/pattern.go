package models

import "time"

type Patterns struct {
	Id        uint `gorm:"primaryKey;default:auto`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

