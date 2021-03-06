package database

import (
	"fmt"
	"github.com/berkayersoyy/go-products-example/pkg/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
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
	DSN := os.Getenv("MYSQL_DSN")

	//conf, err := config.LoadConfig(path)
	//if err != nil {
	//	panic(err)
	//}
	db, err := gorm.Open("mysql", DSN)
	//ctx := context.Background()
	//if err := retry.Fibonacci(ctx, 1*time.Second, func(ctx context.Context) error {
	//	if err := db.DB().Ping(); err != nil {
	//		fmt.Println(err)
	//
	//		return retry.RetryableError(err)
	//	}
	//	return nil
	//}); err != nil {
	//	log.Fatal(err)
	//}
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(5)

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.User{})

	return db
}
