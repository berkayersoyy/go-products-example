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
type MysqlClient interface {
	GetClient() *gorm.DB
}

func ProvideMysqlClient(path string) MysqlClient {
	return &mysqlClient{SingletonMysql: InitDb(path)}
}

func (m *mysqlClient) GetClient() *gorm.DB {
	return m.SingletonMysql
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
