package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/berkay.ersoyy/go-products-example/pkg/dto"
	"github.com/berkay.ersoyy/go-products-example/pkg/mappers"
	"github.com/berkay.ersoyy/go-products-example/pkg/models"
	"github.com/berkay.ersoyy/go-products-example/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductAPI struct {
	ProductService services.ProductService
}

func ProvideProductAPI(p services.ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

func (p *ProductAPI) GetAllProducts(c *gin.Context) {
	products := p.ProductService.GetAllProducts()

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (p *ProductAPI) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.GetProductByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": mappers.ToProductDTO(product)})
}

func (p *ProductAPI) AddProduct(c *gin.Context) {
	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	createdProduct := p.ProductService.AddProduct(product)

	c.JSON(http.StatusOK, gin.H{"product": mappers.ToProductDTO(createdProduct)})
}

func (p *ProductAPI) UpdateProduct(c *gin.Context) {
	var productDTO dto.ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(productDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.GetProductByID(uint(id))
	if product == (models.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	product.Name = productDTO.Name
	product.Price = productDTO.Price
	product.Description = productDTO.Description
	p.ProductService.AddProduct(product)

	c.Status(http.StatusOK)
}

func (p *ProductAPI) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.GetProductByID(uint(id))
	if product == (models.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.ProductService.DeleteProduct(product)

	c.Status(http.StatusOK)
}
