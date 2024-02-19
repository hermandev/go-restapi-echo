package controllers

import (
	"net/http"

	"github.com/hermandev/go-restapi-echo/models"
	"github.com/hermandev/go-restapi-echo/services"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	ProductService services.ProductService
}

func New(productService services.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (p *ProductController) CreateProduct(ctx echo.Context) error {
	product := new(models.Product)

	if err := ctx.Bind(product); err != nil {
		return err
	}

	if err := p.ProductService.Create(product); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, product)
}

func (p *ProductController) FindAll(ctx echo.Context) error {
	products, err := p.ProductService.FindAll()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, products)
}
