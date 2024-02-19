package services

import (
	"github.com/hermandev/go-restapi-echo/dao"
	"github.com/hermandev/go-restapi-echo/models"
)

type ProductService struct {
	ProductDao dao.ProductDao
}

func New(productDao dao.ProductDao) *ProductService {
	return &ProductService{ProductDao: productDao}
}

func (ps *ProductService) Create(product *models.Product) error {
	return ps.ProductDao.Create(product)
}

func (ps *ProductService) FindAll() ([]models.Product, error) {
	return ps.ProductDao.FindAll()
}
