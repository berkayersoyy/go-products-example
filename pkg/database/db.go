package database

import (
	"github.com/berkay.ersoyy/go-products-example/pkg/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "x03121998X"
const DB_NAME = "godb"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

func InitDb() *gorm.DB {
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.User{})

	return db
}
