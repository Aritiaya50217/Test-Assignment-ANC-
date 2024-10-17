package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Products struct {
	Id        uint   `gorm:"primaryKey;default:auto_random()"`
	Style     string `gorm:"style"`
	Gender    *Gender
	GenderId  uint
	Size      *Size
	SizeId    uint
	Color     *Color
	ColorId   uint
	Price     float64
	Pattern   *Patterns
	PatternId uint
	Figures   *Figures
	FigureId  uint
}

func GetAllProducts(genderId int, style string, sizeId, offset, limit, perPage int) (products []Products, total int, err error) {
	// TODO : query

	// if err := db.Find(&products).Error; err != nil {
	// 	return nil, err
	// }
	return products, total, nil
}

func GetProductById(id int) (product *Products, err error) {
	// TODO : query

	// if err = db.First(product, "id = ? ", id).Error; err != nil {
	// 	return nil, err
	// }
	return product, nil
}
