//go:build wireinject
// +build wireinject

package main

import (
	"github.com/berkay.ersoyy/go-products-example/pkg/handlers"
	"github.com/berkay.ersoyy/go-products-example/pkg/repositories"
	"github.com/berkay.ersoyy/go-products-example/pkg/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitProductAPI(db *gorm.DB) handlers.ProductAPI {
	wire.Build(repositories.ProvideProductRepository, services.ProvideProductService, handlers.ProvideProductAPI)

	return handlers.ProductAPI{}
}
