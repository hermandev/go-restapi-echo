package dao

import (
	"github.com/hermandev/go-restapi-echo/models"
	"gorm.io/gorm"
)

type ProductDao struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *ProductDao {
	return &ProductDao{DB: db}
}

func (p *ProductDao) Create(product *models.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductDao) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := p.DB.Find(&products).Error
	return products, err
}
