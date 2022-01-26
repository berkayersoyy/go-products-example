package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/berkay.ersoyy/go-products-example/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Products)
}
func PostProduct(c *gin.Context) {
	var newProduct models.Product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}
	models.Products = append(models.Products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}
func GetProductById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, a := range models.Products {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Product not found"})
}
func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, a := range models.Products {
		if a.Id == id {
			// Db delete operation will be here
			remove(models.Products, i)
			c.Status(http.StatusOK)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Product not found"})
}

func UpdateProduct(c *gin.Context) {
	var newProduct models.Product
	var productToUpdate models.Product
	err := c.BindJSON(&newProduct)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	for _, v := range models.Products {
		if v.Id == id {
			productToUpdate = v
		}
	}
	productToUpdate.Name = newProduct.Name
	productToUpdate.Price = newProduct.Price
	productToUpdate.Description = newProduct.Description
	c.Status(http.StatusOK)
}
func remove(slice []models.Product, s int) []models.Product {
	return append(slice[:s], slice[s+1:]...)
}
