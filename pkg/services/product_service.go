package services

import (
	"github.com/berkayersoyy/go-products-example/pkg/models"
	"github.com/berkayersoyy/go-products-example/pkg/repositories"
)

type productService struct {
	ProductRepository repositories.ProductRepository
}
type ProductService interface {
	GetAllProducts() []models.Product
	GetProductByID(id uint) models.Product
	AddProduct(product models.Product) models.Product
	DeleteProduct(product models.Product)
}

func ProvideProductService(p repositories.ProductRepository) ProductService {
	return &productService{ProductRepository: p}
}

func (p *productService) GetAllProducts() []models.Product {
	return p.ProductRepository.GetAllProducts()
}

func (p *productService) GetProductByID(id uint) models.Product {
	return p.ProductRepository.GetProductByID(id)
}

func (p *productService) AddProduct(product models.Product) models.Product {
	p.ProductRepository.AddProduct(product)

	return product
}

func (p *productService) DeleteProduct(product models.Product) {
	p.ProductRepository.DeleteProduct(product)
}
