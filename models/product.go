package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model

	Name      string `json:"name"`
	Price     int	`json:"price"`
	Weight    int 	`json:"weight"`
	Stock     int `json:"stock"`
	TotalView int `json:"total_view"`

	// Brand has many products, BrandID is the foreign key
	BrandID uint `json:"brand_id"`
	Brand   Brand `gorm:"-" json:"brand"`
	// Category has many products, CategoryID is the foreign key
	Category   Category `gorm:"-" json:"category"`
	CategoryID uint `json:"category_id"`

	// Product has many reviews, images and variant
	// db.Model(&user).Related(&emails)
	// SELECT * FROM emails WHERE user_id = 111;
	// 111 is user's primary key
	Review  []Review `json:"review"`
	Image   []Image `json:"image"`
	Variant []Variant `json:"variant"`
}
