package db

import (
	"github.com/hermandev/go-restapi-echo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	dsn := "root:7qvt6t2738@tcp(127.0.0.1:3306)/dbproduct?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Product{})

	return db, nil
}
