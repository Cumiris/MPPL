package models

import "github.com/jinzhu/gorm"

type Category struct {
gorm.Model

Name string `json:"name"`

// Brand has many products
//db.Model(&user).Related(&emails)
// SELECT * FROM emails WHERE user_id = 111; // 111 is user's primary key
Products []Product `json:"products"`
}