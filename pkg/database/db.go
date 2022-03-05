package database

import (
	"github.com/berkayersoyy/go-products-example/pkg/models"
	config "github.com/berkayersoyy/go-products-example/pkg/utils/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type mysqlClient struct {
	SingletonMysql *gorm.DB
}

var mysqlclient mysqlClient

func GetMysqlClient(path string) mysqlClient {
	if mysqlclient.SingletonMysql == nil {
		mysqlclient.SingletonMysql = InitDb(path)
	}
	return mysqlclient
}

func InitDb(path string) *gorm.DB {
	conf, err := config.LoadConfig(path)
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
