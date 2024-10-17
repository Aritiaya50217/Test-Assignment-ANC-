package models

import "time"

type User struct {
	Id        uint      `gorm:"primaryKey;default:auto_random()"`
	Username  string    `gorm:"username"`
	Address   string    `gorm:"address"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
