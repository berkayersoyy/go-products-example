package models

import "github.com/jinzhu/gorm"

// swagger:model User
type User struct {
	gorm.Model
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
