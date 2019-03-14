package models

import "github.com/jinzhu/gorm"

type Review struct {
	gorm.Model

	ImageURL  string `json:"image_url"`
	RatePoint int `json:"rate_point"`
	Comment   string `json:"comment"`

	// Product has many reviews, ProductID is the foreign key
	ProductID uint `gorm:"foreignkey:ReviewRefer" json:"product_id"`
}

//var review Review
//db.Model(&productid).Related(&review, "Review")