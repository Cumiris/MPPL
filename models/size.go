package models

import "github.com/jinzhu/gorm"

type Size struct {
	gorm.Model

	Name string `json:"name"`

	// Size belongs to Category
	// db.Model(&user).Related(&profile)
	// SELECT * FROM profiles WHERE id = 111;
	// 111 is user's foreign key ProfileID
	Category   Category `json:"category"`
	CategoryID uint `json:"category_id"`
}
