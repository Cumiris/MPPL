package models

import "github.com/jinzhu/gorm"

type Variant struct {
	gorm.Model

	Color   Color `json:"color"`
	ColorID uint `json:"color_id"`
	Size    Size `json:"size"`
	SizeID  uint `json:"size_id"`

	// Product has many variants, ProductID is the foreign key
	ProductID uint `json:"product_id"`
}
