package main

import (
	"log"

	"github.com/berkay.ersoyy/go-products-example/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func setup() *gin.Engine {
	router := gin.Default()
	router.GET("/products", handlers.GetProducts)
	router.POST("/products", handlers.PostProduct)
	router.GET("/products/:id", handlers.GetProductById)
	router.DELETE("/products/:id", handlers.DeleteProduct)
	router.PUT("/products/:id", handlers.UpdateProduct)

	return router
}

func main() {
	r := setup()
	log.Fatal(r.Run("localhost:8080"))
}
