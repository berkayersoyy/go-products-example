package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}
