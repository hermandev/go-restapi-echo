package models

import "time"

type Product struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	ProductName string    `json:"product_name"`
	ProducPrice int       `json:"product_price"`
	ProductDesc string    `json:"product_desc"`
	CreatedAt   time.Time `json:"created_at"`
}
