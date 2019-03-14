package models

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model

	ImageURL  string `json:"image_url"`

	// Product has many images, ProductID is the foreign key
	ProductID uint	`json:"product_id"`
}
