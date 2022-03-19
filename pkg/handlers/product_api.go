package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/berkayersoyy/go-products-example/pkg/dto"
	"github.com/berkayersoyy/go-products-example/pkg/mappers"
	"github.com/berkayersoyy/go-products-example/pkg/models"
	"github.com/berkayersoyy/go-products-example/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type productAPI struct {
	ProductService services.ProductService
}
type ProductAPI interface {
	GetAllProducts(c *gin.Context)
	GetProductByID(c *gin.Context)
	AddProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

func ProvideProductAPI(p services.ProductService) ProductAPI {
	return &productAPI{ProductService: p}
}

// @BasePath /api/v1

// GetAllProducts
// @Summary Fetch all product
// @Schemes
// @Description Fetch all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} models.Product
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/ [get]
func (p *productAPI) GetAllProducts(c *gin.Context) {
	products := p.ProductService.GetAllProducts()

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// @BasePath /api/v1

// GetProductByID
// @Summary Fetch product by id
// @Schemes
// @Description Fetch product by id
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/{id} [get]
func (p *productAPI) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.GetProductByID(uint(id))
	if product == (models.Product{}) {
		c.JSON(http.StatusNotFound, "Product not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": mappers.ToProductDTO(product)})
}

// @BasePath /api/v1

// AddProduct
// @Summary Add Product
// @Schemes
// @Description Add Product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product ID"
// @Success 200 {object} models.Product
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/ [post]
func (p *productAPI) AddProduct(c *gin.Context) {
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

// @BasePath /api/v1

// UpdateProduct
// @Summary Update Product
// @Schemes
// @Description Update Product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body dto.ProductDTO true "Product ID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/ [put]
func (p *productAPI) UpdateProduct(c *gin.Context) {
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

// @BasePath /api/v1

// DeleteProduct
// @Summary Delete Product
// @Schemes
// @Description Delete Product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/{id} [delete]
func (p *productAPI) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.GetProductByID(uint(id))
	if product == (models.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.ProductService.DeleteProduct(product)

	c.Status(http.StatusOK)
}
