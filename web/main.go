package main

import (
	"github.com/berkay.ersoyy/go-products-example/pkg/database"
	"github.com/berkay.ersoyy/go-products-example/pkg/handlers"
	"github.com/berkay.ersoyy/go-products-example/pkg/repositories"
	"github.com/berkay.ersoyy/go-products-example/pkg/services"
	"github.com/berkay.ersoyy/go-products-example/pkg/validators"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setup(db *gorm.DB) *gin.Engine {

	productRepository := repositories.ProductRepository{DB: db}
	productService := services.ProductService{ProductRepository: productRepository}
	productApi := handlers.ProductAPI{ProductService: productService}
	// productApi := InitProductAPI(db)
	//TODO Dependency injection with wire.go but
	router := gin.Default()

	router.GET("/products", productApi.GetAllProducts)
	router.POST("/products", validators.ProductValidator(), productApi.AddProduct)
	router.GET("/products/:id", productApi.GetProductByID)
	router.DELETE("/products/:id", productApi.DeleteProduct)
	router.PUT("/products/:id", validators.ProductValidator(), productApi.UpdateProduct)

	return router
}

func main() {
	//v:=validator.New()
	db := database.InitDb()
	defer db.Close()
	r := setup(db)
	err := r.Run()
	if err != nil {
		panic(err)
	}

}
