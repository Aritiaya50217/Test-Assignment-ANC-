package main

import (
	models "github.com/Aritiaya50217/Test-Assignment-ANC/models"
	"github.com/Aritiaya50217/Test-Assignment-ANC/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	initDatabase()
	r := routes.SetRouter()
	r.Run()
}

func initDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:test@tcp(127.0.0.1:3306)/ecommerce"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// auto migrate
	db.AutoMigrate(&models.Gender{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Products{})
	db.AutoMigrate(&models.Size{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Status{})
	db.AutoMigrate(&models.Color{})
	db.AutoMigrate(&models.Patterns{})
	db.AutoMigrate(&models.Figures{})

	return db
}
