package repository

import (
	"MPPL/models"
	"github.com/jinzhu/gorm"
)
type brandRepo struct {
		DB *gorm.DB
}

func (r *brandRepo) Brands() []models.Brand {

	var bb []models.Brand
	r.DB.Find(&bb)

	return bb
}

func (r *brandRepo) Create(b models.Brand) models.Brand {
	r.DB.Create(&b)

	return b
}

func (r *brandRepo) Update(id int, b models.Brand) models.Brand {
	var brand models.Brand
	r.DB.First(&brand, id)

	r.DB.Model(&brand).Select([]string{"name", "update_at"}).
		Update(map[string]interface{}{"name": b.Name, "update_at" : b.UpdatedAt})

	return brand
}

func (r *brandRepo) Delete(id int) bool {
	var brand models.Brand
	r.DB.First(&brand, id)

	r.DB.Delete(&brand)

	return true
}

type BrandRepository interface {
	Brand(id int) models.Brand
	Brands()[]models.Brand
	Create(b models.Brand) models.Brand
	Update(id int, b models.Brand) models.Brand
	Delete(id int) bool
}

func newBrandRepo (DB *gorm.DB) BrandRepository {
	return &brandRepo{DB}
}

func (r *brandRepo) Brand(id int) models.Brand{
	var b models.Brand
	r.DB.Find(&b, id)

	return b

}
