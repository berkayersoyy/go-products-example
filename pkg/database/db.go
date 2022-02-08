package database

import (
	"github.com/berkayersoyy/go-products-example/pkg/models"
	config "github.com/berkayersoyy/go-products-example/pkg/utils/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// const DB_USERNAME = "root"
// const DB_PASSWORD = "x03121998X"
// const DB_NAME = "localhost"
// const DB_HOST = "127.0.0.1"
// const DB_PORT = "3306"

var singletonMysql *gorm.DB

func GetMysqlClient() *gorm.DB {
	if singletonMysql == nil {
		singletonMysql = InitDb()
	}
	return singletonMysql
}

func InitDb() *gorm.DB {
	conf, err := config.LoadConfig("./")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open("mysql", conf.MysqlDSN)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(5)

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.User{})

	return db
}
