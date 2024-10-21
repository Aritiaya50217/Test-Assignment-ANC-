package models

import "time"

type Order struct {
	Id uint `gorm:"primaryKey;default:auto_random()"`
	// Product   *Products
	ProductId uint `gorm:"product_id"`
	Amount    uint `gorm:"amount"`
	Status    *Status
	StatusId  uint `gorm:"status_id"`
	User      *User
	UserId    uint
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func GetOrders(startDate, endDate string, statusId, offset, max, perPage int) (order []Order, total int, err error) {
	// TODO : query

	return order, total, err
}
