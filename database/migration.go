package database

import (
	"github.com/jinzhu/gorm"
	"main/api/models"
)

func Migrate(DB *gorm.DB) (err error) {
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Brand{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Color{})
	DB.AutoMigrate(&models.Image{})
	DB.AutoMigrate(&models.Review{})
	DB.AutoMigrate(&models.Size{})
	DB.AutoMigrate(&models.Variant{})

return
}