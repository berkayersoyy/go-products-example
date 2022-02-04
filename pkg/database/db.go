package database

import (
	"fmt"
	"os"

	"github.com/berkayersoyy/go-products-example/pkg/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "x03121998X"
const DB_NAME = "mysql"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

func InitDb() *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	fmt.Println(dsn)
	db, err := gorm.Open("mysql", "root:x03121998X@tcp(mysql:3306)/godb")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.User{})

	return db
}
