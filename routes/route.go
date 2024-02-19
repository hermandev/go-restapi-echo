package routes

import (
	"github.com/hermandev/go-restapi-echo/controllers"
	"github.com/hermandev/go-restapi-echo/dao"
	"github.com/hermandev/go-restapi-echo/db"
	"github.com/hermandev/go-restapi-echo/services"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	db, err := db.Init()
	if err != nil {
		panic("DB Connection Failed")
	}

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(200, "Hello, World!")
	})

	// product

	productDao := dao.New(db)
	productService := services.New(*productDao)
	productController := controllers.New(*productService)

	e.GET("/product", productController.FindAll)
	e.POST("/product", productController.CreateProduct)

	return e
}
