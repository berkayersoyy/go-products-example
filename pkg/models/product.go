package models

import (
	_ "github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name" validate:"required,min=2,max=45"`
	Price       float32 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}
